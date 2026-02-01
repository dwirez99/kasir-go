package databases

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(ConnectionString string) (*sql.DB, error) {
	// Open Database
	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	// Test Connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
		return nil, err
	}

	// Set connection pool settings 
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connection established")
	return db, nil
}