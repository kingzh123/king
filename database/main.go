package main

import (
	model2 "king/gorm/model"
)

func main() {
	//sf := sonyflake.NewSonyflake(sonyflake.Settings{
	//	MachineID: func() (uint16, error) {
	//		return uint16(1), nil
	//	},
	//})
	//if sf == nil {
	//	panic("sonyflake not created")
	//}
	//id, err := sf.NextID()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(id)
	//fmt.Println(db)
	//arr := dao.NewArr(db)
	//table := arr.TableName()
	//fmt.Println(table)
	//数据库 orm generate struct  2025-12-13 13:41:23  gorm 1.31 对 gorm-gen 0.3.7 有版本问题
	//#1 go install gorm.io/gen/tools/gentool@latest
	//#2 gentool --db=postgres --dsn="host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable" --outPath="./dao"
	//model.Create()
	//db, err := kg.ConnectPostgresDB()
	//if err != nil {
	//	panic(err)
	//}
	//user := model.User{}
	//ctx := context.Background()
	//user, err = gorm.G[model.User](db).First(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", user)
	model2.ToStrcut()

}
