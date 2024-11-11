package main

import (
	"books_crud/controllers"
	_ "books_crud/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
)

func main() {
	// sqlConn, err := beego.AppConfig.String("sqlconn")
	// if err != nil {
	// 	log.Fatal(err)
	// 	// log.Fatal(http.ListenAndServe(":8080", nil))
	// }
	// orm.RegisterDataBase("default", "postgres", sqlConn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/apidocs/here"] = "swagger"
		beego.ErrorController(&controllers.ErrorController{})
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	httpPort, _ := beego.AppConfig.String("httpport")
	beego.Run(":" + httpPort)
}
