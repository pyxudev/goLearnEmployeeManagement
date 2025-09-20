package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	dsn := "username:pass@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// テーブルのマイグレーション
	db.AutoMigrate(&User{})

	// GORMを使ったデータベース操作をここに記述
	user := User{Name: "new_name", Email: "abc@gmail.com"}
	result := db.Create(&user)
	if result != nil {
		panic("failed to insert database")
	} else {
		var user User
		db.First(&user)
		db.Where("name = ?", "new_name").Find(&user)
		var users []User
		db.Find(&users)

		fmt.Println("Done!")
	}
}
