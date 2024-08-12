package dbstuff

import (
	"time"

	"github.com/nafell/goble_gateway/ble"
	"github.com/nafell/goble_gateway/dbstuff/model"
)

func Tick() error {
	if !ble.HasValidValue {
		return nil
	}
	newEntry := &model.SensorLogEntry{
		Datetime:    time.Now(),
		Co2:         model.Int32Ptr(int(ble.Co2CurrentValue)),
		Temperature: nil,
		Humidity:    nil,
		Pressure:    nil,
	}
	return AddSensorLogEntry(newEntry)
}
