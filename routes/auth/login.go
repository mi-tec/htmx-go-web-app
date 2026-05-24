package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	test := map[string]string{
		"name": "John Doe",
	}

	return c.JSON(http.StatusOK, test)
}
