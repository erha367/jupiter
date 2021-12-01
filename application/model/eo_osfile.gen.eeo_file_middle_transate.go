package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _EeoFileMiddleTransateMgr struct {
	*_BaseMgr
}

// EeoFileMiddleTransateMgr open func
func EeoFileMiddleTransateMgr(db *gorm.DB) *_EeoFileMiddleTransateMgr {
	if db == nil {
		panic(fmt.Errorf("EeoFileMiddleTransateMgr need init by db"))
	}
	return &_EeoFileMiddleTransateMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoFileMiddleTransateMgr) GetTableName() string {
	return "eeo_file_middle_transate"
}

// Get 获取
func (obj *_EeoFileMiddleTransateMgr) Get() (result EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoFileMiddleTransateMgr) Gets() (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_EeoFileMiddleTransateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUId uid获取 上传者id
func (obj *_EeoFileMiddleTransateMgr) WithUId(uid int64) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithType type获取 文件格式
func (obj *_EeoFileMiddleTransateMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCount count获取 待转换文件个数
func (obj *_EeoFileMiddleTransateMgr) WithCount(count int) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// GetByOption 功能选项模式获取
func (obj *_EeoFileMiddleTransateMgr) GetByOption(opts ...Option) (result EeoFileMiddleTransate, err error) {
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
func (obj *_EeoFileMiddleTransateMgr) GetByOptions(opts ...Option) (results []*EeoFileMiddleTransate, err error) {
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
func (obj *_EeoFileMiddleTransateMgr) GetFromID(id int) (result EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID
func (obj *_EeoFileMiddleTransateMgr) GetBatchFromID(ids []int) (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUId 通过uid获取内容 上传者id
func (obj *_EeoFileMiddleTransateMgr) GetFromUId(uid int64) (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uid = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUId 批量唯一主键查找 上传者id
func (obj *_EeoFileMiddleTransateMgr) GetBatchFromUId(uids []int64) (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uid IN (?)", uids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 文件格式
func (obj *_EeoFileMiddleTransateMgr) GetFromType(_type string) (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 文件格式
func (obj *_EeoFileMiddleTransateMgr) GetBatchFromType(_types []string) (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容 待转换文件个数
func (obj *_EeoFileMiddleTransateMgr) GetFromCount(count int) (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("count = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量唯一主键查找 待转换文件个数
func (obj *_EeoFileMiddleTransateMgr) GetBatchFromCount(counts []int) (results []*EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("count IN (?)", counts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoFileMiddleTransateMgr) FetchByPrimaryKey(id int) (result EeoFileMiddleTransate, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
