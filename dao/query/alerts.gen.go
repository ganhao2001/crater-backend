// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/raids-lab/crater/dao/model"
)

func newAlert(db *gorm.DB, opts ...gen.DOOption) alert {
	_alert := alert{}

	_alert.alertDo.UseDB(db, opts...)
	_alert.alertDo.UseModel(&model.Alert{})

	tableName := _alert.alertDo.TableName()
	_alert.ALL = field.NewAsterisk(tableName)
	_alert.ID = field.NewUint(tableName, "id")
	_alert.CreatedAt = field.NewTime(tableName, "created_at")
	_alert.UpdatedAt = field.NewTime(tableName, "updated_at")
	_alert.DeletedAt = field.NewField(tableName, "deleted_at")
	_alert.JobName = field.NewString(tableName, "job_name")
	_alert.AlertType = field.NewString(tableName, "alert_type")
	_alert.AlertTimestamp = field.NewTime(tableName, "alert_timestamp")
	_alert.AllowRepeat = field.NewBool(tableName, "allow_repeat")
	_alert.SendCount = field.NewInt(tableName, "send_count")

	_alert.fillFieldMap()

	return _alert
}

type alert struct {
	alertDo alertDo

	ALL            field.Asterisk
	ID             field.Uint
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	JobName        field.String // 作业名
	AlertType      field.String // 邮件类型
	AlertTimestamp field.Time   // 邮件发送时间
	AllowRepeat    field.Bool   // 是否允许重复发送
	SendCount      field.Int    // 邮件发送次数

	fieldMap map[string]field.Expr
}

func (a alert) Table(newTableName string) *alert {
	a.alertDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a alert) As(alias string) *alert {
	a.alertDo.DO = *(a.alertDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *alert) updateTableName(table string) *alert {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.JobName = field.NewString(table, "job_name")
	a.AlertType = field.NewString(table, "alert_type")
	a.AlertTimestamp = field.NewTime(table, "alert_timestamp")
	a.AllowRepeat = field.NewBool(table, "allow_repeat")
	a.SendCount = field.NewInt(table, "send_count")

	a.fillFieldMap()

	return a
}

func (a *alert) WithContext(ctx context.Context) IAlertDo { return a.alertDo.WithContext(ctx) }

func (a alert) TableName() string { return a.alertDo.TableName() }

func (a alert) Alias() string { return a.alertDo.Alias() }

func (a alert) Columns(cols ...field.Expr) gen.Columns { return a.alertDo.Columns(cols...) }

func (a *alert) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *alert) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["job_name"] = a.JobName
	a.fieldMap["alert_type"] = a.AlertType
	a.fieldMap["alert_timestamp"] = a.AlertTimestamp
	a.fieldMap["allow_repeat"] = a.AllowRepeat
	a.fieldMap["send_count"] = a.SendCount
}

func (a alert) clone(db *gorm.DB) alert {
	a.alertDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a alert) replaceDB(db *gorm.DB) alert {
	a.alertDo.ReplaceDB(db)
	return a
}

type alertDo struct{ gen.DO }

