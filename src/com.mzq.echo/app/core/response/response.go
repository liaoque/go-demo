package response

import (
	"com.mzq.echo/app/http/exception"
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"net/http"
)

func init()  {

}

func Render(tpl string, contextArray ...pongo2.Context) string {
	context := pongo2.Context{}
	if len(contextArray) > 0 {
		context = contextArray[0]
	}
	file, err := pongo2.RenderTemplateFile("src/com.mzq.echo/app/view/" + tpl, context)
	if err == nil{
		return file
	}
	panic(fmt.Sprintln("模板 " + tpl+ " 渲染错误"))
	return ""
}

func MzqHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*exception.ExceptionMzq)
	if !ok {
		he = &exception.ExceptionMzq{
			HttpCode:  http.StatusInternalServerError,
			Code:    http.StatusInternalServerError,
			Mes: http.StatusText(http.StatusInternalServerError),
		}
	}

	// Issue #1426
	code := he.HttpCode


	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, he)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}

func Json(c echo.Context, i interface{}, mp ...map[string]interface{}) error {
	m := map[string]interface{}{
		"HttpCode": http.StatusOK,
		"code":     0,
		"data":     i,
	}
	//r := result
	if len(mp) >0 {
		for k,v := range mp[0]{
			m[k] = v
		}
	}


	return c.JSON(http.StatusOK, m)
}