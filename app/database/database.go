package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "todo_app"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
