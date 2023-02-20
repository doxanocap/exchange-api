package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "eldoseldos"
	dbname   = "bapi"
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	connection, err := sql.Open("postgres", psqlInfo)
	DB = connection

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
