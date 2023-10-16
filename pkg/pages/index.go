package pages

import (
	"net/mail"
	"time"

	"github.com/labstack/echo/v4"
)

type Page struct {
    Email string `form:"email"`
    ErrorMsg string
}

func Index(c echo.Context) error {
    email := c.QueryParam("email")
    return c.Render(200, "index.html", Page {
        Email: email,
        ErrorMsg: "",
    })
}


func valid(email string) bool {
    _, err := mail.ParseAddress(email)

    return err == nil
}

func Subscribed(c echo.Context) error {
    var subscribePage Page
    c.Logger().Error("Subscribe Called")
    if err := c.Bind(&subscribePage); err != nil {
        c.Logger().Error("No email address provided")
        c.Render(200, "index.html", Page{
            Email: "",
            ErrorMsg: "Please provide an email",
        })
    }

    c.Logger().Error("Subscribed called with %s", subscribePage.Email)

    if !valid(subscribePage.Email) {
        c.Logger().Error("invalid email")
        c.Render(200, "index.html", Page{
            Email: subscribePage.Email,
            ErrorMsg: "Invalid Email",
        })
    }

    time.Sleep(1 * time.Second)
    return c.Render(200, "subscribed.html", Page{
        Email: subscribePage.Email,
        ErrorMsg: "",
    })
}

