package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から読み込む
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// テーブルのマイグレーション
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate: " + err.Error())
	}

	// GORMを使ったデータベース操作をここに記述
	newUser := User{Name: "new_name", Email: "User@gmail.com"}
	result := db.Create(&newUser) // レコードの挿入

	if result.Error != nil {
		panic("Failed to insert database. Error Message: " + result.Error.Error())
	} else {
		var u User
		if err := db.First(&u).Error; err != nil {
			panic(err)
		}
		fmt.Printf("First user: %+v\n", u)

		var users []User
		db.Find(&users)
		fmt.Printf("All users: %+v\n", users)

		fmt.Println("Done!")
	}
}
