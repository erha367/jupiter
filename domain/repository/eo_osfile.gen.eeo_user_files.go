package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"jupiter/domain/entity"
)

type _EeoUserFilesMgr struct {
	*_BaseMgr
}

// EeoUserFilesMgr open func
func EeoUserFilesMgr(db *gorm.DB) *_EeoUserFilesMgr {
	if db == nil {
		panic(fmt.Errorf("EeoUserFilesMgr need init by db"))
	}
	return &_EeoUserFilesMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoUserFilesMgr) GetTableName() string {
	return "eeo_user_files"
}

// Get 获取
func (obj *_EeoUserFilesMgr) Get() (result entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EeoUserFilesMgr) Gets() (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_EeoUserFilesMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_EeoUserFilesMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithFileID file_id获取 文件id
func (obj *_EeoUserFilesMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithFolderID folder_id获取 所属文件夹id
func (obj *_EeoUserFilesMgr) WithFolderID(folderID int64) Option {
	return optionFunc(func(o *options) { o.query["folder_id"] = folderID })
}

// WithFileName file_name获取 文件名称
func (obj *_EeoUserFilesMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

// WithFileSize file_size获取 文件大小(单位是KB)
func (obj *_EeoUserFilesMgr) WithFileSize(fileSize int) Option {
	return optionFunc(func(o *options) { o.query["file_size"] = fileSize })
}

// WithFileExtension file_extension获取 文件格式
func (obj *_EeoUserFilesMgr) WithFileExtension(fileExtension string) Option {
	return optionFunc(func(o *options) { o.query["file_extension"] = fileExtension })
}

// WithTransitionState transition_state获取 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的）
func (obj *_EeoUserFilesMgr) WithTransitionState(transitionState bool) Option {
	return optionFunc(func(o *options) { o.query["transition_state"] = transitionState })
}

// WithFileNamePy file_name_py获取 名字拼音
func (obj *_EeoUserFilesMgr) WithFileNamePy(fileNamePy string) Option {
	return optionFunc(func(o *options) { o.query["file_name_py"] = fileNamePy })
}

// WithIsdel isdel获取 是否删除(在我的网盘里不显示)
func (obj *_EeoUserFilesMgr) WithIsdel(isdel bool) Option {
	return optionFunc(func(o *options) { o.query["isdel"] = isdel })
}

// WithExp1 exp1获取 扩展字段
func (obj *_EeoUserFilesMgr) WithExp1(exp1 bool) Option {
	return optionFunc(func(o *options) { o.query["exp1"] = exp1 })
}

// WithAddtime addtime获取 存入我的网盘时间
func (obj *_EeoUserFilesMgr) WithAddtime(addtime int) Option {
	return optionFunc(func(o *options) { o.query["addtime"] = addtime })
}

// WithUptime uptime获取 最后修改时间
func (obj *_EeoUserFilesMgr) WithUptime(uptime int) Option {
	return optionFunc(func(o *options) { o.query["uptime"] = uptime })
}

// WithAllowOperate allow_operate获取 授权文件是否允许操作
func (obj *_EeoUserFilesMgr) WithAllowOperate(allowOperate uint8) Option {
	return optionFunc(func(o *options) { o.query["allow_operate"] = allowOperate })
}

// WithExamTopicSum exam_topic_sum获取 试卷题数
//
func (obj *_EeoUserFilesMgr) WithExamTopicSum(examTopicSum int) Option {
	return optionFunc(func(o *options) { o.query["exam_topic_sum"] = examTopicSum })
}

// WithFullMarks full_marks获取 试卷满分
//
func (obj *_EeoUserFilesMgr) WithFullMarks(fullMarks int) Option {
	return optionFunc(func(o *options) { o.query["full_marks"] = fullMarks })
}

// WithVersion version获取 版本号
func (obj *_EeoUserFilesMgr) WithVersion(version int8) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithType type获取 数据类型 0=课件 1=考试 2=作业
func (obj *_EeoUserFilesMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithOldID old_id获取 迁移前的主键id
func (obj *_EeoUserFilesMgr) WithOldID(oldID int64) Option {
	return optionFunc(func(o *options) { o.query["old_id"] = oldID })
}

// WithEdocAuth edoc_auth获取 在线文档权限 0未设置 1只读 2写
func (obj *_EeoUserFilesMgr) WithEdocAuth(edocAuth int8) Option {
	return optionFunc(func(o *options) { o.query["edoc_auth"] = edocAuth })
}


// GetByOption 功能选项模式获取
func (obj *_EeoUserFilesMgr) GetByOption(opts ...Option) (result entity.EeoUserFiles, err error) {
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
func (obj *_EeoUserFilesMgr) GetByOptions(opts ...Option) (results []*entity.EeoUserFiles, err error) {
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
func (obj *_EeoUserFilesMgr)  GetFromID(id int64) (result entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 ID
func (obj *_EeoUserFilesMgr) GetBatchFromID(ids []int64) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromUserID 通过user_id获取内容 用户id 
func (obj *_EeoUserFilesMgr)  GetFromUserID(userID int64) (result entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error
	
	return
}

// GetBatchFromUserID 批量唯一主键查找 用户id
func (obj *_EeoUserFilesMgr) GetBatchFromUserID(userIDs []int64) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
	
	return
}
 
// GetFromFileID 通过file_id获取内容 文件id 
func (obj *_EeoUserFilesMgr) GetFromFileID(fileID int64) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&results).Error
	
	return
}

// GetBatchFromFileID 批量唯一主键查找 文件id
func (obj *_EeoUserFilesMgr) GetBatchFromFileID(fileIDs []int64) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id IN (?)", fileIDs).Find(&results).Error
	
	return
}
 
// GetFromFolderID 通过folder_id获取内容 所属文件夹id 
func (obj *_EeoUserFilesMgr)  GetFromFolderID(folderID int64) (result entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id = ?", folderID).Find(&result).Error
	
	return
}

// GetBatchFromFolderID 批量唯一主键查找 所属文件夹id
func (obj *_EeoUserFilesMgr) GetBatchFromFolderID(folderIDs []int64) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id IN (?)", folderIDs).Find(&results).Error
	
	return
}
 
// GetFromFileName 通过file_name获取内容 文件名称 
func (obj *_EeoUserFilesMgr) GetFromFileName(fileName string) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error
	
	return
}

// GetBatchFromFileName 批量唯一主键查找 文件名称
func (obj *_EeoUserFilesMgr) GetBatchFromFileName(fileNames []string) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error
	
	return
}
 
