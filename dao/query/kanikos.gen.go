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

func newKaniko(db *gorm.DB, opts ...gen.DOOption) kaniko {
	_kaniko := kaniko{}

	_kaniko.kanikoDo.UseDB(db, opts...)
	_kaniko.kanikoDo.UseModel(&model.Kaniko{})

	tableName := _kaniko.kanikoDo.TableName()
	_kaniko.ALL = field.NewAsterisk(tableName)
	_kaniko.ID = field.NewUint(tableName, "id")
	_kaniko.CreatedAt = field.NewTime(tableName, "created_at")
	_kaniko.UpdatedAt = field.NewTime(tableName, "updated_at")
	_kaniko.DeletedAt = field.NewField(tableName, "deleted_at")
	_kaniko.UserID = field.NewUint(tableName, "user_id")
	_kaniko.ImagePackName = field.NewString(tableName, "image_pack_name")
	_kaniko.ImageLink = field.NewString(tableName, "image_link")
	_kaniko.NameSpace = field.NewString(tableName, "name_space")
	_kaniko.Status = field.NewString(tableName, "status")
	_kaniko.Description = field.NewString(tableName, "description")
	_kaniko.Size = field.NewInt64(tableName, "size")
	_kaniko.Dockerfile = field.NewString(tableName, "dockerfile")
	_kaniko.BuildSource = field.NewString(tableName, "build_source")
	_kaniko.Tags = field.NewField(tableName, "tags")
	_kaniko.Template = field.NewString(tableName, "template")
	_kaniko.Archs = field.NewField(tableName, "archs")
	_kaniko.User = kanikoBelongsToUser{
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

	_kaniko.fillFieldMap()

	return _kaniko
}

type kaniko struct {
	kanikoDo kanikoDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	UserID        field.Uint
	ImagePackName field.String // ImagePack CRD 名称
	ImageLink     field.String // 镜像链接
	NameSpace     field.String // 命名空间
	Status        field.String // 构建状态
	Description   field.String // 描述
	Size          field.Int64  // 镜像大小
	Dockerfile    field.String // Dockerfile内容
	BuildSource   field.String // 构建来源
	Tags          field.Field  // 镜像标签
	Template      field.String // 镜像的模板配置
	Archs         field.Field  // 镜像架构
	User          kanikoBelongsToUser

	fieldMap map[string]field.Expr
}

func (k kaniko) Table(newTableName string) *kaniko {
	k.kanikoDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kaniko) As(alias string) *kaniko {
	k.kanikoDo.DO = *(k.kanikoDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kaniko) updateTableName(table string) *kaniko {
	k.ALL = field.NewAsterisk(table)
	k.ID = field.NewUint(table, "id")
	k.CreatedAt = field.NewTime(table, "created_at")
	k.UpdatedAt = field.NewTime(table, "updated_at")
	k.DeletedAt = field.NewField(table, "deleted_at")
	k.UserID = field.NewUint(table, "user_id")
	k.ImagePackName = field.NewString(table, "image_pack_name")
	k.ImageLink = field.NewString(table, "image_link")
	k.NameSpace = field.NewString(table, "name_space")
	k.Status = field.NewString(table, "status")
	k.Description = field.NewString(table, "description")
	k.Size = field.NewInt64(table, "size")
	k.Dockerfile = field.NewString(table, "dockerfile")
	k.BuildSource = field.NewString(table, "build_source")
	k.Tags = field.NewField(table, "tags")
	k.Template = field.NewString(table, "template")
	k.Archs = field.NewField(table, "archs")

	k.fillFieldMap()

	return k
}

func (k *kaniko) WithContext(ctx context.Context) IKanikoDo { return k.kanikoDo.WithContext(ctx) }

func (k kaniko) TableName() string { return k.kanikoDo.TableName() }

func (k kaniko) Alias() string { return k.kanikoDo.Alias() }

func (k kaniko) Columns(cols ...field.Expr) gen.Columns { return k.kanikoDo.Columns(cols...) }

func (k *kaniko) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kaniko) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 17)
	k.fieldMap["id"] = k.ID
	k.fieldMap["created_at"] = k.CreatedAt
	k.fieldMap["updated_at"] = k.UpdatedAt
	k.fieldMap["deleted_at"] = k.DeletedAt
	k.fieldMap["user_id"] = k.UserID
	k.fieldMap["image_pack_name"] = k.ImagePackName
	k.fieldMap["image_link"] = k.ImageLink
	k.fieldMap["name_space"] = k.NameSpace
	k.fieldMap["status"] = k.Status
	k.fieldMap["description"] = k.Description
	k.fieldMap["size"] = k.Size
	k.fieldMap["dockerfile"] = k.Dockerfile
	k.fieldMap["build_source"] = k.BuildSource
	k.fieldMap["tags"] = k.Tags
	k.fieldMap["template"] = k.Template
	k.fieldMap["archs"] = k.Archs

}

func (k kaniko) clone(db *gorm.DB) kaniko {
	k.kanikoDo.ReplaceConnPool(db.Statement.ConnPool)
	k.User.db = db.Session(&gorm.Session{Initialized: true})
	k.User.db.Statement.ConnPool = db.Statement.ConnPool
	return k
}

func (k kaniko) replaceDB(db *gorm.DB) kaniko {
	k.kanikoDo.ReplaceDB(db)
	k.User.db = db.Session(&gorm.Session{})
	return k
}

type kanikoBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	UserAccounts struct {
		field.RelationField
	}
	UserDatasets struct {
		field.RelationField
	}
}

