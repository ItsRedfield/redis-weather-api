package utils

import (
	"github.com/labstack/echo/v4"
)

func SecurityHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoInstance echo.Context) error {
		echoInstance.Response().Header().Set(ContentType, ContentTypeVal)
		echoInstance.Response().Header().Set(FrameOptions, FrameOptionsVal)
		echoInstance.Response().Header().Set(SecurityPolicy, SecurityPolicyVal)
		echoInstance.Response().Header().Set(XSSProtection, XSSProtectioVal)
		echoInstance.Response().Header().Set(ReferrerPolicy, ReferrerPolicyVal)
		echoInstance.Response().Header().Set(PermissionsPolicy, PermissionsPolicyVal)

		return next(echoInstance)
	}
}
