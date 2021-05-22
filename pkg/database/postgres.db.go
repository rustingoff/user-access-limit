package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var db *gorm.DB

func InitPostgresDB() (*gorm.DB, error) {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
		return nil, e
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable port=%s password=%s", dbHost, username, dbName, dbPort, password) //Build connection string
	fmt.Print(dbUri)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	db = conn
	db.Debug().AutoMigrate() //Database migration

	return db, nil
}
