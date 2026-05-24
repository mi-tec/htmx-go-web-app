package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(c *echo.Context, w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger()) // use the RequestLogger middleware with slog logger
	e.Use(middleware.Recover())       // recover panics as errors for proper error handling

	e.Renderer = newTemplate()

	e.GET("/", func(c *echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/about-us", func(c *echo.Context) error {
		return c.Render(200, "about-us.html", nil)
	})

	if err := e.Start(":1149"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
