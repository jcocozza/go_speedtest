package main

import (
	"os"

	"github.com/jcocozza/go_speedtest/speedtest"
)

func main() {
	data := speedtest.CreateData(50)
	os.WriteFile("data.txt", data, os.ModeAppend)
}
