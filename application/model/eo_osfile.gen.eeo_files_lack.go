package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _EeoFilesLackMgr struct {
	*_BaseMgr
}

// EeoFilesLackMgr open func
func EeoFilesLackMgr(db *gorm.DB) *_EeoFilesLackMgr {
	if db == nil {
		panic(fmt.Errorf("EeoFilesLackMgr need init by db"))
	}
	return &_EeoFilesLackMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoFilesLackMgr) GetTableName() string {
	return "eeo_files_lack"
}

// Get 获取
func (obj *_EeoFilesLackMgr) Get() (result EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoFilesLackMgr) Gets() (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_EeoFilesLackMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFileID file_id获取 文件ID
func (obj *_EeoFilesLackMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithFileSavePath file_save_path获取 文件保存路径
func (obj *_EeoFilesLackMgr) WithFileSavePath(fileSavePath string) Option {
	return optionFunc(func(o *options) { o.query["file_save_path"] = fileSavePath })
}

// WithAddtime addtime获取 添加时间
func (obj *_EeoFilesLackMgr) WithAddtime(addtime int) Option {
	return optionFunc(func(o *options) { o.query["addtime"] = addtime })
}

// GetByOption 功能选项模式获取
func (obj *_EeoFilesLackMgr) GetByOption(opts ...Option) (result EeoFilesLack, err error) {
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
func (obj *_EeoFilesLackMgr) GetByOptions(opts ...Option) (results []*EeoFilesLack, err error) {
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

// GetFromID 通过id获取内容 ID
func (obj *_EeoFilesLackMgr) GetFromID(id int64) (result EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID
func (obj *_EeoFilesLackMgr) GetBatchFromID(ids []int64) (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromFileID 通过file_id获取内容 文件ID
func (obj *_EeoFilesLackMgr) GetFromFileID(fileID int64) (result EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&result).Error

	return
}

// GetBatchFromFileID 批量唯一主键查找 文件ID
func (obj *_EeoFilesLackMgr) GetBatchFromFileID(fileIDs []int64) (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromFileSavePath 通过file_save_path获取内容 文件保存路径
func (obj *_EeoFilesLackMgr) GetFromFileSavePath(fileSavePath string) (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_save_path = ?", fileSavePath).Find(&results).Error

	return
}

// GetBatchFromFileSavePath 批量唯一主键查找 文件保存路径
func (obj *_EeoFilesLackMgr) GetBatchFromFileSavePath(fileSavePaths []string) (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_save_path IN (?)", fileSavePaths).Find(&results).Error

	return
}

// GetFromAddtime 通过addtime获取内容 添加时间
func (obj *_EeoFilesLackMgr) GetFromAddtime(addtime int) (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&results).Error

	return
}

// GetBatchFromAddtime 批量唯一主键查找 添加时间
func (obj *_EeoFilesLackMgr) GetBatchFromAddtime(addtimes []int) (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime IN (?)", addtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoFilesLackMgr) FetchByPrimaryKey(id int64) (result EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByFileID  获取多个内容
func (obj *_EeoFilesLackMgr) FetchIndexByFileID(fileID int64) (results []*EeoFilesLack, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&results).Error

	return
}
