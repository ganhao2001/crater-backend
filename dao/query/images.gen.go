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

func newImage(db *gorm.DB, opts ...gen.DOOption) image {
	_image := image{}

	_image.imageDo.UseDB(db, opts...)
	_image.imageDo.UseModel(&model.Image{})

	tableName := _image.imageDo.TableName()
	_image.ALL = field.NewAsterisk(tableName)
	_image.ID = field.NewUint(tableName, "id")
	_image.CreatedAt = field.NewTime(tableName, "created_at")
	_image.UpdatedAt = field.NewTime(tableName, "updated_at")
	_image.DeletedAt = field.NewField(tableName, "deleted_at")
	_image.UserID = field.NewUint(tableName, "user_id")
	_image.ImageLink = field.NewString(tableName, "image_link")
	_image.ImagePackName = field.NewString(tableName, "image_pack_name")
	_image.Description = field.NewString(tableName, "description")
	_image.IsPublic = field.NewBool(tableName, "is_public")
	_image.TaskType = field.NewString(tableName, "task_type")
	_image.ImageSource = field.NewUint8(tableName, "image_source")
	_image.Size = field.NewInt64(tableName, "size")
	_image.Tags = field.NewField(tableName, "tags")
	_image.Archs = field.NewField(tableName, "archs")
	_image.User = imageBelongsToUser{
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

	_image.fillFieldMap()

	return _image
}

type image struct {
	imageDo imageDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	UserID        field.Uint
	ImageLink     field.String // 镜像链接
	ImagePackName field.String // ImagePack CRD 名称
	Description   field.String // 描述
	IsPublic      field.Bool   // 是否公共
	TaskType      field.String // 镜像任务类型
	ImageSource   field.Uint8  // 镜像来源类型
	Size          field.Int64  // 镜像大小
	Tags          field.Field  // 镜像标签
	Archs         field.Field  // 镜像架构
	User          imageBelongsToUser

	fieldMap map[string]field.Expr
}

func (i image) Table(newTableName string) *image {
	i.imageDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i image) As(alias string) *image {
	i.imageDo.DO = *(i.imageDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *image) updateTableName(table string) *image {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewUint(table, "id")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")
	i.DeletedAt = field.NewField(table, "deleted_at")
	i.UserID = field.NewUint(table, "user_id")
	i.ImageLink = field.NewString(table, "image_link")
	i.ImagePackName = field.NewString(table, "image_pack_name")
	i.Description = field.NewString(table, "description")
	i.IsPublic = field.NewBool(table, "is_public")
	i.TaskType = field.NewString(table, "task_type")
	i.ImageSource = field.NewUint8(table, "image_source")
	i.Size = field.NewInt64(table, "size")
	i.Tags = field.NewField(table, "tags")
	i.Archs = field.NewField(table, "archs")

	i.fillFieldMap()

	return i
}

func (i *image) WithContext(ctx context.Context) IImageDo { return i.imageDo.WithContext(ctx) }

func (i image) TableName() string { return i.imageDo.TableName() }

func (i image) Alias() string { return i.imageDo.Alias() }

func (i image) Columns(cols ...field.Expr) gen.Columns { return i.imageDo.Columns(cols...) }

func (i *image) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *image) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 15)
	i.fieldMap["id"] = i.ID
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
	i.fieldMap["deleted_at"] = i.DeletedAt
	i.fieldMap["user_id"] = i.UserID
	i.fieldMap["image_link"] = i.ImageLink
	i.fieldMap["image_pack_name"] = i.ImagePackName
	i.fieldMap["description"] = i.Description
	i.fieldMap["is_public"] = i.IsPublic
	i.fieldMap["task_type"] = i.TaskType
	i.fieldMap["image_source"] = i.ImageSource
	i.fieldMap["size"] = i.Size
	i.fieldMap["tags"] = i.Tags
	i.fieldMap["archs"] = i.Archs

}

func (i image) clone(db *gorm.DB) image {
	i.imageDo.ReplaceConnPool(db.Statement.ConnPool)
	i.User.db = db.Session(&gorm.Session{Initialized: true})
	i.User.db.Statement.ConnPool = db.Statement.ConnPool
	return i
}

func (i image) replaceDB(db *gorm.DB) image {
	i.imageDo.ReplaceDB(db)
	i.User.db = db.Session(&gorm.Session{})
	return i
}

type imageBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	UserAccounts struct {
		field.RelationField
	}
	UserDatasets struct {
		field.RelationField
	}
}

