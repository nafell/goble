package main

import (
	"fmt"
	"os"
	"log/slog"

	"tinygo.org/x/bluetooth"

	"github.com/nafell/goble_gateway/ble"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	for _, v := range os.Args {
		if v == "--debug" {
			slog.SetLogLoggerLevel(slog.LevelDebug)
		}
	}

	// Start BLE module
	err := ble.Start()
	if err != nil {
		slog.Error("Failed to start BLE module")
		slog.Error("Error: ", err)
	}

	//read text from the console
	for {
		var text string
		fmt.Scanln(&text)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
