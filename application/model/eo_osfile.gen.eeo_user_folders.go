package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _EeoUserFoldersMgr struct {
	*_BaseMgr
}

// EeoUserFoldersMgr open func
func EeoUserFoldersMgr(db *gorm.DB) *_EeoUserFoldersMgr {
	if db == nil {
		panic(fmt.Errorf("EeoUserFoldersMgr need init by db"))
	}
	return &_EeoUserFoldersMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoUserFoldersMgr) GetTableName() string {
	return "eeo_user_folders"
}

// Get 获取
func (obj *_EeoUserFoldersMgr) Get() (result EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoUserFoldersMgr) Gets() (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithFolderID folder_id获取 ID
func (obj *_EeoUserFoldersMgr) WithFolderID(folderID int64) Option {
	return optionFunc(func(o *options) { o.query["folder_id"] = folderID })
}

// WithUserID user_id获取 用户id
func (obj *_EeoUserFoldersMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithFolderName folder_name获取
func (obj *_EeoUserFoldersMgr) WithFolderName(folderName string) Option {
	return optionFunc(func(o *options) { o.query["folder_name"] = folderName })
}

// WithFolderNamePy folder_name_py获取 文件夹名称拼音
func (obj *_EeoUserFoldersMgr) WithFolderNamePy(folderNamePy string) Option {
	return optionFunc(func(o *options) { o.query["folder_name_py"] = folderNamePy })
}

// WithFolderNode folder_node获取 文件层级
func (obj *_EeoUserFoldersMgr) WithFolderNode(folderNode bool) Option {
	return optionFunc(func(o *options) { o.query["folder_node"] = folderNode })
}

// WithParentID parent_id获取 父级id
func (obj *_EeoUserFoldersMgr) WithParentID(parentID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithFolderPath folder_path获取 目录结构
func (obj *_EeoUserFoldersMgr) WithFolderPath(folderPath string) Option {
	return optionFunc(func(o *options) { o.query["folder_path"] = folderPath })
}

// WithSystemLock system_lock获取 是否可分享(根据文件夹下用户存放的文件决定的 0为可分享 1为不可分享)
func (obj *_EeoUserFoldersMgr) WithSystemLock(systemLock bool) Option {
	return optionFunc(func(o *options) { o.query["system_lock"] = systemLock })
}

// WithExp1 exp1获取 备用字段
func (obj *_EeoUserFoldersMgr) WithExp1(exp1 bool) Option {
	return optionFunc(func(o *options) { o.query["exp1"] = exp1 })
}

// WithOperateLevel operate_level获取 操作级别（0 自己创建可以随意删除复制移动，1系统创建不可删除、移动、复制、重命名、分享）
func (obj *_EeoUserFoldersMgr) WithOperateLevel(operateLevel bool) Option {
	return optionFunc(func(o *options) { o.query["operate_level"] = operateLevel })
}

// WithIsdel isdel获取 是否删除(0没删除 1删除)
func (obj *_EeoUserFoldersMgr) WithIsdel(isdel bool) Option {
	return optionFunc(func(o *options) { o.query["isdel"] = isdel })
}

// WithAddtime addtime获取 添加时间
func (obj *_EeoUserFoldersMgr) WithAddtime(addtime int) Option {
	return optionFunc(func(o *options) { o.query["addtime"] = addtime })
}

// WithUptime uptime获取 最后修改时间
func (obj *_EeoUserFoldersMgr) WithUptime(uptime int) Option {
	return optionFunc(func(o *options) { o.query["uptime"] = uptime })
}

// WithFromID from_id获取 来自文件夹id
func (obj *_EeoUserFoldersMgr) WithFromID(fromID int64) Option {
	return optionFunc(func(o *options) { o.query["from_id"] = fromID })
}

// WithFromParent from_parent获取 原来的父级id
func (obj *_EeoUserFoldersMgr) WithFromParent(fromParent int64) Option {
	return optionFunc(func(o *options) { o.query["from_parent"] = fromParent })
}

// WithType type获取 目录类型：0=课件 1=考试 2=作业
func (obj *_EeoUserFoldersMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithOldID old_id获取 迁移前目录
func (obj *_EeoUserFoldersMgr) WithOldID(oldID int64) Option {
	return optionFunc(func(o *options) { o.query["old_id"] = oldID })
}

// WithBatchOperateTime batch_operate_time获取 批量复制微秒时间戳
func (obj *_EeoUserFoldersMgr) WithBatchOperateTime(batchOperateTime int64) Option {
	return optionFunc(func(o *options) { o.query["batch_operate_time"] = batchOperateTime })
}

// GetByOption 功能选项模式获取
func (obj *_EeoUserFoldersMgr) GetByOption(opts ...Option) (result EeoUserFolders, err error) {
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
func (obj *_EeoUserFoldersMgr) GetByOptions(opts ...Option) (results []*EeoUserFolders, err error) {
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
func (obj *_EeoUserFoldersMgr) GetFromFolderID(folderID int64) (result EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id = ?", folderID).Find(&result).Error

	return
}

// GetBatchFromFolderID 批量唯一主键查找 ID
func (obj *_EeoUserFoldersMgr) GetBatchFromFolderID(folderIDs []int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id IN (?)", folderIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_EeoUserFoldersMgr) GetFromUserID(userID int64) (result EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_EeoUserFoldersMgr) GetBatchFromUserID(userIDs []int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromFolderName 通过folder_name获取内容
func (obj *_EeoUserFoldersMgr) GetFromFolderName(folderName string) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_name = ?", folderName).Find(&results).Error

	return
}

// GetBatchFromFolderName 批量唯一主键查找
func (obj *_EeoUserFoldersMgr) GetBatchFromFolderName(folderNames []string) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_name IN (?)", folderNames).Find(&results).Error

	return
}

// GetFromFolderNamePy 通过folder_name_py获取内容 文件夹名称拼音
func (obj *_EeoUserFoldersMgr) GetFromFolderNamePy(folderNamePy string) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_name_py = ?", folderNamePy).Find(&results).Error

	return
}

// GetBatchFromFolderNamePy 批量唯一主键查找 文件夹名称拼音
func (obj *_EeoUserFoldersMgr) GetBatchFromFolderNamePy(folderNamePys []string) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_name_py IN (?)", folderNamePys).Find(&results).Error

	return
}

// GetFromFolderNode 通过folder_node获取内容 文件层级
func (obj *_EeoUserFoldersMgr) GetFromFolderNode(folderNode bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_node = ?", folderNode).Find(&results).Error

	return
}

// GetBatchFromFolderNode 批量唯一主键查找 文件层级
func (obj *_EeoUserFoldersMgr) GetBatchFromFolderNode(folderNodes []bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_node IN (?)", folderNodes).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父级id
func (obj *_EeoUserFoldersMgr) GetFromParentID(parentID int64) (result EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&result).Error

	return
}

// GetBatchFromParentID 批量唯一主键查找 父级id
func (obj *_EeoUserFoldersMgr) GetBatchFromParentID(parentIDs []int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromFolderPath 通过folder_path获取内容 目录结构
func (obj *_EeoUserFoldersMgr) GetFromFolderPath(folderPath string) (result EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_path = ?", folderPath).Find(&result).Error

	return
}

// GetBatchFromFolderPath 批量唯一主键查找 目录结构
func (obj *_EeoUserFoldersMgr) GetBatchFromFolderPath(folderPaths []string) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_path IN (?)", folderPaths).Find(&results).Error

	return
}

// GetFromSystemLock 通过system_lock获取内容 是否可分享(根据文件夹下用户存放的文件决定的 0为可分享 1为不可分享)
func (obj *_EeoUserFoldersMgr) GetFromSystemLock(systemLock bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("system_lock = ?", systemLock).Find(&results).Error

	return
}

// GetBatchFromSystemLock 批量唯一主键查找 是否可分享(根据文件夹下用户存放的文件决定的 0为可分享 1为不可分享)
func (obj *_EeoUserFoldersMgr) GetBatchFromSystemLock(systemLocks []bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("system_lock IN (?)", systemLocks).Find(&results).Error

	return
}

// GetFromExp1 通过exp1获取内容 备用字段
func (obj *_EeoUserFoldersMgr) GetFromExp1(exp1 bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 = ?", exp1).Find(&results).Error

	return
}

// GetBatchFromExp1 批量唯一主键查找 备用字段
func (obj *_EeoUserFoldersMgr) GetBatchFromExp1(exp1s []bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 IN (?)", exp1s).Find(&results).Error

	return
}

// GetFromOperateLevel 通过operate_level获取内容 操作级别（0 自己创建可以随意删除复制移动，1系统创建不可删除、移动、复制、重命名、分享）
func (obj *_EeoUserFoldersMgr) GetFromOperateLevel(operateLevel bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("operate_level = ?", operateLevel).Find(&results).Error

	return
}

// GetBatchFromOperateLevel 批量唯一主键查找 操作级别（0 自己创建可以随意删除复制移动，1系统创建不可删除、移动、复制、重命名、分享）
func (obj *_EeoUserFoldersMgr) GetBatchFromOperateLevel(operateLevels []bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("operate_level IN (?)", operateLevels).Find(&results).Error

	return
}

// GetFromIsdel 通过isdel获取内容 是否删除(0没删除 1删除)
func (obj *_EeoUserFoldersMgr) GetFromIsdel(isdel bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel = ?", isdel).Find(&results).Error

	return
}

// GetBatchFromIsdel 批量唯一主键查找 是否删除(0没删除 1删除)
func (obj *_EeoUserFoldersMgr) GetBatchFromIsdel(isdels []bool) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel IN (?)", isdels).Find(&results).Error

	return
}

// GetFromAddtime 通过addtime获取内容 添加时间
func (obj *_EeoUserFoldersMgr) GetFromAddtime(addtime int) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&results).Error

	return
}