// GetFromFileSize 通过file_size获取内容 文件大小(单位是KB) 
func (obj *_EeoUserFilesMgr) GetFromFileSize(fileSize int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size = ?", fileSize).Find(&results).Error
	
	return
}

// GetBatchFromFileSize 批量唯一主键查找 文件大小(单位是KB)
func (obj *_EeoUserFilesMgr) GetBatchFromFileSize(fileSizes []int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size IN (?)", fileSizes).Find(&results).Error
	
	return
}
 
// GetFromFileExtension 通过file_extension获取内容 文件格式 
func (obj *_EeoUserFilesMgr) GetFromFileExtension(fileExtension string) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_extension = ?", fileExtension).Find(&results).Error
	
	return
}

// GetBatchFromFileExtension 批量唯一主键查找 文件格式
func (obj *_EeoUserFilesMgr) GetBatchFromFileExtension(fileExtensions []string) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_extension IN (?)", fileExtensions).Find(&results).Error
	
	return
}
 
// GetFromTransitionState 通过transition_state获取内容 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的） 
func (obj *_EeoUserFilesMgr) GetFromTransitionState(transitionState bool) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_state = ?", transitionState).Find(&results).Error
	
	return
}

// GetBatchFromTransitionState 批量唯一主键查找 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的）
func (obj *_EeoUserFilesMgr) GetBatchFromTransitionState(transitionStates []bool) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_state IN (?)", transitionStates).Find(&results).Error
	
	return
}
 
// GetFromFileNamePy 通过file_name_py获取内容 名字拼音 
func (obj *_EeoUserFilesMgr) GetFromFileNamePy(fileNamePy string) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name_py = ?", fileNamePy).Find(&results).Error
	
	return
}

// GetBatchFromFileNamePy 批量唯一主键查找 名字拼音
func (obj *_EeoUserFilesMgr) GetBatchFromFileNamePy(fileNamePys []string) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_name_py IN (?)", fileNamePys).Find(&results).Error
	
	return
}
 
// GetFromIsdel 通过isdel获取内容 是否删除(在我的网盘里不显示) 
func (obj *_EeoUserFilesMgr) GetFromIsdel(isdel bool) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel = ?", isdel).Find(&results).Error
	
	return
}

// GetBatchFromIsdel 批量唯一主键查找 是否删除(在我的网盘里不显示)
func (obj *_EeoUserFilesMgr) GetBatchFromIsdel(isdels []bool) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel IN (?)", isdels).Find(&results).Error
	
	return
}
 
// GetFromExp1 通过exp1获取内容 扩展字段 
func (obj *_EeoUserFilesMgr) GetFromExp1(exp1 bool) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 = ?", exp1).Find(&results).Error
	
	return
}

// GetBatchFromExp1 批量唯一主键查找 扩展字段
func (obj *_EeoUserFilesMgr) GetBatchFromExp1(exp1s []bool) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 IN (?)", exp1s).Find(&results).Error
	
	return
}
 
