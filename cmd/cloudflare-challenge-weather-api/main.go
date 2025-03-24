package main

import (
	"cloudflare-challenge-weaher-api/internal/handlers"
	"cloudflare-challenge-weaher-api/pkg/utils"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echoInstance := echo.New()
	echoInstance.Use(utils.SecurityHeaders)
	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Recover())

	echoInstance.POST(utils.LoginEndPoint, handlers.Login)

	echoInstance.GET(utils.WeatherEndPoint, handlers.GetWeather, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(utils.SecSigningKey),
	}))

	echoInstance.Logger.Fatal(echoInstance.Start(utils.APIPort))
}
