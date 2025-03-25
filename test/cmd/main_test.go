package main

import (
	"cloudflare-challenge-weaher-api/internal/handlers"
	"cloudflare-challenge-weaher-api/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *echo.Echo {
	e := echo.New()
	e.Use(utils.SecurityHeaders)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST(utils.LoginEndPoint, handlers.Login)

	e.GET(utils.WeatherEndPoint, handlers.GetWeather, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(utils.SecSigningKey),
	}))

	return e
}

func TestRoutes(t *testing.T) {
	e := setupRouter()

	req := httptest.NewRequest(http.MethodPost, utils.LoginEndPoint, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, utils.WeatherEndPoint, nil)
	rec = httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	token := utils.SecSigningKey
	req = httptest.NewRequest(http.MethodGet, utils.WeatherEndPoint, nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	rec = httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
