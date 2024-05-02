package main

import (
	"html/template"
	"io"
	"rulette_rousse/lib"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Creation du type templates
type templates struct {
	templates *template.Template
	// game.baptiste()
}

func (t *templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *templates {
	return &templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.GET("/public/*", func(c echo.Context) error {
		//get route path
		path := c.Request().URL.Path
		//return file
		return c.File("." + path)
	})

	// Page d'accueil du site
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", 0)
	})

	// Page d'accueil du site
	e.GET("/game", func(c echo.Context) error {
		return c.Render(200, "game", 0)
	})

	// Page d'accueil du site
	e.GET("/homePage", func(c echo.Context) error {
		return c.Render(200, "homePage", 0)
	})

	// Page de jouer avec un ami
	e.GET("/friendInvite", func(c echo.Context) error {
		return c.Render(200, "playFriendPage", 0)
	})

	// Page de jouer avec un ami
	e.GET("/createGame", func(c echo.Context) error {
		// code := Game.getCode()
		return c.Render(200, "createGame", 0)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
