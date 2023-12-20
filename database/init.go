package database

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jcocozza/go_speedtest/sql"
)

const dbFileName = ".go-speedtest.db"

func init() {
	dbPath, err := getDataBaseFilePath(dbFileName)
    if err != nil {
		slog.Error("Failed to get database file path")
        return
    }

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		slog.Debug("File " + dbFileName + " does not exist.")
		slog.Debug("Running init database")

		db, err := connect(dbFileName)
		if err != nil {
			slog.Error(fmt.Sprint(err))
			panic(err)
		}

		RunSQL(sql.SqlInit, db)
		slog.Debug("Database init complete")
		return
	}
}