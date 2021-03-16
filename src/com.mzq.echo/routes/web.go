package router

import (
	admin2 "com.mzq.echo/app/http/controller/admin"
	"com.mzq.echo/app/http/controller/captcha"
	"com.mzq.echo/app/http/controller/index"
	"com.mzq.echo/app/http/exception"
	"com.mzq.echo/app/http/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)


func Router(ec *echo.Echo)  {
	// Routes
	ec.GET("/", index.Hello)


	ec.Static("/js/*", "public/js")
	ec.Static("/json/*", "public/json")
	ec.Static("*.ico", "public")

	//ec.GET("/js/*", func(context echo.Context) error {
	//	//filepath.Base("")
	//	file, err := filepath.Abs("public" + context.Request().URL.Path)
	//	if err == nil {
	//		openFile, err := os.Open(file)
	//		if err == nil {
	//			bytes := make([]byte, 4096)
	//			var b []byte
	//			defer openFile.Close()
	//			for  {
	//				red, err := openFile.Read(bytes)
	//				if err == io.EOF {
	//					break
	//				}
	//				b = append(b, bytes[:red]...)
	//			}
	//			if file[len(file) - 3:] == "css" {
	//				return context.Blob(http.StatusOK,"text/css", b)
	//			}else {
	//				return context.Blob(http.StatusOK,echo.MIMEApplicationJavaScript, b)
	//			}
	//
	//			//return context.String(http.StatusOK, string(b))
	//		}
	//	}
	//	return context.String(http.StatusNotFound, "")
	//})


	ec.GET("/login", index.LoginView)
	ec.PUT("/login", index.LoginIn)

	ec.GET("/captcha", captcha.CodeImage)

	admin := ec.Group("/admin")
	admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(models.AppKey()),
		TokenLookup: "cookie:token",
		ContextKey:models.APP_UID,
		ErrorHandlerWithContext: func(err error, context echo.Context) error {
			return context.Redirect(http.StatusFound, "/login#/user/login")
		},
	}))
	admin.GET("/index", admin2.Index)

	api := ec.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(models.AppKey()),
		TokenLookup: "cookie:token",
		ContextKey:models.APP_UID,
		ErrorHandlerWithContext: func(err error, context echo.Context) error {
			return exception.NewExceptionMzq(exception.ERROR_CREATE_TOKEN, "请登录")
		},
	}))
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			err := models.CheckLogin(context)
			if err != nil {
				return err
			}
			return next(context)
		}
	})
	api.GET("/user", admin2.User)
	api.GET("/menu", index.Menu)
	api.GET("/admin/list", admin2.List)
	api.POST("/admin/user", admin2.Create)
	api.PUT("/admin/user", admin2.Update)

	//
	//g := ec.Group("/admin")
	//g.Use(middleware.BasicAuth(func(username, password string,
	//	c echo.Context) (bool, error) {
	//	if username == "joe" && password == "secret" {
	//		return true, nil
	//	}
	//	return false, nil
	//}))
	//
	//// Route level middleware
	//track := func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		println("request to /users")
	//		return next(c)
	//	}
	//}
	//ec.GET("/users", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "/users")
	//}, track)
}

