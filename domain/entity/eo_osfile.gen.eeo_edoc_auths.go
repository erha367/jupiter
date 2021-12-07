package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _EeoEdocAuthsMgr struct {
	*_BaseMgr
}

// EeoEdocAuthsMgr open func
func EeoEdocAuthsMgr(db *gorm.DB) *_EeoEdocAuthsMgr {
	if db == nil {
		panic(fmt.Errorf("EeoEdocAuthsMgr need init by db"))
	}
	return &_EeoEdocAuthsMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoEdocAuthsMgr) GetTableName() string {
	return "eeo_edoc_auths"
}

// Get 获取
func (obj *_EeoEdocAuthsMgr) Get() (result EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoEdocAuthsMgr) Gets() (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EeoEdocAuthsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFileID file_id获取 协作文档id
func (obj *_EeoEdocAuthsMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithUserID user_id获取 被授权人id
func (obj *_EeoEdocAuthsMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCourseID course_id获取 被授权群id
func (obj *_EeoEdocAuthsMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithEdocAuth edoc_auth获取 文档权限 0无 1读 2写
func (obj *_EeoEdocAuthsMgr) WithEdocAuth(edocAuth int8) Option {
	return optionFunc(func(o *options) { o.query["edoc_auth"] = edocAuth })
}

// WithIsDel is_del获取 是否删除 0否 1是
func (obj *_EeoEdocAuthsMgr) WithIsDel(isDel int8) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_EeoEdocAuthsMgr) WithCreatedAt(createdAt int) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 修改时间
func (obj *_EeoEdocAuthsMgr) WithUpdatedAt(updatedAt int) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_EeoEdocAuthsMgr) GetByOption(opts ...Option) (result EeoEdocAuths, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EeoEdocAuthsMgr) GetByOptions(opts ...Option) (results []*EeoEdocAuths, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_EeoEdocAuthsMgr) GetFromID(id int64) (result EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_EeoEdocAuthsMgr) GetBatchFromID(ids []int64) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromFileID 通过file_id获取内容 协作文档id
func (obj *_EeoEdocAuthsMgr) GetFromFileID(fileID int64) (result EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&result).Error

	return
}

// GetBatchFromFileID 批量唯一主键查找 协作文档id
func (obj *_EeoEdocAuthsMgr) GetBatchFromFileID(fileIDs []int64) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 被授权人id
func (obj *_EeoEdocAuthsMgr) GetFromUserID(userID int64) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 被授权人id
func (obj *_EeoEdocAuthsMgr) GetBatchFromUserID(userIDs []int64) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 被授权群id
func (obj *_EeoEdocAuthsMgr) GetFromCourseID(courseID int64) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量唯一主键查找 被授权群id
func (obj *_EeoEdocAuthsMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromEdocAuth 通过edoc_auth获取内容 文档权限 0无 1读 2写
func (obj *_EeoEdocAuthsMgr) GetFromEdocAuth(edocAuth int8) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("edoc_auth = ?", edocAuth).Find(&results).Error

	return
}

// GetBatchFromEdocAuth 批量唯一主键查找 文档权限 0无 1读 2写
func (obj *_EeoEdocAuthsMgr) GetBatchFromEdocAuth(edocAuths []int8) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("edoc_auth IN (?)", edocAuths).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容 是否删除 0否 1是
func (obj *_EeoEdocAuthsMgr) GetFromIsDel(isDel int8) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_del = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量唯一主键查找 是否删除 0否 1是
func (obj *_EeoEdocAuthsMgr) GetBatchFromIsDel(isDels []int8) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_del IN (?)", isDels).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_EeoEdocAuthsMgr) GetFromCreatedAt(createdAt int) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建时间
func (obj *_EeoEdocAuthsMgr) GetBatchFromCreatedAt(createdAts []int) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 修改时间
func (obj *_EeoEdocAuthsMgr) GetFromUpdatedAt(updatedAt int) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 修改时间
func (obj *_EeoEdocAuthsMgr) GetBatchFromUpdatedAt(updatedAts []int) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoEdocAuthsMgr) FetchByPrimaryKey(id int64) (result EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIFileID  获取多个内容
func (obj *_EeoEdocAuthsMgr) FetchIndexByIFileID(fileID int64) (results []*EeoEdocAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&results).Error

	return
}
