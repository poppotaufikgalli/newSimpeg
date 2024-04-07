package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lpernett/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
)

var MyDB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_DATABASE = os.Getenv("DB_DATABASE")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	MyDB, _ = CreateCon0()
}

type ResultQuery struct {
	RowsAffected int
	Error        error
}

func CreateCon0() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_DATABASE)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db.Close()

	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", DB_USERNAME, DB_PASSWORD, DB_DATABASE))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return
}

func CreateCon() (db *gorm.DB, err error) {
	db = MyDB

	return
}
