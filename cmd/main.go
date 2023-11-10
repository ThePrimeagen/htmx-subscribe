package main

import (
	"html/template"
	"io"
	"net/mail"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	tmpls *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpls.ExecuteTemplate(w, name, data)
}

type IndexPage struct {
	Email     string
	ErrorMsgs map[string]string
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main() {
	e := echo.New()

	e.Renderer = &Template{
		tmpls: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Use(middleware.Logger())
    e.Static("/src", "src")

	e.GET("/", func(c echo.Context) error {
		email := c.QueryParam("email")
		return c.Render(200, "index", IndexPage{Email: email})
	})

	e.POST("/subscribe", func(c echo.Context) error {
        time.Sleep(1 * time.Second)
		email := c.FormValue("email")
		if !validEmail(email) {
            c.Logger().Error("Invalid email address: " + email)
			return c.Render(422, "index", IndexPage{
				Email: email,
				ErrorMsgs: map[string]string{
					"email": "Invalid email address",
				},
			})
		}

		return c.Redirect(303, "/subscribe?email="+email)
	})

	e.GET("/subscribe", func(c echo.Context) error {
		email := c.FormValue("email")
		return c.Render(200, "thanks", IndexPage{Email: email})
	})

	e.Logger.Fatal(e.Start(":42069"))
}
