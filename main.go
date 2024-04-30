package main

import (
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
}
