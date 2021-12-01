package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _EeoCourseFilesMgr struct {
	*_BaseMgr
}

// EeoCourseFilesMgr open func
func EeoCourseFilesMgr(db *gorm.DB) *_EeoCourseFilesMgr {
	if db == nil {
		panic(fmt.Errorf("EeoCourseFilesMgr need init by db"))
	}
	return &_EeoCourseFilesMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoCourseFilesMgr) GetTableName() string {
	return "eeo_course_files"
}

// Get 获取
func (obj *_EeoCourseFilesMgr) Get() (result EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoCourseFilesMgr) Gets() (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_EeoCourseFilesMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSid sid获取 机构id
func (obj *_EeoCourseFilesMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["sid"] = sid })
}

// WithCourseID course_id获取 群id
func (obj *_EeoCourseFilesMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithUserID user_id获取 上传用户id
func (obj *_EeoCourseFilesMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithFileID file_id获取 文件id
func (obj *_EeoCourseFilesMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithFolderID folder_id获取 所属文件夹id
func (obj *_EeoCourseFilesMgr) WithFolderID(folderID int64) Option {
	return optionFunc(func(o *options) { o.query["folder_id"] = folderID })
}

// WithFileName file_name获取 文件名称
func (obj *_EeoCourseFilesMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

// WithUserName user_name获取 上传人名称
func (obj *_EeoCourseFilesMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithFileSize file_size获取 文件大小(单位是KB)
func (obj *_EeoCourseFilesMgr) WithFileSize(fileSize int) Option {
	return optionFunc(func(o *options) { o.query["file_size"] = fileSize })
}

// WithFileExtension file_extension获取 文件格式
func (obj *_EeoCourseFilesMgr) WithFileExtension(fileExtension string) Option {
	return optionFunc(func(o *options) { o.query["file_extension"] = fileExtension })
}

// WithFileNamePy file_name_py获取 名字拼音
func (obj *_EeoCourseFilesMgr) WithFileNamePy(fileNamePy string) Option {
	return optionFunc(func(o *options) { o.query["file_name_py"] = fileNamePy })
}

// WithIsdel isdel获取 是否删除(在我的网盘里不显示)
func (obj *_EeoCourseFilesMgr) WithIsdel(isdel bool) Option {
	return optionFunc(func(o *options) { o.query["isdel"] = isdel })
}

// WithExp1 exp1获取 扩展字段
func (obj *_EeoCourseFilesMgr) WithExp1(exp1 int8) Option {
	return optionFunc(func(o *options) { o.query["exp1"] = exp1 })
}

// WithAddtime addtime获取 创建时间
func (obj *_EeoCourseFilesMgr) WithAddtime(addtime int) Option {
	return optionFunc(func(o *options) { o.query["addtime"] = addtime })
}

// WithUptime uptime获取 修改时间
func (obj *_EeoCourseFilesMgr) WithUptime(uptime int) Option {
	return optionFunc(func(o *options) { o.query["uptime"] = uptime })
}

// WithType type获取 文件类型 0课件1试卷 2作业
func (obj *_EeoCourseFilesMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_EeoCourseFilesMgr) GetByOption(opts ...Option) (result EeoCourseFiles, err error) {
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
func (obj *_EeoCourseFilesMgr) GetByOptions(opts ...Option) (results []*EeoCourseFiles, err error) {
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
func (obj *_EeoCourseFilesMgr) GetFromID(id int64) (result EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID
func (obj *_EeoCourseFilesMgr) GetBatchFromID(ids []int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromSid 通过sid获取内容 机构id
func (obj *_EeoCourseFilesMgr) GetFromSid(sid int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sid = ?", sid).Find(&results).Error

	return
}

// GetBatchFromSid 批量唯一主键查找 机构id
func (obj *_EeoCourseFilesMgr) GetBatchFromSid(sids []int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sid IN (?)", sids).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 群id
func (obj *_EeoCourseFilesMgr) GetFromCourseID(courseID int64) (result EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id = ?", courseID).Find(&result).Error

	return
}

// GetBatchFromCourseID 批量唯一主键查找 群id
func (obj *_EeoCourseFilesMgr) GetBatchFromCourseID(courseIDs []int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 上传用户id
func (obj *_EeoCourseFilesMgr) GetFromUserID(userID int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 上传用户id
func (obj *_EeoCourseFilesMgr) GetBatchFromUserID(userIDs []int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromFileID 通过file_id获取内容 文件id
func (obj *_EeoCourseFilesMgr) GetFromFileID(fileID int64) (result EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&result).Error

	return
}

// GetBatchFromFileID 批量唯一主键查找 文件id
func (obj *_EeoCourseFilesMgr) GetBatchFromFileID(fileIDs []int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromFolderID 通过folder_id获取内容 所属文件夹id
func (obj *_EeoCourseFilesMgr) GetFromFolderID(folderID int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id = ?", folderID).Find(&results).Error

	return
}

// GetBatchFromFolderID 批量唯一主键查找 所属文件夹id
func (obj *_EeoCourseFilesMgr) GetBatchFromFolderID(folderIDs []int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id IN (?)", folderIDs).Find(&results).Error

	return
}

// GetFromFileName 通过file_name获取内容 文件名称
func (obj *_EeoCourseFilesMgr) GetFromFileName(fileName string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error

	return
}

// GetBatchFromFileName 批量唯一主键查找 文件名称
func (obj *_EeoCourseFilesMgr) GetBatchFromFileName(fileNames []string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容 上传人名称
func (obj *_EeoCourseFilesMgr) GetFromUserName(userName string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_name = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量唯一主键查找 上传人名称
func (obj *_EeoCourseFilesMgr) GetBatchFromUserName(userNames []string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_name IN (?)", userNames).Find(&results).Error

	return
}

// GetFromFileSize 通过file_size获取内容 文件大小(单位是KB)
func (obj *_EeoCourseFilesMgr) GetFromFileSize(fileSize int) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size = ?", fileSize).Find(&results).Error

	return
}

// GetBatchFromFileSize 批量唯一主键查找 文件大小(单位是KB)
func (obj *_EeoCourseFilesMgr) GetBatchFromFileSize(fileSizes []int) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size IN (?)", fileSizes).Find(&results).Error

	return
}

// GetFromFileExtension 通过file_extension获取内容 文件格式
func (obj *_EeoCourseFilesMgr) GetFromFileExtension(fileExtension string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_extension = ?", fileExtension).Find(&results).Error

	return
}

// GetBatchFromFileExtension 批量唯一主键查找 文件格式
func (obj *_EeoCourseFilesMgr) GetBatchFromFileExtension(fileExtensions []string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_extension IN (?)", fileExtensions).Find(&results).Error

	return
}

// GetFromFileNamePy 通过file_name_py获取内容 名字拼音
func (obj *_EeoCourseFilesMgr) GetFromFileNamePy(fileNamePy string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name_py = ?", fileNamePy).Find(&results).Error

	return
}

// GetBatchFromFileNamePy 批量唯一主键查找 名字拼音
func (obj *_EeoCourseFilesMgr) GetBatchFromFileNamePy(fileNamePys []string) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name_py IN (?)", fileNamePys).Find(&results).Error

	return
}

// GetFromIsdel 通过isdel获取内容 是否删除(在我的网盘里不显示)
func (obj *_EeoCourseFilesMgr) GetFromIsdel(isdel bool) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel = ?", isdel).Find(&results).Error

	return
}

// GetBatchFromIsdel 批量唯一主键查找 是否删除(在我的网盘里不显示)
func (obj *_EeoCourseFilesMgr) GetBatchFromIsdel(isdels []bool) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel IN (?)", isdels).Find(&results).Error

	return
}

// GetFromExp1 通过exp1获取内容 扩展字段
func (obj *_EeoCourseFilesMgr) GetFromExp1(exp1 int8) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 = ?", exp1).Find(&results).Error

	return
}

// GetBatchFromExp1 批量唯一主键查找 扩展字段
func (obj *_EeoCourseFilesMgr) GetBatchFromExp1(exp1s []int8) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 IN (?)", exp1s).Find(&results).Error

	return
}

// GetFromAddtime 通过addtime获取内容 创建时间
func (obj *_EeoCourseFilesMgr) GetFromAddtime(addtime int) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&results).Error

	return
}

// GetBatchFromAddtime 批量唯一主键查找 创建时间
func (obj *_EeoCourseFilesMgr) GetBatchFromAddtime(addtimes []int) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime IN (?)", addtimes).Find(&results).Error

	return
}

// GetFromUptime 通过uptime获取内容 修改时间
func (obj *_EeoCourseFilesMgr) GetFromUptime(uptime int) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime = ?", uptime).Find(&results).Error

	return
}

// GetBatchFromUptime 批量唯一主键查找 修改时间
func (obj *_EeoCourseFilesMgr) GetBatchFromUptime(uptimes []int) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime IN (?)", uptimes).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 文件类型 0课件1试卷 2作业
func (obj *_EeoCourseFilesMgr) GetFromType(_type int8) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 文件类型 0课件1试卷 2作业
func (obj *_EeoCourseFilesMgr) GetBatchFromType(_types []int8) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoCourseFilesMgr) FetchByPrimaryKey(id int64) (result EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByICourseID  获取多个内容
func (obj *_EeoCourseFilesMgr) FetchIndexByICourseID(courseID int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id = ?", courseID).Find(&results).Error

	return
}

// FetchIndexByIFileID  获取多个内容
func (obj *_EeoCourseFilesMgr) FetchIndexByIFileID(fileID int64) (results []*EeoCourseFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&results).Error

	return
}
