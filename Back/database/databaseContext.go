package database

import (
	"Backend/database/inits"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB(source string) {
	DB, err = sql.Open("sqlite3", source)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}
	inits.DB = DB
	// Execute inits
	inits.CreateUserTable()
	inits.CreateTaskTable()
	inits.CreateMailTable()
	inits.CreateMailServiceTable()
	inits.CreateMailRecipentsTable()
}
