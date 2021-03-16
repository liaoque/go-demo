package index

import (
	"com.mzq.echo/app/core/response"
	"com.mzq.echo/app/http/models"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Index struct {
	
}

// Handler
func Hello(c echo.Context) error {
	//_, err := c.Cookie("token")
	//if err != nil {
	//	return c.Redirect(http.StatusFound, "/login#/user/login")
	//}
	user := models.LoginUser()
	render := response.Render("index/index.tpl.html", pongo2.Context{user})
	return c.HTML(http.StatusOK, render)
}


func Menu(c echo.Context) error {
	render := response.Render("index/menu.tpl.html")
	return c.HTML(http.StatusOK, render)
}