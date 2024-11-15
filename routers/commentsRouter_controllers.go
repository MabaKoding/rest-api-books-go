package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["books_crud/controllers:BookController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BookController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BookController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BookController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BookController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BookController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:bookIsbn",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BookController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BookController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:bookIsbn/remove",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BookController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BookController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:bookIsbn/update",
            AllowHTTPMethods: []string{"PUT"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BooksController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BooksController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BooksController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BooksController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BooksController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BooksController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["books_crud/controllers:BooksController"] = append(beego.GlobalControllerRouter["books_crud/controllers:BooksController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
