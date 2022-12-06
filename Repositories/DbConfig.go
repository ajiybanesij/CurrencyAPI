package Repositories

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	//dsn := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"

	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_Name")
	dsn := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + name + "?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Wallet{})
	db.AutoMigrate(&Balance{})
	db.AutoMigrate(&Transaction{})

	DB = db
	return DB
}
func GetDBInstance() *gorm.DB {
	return DB
}