func (a imageBelongsToUser) Where(conds ...field.Expr) *imageBelongsToUser {
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

func (a imageBelongsToUser) WithContext(ctx context.Context) *imageBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a imageBelongsToUser) Session(session *gorm.Session) *imageBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a imageBelongsToUser) Model(m *model.Image) *imageBelongsToUserTx {
	return &imageBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a imageBelongsToUser) Unscoped() *imageBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type imageBelongsToUserTx struct{ tx *gorm.Association }

func (a imageBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a imageBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a imageBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a imageBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a imageBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a imageBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a imageBelongsToUserTx) Unscoped() *imageBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type imageDo struct{ gen.DO }

type IImageDo interface {
	gen.SubQuery
	Debug() IImageDo
	WithContext(ctx context.Context) IImageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IImageDo
	WriteDB() IImageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IImageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IImageDo
	Not(conds ...gen.Condition) IImageDo
	Or(conds ...gen.Condition) IImageDo
	Select(conds ...field.Expr) IImageDo
	Where(conds ...gen.Condition) IImageDo
	Order(conds ...field.Expr) IImageDo
	Distinct(cols ...field.Expr) IImageDo
	Omit(cols ...field.Expr) IImageDo
	Join(table schema.Tabler, on ...field.Expr) IImageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IImageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IImageDo
	Group(cols ...field.Expr) IImageDo
	Having(conds ...gen.Condition) IImageDo
	Limit(limit int) IImageDo
	Offset(offset int) IImageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IImageDo
	Unscoped() IImageDo
	Create(values ...*model.Image) error
	CreateInBatches(values []*model.Image, batchSize int) error
	Save(values ...*model.Image) error
	First() (*model.Image, error)
	Take() (*model.Image, error)
	Last() (*model.Image, error)
	Find() ([]*model.Image, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Image, err error)
	FindInBatches(result *[]*model.Image, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Image) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IImageDo
	Assign(attrs ...field.AssignExpr) IImageDo
	Joins(fields ...field.RelationField) IImageDo
	Preload(fields ...field.RelationField) IImageDo
	FirstOrInit() (*model.Image, error)
	FirstOrCreate() (*model.Image, error)
	FindByPage(offset int, limit int) (result []*model.Image, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IImageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i imageDo) Debug() IImageDo {
	return i.withDO(i.DO.Debug())
}

func (i imageDo) WithContext(ctx context.Context) IImageDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imageDo) ReadDB() IImageDo {
	return i.Clauses(dbresolver.Read)
}

func (i imageDo) WriteDB() IImageDo {
	return i.Clauses(dbresolver.Write)
}

func (i imageDo) Session(config *gorm.Session) IImageDo {
	return i.withDO(i.DO.Session(config))
}

func (i imageDo) Clauses(conds ...clause.Expression) IImageDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imageDo) Returning(value interface{}, columns ...string) IImageDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imageDo) Not(conds ...gen.Condition) IImageDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imageDo) Or(conds ...gen.Condition) IImageDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imageDo) Select(conds ...field.Expr) IImageDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imageDo) Where(conds ...gen.Condition) IImageDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imageDo) Order(conds ...field.Expr) IImageDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imageDo) Distinct(cols ...field.Expr) IImageDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imageDo) Omit(cols ...field.Expr) IImageDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imageDo) Join(table schema.Tabler, on ...field.Expr) IImageDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IImageDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imageDo) RightJoin(table schema.Tabler, on ...field.Expr) IImageDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imageDo) Group(cols ...field.Expr) IImageDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imageDo) Having(conds ...gen.Condition) IImageDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imageDo) Limit(limit int) IImageDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imageDo) Offset(offset int) IImageDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IImageDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imageDo) Unscoped() IImageDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imageDo) Create(values ...*model.Image) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imageDo) CreateInBatches(values []*model.Image, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imageDo) Save(values ...*model.Image) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imageDo) First() (*model.Image, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) Take() (*model.Image, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) Last() (*model.Image, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) Find() ([]*model.Image, error) {
	result, err := i.DO.Find()
	return result.([]*model.Image), err
}

func (i imageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Image, err error) {
	buf := make([]*model.Image, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imageDo) FindInBatches(result *[]*model.Image, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imageDo) Attrs(attrs ...field.AssignExpr) IImageDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imageDo) Assign(attrs ...field.AssignExpr) IImageDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imageDo) Joins(fields ...field.RelationField) IImageDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imageDo) Preload(fields ...field.RelationField) IImageDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imageDo) FirstOrInit() (*model.Image, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) FirstOrCreate() (*model.Image, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Image), nil
	}
}

func (i imageDo) FindByPage(offset int, limit int) (result []*model.Image, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i imageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imageDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imageDo) Delete(models ...*model.Image) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imageDo) withDO(do gen.Dao) *imageDo {
	i.DO = *do.(*gen.DO)
	return i
}