func (a kanikoBelongsToUser) Where(conds ...field.Expr) *kanikoBelongsToUser {
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

func (a kanikoBelongsToUser) WithContext(ctx context.Context) *kanikoBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a kanikoBelongsToUser) Session(session *gorm.Session) *kanikoBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a kanikoBelongsToUser) Model(m *model.Kaniko) *kanikoBelongsToUserTx {
	return &kanikoBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a kanikoBelongsToUser) Unscoped() *kanikoBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type kanikoBelongsToUserTx struct{ tx *gorm.Association }

func (a kanikoBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a kanikoBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a kanikoBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a kanikoBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a kanikoBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a kanikoBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a kanikoBelongsToUserTx) Unscoped() *kanikoBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type kanikoDo struct{ gen.DO }

type IKanikoDo interface {
	gen.SubQuery
	Debug() IKanikoDo
	WithContext(ctx context.Context) IKanikoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IKanikoDo
	WriteDB() IKanikoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IKanikoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IKanikoDo
	Not(conds ...gen.Condition) IKanikoDo
	Or(conds ...gen.Condition) IKanikoDo
	Select(conds ...field.Expr) IKanikoDo
	Where(conds ...gen.Condition) IKanikoDo
	Order(conds ...field.Expr) IKanikoDo
	Distinct(cols ...field.Expr) IKanikoDo
	Omit(cols ...field.Expr) IKanikoDo
	Join(table schema.Tabler, on ...field.Expr) IKanikoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IKanikoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IKanikoDo
	Group(cols ...field.Expr) IKanikoDo
	Having(conds ...gen.Condition) IKanikoDo
	Limit(limit int) IKanikoDo
	Offset(offset int) IKanikoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IKanikoDo
	Unscoped() IKanikoDo
	Create(values ...*model.Kaniko) error
	CreateInBatches(values []*model.Kaniko, batchSize int) error
	Save(values ...*model.Kaniko) error
	First() (*model.Kaniko, error)
	Take() (*model.Kaniko, error)
	Last() (*model.Kaniko, error)
	Find() ([]*model.Kaniko, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kaniko, err error)
	FindInBatches(result *[]*model.Kaniko, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Kaniko) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IKanikoDo
	Assign(attrs ...field.AssignExpr) IKanikoDo
	Joins(fields ...field.RelationField) IKanikoDo
	Preload(fields ...field.RelationField) IKanikoDo
	FirstOrInit() (*model.Kaniko, error)
	FirstOrCreate() (*model.Kaniko, error)
	FindByPage(offset int, limit int) (result []*model.Kaniko, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IKanikoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (k kanikoDo) Debug() IKanikoDo {
	return k.withDO(k.DO.Debug())
}

func (k kanikoDo) WithContext(ctx context.Context) IKanikoDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kanikoDo) ReadDB() IKanikoDo {
	return k.Clauses(dbresolver.Read)
}

func (k kanikoDo) WriteDB() IKanikoDo {
	return k.Clauses(dbresolver.Write)
}

func (k kanikoDo) Session(config *gorm.Session) IKanikoDo {
	return k.withDO(k.DO.Session(config))
}

func (k kanikoDo) Clauses(conds ...clause.Expression) IKanikoDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kanikoDo) Returning(value interface{}, columns ...string) IKanikoDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kanikoDo) Not(conds ...gen.Condition) IKanikoDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kanikoDo) Or(conds ...gen.Condition) IKanikoDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kanikoDo) Select(conds ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kanikoDo) Where(conds ...gen.Condition) IKanikoDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kanikoDo) Order(conds ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kanikoDo) Distinct(cols ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kanikoDo) Omit(cols ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kanikoDo) Join(table schema.Tabler, on ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kanikoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kanikoDo) RightJoin(table schema.Tabler, on ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kanikoDo) Group(cols ...field.Expr) IKanikoDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kanikoDo) Having(conds ...gen.Condition) IKanikoDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kanikoDo) Limit(limit int) IKanikoDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kanikoDo) Offset(offset int) IKanikoDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kanikoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IKanikoDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kanikoDo) Unscoped() IKanikoDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kanikoDo) Create(values ...*model.Kaniko) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kanikoDo) CreateInBatches(values []*model.Kaniko, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kanikoDo) Save(values ...*model.Kaniko) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kanikoDo) First() (*model.Kaniko, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kaniko), nil
	}
}

func (k kanikoDo) Take() (*model.Kaniko, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kaniko), nil
	}
}

func (k kanikoDo) Last() (*model.Kaniko, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kaniko), nil
	}
}

func (k kanikoDo) Find() ([]*model.Kaniko, error) {
	result, err := k.DO.Find()
	return result.([]*model.Kaniko), err
}

func (k kanikoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kaniko, err error) {
	buf := make([]*model.Kaniko, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kanikoDo) FindInBatches(result *[]*model.Kaniko, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kanikoDo) Attrs(attrs ...field.AssignExpr) IKanikoDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kanikoDo) Assign(attrs ...field.AssignExpr) IKanikoDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kanikoDo) Joins(fields ...field.RelationField) IKanikoDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kanikoDo) Preload(fields ...field.RelationField) IKanikoDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kanikoDo) FirstOrInit() (*model.Kaniko, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kaniko), nil
	}
}

func (k kanikoDo) FirstOrCreate() (*model.Kaniko, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kaniko), nil
	}
}

func (k kanikoDo) FindByPage(offset int, limit int) (result []*model.Kaniko, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k kanikoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k kanikoDo) Scan(result interface{}) (err error) {
	return k.DO.Scan(result)
}

func (k kanikoDo) Delete(models ...*model.Kaniko) (result gen.ResultInfo, err error) {
	return k.DO.Delete(models)
}

func (k *kanikoDo) withDO(do gen.Dao) *kanikoDo {
	k.DO = *do.(*gen.DO)
	return k
}
