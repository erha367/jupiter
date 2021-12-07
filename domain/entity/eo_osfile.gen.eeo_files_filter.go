package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _EeoFilesFilterMgr struct {
	*_BaseMgr
}

// EeoFilesFilterMgr open func
func EeoFilesFilterMgr(db *gorm.DB) *_EeoFilesFilterMgr {
	if db == nil {
		panic(fmt.Errorf("EeoFilesFilterMgr need init by db"))
	}
	return &_EeoFilesFilterMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoFilesFilterMgr) GetTableName() string {
	return "eeo_files_filter"
}

// Get 获取
func (obj *_EeoFilesFilterMgr) Get() (result EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoFilesFilterMgr) Gets() (results []*EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键,自增id
func (obj *_EeoFilesFilterMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFileID file_id获取 文件id
func (obj *_EeoFilesFilterMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithFileMd5 file_md5获取 文件md5
func (obj *_EeoFilesFilterMgr) WithFileMd5(fileMd5 string) Option {
	return optionFunc(func(o *options) { o.query["file_md5"] = fileMd5 })
}

// WithFileSize file_size获取 文件大小，单位为b
func (obj *_EeoFilesFilterMgr) WithFileSize(fileSize int64) Option {
	return optionFunc(func(o *options) { o.query["file_size"] = fileSize })
}

// GetByOption 功能选项模式获取
func (obj *_EeoFilesFilterMgr) GetByOption(opts ...Option) (result EeoFilesFilter, err error) {
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
func (obj *_EeoFilesFilterMgr) GetByOptions(opts ...Option) (results []*EeoFilesFilter, err error) {
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

// GetFromID 通过id获取内容 主键,自增id
func (obj *_EeoFilesFilterMgr) GetFromID(id int64) (result EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键,自增id
func (obj *_EeoFilesFilterMgr) GetBatchFromID(ids []int64) (results []*EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromFileID 通过file_id获取内容 文件id
func (obj *_EeoFilesFilterMgr) GetFromFileID(fileID int64) (results []*EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&results).Error

	return
}

// GetBatchFromFileID 批量唯一主键查找 文件id
func (obj *_EeoFilesFilterMgr) GetBatchFromFileID(fileIDs []int64) (results []*EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromFileMd5 通过file_md5获取内容 文件md5
func (obj *_EeoFilesFilterMgr) GetFromFileMd5(fileMd5 string) (result EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_md5 = ?", fileMd5).Find(&result).Error

	return
}

// GetBatchFromFileMd5 批量唯一主键查找 文件md5
func (obj *_EeoFilesFilterMgr) GetBatchFromFileMd5(fileMd5s []string) (results []*EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_md5 IN (?)", fileMd5s).Find(&results).Error

	return
}

// GetFromFileSize 通过file_size获取内容 文件大小，单位为b
func (obj *_EeoFilesFilterMgr) GetFromFileSize(fileSize int64) (results []*EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size = ?", fileSize).Find(&results).Error

	return
}

// GetBatchFromFileSize 批量唯一主键查找 文件大小，单位为b
func (obj *_EeoFilesFilterMgr) GetBatchFromFileSize(fileSizes []int64) (results []*EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size IN (?)", fileSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoFilesFilterMgr) FetchByPrimaryKey(id int64) (result EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByIFileMd5 primay or index 获取唯一内容
func (obj *_EeoFilesFilterMgr) FetchUniqueByIFileMd5(fileMd5 string) (result EeoFilesFilter, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_md5 = ?", fileMd5).Find(&result).Error

	return
}
