package main

import (
	"fmt"
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
	funcMap := template.FuncMap{
		"dict": func(values ...any) (map[string]any, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("dict requires even number of args")
			}
			m := make(map[string]any, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("dict keys must be strings")
				}
				m[key] = values[i+1]
			}
			return m, nil
		},
	}

	tmpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("views/*.html"))
	template.Must(tmpl.ParseGlob("views/components/*.html"))
	return &Templates{templates: tmpl}
}

func main() {
	e := echo.New()

	e.Debug = os.Getenv("DEBUG") == "true"
	e.HideBanner = true

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.Renderer = newTemplate()

	e.Static("/resources", "resources")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", nil)
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
