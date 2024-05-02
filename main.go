package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

// Creation du type templates
type templates struct {
	templates *template.Template
}

func (t *templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *templates {
	return &templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type User struct {
	id int
}

func getUser() User {
	return User{
		id: 10,
	}
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	userId := getUser()

	e.Renderer = newTemplate()

	// Page d'accueil du site
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", userId)

	})

	e.Logger.Fatal(e.Start(":8080"))
}
