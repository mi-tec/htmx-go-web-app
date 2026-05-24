package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"htmx-go-web-app/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()

	e.Debug = os.Getenv("DEBUG") == "true"
	e.HideBanner = true

	e.Use(middleware.RequestLogger()) // use the RequestLogger middleware with slog logger
	e.Use(middleware.Recover())       // recover panics as errors for proper error handling

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/about-us", func(c echo.Context) error {
		return c.Render(200, "about-us.html", nil)
	})

	// Routes

	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "200",
			"mesage": "Hello World!",
		})
	})

	handler.RoutesHandler(e)

	if err := e.Start(":3000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
