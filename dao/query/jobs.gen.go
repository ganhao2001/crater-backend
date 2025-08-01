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

func newJob(db *gorm.DB, opts ...gen.DOOption) job {
	_job := job{}

	_job.jobDo.UseDB(db, opts...)
	_job.jobDo.UseModel(&model.Job{})

	tableName := _job.jobDo.TableName()
	_job.ALL = field.NewAsterisk(tableName)
	_job.ID = field.NewUint(tableName, "id")
	_job.CreatedAt = field.NewTime(tableName, "created_at")
	_job.UpdatedAt = field.NewTime(tableName, "updated_at")
	_job.DeletedAt = field.NewField(tableName, "deleted_at")
	_job.Name = field.NewString(tableName, "name")
	_job.JobName = field.NewString(tableName, "job_name")
	_job.UserID = field.NewUint(tableName, "user_id")
	_job.AccountID = field.NewUint(tableName, "account_id")
	_job.JobType = field.NewString(tableName, "job_type")
	_job.Status = field.NewString(tableName, "status")
	_job.CreationTimestamp = field.NewTime(tableName, "creation_timestamp")
	_job.RunningTimestamp = field.NewTime(tableName, "running_timestamp")
	_job.CompletedTimestamp = field.NewTime(tableName, "completed_timestamp")
	_job.Nodes = field.NewField(tableName, "nodes")
	_job.Resources = field.NewField(tableName, "resources")
	_job.Attributes = field.NewField(tableName, "attributes")
	_job.Template = field.NewString(tableName, "template")
	_job.AlertEnabled = field.NewBool(tableName, "alert_enabled")
	_job.Reminded = field.NewBool(tableName, "reminded")
	_job.KeepWhenLowResourceUsage = field.NewBool(tableName, "keep_when_low_resource_usage")
	_job.LockedTimestamp = field.NewTime(tableName, "locked_timestamp")
	_job.ProfileData = field.NewField(tableName, "profile_data")
	_job.ScheduleData = field.NewField(tableName, "schedule_data")
	_job.Events = field.NewField(tableName, "events")
	_job.TerminatedStates = field.NewField(tableName, "terminated_states")
	_job.User = jobBelongsToUser{
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

	_job.Account = jobBelongsToAccount{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Account", "model.Account"),
		UserAccounts: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Account.UserAccounts", "model.UserAccount"),
		},
		AccountDatasets: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Account.AccountDatasets", "model.AccountDataset"),
		},
	}

	_job.fillFieldMap()

	return _job
}

type job struct {
	jobDo jobDo

	ALL                      field.Asterisk
	ID                       field.Uint
	CreatedAt                field.Time
	UpdatedAt                field.Time
	DeletedAt                field.Field
	Name                     field.String // 作业名称
	JobName                  field.String // 作业名称
	UserID                   field.Uint
	AccountID                field.Uint
	JobType                  field.String // 作业类型
	Status                   field.String // 作业状态
	CreationTimestamp        field.Time   // 作业创建时间
	RunningTimestamp         field.Time   // 作业开始运行时间
	CompletedTimestamp       field.Time   // 作业完成时间
	Nodes                    field.Field  // 作业运行的节点
	Resources                field.Field  // 作业的资源需求
	Attributes               field.Field  // 作业的原始属性
	Template                 field.String // 作业的模板配置
	AlertEnabled             field.Bool   // 是否启用通知
	Reminded                 field.Bool   // 是否已经处于发送了提醒的状态
	KeepWhenLowResourceUsage field.Bool   // 当资源利用率低时是否保留
	LockedTimestamp          field.Time   // 作业锁定时间
	ProfileData              field.Field  // 作业的性能数据
	ScheduleData             field.Field  // 作业的调度数据
	Events                   field.Field  // 作业的事件 (运行时、失败时采集)
	TerminatedStates         field.Field  // 作业的终止状态 (运行时、失败时采集)
	User                     jobBelongsToUser

	Account jobBelongsToAccount

	fieldMap map[string]field.Expr
}

