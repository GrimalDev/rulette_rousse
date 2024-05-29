package main

import (
	"html/template"
	"io"
	game "rulette_rousse/lib"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Renderer = newTemplate()

	e.GET("/public/*", func(c echo.Context) error {
		//get route path
		path := c.Request().URL.Path
		//return file
		return c.File("." + path)
	})

	// Page d'accueil du site
	e.GET("/", func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		if sess.IsNew {
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400 * 7,
				HttpOnly: true,
			}
			sess.Values["player"] = "id"
			sess.Values["code"], _ = game.RandomHex(3)
			if err := sess.Save(c.Request(), c.Response()); err != nil {
				return err
			}
		}

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
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		game := game.CreateGame(sess)
		return c.Render(200, "createGame", game)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
