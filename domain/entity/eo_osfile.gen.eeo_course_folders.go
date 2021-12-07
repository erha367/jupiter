package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _EeoCourseFoldersMgr struct {
	*_BaseMgr
}

// EeoCourseFoldersMgr open func
func EeoCourseFoldersMgr(db *gorm.DB) *_EeoCourseFoldersMgr {
	if db == nil {
		panic(fmt.Errorf("EeoCourseFoldersMgr need init by db"))
	}
	return &_EeoCourseFoldersMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoCourseFoldersMgr) GetTableName() string {
	return "eeo_course_folders"
}

// Get 获取
func (obj *_EeoCourseFoldersMgr) Get() (result EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoCourseFoldersMgr) Gets() (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithFolderID folder_id获取 ID
func (obj *_EeoCourseFoldersMgr) WithFolderID(folderID int64) Option {
	return optionFunc(func(o *options) { o.query["folder_id"] = folderID })
}

// WithSid sid获取 机构id
func (obj *_EeoCourseFoldersMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["sid"] = sid })
}

// WithCourseID course_id获取 群id
func (obj *_EeoCourseFoldersMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithFolderName folder_name获取 目录名
func (obj *_EeoCourseFoldersMgr) WithFolderName(folderName string) Option {
	return optionFunc(func(o *options) { o.query["folder_name"] = folderName })
}

// WithFolderNode folder_node获取 文件层级
func (obj *_EeoCourseFoldersMgr) WithFolderNode(folderNode int8) Option {
	return optionFunc(func(o *options) { o.query["folder_node"] = folderNode })
}

// WithParentID parent_id获取 父级id
func (obj *_EeoCourseFoldersMgr) WithParentID(parentID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithFolderPath folder_path获取 目录结构
func (obj *_EeoCourseFoldersMgr) WithFolderPath(folderPath string) Option {
	return optionFunc(func(o *options) { o.query["folder_path"] = folderPath })
}

// WithExp1 exp1获取 备用字段
func (obj *_EeoCourseFoldersMgr) WithExp1(exp1 bool) Option {
	return optionFunc(func(o *options) { o.query["exp1"] = exp1 })
}

// WithOperateLevel operate_level获取 操作级别（0 自己创建可以随意删除复制移动，1系统创建不可删除、移动、复制、重命名、分享）
func (obj *_EeoCourseFoldersMgr) WithOperateLevel(operateLevel bool) Option {
	return optionFunc(func(o *options) { o.query["operate_level"] = operateLevel })
}

// WithIsdel isdel获取 是否删除(0没删除 1删除)
func (obj *_EeoCourseFoldersMgr) WithIsdel(isdel bool) Option {
	return optionFunc(func(o *options) { o.query["isdel"] = isdel })
}

// WithAddtime addtime获取 添加时间
func (obj *_EeoCourseFoldersMgr) WithAddtime(addtime int) Option {
	return optionFunc(func(o *options) { o.query["addtime"] = addtime })
}

// WithUptime uptime获取 修改时间
func (obj *_EeoCourseFoldersMgr) WithUptime(uptime int) Option {
	return optionFunc(func(o *options) { o.query["uptime"] = uptime })
}

// WithUserName user_name获取 创建人name
func (obj *_EeoCourseFoldersMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithUserID user_id获取 创建人uid
func (obj *_EeoCourseFoldersMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// GetByOption 功能选项模式获取
func (obj *_EeoCourseFoldersMgr) GetByOption(opts ...Option) (result EeoCourseFolders, err error) {
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
func (obj *_EeoCourseFoldersMgr) GetByOptions(opts ...Option) (results []*EeoCourseFolders, err error) {
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

// GetFromFolderID 通过folder_id获取内容 ID
func (obj *_EeoCourseFoldersMgr) GetFromFolderID(folderID int64) (result EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id = ?", folderID).Find(&result).Error

	return
}

// GetBatchFromFolderID 批量唯一主键查找 ID
func (obj *_EeoCourseFoldersMgr) GetBatchFromFolderID(folderIDs []int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id IN (?)", folderIDs).Find(&results).Error

	return
}

// GetFromSid 通过sid获取内容 机构id
func (obj *_EeoCourseFoldersMgr) GetFromSid(sid int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sid = ?", sid).Find(&results).Error

	return
}

// GetBatchFromSid 批量唯一主键查找 机构id
func (obj *_EeoCourseFoldersMgr) GetBatchFromSid(sids []int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sid IN (?)", sids).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 群id
func (obj *_EeoCourseFoldersMgr) GetFromCourseID(courseID int64) (result EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id = ?", courseID).Find(&result).Error

	return
}

// GetBatchFromCourseID 批量唯一主键查找 群id
func (obj *_EeoCourseFoldersMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromFolderName 通过folder_name获取内容 目录名
func (obj *_EeoCourseFoldersMgr) GetFromFolderName(folderName string) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_name = ?", folderName).Find(&results).Error

	return
}

// GetBatchFromFolderName 批量唯一主键查找 目录名
func (obj *_EeoCourseFoldersMgr) GetBatchFromFolderName(folderNames []string) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_name IN (?)", folderNames).Find(&results).Error

	return
}

// GetFromFolderNode 通过folder_node获取内容 文件层级
func (obj *_EeoCourseFoldersMgr) GetFromFolderNode(folderNode int8) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_node = ?", folderNode).Find(&results).Error

	return
}

// GetBatchFromFolderNode 批量唯一主键查找 文件层级
func (obj *_EeoCourseFoldersMgr) GetBatchFromFolderNode(folderNodes []int8) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_node IN (?)", folderNodes).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父级id
func (obj *_EeoCourseFoldersMgr) GetFromParentID(parentID int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量唯一主键查找 父级id
func (obj *_EeoCourseFoldersMgr) GetBatchFromParentID(parentIDs []int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromFolderPath 通过folder_path获取内容 目录结构
func (obj *_EeoCourseFoldersMgr) GetFromFolderPath(folderPath string) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_path = ?", folderPath).Find(&results).Error

	return
}

// GetBatchFromFolderPath 批量唯一主键查找 目录结构
func (obj *_EeoCourseFoldersMgr) GetBatchFromFolderPath(folderPaths []string) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_path IN (?)", folderPaths).Find(&results).Error

	return
}

// GetFromExp1 通过exp1获取内容 备用字段
func (obj *_EeoCourseFoldersMgr) GetFromExp1(exp1 bool) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 = ?", exp1).Find(&results).Error

	return
}

// GetBatchFromExp1 批量唯一主键查找 备用字段
func (obj *_EeoCourseFoldersMgr) GetBatchFromExp1(exp1s []bool) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 IN (?)", exp1s).Find(&results).Error

	return
}

// GetFromOperateLevel 通过operate_level获取内容 操作级别（0 自己创建可以随意删除复制移动，1系统创建不可删除、移动、复制、重命名、分享）
func (obj *_EeoCourseFoldersMgr) GetFromOperateLevel(operateLevel bool) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("operate_level = ?", operateLevel).Find(&results).Error

	return
}

// GetBatchFromOperateLevel 批量唯一主键查找 操作级别（0 自己创建可以随意删除复制移动，1系统创建不可删除、移动、复制、重命名、分享）
func (obj *_EeoCourseFoldersMgr) GetBatchFromOperateLevel(operateLevels []bool) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("operate_level IN (?)", operateLevels).Find(&results).Error

	return
}

// GetFromIsdel 通过isdel获取内容 是否删除(0没删除 1删除)
func (obj *_EeoCourseFoldersMgr) GetFromIsdel(isdel bool) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel = ?", isdel).Find(&results).Error

	return
}

