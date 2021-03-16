package product

import (
	"com.mzq.echo/app/core"
	"context"
	"github.com/labstack/echo/v4"
)


func List(c echo.Context) error {
	client := core.Elastic()
	p1 := Person{Name: "lmh", Age: 18, Married: false}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	//listUser, count, err := models.ListUser(models.UserQuery{
	//	UserName: c.Request().FormValue("username"),
	//})
	//if err != nil {
	//	return response.Json(c, listUser, map[string]interface{}{
	//		"count" : count,
	//	})
	//}
	//return response.Json(c, listUser, map[string]interface{}{
	//	"count" : count,
	//})
}