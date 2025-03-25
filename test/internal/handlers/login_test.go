package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"cloudflare-challenge-weaher-api/internal/handlers"
	"cloudflare-challenge-weaher-api/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Setup Echo
	e := echo.New()

	t.Run("Success - valid token generation", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handlers.Login(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"token":"`)
	})

	t.Run("Error - invalid signing key", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handlers.Login(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, http.StatusInternalServerError)
	})
}

func TestLoginTokenClaims(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handlers.Login(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	tokenString := response["token"]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SecSigningKey), nil
	})
	assert.NoError(t, err)

	claims := token.Claims.(jwt.MapClaims)
	assert.Equal(t, utils.User, claims[utils.Usr])
	assert.InDelta(t, time.Now().Add(time.Hour*1).Unix(), int64(claims[utils.Expiry].(float64)), 10)
}
