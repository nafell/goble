package dbstuff

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gen"
)

func Generate() error {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./dbstuff/query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	db, err := gorm.Open(sqlite.Open("../sensor_log.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModelAs("sensor", "SensorLogEntry"))

	g.Execute()

	return nil
}
