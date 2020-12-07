package helpers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbURL := "us-cdbr-east-02.cleardb.com"
	dbName := "heroku_ac60454d09cc541"
	db, err := sql.Open("mysql", "b7da67a7641056:91838c9c@tcp("+dbURL+")/"+dbName)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