func (j job) Table(newTableName string) *job {
	j.jobDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j job) As(alias string) *job {
	j.jobDo.DO = *(j.jobDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *job) updateTableName(table string) *job {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewUint(table, "id")
	j.CreatedAt = field.NewTime(table, "created_at")
	j.UpdatedAt = field.NewTime(table, "updated_at")
	j.DeletedAt = field.NewField(table, "deleted_at")
	j.Name = field.NewString(table, "name")
	j.JobName = field.NewString(table, "job_name")
	j.UserID = field.NewUint(table, "user_id")
	j.AccountID = field.NewUint(table, "account_id")
	j.JobType = field.NewString(table, "job_type")
	j.Status = field.NewString(table, "status")
	j.CreationTimestamp = field.NewTime(table, "creation_timestamp")
	j.RunningTimestamp = field.NewTime(table, "running_timestamp")
	j.CompletedTimestamp = field.NewTime(table, "completed_timestamp")
	j.Nodes = field.NewField(table, "nodes")
	j.Resources = field.NewField(table, "resources")
	j.Attributes = field.NewField(table, "attributes")
	j.Template = field.NewString(table, "template")
	j.AlertEnabled = field.NewBool(table, "alert_enabled")
	j.Reminded = field.NewBool(table, "reminded")
	j.KeepWhenLowResourceUsage = field.NewBool(table, "keep_when_low_resource_usage")
	j.LockedTimestamp = field.NewTime(table, "locked_timestamp")
	j.ProfileData = field.NewField(table, "profile_data")
	j.ScheduleData = field.NewField(table, "schedule_data")
	j.Events = field.NewField(table, "events")
	j.TerminatedStates = field.NewField(table, "terminated_states")

	j.fillFieldMap()

	return j
}

func (j *job) WithContext(ctx context.Context) IJobDo { return j.jobDo.WithContext(ctx) }

func (j job) TableName() string { return j.jobDo.TableName() }

func (j job) Alias() string { return j.jobDo.Alias() }

func (j job) Columns(cols ...field.Expr) gen.Columns { return j.jobDo.Columns(cols...) }

func (j *job) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *job) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 27)
	j.fieldMap["id"] = j.ID
	j.fieldMap["created_at"] = j.CreatedAt
	j.fieldMap["updated_at"] = j.UpdatedAt
	j.fieldMap["deleted_at"] = j.DeletedAt
	j.fieldMap["name"] = j.Name
	j.fieldMap["job_name"] = j.JobName
	j.fieldMap["user_id"] = j.UserID
	j.fieldMap["account_id"] = j.AccountID
	j.fieldMap["job_type"] = j.JobType
	j.fieldMap["status"] = j.Status
	j.fieldMap["creation_timestamp"] = j.CreationTimestamp
	j.fieldMap["running_timestamp"] = j.RunningTimestamp
	j.fieldMap["completed_timestamp"] = j.CompletedTimestamp
	j.fieldMap["nodes"] = j.Nodes
	j.fieldMap["resources"] = j.Resources
	j.fieldMap["attributes"] = j.Attributes
	j.fieldMap["template"] = j.Template
	j.fieldMap["alert_enabled"] = j.AlertEnabled
	j.fieldMap["reminded"] = j.Reminded
	j.fieldMap["keep_when_low_resource_usage"] = j.KeepWhenLowResourceUsage
	j.fieldMap["locked_timestamp"] = j.LockedTimestamp
	j.fieldMap["profile_data"] = j.ProfileData
	j.fieldMap["schedule_data"] = j.ScheduleData
	j.fieldMap["events"] = j.Events
	j.fieldMap["terminated_states"] = j.TerminatedStates

}

func (j job) clone(db *gorm.DB) job {
	j.jobDo.ReplaceConnPool(db.Statement.ConnPool)
	j.User.db = db.Session(&gorm.Session{Initialized: true})
	j.User.db.Statement.ConnPool = db.Statement.ConnPool
	j.Account.db = db.Session(&gorm.Session{Initialized: true})
	j.Account.db.Statement.ConnPool = db.Statement.ConnPool
	return j
}

func (j job) replaceDB(db *gorm.DB) job {
	j.jobDo.ReplaceDB(db)
	j.User.db = db.Session(&gorm.Session{})
	j.Account.db = db.Session(&gorm.Session{})
	return j
}

type jobBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	UserAccounts struct {
		field.RelationField
	}
	UserDatasets struct {
		field.RelationField
	}
}

