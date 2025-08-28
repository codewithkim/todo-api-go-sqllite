package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDB() {
    dbPath := os.Getenv("DB_PATH")
    if dbPath == "" {
        dbPath = "./todos.db"
    }

    dsn := fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", dbPath)

    var err error
    DB, err = sql.Open("sqlite", dsn)
    if err != nil {
        log.Fatalf("❌ Failed to connect to SQLite: %v", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalf("❌ SQLite unreachable: %v", err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        completed BOOLEAN DEFAULT FALSE
    );`
    if _, err := DB.Exec(createTable); err != nil {
        log.Fatalf("❌ Failed to create todos table: %v", err)
    }

    log.Println("✅ SQLite connected and table ensured")
}