// GetBatchFromAddtime 批量唯一主键查找 添加时间
func (obj *_EeoUserFoldersMgr) GetBatchFromAddtime(addtimes []int) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime IN (?)", addtimes).Find(&results).Error

	return
}

// GetFromUptime 通过uptime获取内容 最后修改时间
func (obj *_EeoUserFoldersMgr) GetFromUptime(uptime int) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime = ?", uptime).Find(&results).Error

	return
}

// GetBatchFromUptime 批量唯一主键查找 最后修改时间
func (obj *_EeoUserFoldersMgr) GetBatchFromUptime(uptimes []int) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime IN (?)", uptimes).Find(&results).Error

	return
}

// GetFromFromID 通过from_id获取内容 来自文件夹id
func (obj *_EeoUserFoldersMgr) GetFromFromID(fromID int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("from_id = ?", fromID).Find(&results).Error

	return
}

// GetBatchFromFromID 批量唯一主键查找 来自文件夹id
func (obj *_EeoUserFoldersMgr) GetBatchFromFromID(fromIDs []int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("from_id IN (?)", fromIDs).Find(&results).Error

	return
}

// GetFromFromParent 通过from_parent获取内容 原来的父级id
func (obj *_EeoUserFoldersMgr) GetFromFromParent(fromParent int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("from_parent = ?", fromParent).Find(&results).Error

	return
}

