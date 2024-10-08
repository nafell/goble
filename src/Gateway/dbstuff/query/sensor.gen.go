// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/nafell/goble_gateway/dbstuff/model"
)

func newSensorLogEntry(db *gorm.DB, opts ...gen.DOOption) sensorLogEntry {
	_sensorLogEntry := sensorLogEntry{}

	_sensorLogEntry.sensorLogEntryDo.UseDB(db, opts...)
	_sensorLogEntry.sensorLogEntryDo.UseModel(&model.SensorLogEntry{})

	tableName := _sensorLogEntry.sensorLogEntryDo.TableName()
	_sensorLogEntry.ALL = field.NewAsterisk(tableName)
	_sensorLogEntry.Datetime = field.NewTime(tableName, "datetime")
	_sensorLogEntry.Co2 = field.NewInt32(tableName, "co2")
	_sensorLogEntry.Temperature = field.NewInt32(tableName, "temperature")
	_sensorLogEntry.Humidity = field.NewInt32(tableName, "humidity")
	_sensorLogEntry.Pressure = field.NewInt32(tableName, "pressure")

	_sensorLogEntry.fillFieldMap()

	return _sensorLogEntry
}

type sensorLogEntry struct {
	sensorLogEntryDo

	ALL         field.Asterisk
	Datetime    field.Time
	Co2         field.Int32
	Temperature field.Int32
	Humidity    field.Int32
	Pressure    field.Int32

	fieldMap map[string]field.Expr
}

func (s sensorLogEntry) Table(newTableName string) *sensorLogEntry {
	s.sensorLogEntryDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sensorLogEntry) As(alias string) *sensorLogEntry {
	s.sensorLogEntryDo.DO = *(s.sensorLogEntryDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sensorLogEntry) updateTableName(table string) *sensorLogEntry {
	s.ALL = field.NewAsterisk(table)
	s.Datetime = field.NewTime(table, "datetime")
	s.Co2 = field.NewInt32(table, "co2")
	s.Temperature = field.NewInt32(table, "temperature")
	s.Humidity = field.NewInt32(table, "humidity")
	s.Pressure = field.NewInt32(table, "pressure")

	s.fillFieldMap()

	return s
}

func (s *sensorLogEntry) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sensorLogEntry) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["datetime"] = s.Datetime
	s.fieldMap["co2"] = s.Co2
	s.fieldMap["temperature"] = s.Temperature
	s.fieldMap["humidity"] = s.Humidity
	s.fieldMap["pressure"] = s.Pressure
}

func (s sensorLogEntry) clone(db *gorm.DB) sensorLogEntry {
	s.sensorLogEntryDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sensorLogEntry) replaceDB(db *gorm.DB) sensorLogEntry {
	s.sensorLogEntryDo.ReplaceDB(db)
	return s
}

type sensorLogEntryDo struct{ gen.DO }

type ISensorLogEntryDo interface {
	gen.SubQuery
	Debug() ISensorLogEntryDo
	WithContext(ctx context.Context) ISensorLogEntryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISensorLogEntryDo
	WriteDB() ISensorLogEntryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISensorLogEntryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISensorLogEntryDo
	Not(conds ...gen.Condition) ISensorLogEntryDo
	Or(conds ...gen.Condition) ISensorLogEntryDo
	Select(conds ...field.Expr) ISensorLogEntryDo
	Where(conds ...gen.Condition) ISensorLogEntryDo
	Order(conds ...field.Expr) ISensorLogEntryDo
	Distinct(cols ...field.Expr) ISensorLogEntryDo
	Omit(cols ...field.Expr) ISensorLogEntryDo
	Join(table schema.Tabler, on ...field.Expr) ISensorLogEntryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISensorLogEntryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISensorLogEntryDo
	Group(cols ...field.Expr) ISensorLogEntryDo
	Having(conds ...gen.Condition) ISensorLogEntryDo
	Limit(limit int) ISensorLogEntryDo
	Offset(offset int) ISensorLogEntryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISensorLogEntryDo
	Unscoped() ISensorLogEntryDo
	Create(values ...*model.SensorLogEntry) error
	CreateInBatches(values []*model.SensorLogEntry, batchSize int) error
	Save(values ...*model.SensorLogEntry) error
	First() (*model.SensorLogEntry, error)
	Take() (*model.SensorLogEntry, error)
	Last() (*model.SensorLogEntry, error)
	Find() ([]*model.SensorLogEntry, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SensorLogEntry, err error)
	FindInBatches(result *[]*model.SensorLogEntry, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SensorLogEntry) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISensorLogEntryDo
	Assign(attrs ...field.AssignExpr) ISensorLogEntryDo
	Joins(fields ...field.RelationField) ISensorLogEntryDo
	Preload(fields ...field.RelationField) ISensorLogEntryDo
	FirstOrInit() (*model.SensorLogEntry, error)
	FirstOrCreate() (*model.SensorLogEntry, error)
	FindByPage(offset int, limit int) (result []*model.SensorLogEntry, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISensorLogEntryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sensorLogEntryDo) Debug() ISensorLogEntryDo {
	return s.withDO(s.DO.Debug())
}

func (s sensorLogEntryDo) WithContext(ctx context.Context) ISensorLogEntryDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sensorLogEntryDo) ReadDB() ISensorLogEntryDo {
	return s.Clauses(dbresolver.Read)
}

