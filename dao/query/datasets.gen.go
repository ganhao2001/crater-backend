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

func newDataset(db *gorm.DB, opts ...gen.DOOption) dataset {
	_dataset := dataset{}

	_dataset.datasetDo.UseDB(db, opts...)
	_dataset.datasetDo.UseModel(&model.Dataset{})

	tableName := _dataset.datasetDo.TableName()
	_dataset.ALL = field.NewAsterisk(tableName)
	_dataset.ID = field.NewUint(tableName, "id")
	_dataset.CreatedAt = field.NewTime(tableName, "created_at")
	_dataset.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dataset.DeletedAt = field.NewField(tableName, "deleted_at")
	_dataset.Name = field.NewString(tableName, "name")
	_dataset.URL = field.NewString(tableName, "url")
	_dataset.Describe = field.NewString(tableName, "describe")
	_dataset.Type = field.NewString(tableName, "type")
	_dataset.Extra = field.NewField(tableName, "extra")
	_dataset.UserID = field.NewUint(tableName, "user_id")
	_dataset.UserDatasets = datasetHasManyUserDatasets{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("UserDatasets", "model.UserDataset"),
	}

	_dataset.AccountDatasets = datasetHasManyAccountDatasets{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AccountDatasets", "model.AccountDataset"),
	}

	_dataset.User = datasetBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
		UserAccounts: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.UserAccounts", "model.UserAccount"),
		},
		UserDatasets: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.UserDatasets", "model.UserDataset"),
		},
	}

	_dataset.fillFieldMap()

	return _dataset
}

type dataset struct {
	datasetDo datasetDo

	ALL          field.Asterisk
	ID           field.Uint
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	Name         field.String // 数据集名
	URL          field.String // 数据集空间路径
	Describe     field.String // 数据集描述
	Type         field.String // 数据类型
	Extra        field.Field  // 额外信息(tags、weburl等)
	UserID       field.Uint
	UserDatasets datasetHasManyUserDatasets

	AccountDatasets datasetHasManyAccountDatasets

	User datasetBelongsToUser

	fieldMap map[string]field.Expr
}

func (d dataset) Table(newTableName string) *dataset {
	d.datasetDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dataset) As(alias string) *dataset {
	d.datasetDo.DO = *(d.datasetDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dataset) updateTableName(table string) *dataset {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewUint(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Name = field.NewString(table, "name")
	d.URL = field.NewString(table, "url")
	d.Describe = field.NewString(table, "describe")
	d.Type = field.NewString(table, "type")
	d.Extra = field.NewField(table, "extra")
	d.UserID = field.NewUint(table, "user_id")

	d.fillFieldMap()

	return d
}

func (d *dataset) WithContext(ctx context.Context) IDatasetDo { return d.datasetDo.WithContext(ctx) }

func (d dataset) TableName() string { return d.datasetDo.TableName() }

func (d dataset) Alias() string { return d.datasetDo.Alias() }

func (d dataset) Columns(cols ...field.Expr) gen.Columns { return d.datasetDo.Columns(cols...) }

func (d *dataset) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dataset) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 13)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["name"] = d.Name
	d.fieldMap["url"] = d.URL
	d.fieldMap["describe"] = d.Describe
	d.fieldMap["type"] = d.Type
	d.fieldMap["extra"] = d.Extra
	d.fieldMap["user_id"] = d.UserID

}

func (d dataset) clone(db *gorm.DB) dataset {
	d.datasetDo.ReplaceConnPool(db.Statement.ConnPool)
	d.UserDatasets.db = db.Session(&gorm.Session{Initialized: true})
	d.UserDatasets.db.Statement.ConnPool = db.Statement.ConnPool
	d.AccountDatasets.db = db.Session(&gorm.Session{Initialized: true})
	d.AccountDatasets.db.Statement.ConnPool = db.Statement.ConnPool
	d.User.db = db.Session(&gorm.Session{Initialized: true})
	d.User.db.Statement.ConnPool = db.Statement.ConnPool
	return d
}

func (d dataset) replaceDB(db *gorm.DB) dataset {
	d.datasetDo.ReplaceDB(db)
	d.UserDatasets.db = db.Session(&gorm.Session{})
	d.AccountDatasets.db = db.Session(&gorm.Session{})
	d.User.db = db.Session(&gorm.Session{})
	return d
}

type datasetHasManyUserDatasets struct {
	db *gorm.DB

	field.RelationField
}

