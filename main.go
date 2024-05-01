package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

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

type Contact struct {
	Name  string
	Email string
}

func NewContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func NewData() Data {
	return Data{
		Contacts: []Contact{
			NewContact("john", "aaa@b.com"),
			NewContact("billy", "bbb@b.com"),
			NewContact("ewan", "ccc@b.com"),
			NewContact("bob", "ddd@b.com"),
		},
	}
}

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	data := NewData()
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		data.Contacts = append(data.Contacts, NewContact(name, email))

		return c.Render(200, "index", data)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
