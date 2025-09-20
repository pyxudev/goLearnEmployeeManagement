package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
	// Load .env File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Reade Environment Variables
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

	// Table Migration
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate: " + err.Error())
	}

	newUser := []User{}

	// --- CLI Enter Start ---
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("Enter UserName: ")
		userName, _ := reader.ReadString('\n')
		userName = strings.TrimSpace(userName)

		if userName == "q" {
			fmt.Println("User input ended.")
			break
		}

		fmt.Print("Enter UserEmail: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		newUser = append(newUser, User{Name: userName, Email: email})
	}
	// --- CLI Enter End ---

	// CUDA Operations
	result := db.Create(&newUser) // Insert Records

	if result.Error != nil {
		panic("Failed to insert database. Error Message: " + result.Error.Error())
	} else {
		var u User
		if err := db.First(&u).Error; err != nil {
			panic(err)
		}

		// --- Activate if you need to see the first record ---
		// fmt.Printf("First user: %+v\n", u)

		// --- Activate if you need to see all records ---
		// var users []User
		// db.Find(&users)
		// fmt.Printf("All users: %+v\n", users)

		fmt.Println("User insert ended!")
	}
}
