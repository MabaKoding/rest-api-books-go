package main

import (
	_ "books_crud/routers"
	"log"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func main() {
	sqlConn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		log.Fatal(err)
		// log.Fatal(http.ListenAndServe(":8080", nil))
	}
	orm.RegisterDataBase("default", "postgres", sqlConn)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/apidocs/here"] = "swagger"
	}
	beego.Run()
}
