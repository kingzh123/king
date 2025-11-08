package main

import (
	"king/gin"
)

func main() {
	/*----数据库操作----*/
	//d := postgres.DB{}
	//d.SetTable("employees")
	//d.Connect()
	//r := rand.Intn(1000)
	//m := map[string]interface{}{"name": "king", "email": strconv.Itoa(r) + "@1.com", "date": time.Now()}
	//id := d.InsertGetId(m)
	//fmt.Println(id)
	//sql := "select * from employees"
	//postgres.Query(sql)
	//insertRowGetId()
	//postgres.Insert()

	/*----web service----*/
	gin.Run()

}
