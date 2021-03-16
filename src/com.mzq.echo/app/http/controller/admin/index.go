package admin

import (
	"com.mzq.echo/app/core/response"
	"com.mzq.echo/app/http/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {

	render := response.Render("admin/index.tpl.html")
	return c.HTML(http.StatusOK, render)
}

func User(c echo.Context) error {
	user := models.LoginUser()
	//return response.Json(c, map[string]interface{
	//}{
	//	"username" : user.UserName,
	//	"Id" : user.Id,
	//})
	return response.Json(c, user)
}

func List(c echo.Context) error {
	listUser, count, err := models.ListUser(models.UserQuery{
		UserName: c.Request().FormValue("username"),
	})
	if err != nil {
		return response.Json(c, listUser, map[string]interface{}{
			"count" : count,
		})
	}
	return response.Json(c, listUser, map[string]interface{}{
		"count" : count,
	})
}

func Create(c echo.Context) error {
	user, err := models.CreateUser(
		c.Request().FormValue("username"),
		c.Request().FormValue("password"),
		)
	if err != nil {
		return response.Json(c, user)
	}
	return response.Json(c, user)
}


func Update(c echo.Context) error {
	uid := c.Request().FormValue("uid")
	user, err2 := models.FindUidString(uid)
	if err2 != nil{
		return err2
	}
	user, err := user.UpdateUser(
		c.Request().FormValue("username"),
		c.Request().FormValue("password"),
	)
	if err != nil {
		return err
	}
	return response.Json(c, user)
}