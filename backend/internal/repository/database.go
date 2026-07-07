package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(storagePath string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при открытии базы данных %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Ошибка при пинге базы данных: %v", err)
	}

	// подключение внешних ключей и WAL режима
	pragmas := []string{
		"PRAGMA foreign_keys = ON;",
		"PRAGMA journal_mode = WAL;",
	}
	for _, pragma := range pragmas {
		if _, err := db.Exec(pragma); err != nil {
			return nil, err
		}
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil

}

func createTables(db *sql.DB) error {

	schema := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE IF NOT EXISTS subscriptions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,         
		expire_date DATETIME NOT NULL, 
		is_active BOOLEAN DEFAULT 1,  
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	
	CREATE INDEX IF NOT EXISTS idx_subscriptions_expire_date ON subscriptions(expire_date);`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("Ошибка при создании таблиц: %v", err)
	}

	return nil
}
