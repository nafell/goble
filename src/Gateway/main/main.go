package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/nafell/goble_gateway/ble"
	"github.com/nafell/goble_gateway/dbstuff"
	"github.com/nafell/goble_gateway/dbstuff/model"
)

func main() {
	args()

	// Start BLE module
	err := ble.Start()
	if err != nil {
		slog.Error("Failed to start BLE module")
		slog.Error("Error: ", err)
	}

	// Periodically log sensor data to sqlite
	ticker := time.NewTicker(10 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return

			case t := <-ticker.C:
				slog.Info("Logged sensor data at:", t.GoString())
			}
		}
	}()

	//read text from the console
	for {
		var text string
		fmt.Scanln(&text)
	}
}

func args() {
	for _, v := range os.Args {
		if v == "--debug" {
			slog.SetLogLoggerLevel(slog.LevelDebug)
		}
		if v == "--gen" {
			slog.SetLogLoggerLevel(slog.LevelDebug)
			slog.Debug("Generating database queries...")
			must("Gen generate", dbstuff.Generate())
			slog.Debug("Generated database queries.")
			os.Exit(0)
		}
		if v == "--create" {
			slog.SetLogLoggerLevel(slog.LevelDebug)
			slog.Debug("Creating database connection...")
			must("Connect", dbstuff.Connect())
			slog.Debug("Created database connection.")
			newEntry := &model.SensorLogEntry{Datetime: time.Now(), Co2: model.Int32Ptr(900), Temperature: model.Int32Ptr(25), Humidity: model.Int32Ptr(60), Pressure: model.Int32Ptr(1000)}
			must("AddSensorLogEntry", dbstuff.AddSensorLogEntry(newEntry))
			slog.Debug("Added sensor log entry.")
			os.Exit(0)
		}
		if v == "--read" {
			slog.SetLogLoggerLevel(slog.LevelDebug)
			slog.Debug("Creating database connection...")
			must("Connect", dbstuff.Connect())
			slog.Debug("Created database connection.")
			entries, err := dbstuff.GetAllSensorLogEntries()
			if err != nil {
				slog.Error("Failed to get all sensor log entries")
				slog.Error("Error: ", err)
			} else {
				for _, entry := range entries {
					slog.Debug("Entry: ", entry)
				}
			}
			os.Exit(0)
		}
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
