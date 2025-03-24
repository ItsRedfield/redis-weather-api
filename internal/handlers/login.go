package handlers

import (
	"cloudflare-challenge-weaher-api/pkg/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(request echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[utils.Usr] = utils.User
	claims[utils.Expiry] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(utils.SecSigningKey))
	if err != nil {
		return request.JSON(http.StatusInternalServerError, map[string]string{
			utils.Error: utils.ErrorFailTokenGen,
		})
	}

	return request.JSON(http.StatusOK, map[string]string{
		utils.Token: tokenString,
	})
}
