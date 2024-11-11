package models

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego/utils"
	"github.com/beego/beego/v2/server/web"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/grpc"
)

type (
	BaseModelPG struct {
		Distinct bool
		Limit    int
		Offset   int
		SortBy   string
		SortDir  string
		Search   string
		Filter   string
		BindVars map[string]interface{}
	}
)

var (
	// ZapLogger *zap.Logger
	conn      *grpc.ClientConn
	dbConnect *pgxpool.Pool
)

func Connect() (*pgxpool.Pool, error) {
	// // ZapLogger = logger.ZapLogger
	configPGUser, _ := web.AppConfig.String("pgUser")
	configPGPass, _ := web.AppConfig.String("pgPass")
	configPGHost, _ := web.AppConfig.String("pgHost")
	configPGPort, _ := web.AppConfig.String("pgPort")
	configPGDbname, _ := web.AppConfig.String("pgDbname")
	url := "postgres://" + configPGUser + ":" + configPGPass + "@" + configPGHost + ":" + configPGPort + "/" + configPGDbname
	return pgxpool.Connect(context.Background(), url)
}

// https://github.com/jackc/pgx/issues/387#issuecomment-798348824
func PrepareBindVars(sql string, namedArgs ...map[string]interface{}) (string, []interface{}) {
	var args []interface{}

	// Loop the named args and replace with placeholders
	if len(namedArgs) > 0 {
		var i int = 1
		for pname, pval := range namedArgs[0] {
			m := regexp.MustCompile(`@\b` + pname + `\b`)
			// don't replace @ with $ in here, to avoid clash with regex's capture symbol
			// replace @ with other signature, to avoid clash with
			// postgres' array operators @> (contains) and <@ (is contained by)
			r := `@@@` + strconv.Itoa(i)
			sql = m.ReplaceAllString(sql, r)
			// replace @@@ with $ in here
			sql = strings.ReplaceAll(sql, "@@@", "$")
			args = append(args, pval)
			i++
		}
	}

	return sql, args
}

func GetColumnNames(tableName string) ([]string, error) {
	conn, err := Connect()
	if err != nil {
		// ZapLogger.Error("Unable to connect to database: " + err.Error())
		return nil, err
	}

	pgQuery := `SELECT column_name FROM information_schema.columns WHERE table_name = $1`
	args := make([]interface{}, 0)
	args = append(args, tableName)
	rows, err := conn.Query(context.Background(), pgQuery, args...)
	if err != nil {
		return nil, err
	}

	columns := make([]string, 0)
	for rows.Next() {
		var columnName string
		err := rows.Scan(&columnName)
		if err != nil {
			return nil, err
		}
		columns = append(columns, columnName)
	}

	return columns, nil
}

func (m *BaseModelPG) CreateObject(tableName string, dataMap map[string]interface{}) (pgconn.CommandTag, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	columns, err := GetColumnNames(tableName)
	if err != nil {
		return nil, err
	}
	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)

	pgQuery := `INSERT INTO "` + tableName + `" (`
	keys := make([]string, 0)
	insertArr := make([]string, 0)
	if dataMap["_id"] != nil && dataMap["_id"].(string) != "" {
		dataMap["elastic_id"] = dataMap["_id"]
	}

	values := make([]interface{}, 0)
	var i int = 1
	for pK, pV := range dataMap {
		if pK == "_id" || pK == "_key" || pK == "id" {
			continue
		}
		if !utils.InSlice(pK, columns) {
			continue
		}
		keys = append(keys, `"`+pK+`"`)
		insertArr = append(insertArr, `$`+strconv.Itoa(i))
		values = append(values, pV)
		i++
	}
	pgQuery += strings.Join(keys, `,`)
	pgQuery += `) values (`
	pgQuery += strings.Join(insertArr, `,`)
	pgQuery += `)`

	return conn.Exec(context.Background(), pgQuery, values...)
}

func (m *BaseModelPG) UpdateObject(tableName string, dataMap map[string]interface{}, primaryKey string) (pgconn.CommandTag, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	columns, err := GetColumnNames(tableName)
	if err != nil {
		return nil, err
	}
	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)

	pgQuery := `UPDATE "` + tableName + `" SET`
	updateArr := make([]string, 0)
	var i int = 1
	args := make([]interface{}, 0)
	for pK, pV := range dataMap {
		if pK == "_id" || pK == "_key" || pK == "_rev" || pK == "_from" || pK == "_to" || pK == "id" {
			continue
		}
		if !utils.InSlice(pK, columns) {
			continue
		}
		updateArr = append(updateArr, `"`+pK+`"`+` = $`+strconv.Itoa(i))
		args = append(args, pV)
		i++
	}
	pgQuery += strings.Join(updateArr, `,`)
	pgQuery += ` WHERE elastic_id = $` + strconv.Itoa(i)
	args = append(args, primaryKey)

	return conn.Exec(context.Background(), pgQuery, args...)

}

