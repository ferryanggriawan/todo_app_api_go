package database

import (
	"database/sql"
	"fmt"
	"todo_app_api_go/config"

	_ "github.com/lib/pq"
)

type Connection struct {
	DB *sql.DB
}

func (c *Connection) Initialize() {
	db, err := c.config()
	c.DB = db

	if err != nil {
		panic(err)
	}

	c.checkDb()
}

func (c *Connection) config() (*sql.DB, error) {
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

	return db, err
}

func (c *Connection) checkDb() {
	defer c.DB.Close()

	err := c.DB.Ping()
	if err != nil {
		panic(err)
	}
}