// GetBatchFromIsdel 批量唯一主键查找 是否删除(0没删除 1删除)
func (obj *_EeoCourseFoldersMgr) GetBatchFromIsdel(isdels []bool) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel IN (?)", isdels).Find(&results).Error

	return
}

// GetFromAddtime 通过addtime获取内容 添加时间
func (obj *_EeoCourseFoldersMgr) GetFromAddtime(addtime int) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&results).Error

	return
}

// GetBatchFromAddtime 批量唯一主键查找 添加时间
func (obj *_EeoCourseFoldersMgr) GetBatchFromAddtime(addtimes []int) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime IN (?)", addtimes).Find(&results).Error

	return
}

// GetFromUptime 通过uptime获取内容 修改时间
func (obj *_EeoCourseFoldersMgr) GetFromUptime(uptime int) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime = ?", uptime).Find(&results).Error

	return
}

// GetBatchFromUptime 批量唯一主键查找 修改时间
func (obj *_EeoCourseFoldersMgr) GetBatchFromUptime(uptimes []int) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime IN (?)", uptimes).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容 创建人name
func (obj *_EeoCourseFoldersMgr) GetFromUserName(userName string) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_name = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量唯一主键查找 创建人name
func (obj *_EeoCourseFoldersMgr) GetBatchFromUserName(userNames []string) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_name IN (?)", userNames).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 创建人uid
func (obj *_EeoCourseFoldersMgr) GetFromUserID(userID int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 创建人uid
func (obj *_EeoCourseFoldersMgr) GetBatchFromUserID(userIDs []int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoCourseFoldersMgr) FetchByPrimaryKey(folderID int64) (result EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id = ?", folderID).Find(&result).Error

	return
}

// FetchIndexByICourseID  获取多个内容
func (obj *_EeoCourseFoldersMgr) FetchIndexByICourseID(courseID int64) (results []*EeoCourseFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id = ?", courseID).Find(&results).Error

	return
}
