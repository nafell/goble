package main

import (
	"fmt"
	"os"
	"log/slog"

	"tinygo.org/x/bluetooth"

	"github.com/nafell/goble_gateway/ble"
	"github.com/nafell/goble_gateway/db"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	fmt.Println("nanjakoryaaaaa")
	for _, v := range os.Args {

		if v == "--debug" {
			slog.SetLogLoggerLevel(slog.LevelDebug)
		}
		if v == "--gen" {
			slog.SetLogLoggerLevel(slog.LevelDebug)
			fmt.Println("gen option")
			gen()
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

func gen() {
	slog.Info("Generating db interface...")
	must("generate db interface", db.init())
	slog,Info("Generated db interface.")
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
