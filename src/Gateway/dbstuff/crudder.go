package dbstuff

import (
	"github.com/nafell/goble_gateway/dbstuff/model"
	"github.com/nafell/goble_gateway/dbstuff/query"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var initialized bool
var Q query.Query

func Connect() error {
	dbInstance, err := gorm.Open(sqlite.Open("../sensor_log.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	Q = *query.Use(dbInstance)
	initialized = true

	return nil
}

func AddSensorLogEntry(entry *model.SensorLogEntry) error {
	if err := initialize(); err != nil {
		return err
	}

	return Q.SensorLogEntry.Create(entry)
}

func GetAllSensorLogEntries() ([]*model.SensorLogEntry, error) {
	if err := initialize(); err != nil {
		return nil, err
	}

	return Q.SensorLogEntry.Find()
}

func initialize() error {
	if !initialized {
		if err := Connect(); err != nil {
			return err
		}
	}

	return nil
}
