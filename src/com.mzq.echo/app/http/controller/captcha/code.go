package captcha

import (
	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CodeImage(c echo.Context) error  {
	id := captcha.New()
	c.SetCookie(&http.Cookie{
		Name:"captchaId",
		Value:id,
	})
	c.Request().URL.Path = c.Request().URL.Path + "/" + id + ".png"
	//c.Request()
	width, err := strconv.Atoi(c.QueryParam("w"))
	if err != nil {
		width = captcha.StdWidth
	}

	heigth, err := strconv.Atoi(c.QueryParam("h"))
	if err != nil {
		heigth = captcha.StdHeight
	}

	server := captcha.Server(width, heigth)
	server.ServeHTTP(c.Response().Writer, c.Request())
	return nil
}