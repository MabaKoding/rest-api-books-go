package controllers

import (
	"books_crud/models"
	"database/sql"
	"encoding/json"

	"github.com/beego/beego/validation"
)

type (
	BookController struct {
		MainController
	}
	BookRequestUpdate struct {
		BooksTitle       string `json:"books_title"`
		BooksSubtitle    string `json:"books_subtitle"`
		BooksAuthor      string `json:"books_author"`
		BooksDescription string `json:"books_description"`
		BooksPublished   string `json:"books_published"`
		BooksPublisher   string `json:"books_publisher"`
	}
)

// GET GetAll
// @Title Get All
// @Description API to get all of list book
// @Success 200 {object} models.BookData
// @Failure 403
// @router / [get]
func (c *BookController) Get() {
	defer c.ServeJSON()
	output := make(map[string]interface{}, 0)
	output["success"] = false

	var err error

	bookModel := new(models.BookModel)
	params := make(map[string]interface{})
	// params["merchant_id"] = merchantIdInteger
	res, err := bookModel.GetAllCollection(params)

	if err != nil {
		if err == sql.ErrNoRows {
			// ZapLogger.Error("user not found with id " + userId + " " + err.Error())
			c.Ctx.Output.SetStatus(404)
			output["error"] = "Data tidak ditemukan"
			c.Data["json"] = output
			return
		}

		// ZapLogger.Error("failed get user with id " + userId + " " + err.Error())
		c.Ctx.Output.SetStatus(500)
		output["error"] = "Kesalahan server"
		output["errors"] = err
		c.Data["json"] = output
		return
	}

	c.Ctx.Output.SetStatus(200)
	output["success"] = true
	output["result"] = res
	c.Data["json"] = output
}

// Get GetOne
// @Title Get One
// @Description API to get one data book
// @Param	bookIsbn		path 	string	true  "the book isbn"
// @Success 200 {string} models.BookData
// @Failure 403 body is empty
// @router /:bookIsbn [get]
func (c *BookController) GetOne() {
	defer c.ServeJSON()
	output := make(map[string]interface{}, 0)
	output["success"] = false
	var err error

	//URI VALIDATE
	bookId := c.Ctx.Input.Param(":bookIsbn")
	if bookId == "" {
		// ZapLogger.Error("merchant id " + strconv.FormatInt(merchantIdInteger, 64) + " is invalid")
		c.Ctx.Output.SetStatus(400)
		output["error"] = "book isbn salah"
		c.Data["json"] = output
		return
	}

	bookModel := new(models.BookModel)

	obj, err := bookModel.GetObjectByParams(map[string]interface{}{
		"isbn": bookId,
	})

	if err != nil && obj == nil {
		// ZapLogger.Error("error get data : " + err.Error())
		c.Ctx.Output.SetStatus(404)
		output["error"] = "Data tidak ditemukan"
		c.Data["json"] = output
		return
	}

	output["success"] = true
	output["result"] = obj
	c.Data["json"] = output
}

// POST PostData
// @Title Post
// @Description create Books
// @Param	body		body 	models.BookData	true		"body for Books content"
// @Success 201 {int} models.BookData
// @Failure 403 body is empty
// @router / [post]
func (c *BookController) Post() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false

	inputData := new(models.BookData)
	inputBody := c.Ctx.Input.RequestBody

	json.Unmarshal(inputBody, &inputData)

	valid := validation.Validation{}
	valid.Required(inputData.Id, "ISBN")
	valid.Required(inputData.BooksTitle, "Title")
	valid.Required(inputData.BooksSubtitle, "Sub Title")
	valid.Required(inputData.BooksAuthor, "Author")
	valid.Required(inputData.BooksDescription, "Description")
	valid.Required(inputData.BooksPublisher, "Publisher")
	valid.Required(inputData.BooksPublished, "Published")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Ctx.Output.SetStatus(404)
			output["error"] = err.Key + "" + err.Message
			c.Data["json"] = output
			return
		}
	}

	bookModel := new(models.BookModel)

	existing, err := bookModel.GetCountByIsbn(inputData.Id)

	if err != nil {
		output["error"] = err
		output["success"] = false
		c.Data["json"] = output
		return
	}
	if existing > 0 {
		output["error"] = "ISBN " + inputData.Id + " sudah terdaftar"
		output["success"] = false
		c.Data["json"] = output
		return
	}

	saved, err := bookModel.CreateObject(inputData)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		// ZapLogger.Error(err.Error())
		output["error"] = "Gagal menyimpan data"
		c.Data["json"] = output
		return
	}

	output["object"] = saved
	output["success"] = true
	c.Data["json"] = output
}

