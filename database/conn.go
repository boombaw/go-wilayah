package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() *sqlx.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db *sqlx.DB
	urlConnection := "user=" + fmt.Sprint(os.Getenv("db_user")) + " "
	if os.Getenv("db_password") != "" {
		urlConnection += "password=" + fmt.Sprint(os.Getenv("db_password")) + " "
	}
	urlConnection += "host=" + fmt.Sprint(os.Getenv("db_host")) + " "
	urlConnection += "port=" + fmt.Sprint(os.Getenv("db_port")) + " "
	urlConnection += "dbname=" + fmt.Sprint(os.Getenv("db_name")) + " "
	urlConnection += "sslmode=" + fmt.Sprint(os.Getenv("db_sslmode"))

	log.Println("Connecting to DB Server " + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + "...With DB Name " + fmt.Sprint(os.Getenv("db_name")))

	db, err = sqlx.Open("postgres", urlConnection)
	if err != nil {
		log.Printf("Error while connecting to database : %v", err)
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