func (s sensorLogEntryDo) WriteDB() ISensorLogEntryDo {
	return s.Clauses(dbresolver.Write)
}

func (s sensorLogEntryDo) Session(config *gorm.Session) ISensorLogEntryDo {
	return s.withDO(s.DO.Session(config))
}

func (s sensorLogEntryDo) Clauses(conds ...clause.Expression) ISensorLogEntryDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sensorLogEntryDo) Returning(value interface{}, columns ...string) ISensorLogEntryDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sensorLogEntryDo) Not(conds ...gen.Condition) ISensorLogEntryDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sensorLogEntryDo) Or(conds ...gen.Condition) ISensorLogEntryDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sensorLogEntryDo) Select(conds ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sensorLogEntryDo) Where(conds ...gen.Condition) ISensorLogEntryDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sensorLogEntryDo) Order(conds ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sensorLogEntryDo) Distinct(cols ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sensorLogEntryDo) Omit(cols ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sensorLogEntryDo) Join(table schema.Tabler, on ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sensorLogEntryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sensorLogEntryDo) RightJoin(table schema.Tabler, on ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sensorLogEntryDo) Group(cols ...field.Expr) ISensorLogEntryDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sensorLogEntryDo) Having(conds ...gen.Condition) ISensorLogEntryDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sensorLogEntryDo) Limit(limit int) ISensorLogEntryDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sensorLogEntryDo) Offset(offset int) ISensorLogEntryDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sensorLogEntryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISensorLogEntryDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sensorLogEntryDo) Unscoped() ISensorLogEntryDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sensorLogEntryDo) Create(values ...*model.SensorLogEntry) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sensorLogEntryDo) CreateInBatches(values []*model.SensorLogEntry, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sensorLogEntryDo) Save(values ...*model.SensorLogEntry) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sensorLogEntryDo) First() (*model.SensorLogEntry, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SensorLogEntry), nil
	}
}

func (s sensorLogEntryDo) Take() (*model.SensorLogEntry, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SensorLogEntry), nil
	}
}

func (s sensorLogEntryDo) Last() (*model.SensorLogEntry, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SensorLogEntry), nil
	}
}

func (s sensorLogEntryDo) Find() ([]*model.SensorLogEntry, error) {
	result, err := s.DO.Find()
	return result.([]*model.SensorLogEntry), err
}

func (s sensorLogEntryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SensorLogEntry, err error) {
	buf := make([]*model.SensorLogEntry, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sensorLogEntryDo) FindInBatches(result *[]*model.SensorLogEntry, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sensorLogEntryDo) Attrs(attrs ...field.AssignExpr) ISensorLogEntryDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sensorLogEntryDo) Assign(attrs ...field.AssignExpr) ISensorLogEntryDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sensorLogEntryDo) Joins(fields ...field.RelationField) ISensorLogEntryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sensorLogEntryDo) Preload(fields ...field.RelationField) ISensorLogEntryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sensorLogEntryDo) FirstOrInit() (*model.SensorLogEntry, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SensorLogEntry), nil
	}
}

func (s sensorLogEntryDo) FirstOrCreate() (*model.SensorLogEntry, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SensorLogEntry), nil
	}
}

func (s sensorLogEntryDo) FindByPage(offset int, limit int) (result []*model.SensorLogEntry, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sensorLogEntryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sensorLogEntryDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sensorLogEntryDo) Delete(models ...*model.SensorLogEntry) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sensorLogEntryDo) withDO(do gen.Dao) *sensorLogEntryDo {
	s.DO = *do.(*gen.DO)
	return s
}
