package main

import (
	kg "king/gorm"
)

func main() {
	// ---生成雪花id---
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
	//---dao生成---
	//arr := dao.NewArr(db)
	//table := arr.TableName()
	//fmt.Println(table)
	//---数据库 orm generate struct  2025-12-13 13:41:23  gorm 1.31 对 gorm-gen 0.3.7 有版本问题---
	//#1 go install gorm.io/gen/tools/gentool@latest
	//#2 gentool --db=postgres --dsn="host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable" --outPath="./dao"
	//model.Create()
	//---gorm 数据查询---
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
	// ---gen postgres 表 自动生成 结构体---
	//model2.ToStrcut()
	//db, err := kg.ConnectPostgresDB()
	//if err != nil {
	//	panic(err)
	//}
	//---generics api 泛型查询数据---
	//ctx := context.Background()
	//var user []model.User
	//user, err = gorm.G[model.User](db).Limit(1).Find(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(user)
	//---事务执行---
	kg.TransDemo()
}
