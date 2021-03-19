package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	firebase "firebase.google.com/go/v4"
)

const authHeaderName = "Authorization"

func KeyAuth(firebase *firebase.App) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:" + authHeaderName,
		Skipper: func(c echo.Context) bool {
			return c.Request().Header.Get(authHeaderName) == ""
		},
		Validator: func(idToken string, c echo.Context) (bool, error) {
			auth, err := firebase.Auth(c.Request().Context())
			if err != nil {
				return false, err
			}

			token, err := auth.VerifyIDToken(c.Request().Context(), idToken)

			if err != nil {
				return false, nil
			}

			c.Set("token", token)

			return true, nil
		},
	})
}