// GetFromAddtime 通过addtime获取内容 存入我的网盘时间 
func (obj *_EeoUserFilesMgr) GetFromAddtime(addtime int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&results).Error
	
	return
}

// GetBatchFromAddtime 批量唯一主键查找 存入我的网盘时间
func (obj *_EeoUserFilesMgr) GetBatchFromAddtime(addtimes []int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime IN (?)", addtimes).Find(&results).Error
	
	return
}
 
// GetFromUptime 通过uptime获取内容 最后修改时间 
func (obj *_EeoUserFilesMgr) GetFromUptime(uptime int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime = ?", uptime).Find(&results).Error
	
	return
}

// GetBatchFromUptime 批量唯一主键查找 最后修改时间
func (obj *_EeoUserFilesMgr) GetBatchFromUptime(uptimes []int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime IN (?)", uptimes).Find(&results).Error
	
	return
}
 
// GetFromAllowOperate 通过allow_operate获取内容 授权文件是否允许操作 
func (obj *_EeoUserFilesMgr) GetFromAllowOperate(allowOperate uint8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("allow_operate = ?", allowOperate).Find(&results).Error
	
	return
}

// GetBatchFromAllowOperate 批量唯一主键查找 授权文件是否允许操作
func (obj *_EeoUserFilesMgr) GetBatchFromAllowOperate(allowOperates []uint8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("allow_operate IN (?)", allowOperates).Find(&results).Error
	
	return
}
 
// GetFromExamTopicSum 通过exam_topic_sum获取内容 试卷题数
// 
func (obj *_EeoUserFilesMgr) GetFromExamTopicSum(examTopicSum int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exam_topic_sum = ?", examTopicSum).Find(&results).Error
	
	return
}

// GetBatchFromExamTopicSum 批量唯一主键查找 试卷题数
//
func (obj *_EeoUserFilesMgr) GetBatchFromExamTopicSum(examTopicSums []int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exam_topic_sum IN (?)", examTopicSums).Find(&results).Error
	
	return
}
 
// GetFromFullMarks 通过full_marks获取内容 试卷满分
// 
func (obj *_EeoUserFilesMgr) GetFromFullMarks(fullMarks int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("full_marks = ?", fullMarks).Find(&results).Error
	
	return
}

// GetBatchFromFullMarks 批量唯一主键查找 试卷满分
//
func (obj *_EeoUserFilesMgr) GetBatchFromFullMarks(fullMarkss []int) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("full_marks IN (?)", fullMarkss).Find(&results).Error
	
	return
}
 
// GetFromVersion 通过version获取内容 版本号 
func (obj *_EeoUserFilesMgr) GetFromVersion(version int8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("version = ?", version).Find(&results).Error
	
	return
}

// GetBatchFromVersion 批量唯一主键查找 版本号
func (obj *_EeoUserFilesMgr) GetBatchFromVersion(versions []int8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("version IN (?)", versions).Find(&results).Error
	
	return
}
 
// GetFromType 通过type获取内容 数据类型 0=课件 1=考试 2=作业 
func (obj *_EeoUserFilesMgr) GetFromType(_type int8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量唯一主键查找 数据类型 0=课件 1=考试 2=作业
func (obj *_EeoUserFilesMgr) GetBatchFromType(_types []int8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromOldID 通过old_id获取内容 迁移前的主键id 
func (obj *_EeoUserFilesMgr) GetFromOldID(oldID int64) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("old_id = ?", oldID).Find(&results).Error
	
	return
}

// GetBatchFromOldID 批量唯一主键查找 迁移前的主键id
func (obj *_EeoUserFilesMgr) GetBatchFromOldID(oldIDs []int64) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("old_id IN (?)", oldIDs).Find(&results).Error
	
	return
}
 
// GetFromEdocAuth 通过edoc_auth获取内容 在线文档权限 0未设置 1只读 2写 
func (obj *_EeoUserFilesMgr) GetFromEdocAuth(edocAuth int8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("edoc_auth = ?", edocAuth).Find(&results).Error
	
	return
}

// GetBatchFromEdocAuth 批量唯一主键查找 在线文档权限 0未设置 1只读 2写
func (obj *_EeoUserFilesMgr) GetBatchFromEdocAuth(edocAuths []int8) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("edoc_auth IN (?)", edocAuths).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EeoUserFilesMgr) FetchByPrimaryKey(id int64 ) (result entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 


 
 // FetchIndexByFolderID  获取多个内容
 func (obj *_EeoUserFilesMgr) FetchIndexByFolderID(folderID int64 ) (results []*entity.EeoUserFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("folder_id = ?", folderID).Find(&results).Error
	
	return
}
 

	

