package database

import (
	"fmt"
	"log/slog"
	"github.com/jcocozza/go_speedtest/sql"
)

const dbFileName = "go-speedtest.db"

func init() {
	slog.Info("Running init database")
	db, err := Connect(dbFileName)
	if err != nil {
		slog.Error(fmt.Sprint(err))
		panic(err)
	}


    RunSQL(sql.SqlInit, db)
	slog.Info("database init complete")
}