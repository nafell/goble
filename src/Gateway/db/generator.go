package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gen"
)

func init() error {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode: gen.WithoutContext | gen,WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag: true,
		FieldNullable: true,
	})

	db, err := gorm.Open(sqlite.Open("../sensor.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModelAs("sensor", "SensorLogEntry")
	)


	g.Execute()

	return nil
}
