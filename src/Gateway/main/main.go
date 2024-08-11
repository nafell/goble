package main

import (
	"fmt"

	"tinygo.org/x/bluetooth"

	"github.com/nafell/goble_gateway/ble"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Start BLE module
	err := ble.Start()
	if err != nil {
		fmt.Println("Failed to start BLE module")
		fmt.Println("Error: ", err)
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
