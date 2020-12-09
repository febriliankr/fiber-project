package helpers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbURL+")/"+dbName)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