// GetBatchFromFromParent 批量唯一主键查找 原来的父级id
func (obj *_EeoUserFoldersMgr) GetBatchFromFromParent(fromParents []int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("from_parent IN (?)", fromParents).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 目录类型：0=课件 1=考试 2=作业
func (obj *_EeoUserFoldersMgr) GetFromType(_type int8) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 目录类型：0=课件 1=考试 2=作业
func (obj *_EeoUserFoldersMgr) GetBatchFromType(_types []int8) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromOldID 通过old_id获取内容 迁移前目录
func (obj *_EeoUserFoldersMgr) GetFromOldID(oldID int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("old_id = ?", oldID).Find(&results).Error

	return
}

// GetBatchFromOldID 批量唯一主键查找 迁移前目录
func (obj *_EeoUserFoldersMgr) GetBatchFromOldID(oldIDs []int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("old_id IN (?)", oldIDs).Find(&results).Error

	return
}

// GetFromBatchOperateTime 通过batch_operate_time获取内容 批量复制微秒时间戳
func (obj *_EeoUserFoldersMgr) GetFromBatchOperateTime(batchOperateTime int64) (result EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("batch_operate_time = ?", batchOperateTime).Find(&result).Error

	return
}

// GetBatchFromBatchOperateTime 批量唯一主键查找 批量复制微秒时间戳
func (obj *_EeoUserFoldersMgr) GetBatchFromBatchOperateTime(batchOperateTimes []int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("batch_operate_time IN (?)", batchOperateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoUserFoldersMgr) FetchByPrimaryKey(folderID int64) (result EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id = ?", folderID).Find(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_EeoUserFoldersMgr) FetchIndexByUserID(userID int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// FetchIndexByParenterID  获取多个内容
func (obj *_EeoUserFoldersMgr) FetchIndexByParenterID(parentID int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error

	return
}

// FetchIndexByFolderPath  获取多个内容
func (obj *_EeoUserFoldersMgr) FetchIndexByFolderPath(folderPath string) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_path = ?", folderPath).Find(&results).Error

	return
}

// FetchIndexByIBatchOperateTime  获取多个内容
func (obj *_EeoUserFoldersMgr) FetchIndexByIBatchOperateTime(batchOperateTime int64) (results []*EeoUserFolders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("batch_operate_time = ?", batchOperateTime).Find(&results).Error

	return
}
