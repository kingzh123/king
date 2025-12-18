package model

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func ToStrcut() {
	// 数据库配置
	dsn := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./query", // 相对于当前工作目录
		ModelPkgPath: "./model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 使用数据库连接
	g.UseDB(db)

	// 生成所有表（排除系统表）只生成 model，不会生成query
	g.GenerateAllTable()
	//g.ApplyBasic(g.GenerateAllTable()...) 同时生成model和query

	// 自定义字段类型映射
	//g.WithDataTypeMap(map[string]func(detailType gorm.ColumnType) (dataType string){
	//	// UUID 类型
	//	"uuid": func(detailType gorm.ColumnType) (dataType string) {
	//		return "uuid.UUID"
	//	},
	//
	//	// JSON 类型
	//	"json": func(detailType gorm.ColumnType) (dataType string) {
	//		return "datatypes.JSON"
	//	},
	//	"jsonb": func(detailType gorm.ColumnType) (dataType string) {
	//		return "datatypes.JSON"
	//	},
	//
	//	// 数组类型
	//	"text[]": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pq.StringArray"
	//	},
	//	"integer[]": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pq.Int64Array"
	//	},
	//	"bigint[]": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pq.Int64Array"
	//	},
	//	"boolean[]": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pq.BoolArray"
	//	},
	//	"numeric[]": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pq.Float64Array"
	//	},
	//
	//	// 精确数值类型
	//	"numeric": func(detailType gorm.ColumnType) (dataType string) {
	//		return "decimal.Decimal"
	//	},
	//	"decimal": func(detailType gorm.ColumnType) (dataType string) {
	//		return "decimal.Decimal"
	//	},
	//	// 网络类型
	//	"inet": func(detailType gorm.ColumnType) (dataType string) {
	//		return "net.IP"
	//	},
	//	"cidr": func(detailType gorm.ColumnType) (dataType string) {
	//		return "net.IPNet"
	//	},
	//
	//	// 几何类型
	//	"point": func(detailType gorm.ColumnType) (dataType string) {
	//		return "postgis.Point"
	//	},
	//	"geometry": func(detailType gorm.ColumnType) (dataType string) {
	//		return "postgis.Geometry"
	//	},
	//
	//	// 时间类型
	//	"timestamptz": func(detailType gorm.ColumnType) (dataType string) {
	//		return "time.Time"
	//	},
	//	"timestamp": func(detailType gorm.ColumnType) (dataType string) {
	//		return "time.Time"
	//	},
	//	"date": func(detailType gorm.ColumnType) (dataType string) {
	//		return "time.Time"
	//	},
	//	"time": func(detailType gorm.ColumnType) (dataType string) {
	//		return "time.Time"
	//	},
	//	"timetz": func(detailType gorm.ColumnType) (dataType string) {
	//		return "time.Time"
	//	},
	//
	//	// 区间类型
	//	"int4range": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pgtype.Int4range"
	//	},
	//	"int8range": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pgtype.Int8range"
	//	},
	//	"tsrange": func(detailType gorm.ColumnType) (dataType string) {
	//		return "pgtype.Tsrange"
	//	},
	//
	//	// 其他 PostgreSQL 特有类型
	//	"money": func(detailType gorm.ColumnType) (dataType string) {
	//		return "string"
	//	},
	//	"bytea": func(detailType gorm.ColumnType) (dataType string) {
	//		return "[]byte"
	//	},
	//	"bit": func(detailType gorm.ColumnType) (dataType string) {
	//		return "string"
	//	},
	//	"varbit": func(detailType gorm.ColumnType) (dataType string) {
	//		return "string"
	//	},
	//	"xml": func(detailType gorm.ColumnType) (dataType string) {
	//		return "string"
	//	},
	//	"tsvector": func(detailType gorm.ColumnType) (dataType string) {
	//		return "string"
	//	},
	//	"citext": func(detailType gorm.ColumnType) (dataType string) {
	//		return "string"
	//	},
	//})

	// 执行生成
	g.Execute()

	fmt.Println("模型生成成功")
}
