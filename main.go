package main

import (
	//"github.com/jcocozza/go_speedtest/database"
	"fmt"

	"github.com/jcocozza/go_speedtest/speedtest"
)

const dbFileName = "go-speedtest.db"

func main() {
	/*
	db, err := database.Connect(dbFileName)
	if err != nil {
		panic(err)
	}
	*/

	result := speedtest.SpeedTestJSON()
	fmt.Println(result)
}