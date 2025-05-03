package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"gpt-service-go/client"

	"github.com/labstack/echo/v4"
)

func JWTValidator(javaClient *client.JavaClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header required")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header format")
			}

			token := parts[1]
			isValid, err := javaClient.ValidateToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
			}

			if !isValid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			return next(c)
		}
	}
}
