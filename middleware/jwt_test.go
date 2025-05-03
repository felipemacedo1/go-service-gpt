package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJWTValidator(t *testing.T) {
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" // Example valid token
	invalidToken := "invalid-token"
	emptyToken := ""

	testCases := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "Valid Token",
			token:          validToken,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Token",
			token:          invalidToken,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Missing Token",
			token:          emptyToken,
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Create a handler to be protected by the middleware
			handler := func(c echo.Context) error {
				return c.String(http.StatusOK, "OK")
			}

			// Apply the middleware and test if the handler gets called or not
			middleware := JWTValidator(strings.NewReader(`{ "url": "http://example.com" }`))
			err := middleware(handler)(c)

			if err != nil {
				httpError, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, tc.expectedStatus, httpError.Code)
				}
			} else {
				assert.Equal(t, tc.expectedStatus, rec.Code)
			}
		})
	}
}