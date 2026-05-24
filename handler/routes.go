package handler

import (
	"htmx-go-web-app/routes/auth"

	"github.com/labstack/echo/v4"
)

func RoutesHandler(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")

	authGroup.POST("/login", echo.HandlerFunc(auth.LoginHandler))
}
