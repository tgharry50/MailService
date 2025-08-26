package inits

import (
	"Backend/functions"
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB
var err error

func CreateUserTable() {
	fmt.Println("Creating User table")
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

func CreateMailTable() {
	fmt.Println("Creating Mail table")
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS mails (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid TEXT NOT NULL,
	topic TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create mails table: ", err)
	}
}

func CreateMailRecipentsTable() {
	fmt.Println("Creating Mail Recipents table")
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS mail_recipents (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid TEXT NOT NULL,
	task_uuid TEXT NOT NULL,
	recipent TEXT NOT NULL,
	active BOOLEAN NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create mail_recipents table: ", err)
	}
}

func CreateMailServiceTable() {
	fmt.Println("Creating Mail Service table")
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS mail_services (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid TEXT NOT NULL,
	name TEXT NOT NULL,
	address TEXT NOT NULL,
	port TEXT NOT NULL,
	login TEXT NOT NULL,
	password TEXT NOT NULL,
	active BOOLEAN NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create mail_services table: ", err)
	}
}

func CreateTaskTable() {
	fmt.Println("Creating Task table")
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid TEXT NOT NULL,
	mail_uuid TEXT NOT NULL,
	mail_service_uuid TEXT NOT NULL,
	week_days TEXT,
	every_day INTEGER,
	every_week INTEGER,
	every_year INTEGER,
	specific_day DATETIME,
	repeat BOOLEAN NOT NULL,
	active BOOLEAN NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create tasks table: ", err)
	}
}

////

func InsertTestUserIfNotExists() {
	fmt.Println("Checking if test user is in table")
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