func (a datasetHasManyUserDatasets) Where(conds ...field.Expr) *datasetHasManyUserDatasets {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a datasetHasManyUserDatasets) WithContext(ctx context.Context) *datasetHasManyUserDatasets {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a datasetHasManyUserDatasets) Session(session *gorm.Session) *datasetHasManyUserDatasets {
	a.db = a.db.Session(session)
	return &a
}

func (a datasetHasManyUserDatasets) Model(m *model.Dataset) *datasetHasManyUserDatasetsTx {
	return &datasetHasManyUserDatasetsTx{a.db.Model(m).Association(a.Name())}
}

func (a datasetHasManyUserDatasets) Unscoped() *datasetHasManyUserDatasets {
	a.db = a.db.Unscoped()
	return &a
}

type datasetHasManyUserDatasetsTx struct{ tx *gorm.Association }

func (a datasetHasManyUserDatasetsTx) Find() (result []*model.UserDataset, err error) {
	return result, a.tx.Find(&result)
}

func (a datasetHasManyUserDatasetsTx) Append(values ...*model.UserDataset) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a datasetHasManyUserDatasetsTx) Replace(values ...*model.UserDataset) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a datasetHasManyUserDatasetsTx) Delete(values ...*model.UserDataset) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a datasetHasManyUserDatasetsTx) Clear() error {
	return a.tx.Clear()
}

func (a datasetHasManyUserDatasetsTx) Count() int64 {
	return a.tx.Count()
}

func (a datasetHasManyUserDatasetsTx) Unscoped() *datasetHasManyUserDatasetsTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type datasetHasManyAccountDatasets struct {
	db *gorm.DB

	field.RelationField
}

func (a datasetHasManyAccountDatasets) Where(conds ...field.Expr) *datasetHasManyAccountDatasets {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a datasetHasManyAccountDatasets) WithContext(ctx context.Context) *datasetHasManyAccountDatasets {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a datasetHasManyAccountDatasets) Session(session *gorm.Session) *datasetHasManyAccountDatasets {
	a.db = a.db.Session(session)
	return &a
}

func (a datasetHasManyAccountDatasets) Model(m *model.Dataset) *datasetHasManyAccountDatasetsTx {
	return &datasetHasManyAccountDatasetsTx{a.db.Model(m).Association(a.Name())}
}

func (a datasetHasManyAccountDatasets) Unscoped() *datasetHasManyAccountDatasets {
	a.db = a.db.Unscoped()
	return &a
}

type datasetHasManyAccountDatasetsTx struct{ tx *gorm.Association }

func (a datasetHasManyAccountDatasetsTx) Find() (result []*model.AccountDataset, err error) {
	return result, a.tx.Find(&result)
}

func (a datasetHasManyAccountDatasetsTx) Append(values ...*model.AccountDataset) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a datasetHasManyAccountDatasetsTx) Replace(values ...*model.AccountDataset) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a datasetHasManyAccountDatasetsTx) Delete(values ...*model.AccountDataset) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a datasetHasManyAccountDatasetsTx) Clear() error {
	return a.tx.Clear()
}

func (a datasetHasManyAccountDatasetsTx) Count() int64 {
	return a.tx.Count()
}

func (a datasetHasManyAccountDatasetsTx) Unscoped() *datasetHasManyAccountDatasetsTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type datasetBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	UserAccounts struct {
		field.RelationField
	}
	UserDatasets struct {
		field.RelationField
	}
}

func (a datasetBelongsToUser) Where(conds ...field.Expr) *datasetBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a datasetBelongsToUser) WithContext(ctx context.Context) *datasetBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a datasetBelongsToUser) Session(session *gorm.Session) *datasetBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a datasetBelongsToUser) Model(m *model.Dataset) *datasetBelongsToUserTx {
	return &datasetBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a datasetBelongsToUser) Unscoped() *datasetBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type datasetBelongsToUserTx struct{ tx *gorm.Association }

func (a datasetBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a datasetBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a datasetBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a datasetBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a datasetBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a datasetBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a datasetBelongsToUserTx) Unscoped() *datasetBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type datasetDo struct{ gen.DO }

