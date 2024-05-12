package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Cookie(Key string, Value string) *http.Cookie {
	// Create a new cookie
	return &http.Cookie{
		Name:     Key,
		Value:    Value,
		MaxAge:   3600,                    // Cookie expiration time in seconds
		HttpOnly: true,                    // HttpOnly flag
		Secure:   true,                    // Secure flag (HTTPS only)
		SameSite: http.SameSiteStrictMode, // SameSite attribute
	}
}

func GetCookie(c echo.Context, Key string, Default string) string {
	v, e := c.Cookie(Key)
	if e != nil {
		return Default
	}
	return v.Value
	//
}
