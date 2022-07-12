package database

import (
	"database/sql"
	"fmt"
	"todo_app_api_go/config"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	db, err := getDB()

	if err != nil {
		panic(err)
	}

	checkDb(db)

	Db = db
}

func Connect() {
	if Db == nil {
		db, err := getDB()

		if err != nil {
			panic(err)
		}

		checkDb(db)

		Db = db
	}
}

func getDB() (*sql.DB, error) {
	config := config.GetConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
	)

	db, err := sql.Open("postgres", psqlInfo)
	Db = db

	return db, err
}

func checkDb(db *sql.DB) {
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
