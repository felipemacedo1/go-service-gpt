package middleware

import (
	"fmt"
	"gpt-service-go/client"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type JWTValidatorMiddleware struct {
	JavaClient *client.JavaClient
	Logger     *logrus.Logger
}

func NewJWTValidator(javaClient *client.JavaClient, logger *logrus.Logger) *JWTValidatorMiddleware {
	return &JWTValidatorMiddleware{JavaClient: javaClient, Logger: logger}
}

func (m *JWTValidatorMiddleware) JWTValidator() echo.MiddlewareFunc {
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
			isValid, err := m.JavaClient.ValidateToken(token)
			if err != nil {
				m.Logger.WithError(err).Error("Failed to validate token")
				return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
			}

			if !isValid {
				m.Logger.Error("Invalid token")
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			return next(c)
		}
	}
}
