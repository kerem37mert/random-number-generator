package main

import (
	"html/template"
	"io"
	"randomnumber/handlers"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Static("/static", "static")
	renderer := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", handlers.Home)
	e.GET("/api/randomNumber", handlers.RandomNumber)

	e.Logger.Fatal(e.Start(":8000"))
}
