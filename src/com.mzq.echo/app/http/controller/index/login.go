package index

import (
	"com.mzq.echo/app/core/response"
	"com.mzq.echo/app/http/exception"
	"com.mzq.echo/app/http/models"
	"com.mzq.echo/app/http/rule"
	"github.com/dchest/captcha"
	"github.com/flosch/pongo2"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/labstack/echo/v4"
	"net/http"
)



// Handler
func LoginView(c echo.Context) error {


	//return c.Redirect(http.StatusMovedPermanently, "<URL>")

	render := response.Render("login/login.tpl.html",pongo2.Context{
		"captchaId": captcha.New(),
	})
	return c.HTML(http.StatusOK, render)
}

func LoginInValidate(c echo.Context) (struct {
	username string;
	password string;
	vercode  string
}, error) {
	username := c.Request().PostFormValue("username")
	password := c.Request().PostFormValue("password")
	vercode := c.Request().PostFormValue("vercode")
	requestInfo := struct {
		username string
		password string
		vercode  string
	}{
		username,
		password,
		vercode,
	}
	err := validation.ValidateStruct(&requestInfo,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&requestInfo.username, validation.Required, validation.Length(5, 20)),
		// Gender is optional, and should be either "Female" or "Male".
		validation.Field(&requestInfo.password, validation.Required, validation.Length(3, 20)),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&requestInfo.vercode, validation.Required, validation.By(rule.CaptchaEquals(c))),
	)

	return requestInfo, err
}


func LoginIn(c echo.Context)error  {
	validate, err2 := LoginInValidate(c)
	if err2 != nil {
		return c.JSON(http.StatusUnauthorized, exception.NewExceptionMzq(exception.ERROR_CAPTCHA, err2.Error()))
	}
	user, err := models.Verify(validate.username, validate.password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}
	token, err2 := user.CreateToken()
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	c.SetCookie(&http.Cookie{
		Name:"token",
		Value:token,
	})
	return response.Json(c, map[string]interface{}{
		"HttpCode": http.StatusOK,
		"Code" :0,
	})
	//return c.JSON(http.StatusOK)
	//return c.Redirect(http.StatusFound, "/")
}