type IAlertDo interface {
	gen.SubQuery
	Debug() IAlertDo
	WithContext(ctx context.Context) IAlertDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAlertDo
	WriteDB() IAlertDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAlertDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAlertDo
	Not(conds ...gen.Condition) IAlertDo
	Or(conds ...gen.Condition) IAlertDo
	Select(conds ...field.Expr) IAlertDo
	Where(conds ...gen.Condition) IAlertDo
	Order(conds ...field.Expr) IAlertDo
	Distinct(cols ...field.Expr) IAlertDo
	Omit(cols ...field.Expr) IAlertDo
	Join(table schema.Tabler, on ...field.Expr) IAlertDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAlertDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAlertDo
	Group(cols ...field.Expr) IAlertDo
	Having(conds ...gen.Condition) IAlertDo
	Limit(limit int) IAlertDo
	Offset(offset int) IAlertDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAlertDo
	Unscoped() IAlertDo
	Create(values ...*model.Alert) error
	CreateInBatches(values []*model.Alert, batchSize int) error
	Save(values ...*model.Alert) error
	First() (*model.Alert, error)
	Take() (*model.Alert, error)
	Last() (*model.Alert, error)
	Find() ([]*model.Alert, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Alert, err error)
	FindInBatches(result *[]*model.Alert, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Alert) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAlertDo
	Assign(attrs ...field.AssignExpr) IAlertDo
	Joins(fields ...field.RelationField) IAlertDo
	Preload(fields ...field.RelationField) IAlertDo
	FirstOrInit() (*model.Alert, error)
	FirstOrCreate() (*model.Alert, error)
	FindByPage(offset int, limit int) (result []*model.Alert, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAlertDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a alertDo) Debug() IAlertDo {
	return a.withDO(a.DO.Debug())
}

func (a alertDo) WithContext(ctx context.Context) IAlertDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a alertDo) ReadDB() IAlertDo {
	return a.Clauses(dbresolver.Read)
}

func (a alertDo) WriteDB() IAlertDo {
	return a.Clauses(dbresolver.Write)
}

func (a alertDo) Session(config *gorm.Session) IAlertDo {
	return a.withDO(a.DO.Session(config))
}

func (a alertDo) Clauses(conds ...clause.Expression) IAlertDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a alertDo) Returning(value interface{}, columns ...string) IAlertDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a alertDo) Not(conds ...gen.Condition) IAlertDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a alertDo) Or(conds ...gen.Condition) IAlertDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a alertDo) Select(conds ...field.Expr) IAlertDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a alertDo) Where(conds ...gen.Condition) IAlertDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a alertDo) Order(conds ...field.Expr) IAlertDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a alertDo) Distinct(cols ...field.Expr) IAlertDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a alertDo) Omit(cols ...field.Expr) IAlertDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a alertDo) Join(table schema.Tabler, on ...field.Expr) IAlertDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a alertDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAlertDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a alertDo) RightJoin(table schema.Tabler, on ...field.Expr) IAlertDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a alertDo) Group(cols ...field.Expr) IAlertDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a alertDo) Having(conds ...gen.Condition) IAlertDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a alertDo) Limit(limit int) IAlertDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a alertDo) Offset(offset int) IAlertDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a alertDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAlertDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a alertDo) Unscoped() IAlertDo {
	return a.withDO(a.DO.Unscoped())
}

func (a alertDo) Create(values ...*model.Alert) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a alertDo) CreateInBatches(values []*model.Alert, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a alertDo) Save(values ...*model.Alert) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a alertDo) First() (*model.Alert, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Alert), nil
	}
}

func (a alertDo) Take() (*model.Alert, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Alert), nil
	}
}

func (a alertDo) Last() (*model.Alert, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Alert), nil
	}
}

func (a alertDo) Find() ([]*model.Alert, error) {
	result, err := a.DO.Find()
	return result.([]*model.Alert), err
}

func (a alertDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Alert, err error) {
	buf := make([]*model.Alert, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a alertDo) FindInBatches(result *[]*model.Alert, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a alertDo) Attrs(attrs ...field.AssignExpr) IAlertDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a alertDo) Assign(attrs ...field.AssignExpr) IAlertDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a alertDo) Joins(fields ...field.RelationField) IAlertDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a alertDo) Preload(fields ...field.RelationField) IAlertDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a alertDo) FirstOrInit() (*model.Alert, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Alert), nil
	}
}

func (a alertDo) FirstOrCreate() (*model.Alert, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Alert), nil
	}
}

func (a alertDo) FindByPage(offset int, limit int) (result []*model.Alert, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a alertDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a alertDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a alertDo) Delete(models ...*model.Alert) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *alertDo) withDO(do gen.Dao) *alertDo {
	a.DO = *do.(*gen.DO)
	return a
}
