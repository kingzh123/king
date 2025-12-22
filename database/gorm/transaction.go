package gorm

import (
	"context"
	"king/database/model"
	"math/rand"

	"gorm.io/gorm"
)

func TransDemo() {
	ctx := context.Background()
	db, err := ConnectPostgresDB()
	if err != nil {
		panic(err)
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		u := model.User{Name: "zhang", Age: 30, Email: "11@163.com"}
		if err := u.Validate(); err != nil {
			return err
		}
		err = gorm.G[model.User](tx).Create(ctx, &u)
		if err != nil {
			return err
		}
		product := "iphone"
		customerId := int64(rand.Intn(100)) + 1
		amount := 12.00
		err = gorm.G[model.Order](tx).Create(ctx, &model.Order{CustomerID: customerId, Product: &product, Status: "pending", Amount: amount, IsOver: true})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
