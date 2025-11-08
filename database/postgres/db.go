package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	_ "github.com/lib/pq"
)

type DB struct {
	drive *sql.DB
	table string
}

func (db *DB) SetTable(table string) {
	db.table = table
}

func (db *DB) Connect() {
	// 打开数据库连接
	var err error
	var connStr = "user=postgres password=123456 dbname=test host=localhost port=5432 sslmode=disable"
	db.drive, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// 测试连接
	err = db.drive.Ping()
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	fmt.Println("成功连接到PostgreSQL数据库")
}

func (db *DB) Insert() {
	sql := "insert into employees (name,email,hire_date) values ($1,$2,$3)"
	stmt, err := db.drive.Prepare(sql)
	if err != nil {
		panic(err)
	}
	n := rand.Intn(100)
	res, err := stmt.Exec("king", fmt.Sprintf("%d@a.com", n), "2025-10-21")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
}

func (db *DB) InsertGetId(field map[string]interface{}) int {
	var id int
	sqlStr, vals := db.BuildInsertSql(field)
	sqlStr = fmt.Sprintf("%s RETURNING id", sqlStr)
	stmt, err := db.drive.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	err = stmt.QueryRow(vals...).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func (db *DB) BuildInsertSql(fields map[string]interface{}) (string, []interface{}) {
	if fields == nil {
		panic("insert data is nil")
	}
	var field []string
	var fieldVal []interface{}
	for k, v := range fields {
		field = append(field, k)
		fieldVal = append(fieldVal, v)
	}
	var nFields []string
	for k, _ := range fieldVal {
		nFields = append(nFields, fmt.Sprintf("$%d", k+1))
	}
	sqlStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", db.table, strings.Join(field, ","), strings.Join(nFields, ","))
	return sqlStr, fieldVal
}

func (db *DB) Select(sqlStr string) []interface{} {
	rows, err := db.drive.Query(sqlStr)
	if err != nil {
		panic(err)
	}
	return buildQuery(rows)
}

func buildQuery(rows *sql.Rows) []interface{} {
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	var arr []interface{}
	for rows.Next() {
		val := make([]interface{}, len(cols))  // 用来保存实际值的位置
		valP := make([]interface{}, len(cols)) //用来接收scan返回的地址数据
		for i := range val {
			valP[i] = &val[i]
		}
		err = rows.Scan(valP...)
		if err != nil {
			panic(err)
		}
		m := make(map[string]interface{})
		for i, col := range cols {
			m[col] = val[i]
		}
		arr = append(arr, m)
	}
	return arr
}

func buildWhere(where map[string]interface{}) {

}

//func main() {
//	connect()
//	sql := "select * from employees"
//	Query(sql)
//	//insertRowGetId()
//}
