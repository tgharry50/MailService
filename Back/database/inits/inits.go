package inits

import (
	"Backend/functions"
	"database/sql"
	"log"
)

var DB *sql.DB
var err error

func CreateUserTable() {
	createTableSQL := `
         CREATE TABLE IF NOT EXISTS users (
         id INTEGER PRIMARY KEY AUTOINCREMENT,
         uuid TEXT NOT NULL,
         firstname TEXT NOT NULL,
         lastname TEXT NOT NULL,
         username TEXT NOT NULL,
         password TEXT NOT NULL,
         email TEXT NOT NULL,
         phone TEXT,
         active BOOLEAN NOT NULL,
         created_at DATETIME NOT NULL,
         updated_at DATETIME NOT NULL
        );`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create tables: ", err)
	}
	InsertTestUserIfNotExists()
}

func InsertTestUserIfNotExists() {
	var count int
	err = DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "testuser").Scan(&count)
	if err != nil {
		log.Fatal("Failed to check for test user: ", err)
	}
	if count == 0 {
		hashed, err := functions.HashPassword("testpass") // hashing
		uuid := functions.GenerateUUID()
		if err != nil {
			log.Fatal("Failed to hash password: ", err)
		}
		_, err = DB.Exec(`INSERT INTO users (uuid, firstname, lastname, username, password, email, phone, active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, datetime('now'), datetime('now'))`,
			uuid, "Test", "User", "testuser", hashed, "testuser@example.com", "123456789", true)
		if err != nil {
			log.Fatal("Failed to insert test user: ", err)
		}
		log.Println("Test user inserted.")
	} else {
		log.Println("Test user already exists.")
	}
}