func (m *BaseModelPG) UpdateDataObject(tableName string, primaryKey int64, dataMap map[string]interface{}) (pgconn.CommandTag, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	columns, err := GetColumnNames(tableName)
	if err != nil {
		return nil, err
	}
	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)

	pgQuery := `UPDATE "` + tableName + `" SET`
	updateArr := make([]string, 0)
	var i int = 1
	args := make([]interface{}, 0)
	for pK, pV := range dataMap {
		if pK == "_id" || pK == "_key" || pK == "id" || pK == "key" {
			continue
		}
		if !utils.InSlice(pK, columns) {
			continue
		}
		updateArr = append(updateArr, `"`+pK+`"`+` = $`+strconv.Itoa(i))
		args = append(args, pV)
		i++
	}
	pgQuery += strings.Join(updateArr, `,`)
	pgQuery += ` WHERE id = $` + strconv.Itoa(i)
	args = append(args, primaryKey)

	return conn.Exec(context.Background(), pgQuery, args...)
}

func (m *BaseModelPG) CreateObjectIfExist(tableName string, dataMap map[string]interface{}) (pgconn.CommandTag, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	columns, err := GetColumnNames(tableName)
	if err != nil {
		return nil, err
	}
	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)

	pgQuery := `INSERT INTO "` + tableName + `" (`
	keys := make([]string, 0)
	insertArr := make([]string, 0)
	dataMap["elastic_id"] = dataMap["_id"]

	values := make([]interface{}, 0)
	var i int = 1
	for pK, pV := range dataMap {
		if pK == "_id" || pK == "_key" || pK == "_rev" || pK == "_from" || pK == "_to" || pK == "id" {
			continue
		}
		if !utils.InSlice(pK, columns) {
			continue
		}
		keys = append(keys, `"`+pK+`"`)
		insertArr = append(insertArr, `$`+strconv.Itoa(i))
		values = append(values, pV)
		i++
	}
	pgQuery += strings.Join(keys, `,`)
	pgQuery += `) values (`
	pgQuery += strings.Join(insertArr, `,`)
	pgQuery += `)`
	pgQuery += ` ON CONFLICT DO NOTHING`

	return conn.Exec(context.Background(), pgQuery, values...)

}

func (m *BaseModelPG) GetObject(tableName string, primaryKey string) (interface{}, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)

	pgQuery := `SELECT * FROM ` + tableName + ` WHERE elastic_id = $1`
	args := make([]interface{}, 0)
	args = append(args, primaryKey)
	rows, err := conn.Query(context.Background(), pgQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	instance := make(map[string]interface{})

	if err := pgxscan.ScanOne(&instance, rows); err != nil {
		// ZapLogger.Error("pgxscan.ScanOne: " + err.Error())
		return nil, err
	}

	return instance, nil
}

func (m *BaseModelPG) GetObjectByField(tableName string, columnName string, columnValue interface{}) (interface{}, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)
	columnName = p.Sanitize(columnName)

	pgQuery := `SELECT * FROM public.` + tableName + ` WHERE ` + columnName + ` = $1`
	args := make([]interface{}, 0)
	args = append(args, columnValue)
	rows, err := conn.Query(context.Background(), pgQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	instance := make(map[string]interface{})

	if err := pgxscan.ScanOne(&instance, rows); err != nil {
		if err.Error() == "no rows in result set" {
			// ZapLogger.Error("pgxscan.ScanOne: " + err.Error())

			return nil, err
		}
		// ZapLogger.Error("pgxscan.ScanOne: ")
		return nil, err
	}

	return instance, nil
}

func (m *BaseModelPG) GetObjectV2(output *map[string]interface{}, tableName string, columnName string, columnValue interface{}) error {

	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)
	columnName = p.Sanitize(columnName)

	pgQuery := `SELECT * FROM public.` + tableName + ` WHERE ` + columnName + ` = $1`
	args := make([]interface{}, 0)
	args = append(args, columnValue)
	rows, err := conn.Query(context.Background(), pgQuery, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	var instance map[string]interface{}
	if err := pgxscan.ScanOne(&instance, rows); err != nil {
		if err.Error() == "no rows in result set" {
			// ZapLogger.Error("pgxscan.ScanOne: " + err.Error())

			return err
		}
		// ZapLogger.Error("pgxscan.ScanOne: ")
		return err
	}
	*output = instance
	return nil
}

