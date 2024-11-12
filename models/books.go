package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type (
	BookModel struct {
		Limit   int
		Offset  int
		SortBy  string
		SortDir string
	}
	BookData struct {
		Id               string `json:"books_isbn"`
		BooksTitle       string `json:"books_title"`
		BooksSubtitle    string `json:"books_subtitle"`
		BooksAuthor      string `json:"books_author"`
		BooksDescription string `json:"books_description"`
		BooksPublished   string `json:"books_published"`
		BooksPublisher   string `json:"books_publisher"`
	}
)

func (m *BookModel) CreateObject(data *BookData) (interface{}, error) {
	var dataMap map[string]interface{}
	j, _ := json.Marshal(data)
	err := json.Unmarshal(j, &dataMap)
	if err != nil {
		return nil, err
	}

	pgHandler := new(BaseModelPG)
	_, err = pgHandler.CreateObject("books", dataMap)

	if err != nil {
		return nil, err
	}
	res, err := pgHandler.GetObjectByField("books", "books_isbn", data.Id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *BookModel) UpdateObject(primaryKey string, dataMap map[string]interface{}) (interface{}, error) {
	// if primaryKey == "" {
	// 	return nil, errors.New("Missing primaryKey")
	// }
	// var output map[string]interface{}

	pgHandler := new(BaseModelPG)

	_, err := pgHandler.UpdateDataObject("books", "books_isbn", primaryKey, dataMap)
	if err != nil {
		return nil, err
	}

	output, err := pgHandler.GetObjectByField("books", "books_isbn", primaryKey)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (m *BookModel) DeleteObject(id string) (interface{}, error) {
	// if id == "" {
	// 	return errors.New("Missing ID object")
	// }

	dbHandler := new(BaseModelPG)

	res, err := dbHandler.DeleteObject("books", "books_isbn", id)
	if err != nil {
		return nil, err
	}

	return res, err
}
func (db *BookModel) GetObject(id string) (interface{}, error) {
	// var book map[string]interface{}
	var err error
	bookId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid book isbn")
	}
	dbHandler := new(BaseModelPG)
	res, err := dbHandler.GetObjectByField("books", "books_isbn", bookId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	// if book != nil {
	// delete(user, "password")
	// }

	return res, nil
}
func (db *BookModel) GetObjectByParams(params ...map[string]interface{}) (interface{}, error) {
	var paramsMap map[string]interface{}
	if len(params) > 0 {
		paramsMap = params[0]
	}

	var query string
	bindVars := make(map[string]interface{})
	filterData := make([]string, 0)
	joins := make([]string, 0)

	query += `SELECT *`

	// if paramsMap["isJoin"] != nil && paramsMap["isJoin"].(bool) {
	// 	query += `, wpb.va_number `
	// 	joins = append(joins, ` JOIN "wallet_public_user" wpb ON wpb.public_user_id = pb.id `)
	// }

	// if paramsMap["user_id"] != nil && paramsMap["user_id"].(string) != "" {
	// 	filterData = append(filterData, `pb.id = @userID`)
	// 	bindVars["userID"] = paramsMap["user_id"].(string)
	// }

	if paramsMap["isbn"] != nil && paramsMap["isbn"].(string) != "" {
		filterData = append(filterData, ` books_isbn = @bookIsbn `)
		bindVars["bookIsbn"] = paramsMap["isbn"].(string)
	}

	query += ` FROM "books" `

	var joinsString string
	if len(joins) > 0 {
		joinsString = strings.Join(joins, ` `)
	}

	var where string
	if len(filterData) > 0 {
		where = `WHERE ` + strings.Join(filterData, " AND ")
	}

	query += joinsString
	query += where
	query += ` LIMIT 1`

	dbHandler := new(BaseModelPG)

	var output interface{}

	output, err := dbHandler.GetObjectByQuery(query, bindVars)
	if err != nil {
		return nil, errors.New("invalid book isbn")
		// ZapLogger.Error(err.Error())
	}

	return output, nil
}
func (m *BookModel) GetAllCollection(params ...map[string]interface{}) ([]interface{}, error) {
	var paramsMap map[string]interface{}
	if params != nil && len(params) > 0 {
		paramsMap = params[0]
	}

	var query string
	bindVars := make(map[string]interface{})
	filterData := make([]string, 0)

	query += `SELECT COUNT(b.books_isbn) as count, json_agg(b.*) `

	if paramsMap["isbn"] != nil && paramsMap["isbn"].(string) != "" {
		filterData = append(filterData, ` b.books_isbn = @bookIsbn `)
		bindVars["bookIsbn"] = paramsMap["isbn"].(string)
	}
	query += ` FROM "books" b `

	var where string
	if len(filterData) > 0 {
		where = ` WHERE ` + strings.Join(filterData, " AND ")
	}

	if where != "" {
		query += where
	}

	if m.SortBy != "" {
		query += ` ORDER BY ` + m.SortBy
		if m.SortDir != "" {
			query += " " + m.SortDir + " "
		}
	}

	/*
		virtually unlimited, this is set to secure server
	*/
	if m.Limit == 0 || m.Limit > 1000 {
		m.Limit = 1000
	}
	if m.Limit > 0 {
		query += ` LIMIT ` + strconv.Itoa(m.Limit)
	}

	dbHandler := new(BaseModelPG)
	// dataRes := UserMerchantData{}

	results, _, err := dbHandler.GetCollectionWithCountByQuery(query, bindVars)
	if err != nil {
		// ZapLogger.Error(err.Error())
	}

	return results, nil
}
func (m *BookModel) GetCountByIsbn(bookIsbn string) (int64, error) {
	if bookIsbn == "" {
		return 0, errors.New("Missing Books ISBN")
	}

	filterData := make([]string, 0)
	bindVars := make(map[string]interface{})

	query := `SELECT COUNT(ipd.*) as counting`

	query += ` FROM "books" ipd `

	filterData = append(filterData, ` ipd.books_isbn = @bookIsbn`)
	bindVars["bookIsbn"] = bookIsbn

	var where string
	if len(filterData) > 0 {
		where = `WHERE ` + strings.Join(filterData, " AND ")
	}

	query += where

	dbHandler := new(BaseModelPG)

	count, err := dbHandler.GetCountByQuery(query, bindVars)
	if err != nil {
		return 0, err
	}

	return count, nil
}
