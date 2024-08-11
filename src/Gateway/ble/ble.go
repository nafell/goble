package ble

import (
	"fmt"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter
var co2CurrentValue uint16

func Start() error {
	// Enable BLE interface.
	err := Enable()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	// Start scanning.
	esp32_scan, err := Scan()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	// Connect to the device
	esp32, err := Connect(esp32_scan)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	errNotify := StartCO2Notify(esp32)
	if errNotify != nil {
		fmt.Println("Error: ", errNotify)
		return err
	}

	return nil
}

func Enable() error {
	// Enable BLE interface.
	err := adapter.Enable()

	if err != nil {
		return err
	}

	return nil
}

func Scan() (bluetooth.ScanResult, error) {
	var esp32_scan bluetooth.ScanResult
	// Start scanning.
	println("Scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("Found device:", device.Address.String(), device.RSSI, device.LocalName())
		if device.LocalName() == "Nano ESP32" { // Preferreably use service UUID instaed to identify the device
			adapter.StopScan()
			esp32_scan = device
		}
	})
	if err != nil {
		return esp32_scan, err
	}

	return esp32_scan, nil
}

func Connect(device bluetooth.ScanResult) (bluetooth.Device, error) {
	println("Connecting to", device.Address.String(), device.LocalName())
	esp32, err := adapter.Connect(device.Address, bluetooth.ConnectionParams{}) // Adjust parameters for battery optimisation
	if err != nil {
		return esp32, err
	}
	println("Connection successful.")
	return esp32, nil
}

func StartCO2Notify(esp32 bluetooth.Device) error {
	service_uuid, _ := bluetooth.ParseUUID("E3200001-c577-4615-bc4a-44feb3e806fd")
	// chara_write_uuid, _ := bluetooth.ParseUUID("E3200002-c577-4615-bc4a-44feb3e806fd")
	chara_notify_uuid, _ := bluetooth.ParseUUID("E3200003-c577-4615-bc4a-44feb3e806fd")
	services, err := esp32.DiscoverServices([]bluetooth.UUID{service_uuid})
	if err != nil {
		return err
	}

	if len(services) != 1 {
		return fmt.Errorf("unexpected number of services found")
	}

	service := services[0]

	charas, err := service.DiscoverCharacteristics([]bluetooth.UUID{chara_notify_uuid})
	if err != nil {
		return err
	}

	if len(charas) != 1 {
		return fmt.Errorf("unexpected number of characteristics found")
	}

	chara_notify := charas[0]

	chara_notify.EnableNotifications(notificationCO2Callback)

	return nil
}

func notificationCO2Callback(buf []byte) {
	if len(buf) != 2 {
		return
	}

	co2CurrentValue = uint16(buf[0])<<8 | uint16(buf[1])
	fmt.Printf("notification: %08b %d\n", buf, co2CurrentValue)
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