type IDatasetDo interface {
	gen.SubQuery
	Debug() IDatasetDo
	WithContext(ctx context.Context) IDatasetDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDatasetDo
	WriteDB() IDatasetDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDatasetDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDatasetDo
	Not(conds ...gen.Condition) IDatasetDo
	Or(conds ...gen.Condition) IDatasetDo
	Select(conds ...field.Expr) IDatasetDo
	Where(conds ...gen.Condition) IDatasetDo
	Order(conds ...field.Expr) IDatasetDo
	Distinct(cols ...field.Expr) IDatasetDo
	Omit(cols ...field.Expr) IDatasetDo
	Join(table schema.Tabler, on ...field.Expr) IDatasetDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDatasetDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDatasetDo
	Group(cols ...field.Expr) IDatasetDo
	Having(conds ...gen.Condition) IDatasetDo
	Limit(limit int) IDatasetDo
	Offset(offset int) IDatasetDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDatasetDo
	Unscoped() IDatasetDo
	Create(values ...*model.Dataset) error
	CreateInBatches(values []*model.Dataset, batchSize int) error
	Save(values ...*model.Dataset) error
	First() (*model.Dataset, error)
	Take() (*model.Dataset, error)
	Last() (*model.Dataset, error)
	Find() ([]*model.Dataset, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dataset, err error)
	FindInBatches(result *[]*model.Dataset, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Dataset) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDatasetDo
	Assign(attrs ...field.AssignExpr) IDatasetDo
	Joins(fields ...field.RelationField) IDatasetDo
	Preload(fields ...field.RelationField) IDatasetDo
	FirstOrInit() (*model.Dataset, error)
	FirstOrCreate() (*model.Dataset, error)
	FindByPage(offset int, limit int) (result []*model.Dataset, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDatasetDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d datasetDo) Debug() IDatasetDo {
	return d.withDO(d.DO.Debug())
}

func (d datasetDo) WithContext(ctx context.Context) IDatasetDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d datasetDo) ReadDB() IDatasetDo {
	return d.Clauses(dbresolver.Read)
}

func (d datasetDo) WriteDB() IDatasetDo {
	return d.Clauses(dbresolver.Write)
}

func (d datasetDo) Session(config *gorm.Session) IDatasetDo {
	return d.withDO(d.DO.Session(config))
}

func (d datasetDo) Clauses(conds ...clause.Expression) IDatasetDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d datasetDo) Returning(value interface{}, columns ...string) IDatasetDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d datasetDo) Not(conds ...gen.Condition) IDatasetDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d datasetDo) Or(conds ...gen.Condition) IDatasetDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d datasetDo) Select(conds ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d datasetDo) Where(conds ...gen.Condition) IDatasetDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d datasetDo) Order(conds ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d datasetDo) Distinct(cols ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d datasetDo) Omit(cols ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d datasetDo) Join(table schema.Tabler, on ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d datasetDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d datasetDo) RightJoin(table schema.Tabler, on ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d datasetDo) Group(cols ...field.Expr) IDatasetDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d datasetDo) Having(conds ...gen.Condition) IDatasetDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d datasetDo) Limit(limit int) IDatasetDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d datasetDo) Offset(offset int) IDatasetDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d datasetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDatasetDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d datasetDo) Unscoped() IDatasetDo {
	return d.withDO(d.DO.Unscoped())
}

func (d datasetDo) Create(values ...*model.Dataset) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d datasetDo) CreateInBatches(values []*model.Dataset, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d datasetDo) Save(values ...*model.Dataset) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d datasetDo) First() (*model.Dataset, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dataset), nil
	}
}

func (d datasetDo) Take() (*model.Dataset, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dataset), nil
	}
}

func (d datasetDo) Last() (*model.Dataset, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dataset), nil
	}
}

func (d datasetDo) Find() ([]*model.Dataset, error) {
	result, err := d.DO.Find()
	return result.([]*model.Dataset), err
}

func (d datasetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dataset, err error) {
	buf := make([]*model.Dataset, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d datasetDo) FindInBatches(result *[]*model.Dataset, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d datasetDo) Attrs(attrs ...field.AssignExpr) IDatasetDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d datasetDo) Assign(attrs ...field.AssignExpr) IDatasetDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d datasetDo) Joins(fields ...field.RelationField) IDatasetDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d datasetDo) Preload(fields ...field.RelationField) IDatasetDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d datasetDo) FirstOrInit() (*model.Dataset, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dataset), nil
	}
}

func (d datasetDo) FirstOrCreate() (*model.Dataset, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dataset), nil
	}
}

func (d datasetDo) FindByPage(offset int, limit int) (result []*model.Dataset, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d datasetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d datasetDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d datasetDo) Delete(models ...*model.Dataset) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *datasetDo) withDO(do gen.Dao) *datasetDo {
	d.DO = *do.(*gen.DO)
	return d
}