// Update UpdateData
// @Title Update user data
// @Description Update user data
// @Param	bookIsbn		path 	string	true  "the book isbn"
// @Param   body        body    controllers.BookRequestUpdate   true        "body for Books content"
// @Success 200 {string} models.BookData
// @Failure 403 body is empty
// @router /:bookIsbn/update [PUT]
func (c *BookController) Update() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false

	var err error

	//URI VALIDATE
	bookId := c.Ctx.Input.Param(":bookIsbn")
	if bookId == "" {
		// ZapLogger.Error("merchant id " + strconv.FormatInt(merchantIdInteger, 64) + " is invalid")
		c.Ctx.Output.SetStatus(400)
		output["error"] = "book isbn salah"
		c.Data["json"] = output
		return
	}

	param := c.Ctx.Input.RequestBody
	reqData := BookRequestUpdate{}
	// get data from request
	json.Unmarshal(param, &reqData)

	valid := validation.Validation{}
	valid.Required(reqData.BooksTitle, "Title")
	valid.Required(reqData.BooksSubtitle, "Sub Title")
	valid.Required(reqData.BooksAuthor, "Author")
	valid.Required(reqData.BooksDescription, "Description")
	valid.Required(reqData.BooksPublisher, "Publisher")
	valid.Required(reqData.BooksPublished, "Published")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Ctx.Output.SetStatus(404)
			output["error"] = err.Key + "" + err.Message
			c.Data["json"] = output
			return
		}
	}
	bookModel := new(models.BookModel)
	res, err := bookModel.UpdateObject(bookId, map[string]interface{}{
		"books_author":      reqData.BooksAuthor,
		"books_title":       reqData.BooksTitle,
		"books_subtitle":    reqData.BooksSubtitle,
		"books_publisher":   reqData.BooksPublisher,
		"books_published":   reqData.BooksPublished,
		"books_description": reqData.BooksDescription,
	})

	if err != nil {
		// ZapLogger.Error("invalid updated user " + err.Error())
		c.Ctx.Output.SetStatus(500)
		output["error"] = err.Error()
		c.Data["json"] = output
		return
	}

	c.Ctx.Output.SetStatus(200)
	output["success"] = true
	output["object"] = res
	c.Data["json"] = output
}

// Delete DeleteData
// @Title Remove Book
// @Description Remove books from the list
// @Param	bookIsbn		path 	string	true  "the book isbn"
// @Success 200 {object} controllers.OutputDeleteResponse
// @Failure 404 :id is not found
// @router /:bookIsbn/remove [delete]
func (c *BookController) Delete() {
	defer c.ServeJSON()
	output := make(map[string]interface{})
	output["success"] = false

	//URI VALIDATE
	bookId := c.Ctx.Input.Param(":bookIsbn")
	if bookId == "" {
		c.Ctx.Output.SetStatus(500)
		output["error"] = "ISBN tidak ditemukan"
		c.Data["json"] = output
		return
	}

	bookModel := new(models.BookModel)

	res, errDel := bookModel.DeleteObject(bookId)
	if errDel != nil {
		c.Ctx.Output.SetStatus(500)
		output["error"] = "Gagal menghapus data"
		c.Data["json"] = output
		return
	}
	c.Ctx.Output.SetStatus(200)
	output["success"] = true
	output["object"] = res
	c.Data["json"] = output
}