func (m *BaseModelPG) DeleteObject(tableName string, columnName string, primaryKey int64) (pgconn.CommandTag, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)
	columnName = p.Sanitize(columnName)
	var i int = 1
	args := make([]interface{}, 0)

	pgQuery := `DELETE FROM "` + tableName + `" WHERE "` + columnName + `" = $` + strconv.Itoa(i)
	args = append(args, primaryKey)

	return conn.Exec(context.Background(), pgQuery, args...)
}

func (m *BaseModelPG) DeleteObjectByField(tableName string, columnName string, primaryKey string) (pgconn.CommandTag, error) {

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	/*
		Process the data
	*/
	p := bluemonday.UGCPolicy()
	tableName = p.Sanitize(tableName)
	columnName = p.Sanitize(columnName)
	var i int = 1
	args := make([]interface{}, 0)

	pgQuery := `DELETE FROM "` + tableName + `" WHERE "` + columnName + `" = $` + strconv.Itoa(i)
	args = append(args, primaryKey)

	return conn.Exec(context.Background(), pgQuery, args...)
}

// func (m *BaseModelPG) GetCountByQuery(query string, args []interface{}) int64 {
// 	conn, err := Connect()
// 	if err != nil {
// 		return 0
// 	}
// 	rows, err := conn.Exec(context.Background(), query, args...)

// 	if err != nil {
// 		return 0
// 	}

// 	rowCount := rows.RowsAffected()

// 	return rowCount
// }

func (m *BaseModelPG) GetCollectionWithCountByQuery(query string, bindData ...map[string]interface{}) ([]interface{}, int64, error) {
	var bindVars map[string]interface{}
	if bindData != nil && len(bindData) > 0 {
		bindVars = bindData[0]
	}

	conn, err := Connect()
	if err != nil {
		return nil, 0, errors.New(err.Error())
	}

	preparedSql, args := PrepareBindVars(query, bindVars)
	ctx := context.TODO()
	rows, err := conn.Query(ctx, preparedSql, args...)
	if err != nil {
		// ZapLogger.Error(err.Error())
		return nil, 0, errors.New("SQL error: " + err.Error())
	}
	defer rows.Close()

	var count int64
	var results []interface{}
	for rows.Next() {
		rows.Scan(&count, &results)
	}

	/*
	 Validation if data is empty becomes an array interface (Required, to avoid error panic)
	*/
	if results == nil {
		return make([]interface{}, 0), 0, nil
	}
	// END

	return results, count, nil
}

func (m *BaseModelPG) GetObjectByQuery(query string, bindData ...map[string]interface{}) (interface{}, error) {
	var bindVars map[string]interface{}
	if bindData != nil && len(bindData) > 0 {
		bindVars = bindData[0]
	}

	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	pgQuery, args := PrepareBindVars(query, bindVars)

	ctx := context.Background()
	rows, err := conn.Query(ctx, pgQuery, args...)
	if err != nil {
		return err.Error(), nil
		// ZapLogger.Error(err.Error())
	}

	defer rows.Close()

	output := make(map[string]interface{})

	if err := pgxscan.ScanOne(&output, rows); err != nil {
		if err.Error() == "no rows in result set" {
			// ZapLogger.Error("pgxscan.ScanOne: " + err.Error())

			return nil, err
		}
		// ZapLogger.Error("pgxscan.ScanOne: ")
		return nil, err
	}

	return output, err
}

func (m *BaseModelPG) GetCollectionByQuery(query string, structPointerReference interface{}, bindData ...map[string]interface{}) ([]interface{}, error) {
	var bindVars map[string]interface{}
	if bindData != nil && len(bindData) > 0 {
		bindVars = bindData[0]
	}

	conn, err := Connect()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer conn.Close()

	preparedSql, args := PrepareBindVars(query, bindVars)
	ctx := context.TODO()
	rows, err := conn.Query(ctx, preparedSql, args...)
	if err != nil {
		// ZapLogger.Error(err.Error())
		return nil, errors.New("SQL error: " + err.Error())
	}
	defer rows.Close()
	var results []interface{}

	if rows.Next() {
		err := rows.Scan(&results)
		if err != nil {
			return results, err
		}
	}

	// results, _ := m.ScanAll(rows, structPointerReference)

	return results, nil
}

func (m *BaseModelPG) GetCountByQuery(query string, params map[string]interface{}) (int64, error) {
	if query == "" {
		return 0, errors.New("Missing query")
	}

	conn, err := Connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	pgQuery, args := PrepareBindVars(query, params)

	ctx := context.Background()

	var counter int64
	err = conn.QueryRow(ctx, pgQuery, args...).Scan(&counter)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "no rows in result set") {
			// Don't show message there, because data is empty
		} else {
			// ZapLogger.Error(err.Error())
		}
	}

	return counter, nil
}