func (a jobBelongsToUser) Where(conds ...field.Expr) *jobBelongsToUser {
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

func (a jobBelongsToUser) WithContext(ctx context.Context) *jobBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a jobBelongsToUser) Session(session *gorm.Session) *jobBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a jobBelongsToUser) Model(m *model.Job) *jobBelongsToUserTx {
	return &jobBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

func (a jobBelongsToUser) Unscoped() *jobBelongsToUser {
	a.db = a.db.Unscoped()
	return &a
}

type jobBelongsToUserTx struct{ tx *gorm.Association }

func (a jobBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a jobBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a jobBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a jobBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a jobBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a jobBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

func (a jobBelongsToUserTx) Unscoped() *jobBelongsToUserTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type jobBelongsToAccount struct {
	db *gorm.DB

	field.RelationField

	UserAccounts struct {
		field.RelationField
	}
	AccountDatasets struct {
		field.RelationField
	}
}

func (a jobBelongsToAccount) Where(conds ...field.Expr) *jobBelongsToAccount {
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

func (a jobBelongsToAccount) WithContext(ctx context.Context) *jobBelongsToAccount {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a jobBelongsToAccount) Session(session *gorm.Session) *jobBelongsToAccount {
	a.db = a.db.Session(session)
	return &a
}

func (a jobBelongsToAccount) Model(m *model.Job) *jobBelongsToAccountTx {
	return &jobBelongsToAccountTx{a.db.Model(m).Association(a.Name())}
}

func (a jobBelongsToAccount) Unscoped() *jobBelongsToAccount {
	a.db = a.db.Unscoped()
	return &a
}

type jobBelongsToAccountTx struct{ tx *gorm.Association }

func (a jobBelongsToAccountTx) Find() (result *model.Account, err error) {
	return result, a.tx.Find(&result)
}

func (a jobBelongsToAccountTx) Append(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a jobBelongsToAccountTx) Replace(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a jobBelongsToAccountTx) Delete(values ...*model.Account) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a jobBelongsToAccountTx) Clear() error {
	return a.tx.Clear()
}

func (a jobBelongsToAccountTx) Count() int64 {
	return a.tx.Count()
}

func (a jobBelongsToAccountTx) Unscoped() *jobBelongsToAccountTx {
	a.tx = a.tx.Unscoped()
	return &a
}

type jobDo struct{ gen.DO }

type IJobDo interface {
	gen.SubQuery
	Debug() IJobDo
	WithContext(ctx context.Context) IJobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IJobDo
	WriteDB() IJobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IJobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IJobDo
	Not(conds ...gen.Condition) IJobDo
	Or(conds ...gen.Condition) IJobDo
	Select(conds ...field.Expr) IJobDo
	Where(conds ...gen.Condition) IJobDo
	Order(conds ...field.Expr) IJobDo
	Distinct(cols ...field.Expr) IJobDo
	Omit(cols ...field.Expr) IJobDo
	Join(table schema.Tabler, on ...field.Expr) IJobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IJobDo
	RightJoin(table schema.Tabler, on ...field.Expr) IJobDo
	Group(cols ...field.Expr) IJobDo
	Having(conds ...gen.Condition) IJobDo
	Limit(limit int) IJobDo
	Offset(offset int) IJobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IJobDo
	Unscoped() IJobDo
	Create(values ...*model.Job) error
	CreateInBatches(values []*model.Job, batchSize int) error
	Save(values ...*model.Job) error
	First() (*model.Job, error)
	Take() (*model.Job, error)
	Last() (*model.Job, error)
	Find() ([]*model.Job, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Job, err error)
	FindInBatches(result *[]*model.Job, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Job) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IJobDo
	Assign(attrs ...field.AssignExpr) IJobDo
	Joins(fields ...field.RelationField) IJobDo
	Preload(fields ...field.RelationField) IJobDo
	FirstOrInit() (*model.Job, error)
	FirstOrCreate() (*model.Job, error)
	FindByPage(offset int, limit int) (result []*model.Job, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IJobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (j jobDo) Debug() IJobDo {
	return j.withDO(j.DO.Debug())
}

func (j jobDo) WithContext(ctx context.Context) IJobDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jobDo) ReadDB() IJobDo {
	return j.Clauses(dbresolver.Read)
}

func (j jobDo) WriteDB() IJobDo {
	return j.Clauses(dbresolver.Write)
}

func (j jobDo) Session(config *gorm.Session) IJobDo {
	return j.withDO(j.DO.Session(config))
}

func (j jobDo) Clauses(conds ...clause.Expression) IJobDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jobDo) Returning(value interface{}, columns ...string) IJobDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jobDo) Not(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jobDo) Or(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jobDo) Select(conds ...field.Expr) IJobDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jobDo) Where(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jobDo) Order(conds ...field.Expr) IJobDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jobDo) Distinct(cols ...field.Expr) IJobDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jobDo) Omit(cols ...field.Expr) IJobDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jobDo) Join(table schema.Tabler, on ...field.Expr) IJobDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jobDo) LeftJoin(table schema.Tabler, on ...field.Expr) IJobDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jobDo) RightJoin(table schema.Tabler, on ...field.Expr) IJobDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jobDo) Group(cols ...field.Expr) IJobDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jobDo) Having(conds ...gen.Condition) IJobDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jobDo) Limit(limit int) IJobDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jobDo) Offset(offset int) IJobDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IJobDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jobDo) Unscoped() IJobDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jobDo) Create(values ...*model.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jobDo) CreateInBatches(values []*model.Job, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jobDo) Save(values ...*model.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jobDo) First() (*model.Job, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Take() (*model.Job, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Last() (*model.Job, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Find() ([]*model.Job, error) {
	result, err := j.DO.Find()
	return result.([]*model.Job), err
}

func (j jobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Job, err error) {
	buf := make([]*model.Job, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jobDo) FindInBatches(result *[]*model.Job, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jobDo) Attrs(attrs ...field.AssignExpr) IJobDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jobDo) Assign(attrs ...field.AssignExpr) IJobDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jobDo) Joins(fields ...field.RelationField) IJobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jobDo) Preload(fields ...field.RelationField) IJobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jobDo) FirstOrInit() (*model.Job, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) FirstOrCreate() (*model.Job, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) FindByPage(offset int, limit int) (result []*model.Job, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jobDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jobDo) Delete(models ...*model.Job) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jobDo) withDO(do gen.Dao) *jobDo {
	j.DO = *do.(*gen.DO)
	return j
}
