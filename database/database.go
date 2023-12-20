package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// connect to the application's database
func ConnectToDatabase() *sql.DB {
    db, err := connect(dbFileName)

    if err != nil {
        panic(err)
    }
    return db
}

// get the file path for the sqlite database
func getDataBaseFilePath(databaseFileName string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	dbPath := filepath.Join(filepath.Dir(exePath), databaseFileName)

    slog.Debug("Database Path: " + dbPath)
	return dbPath, nil
}

// connect to the sqlite database
func connect(databaseFileName string) (*sql.DB, error) {
    dbPath, err := getDataBaseFilePath(databaseFileName)
    if err != nil {
        return nil, err
    }

    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        slog.Error("Failed to open database")
        return nil, err
    }

    // Ping the database to ensure connectivity
    err = db.Ping()
    if err != nil {
        slog.Error("Failed to connect to database - ping failed")
        return nil, err
    }
    slog.Debug("Connected to SQLite Database")
    return db, nil
}

// run a sql string
func RunSQL(sql string, db *sql.DB) {
    // Execute the SQL statements in the file
    slog.Debug("Running sql: \n" + sql)
    _, err := db.Exec(sql)
    if err != nil {
        slog.Error(fmt.Sprint(err))
    }
    slog.Debug("Sql completed")
}