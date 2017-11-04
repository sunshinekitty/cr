package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Version returns server version
func Version(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte(`{"version": "0.0.0"}`))
}
