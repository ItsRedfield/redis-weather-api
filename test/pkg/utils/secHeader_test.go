package utils_test

import (
	"cloudflare-challenge-weaher-api/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSecurityHeaders(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	next := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	handler := utils.SecurityHeaders(next)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, utils.ContentTypeVal, rec.Header().Get(utils.ContentType))
		assert.Equal(t, utils.FrameOptionsVal, rec.Header().Get(utils.FrameOptions))
		assert.Equal(t, utils.SecurityPolicyVal, rec.Header().Get(utils.SecurityPolicy))
		assert.Equal(t, utils.XSSProtectioVal, rec.Header().Get(utils.XSSProtection))
		assert.Equal(t, utils.ReferrerPolicyVal, rec.Header().Get(utils.ReferrerPolicy))
		assert.Equal(t, utils.PermissionsPolicyVal, rec.Header().Get(utils.PermissionsPolicy))
	}
}
