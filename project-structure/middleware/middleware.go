package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func APIKeyMiddleware(expectedKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get("X-API-Key")
			if apiKey == "" {
				apiKey = c.FormValue("api_key")
			}

			if apiKey == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "missing api key",
				})
			}

			if apiKey != expectedKey {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid api key",
				})
			}

			// key is valid → continue to handler
			return next(c)
		}
	}
}
