package main

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	service_uuid, _ := bluetooth.ParseUUID("E3200001-c577-4615-bc4a-44feb3e806fd")
	// chara_write_uuid, _ := bluetooth.ParseUUID("E3200002-c577-4615-bc4a-44feb3e806fd")
	chara_notify_uuid, _ := bluetooth.ParseUUID("E3200003-c577-4615-bc4a-44feb3e806fd")
	var esp32_scan bluetooth.ScanResult
	// Start scanning.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName(), device.HasServiceUUID(service_uuid))
		if device.LocalName() == "Nano ESP32" { // Preferreably use service UUID instaed to identify the device
			adapter.StopScan()
			esp32_scan = device
			println("connecting to", device.Address.String(), device.LocalName())
		}
	})
	must("start scan", err)

	esp32, errConnect := adapter.Connect(esp32_scan.Address, bluetooth.ConnectionParams{}) // Adjust parameters for battery optimisation
	must("connect to device", errConnect)
	defer esp32.Disconnect()

	services, err := esp32.DiscoverServices([]bluetooth.UUID{service_uuid})
	must("discover services", err)

	if len(services) != 1 {
		panic("unexpected number of services found")
	}

	service := services[0]

	charas, err := service.DiscoverCharacteristics([]bluetooth.UUID{chara_notify_uuid})
	must("discover characteristics", err)

	if len(charas) != 1 {
		panic("unexpected number of characteristics found")
	}

	chara_notify := charas[0]

	chara_notify.EnableNotifications(func(buf []byte) {
		fmt.Printf("notification: %08b\n", buf)
	})

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
