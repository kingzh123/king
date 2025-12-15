package model

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ToStrcut() {
	// 数据库配置
	dsn := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// 创建输出目录
	outputPath := "./internal/models"
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		log.Fatal("创建目录失败:", err)
	}

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:        outputPath,
		ModelPkgPath:   filepath.Join(outputPath, "model"),
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:  true,  // 可为空字段生成指针类型
		FieldCoverable: false, // 没有默认值的字段生成指针类型
		FieldSignable:  true,  // 检测整型字段的无符号类型
	})

	// 使用数据库连接
	g.UseDB(db)

	// 生成所有表（排除系统表）
	g.ApplyBasic(g.GenerateAllTable()...)

	// 自定义字段类型映射
	g.WithDataTypeMap(map[string]func(detailType gorm.ColumnType) (dataType string){
		"timestamp": func(detailType gorm.ColumnType) (dataType string) {
			return "time.Time"
		},
		"timestamptz": func(detailType gorm.ColumnType) (dataType string) {
			return "time.Time"
		},
		"uuid": func(detailType gorm.ColumnType) (dataType string) {
			return "uuid.UUID"
		},
		"jsonb": func(detailType gorm.ColumnType) (dataType string) {
			return "datatypes.JSON"
		},
		"numeric": func(detailType gorm.ColumnType) (dataType string) {
			return "decimal.Decimal"
		},
	})

	// 执行生成
	g.Execute()

	fmt.Println("模型生成成功！输出目录:", outputPath)
}
