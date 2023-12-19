package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func getDataBaseFilePath(databaseFileName string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	dbPath := filepath.Join(filepath.Dir(exePath), databaseFileName)

    slog.Info("Database found at: " + dbPath)
	return dbPath, nil
}

func Connect(databaseFileName string) (*sql.DB, error) {
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
    slog.Info("Connected to SQLite Database")
    return db, nil
}

func RunSQL(sql string, db *sql.DB) {
    // Execute the SQL statements in the file
    slog.Debug("Running sql: " + sql)
    _, err := db.Exec(sql)
    if err != nil {
        slog.Error(fmt.Sprint(err))
    }
    slog.Debug("Sql completed")
}