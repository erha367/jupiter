package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"jupiter/domain/entity"
)

type _EeoFilesMgr struct {
	*_BaseMgr
}

// EeoFilesMgr open func
func EeoFilesMgr(db *gorm.DB) *_EeoFilesMgr {
	if db == nil {
		panic(fmt.Errorf("EeoFilesMgr need init by db"))
	}
	return &_EeoFilesMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoFilesMgr) GetTableName() string {
	return "eeo_files"
}

// Get 获取
func (obj *_EeoFilesMgr) Get() (result entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EeoFilesMgr) Gets() (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithFileID file_id获取 文件ID
func (obj *_EeoFilesMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithFileSavePath file_save_path获取 文件保存路径
func (obj *_EeoFilesMgr) WithFileSavePath(fileSavePath string) Option {
	return optionFunc(func(o *options) { o.query["file_save_path"] = fileSavePath })
}

// WithFileSaveName file_save_name获取 文件保存名字
func (obj *_EeoFilesMgr) WithFileSaveName(fileSaveName string) Option {
	return optionFunc(func(o *options) { o.query["file_save_name"] = fileSaveName })
}

// WithFileExtension file_extension获取 文件格式
func (obj *_EeoFilesMgr) WithFileExtension(fileExtension string) Option {
	return optionFunc(func(o *options) { o.query["file_extension"] = fileExtension })
}

// WithFileSize file_size获取 文件大小(单位是KB)
func (obj *_EeoFilesMgr) WithFileSize(fileSize int) Option {
	return optionFunc(func(o *options) { o.query["file_size"] = fileSize })
}

// WithTransitionFiles transition_files获取 转换后保存文件地址
func (obj *_EeoFilesMgr) WithTransitionFiles(transitionFiles string) Option {
	return optionFunc(func(o *options) { o.query["transition_files"] = transitionFiles })
}

// WithTransitionData transition_data获取 json格式数据
func (obj *_EeoFilesMgr) WithTransitionData(transitionData string) Option {
	return optionFunc(func(o *options) { o.query["transition_data"] = transitionData })
}

// WithTransitionState transition_state获取 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的）
func (obj *_EeoFilesMgr) WithTransitionState(transitionState bool) Option {
	return optionFunc(func(o *options) { o.query["transition_state"] = transitionState })
}

// WithModerationState moderation_state获取 内容审核状态（0.正常, 1.违规）
func (obj *_EeoFilesMgr) WithModerationState(moderationState int8) Option {
	return optionFunc(func(o *options) { o.query["moderation_state"] = moderationState })
}

// WithUploadUserid upload_userid获取 上传者id
func (obj *_EeoFilesMgr) WithUploadUserid(uploadUserid int64) Option {
	return optionFunc(func(o *options) { o.query["upload_userid"] = uploadUserid })
}

// WithFileType file_type获取 文件状态(0为个人上传，1为平台上传，2为第三方资料,3作业上传 ,4 im聊天图片)
func (obj *_EeoFilesMgr) WithFileType(fileType bool) Option {
	return optionFunc(func(o *options) { o.query["file_type"] = fileType })
}

// WithSystemDel system_del获取 系统删除(0为未删除 1为已删除)
func (obj *_EeoFilesMgr) WithSystemDel(systemDel bool) Option {
	return optionFunc(func(o *options) { o.query["system_del"] = systemDel })
}

// WithSystemLock system_lock获取 系统屏蔽(0为未屏蔽可以分享 1为屏蔽禁止分享)
func (obj *_EeoFilesMgr) WithSystemLock(systemLock bool) Option {
	return optionFunc(func(o *options) { o.query["system_lock"] = systemLock })
}

// WithSaveNum save_num获取 被保存次数
func (obj *_EeoFilesMgr) WithSaveNum(saveNum int) Option {
	return optionFunc(func(o *options) { o.query["save_num"] = saveNum })
}

// WithDownNum down_num获取 被下载次数
func (obj *_EeoFilesMgr) WithDownNum(downNum int) Option {
	return optionFunc(func(o *options) { o.query["down_num"] = downNum })
}

// WithExp1 exp1获取 扩展字段
func (obj *_EeoFilesMgr) WithExp1(exp1 bool) Option {
	return optionFunc(func(o *options) { o.query["exp1"] = exp1 })
}

// WithMigrateStatus migrate_status获取 文件迁移cos状态(0为未迁移 1为已迁移 2迁移失败)
func (obj *_EeoFilesMgr) WithMigrateStatus(migrateStatus bool) Option {
	return optionFunc(func(o *options) { o.query["migrate_status"] = migrateStatus })
}

// WithMigrateTime migrate_time获取 文件迁移cos时间
func (obj *_EeoFilesMgr) WithMigrateTime(migrateTime int) Option {
	return optionFunc(func(o *options) { o.query["migrate_time"] = migrateTime })
}

// WithAddtime addtime获取 上传时间
func (obj *_EeoFilesMgr) WithAddtime(addtime int) Option {
	return optionFunc(func(o *options) { o.query["addtime"] = addtime })
}

// WithUptime uptime获取 状态更新时间
func (obj *_EeoFilesMgr) WithUptime(uptime int) Option {
	return optionFunc(func(o *options) { o.query["uptime"] = uptime })
}

// WithTransitionType transition_type获取 
func (obj *_EeoFilesMgr) WithTransitionType(transitionType string) Option {
	return optionFunc(func(o *options) { o.query["transition_type"] = transitionType })
}

// WithTransitionStartTime transition_start_time获取 转换开始时间
func (obj *_EeoFilesMgr) WithTransitionStartTime(transitionStartTime int) Option {
	return optionFunc(func(o *options) { o.query["transition_start_time"] = transitionStartTime })
}

// WithTransitionLongs transition_longs获取 转换时长,单位（秒）
func (obj *_EeoFilesMgr) WithTransitionLongs(transitionLongs int64) Option {
	return optionFunc(func(o *options) { o.query["transition_longs"] = transitionLongs })
}


// GetByOption 功能选项模式获取
func (obj *_EeoFilesMgr) GetByOption(opts ...Option) (result entity.EeoFiles, err error) {
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
func (obj *_EeoFilesMgr) GetByOptions(opts ...Option) (results []*entity.EeoFiles, err error) {
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


// GetFromFileID 通过file_id获取内容 文件ID 
func (obj *_EeoFilesMgr)  GetFromFileID(fileID int64) (result entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&result).Error
	
	return
}

// GetBatchFromFileID 批量唯一主键查找 文件ID
func (obj *_EeoFilesMgr) GetBatchFromFileID(fileIDs []int64) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id IN (?)", fileIDs).Find(&results).Error
	
	return
}
 
// GetFromFileSavePath 通过file_save_path获取内容 文件保存路径 
func (obj *_EeoFilesMgr) GetFromFileSavePath(fileSavePath string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_save_path = ?", fileSavePath).Find(&results).Error
	
	return
}

// GetBatchFromFileSavePath 批量唯一主键查找 文件保存路径
func (obj *_EeoFilesMgr) GetBatchFromFileSavePath(fileSavePaths []string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_save_path IN (?)", fileSavePaths).Find(&results).Error
	
	return
}
 
// GetFromFileSaveName 通过file_save_name获取内容 文件保存名字 
func (obj *_EeoFilesMgr) GetFromFileSaveName(fileSaveName string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_save_name = ?", fileSaveName).Find(&results).Error
	
	return
}

// GetBatchFromFileSaveName 批量唯一主键查找 文件保存名字
func (obj *_EeoFilesMgr) GetBatchFromFileSaveName(fileSaveNames []string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_save_name IN (?)", fileSaveNames).Find(&results).Error
	
	return
}
 
// GetFromFileExtension 通过file_extension获取内容 文件格式 
func (obj *_EeoFilesMgr) GetFromFileExtension(fileExtension string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_extension = ?", fileExtension).Find(&results).Error
	
	return
}

// GetBatchFromFileExtension 批量唯一主键查找 文件格式
func (obj *_EeoFilesMgr) GetBatchFromFileExtension(fileExtensions []string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_extension IN (?)", fileExtensions).Find(&results).Error
	
	return
}
 
// GetFromFileSize 通过file_size获取内容 文件大小(单位是KB) 
func (obj *_EeoFilesMgr)  GetFromFileSize(fileSize int) (result entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size = ?", fileSize).Find(&result).Error
	
	return
}

// GetBatchFromFileSize 批量唯一主键查找 文件大小(单位是KB)
func (obj *_EeoFilesMgr) GetBatchFromFileSize(fileSizes []int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size IN (?)", fileSizes).Find(&results).Error
	
	return
}
 
// GetFromTransitionFiles 通过transition_files获取内容 转换后保存文件地址 
func (obj *_EeoFilesMgr) GetFromTransitionFiles(transitionFiles string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_files = ?", transitionFiles).Find(&results).Error
	
	return
}

// GetBatchFromTransitionFiles 批量唯一主键查找 转换后保存文件地址
func (obj *_EeoFilesMgr) GetBatchFromTransitionFiles(transitionFiless []string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_files IN (?)", transitionFiless).Find(&results).Error
	
	return
}
 
// GetFromTransitionData 通过transition_data获取内容 json格式数据 
func (obj *_EeoFilesMgr) GetFromTransitionData(transitionData string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_data = ?", transitionData).Find(&results).Error
	
	return
}

// GetBatchFromTransitionData 批量唯一主键查找 json格式数据
func (obj *_EeoFilesMgr) GetBatchFromTransitionData(transitionDatas []string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_data IN (?)", transitionDatas).Find(&results).Error
	
	return
}
 
// GetFromTransitionState 通过transition_state获取内容 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的） 
func (obj *_EeoFilesMgr) GetFromTransitionState(transitionState bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_state = ?", transitionState).Find(&results).Error
	
	return
}

// GetBatchFromTransitionState 批量唯一主键查找 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的）
func (obj *_EeoFilesMgr) GetBatchFromTransitionState(transitionStates []bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_state IN (?)", transitionStates).Find(&results).Error
	
	return
}
 
// GetFromModerationState 通过moderation_state获取内容 内容审核状态（0.正常, 1.违规） 
func (obj *_EeoFilesMgr) GetFromModerationState(moderationState int8) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("moderation_state = ?", moderationState).Find(&results).Error
	
	return
}

// GetBatchFromModerationState 批量唯一主键查找 内容审核状态（0.正常, 1.违规）
func (obj *_EeoFilesMgr) GetBatchFromModerationState(moderationStates []int8) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("moderation_state IN (?)", moderationStates).Find(&results).Error
	
	return
}
 
// GetFromUploadUserid 通过upload_userid获取内容 上传者id 
func (obj *_EeoFilesMgr)  GetFromUploadUserid(uploadUserid int64) (result entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("upload_userid = ?", uploadUserid).Find(&result).Error
	
	return
}

// GetBatchFromUploadUserid 批量唯一主键查找 上传者id
func (obj *_EeoFilesMgr) GetBatchFromUploadUserid(uploadUserids []int64) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("upload_userid IN (?)", uploadUserids).Find(&results).Error
	
	return
}
 
// GetFromFileType 通过file_type获取内容 文件状态(0为个人上传，1为平台上传，2为第三方资料,3作业上传 ,4 im聊天图片) 
func (obj *_EeoFilesMgr) GetFromFileType(fileType bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_type = ?", fileType).Find(&results).Error
	
	return
}

// GetBatchFromFileType 批量唯一主键查找 文件状态(0为个人上传，1为平台上传，2为第三方资料,3作业上传 ,4 im聊天图片)
func (obj *_EeoFilesMgr) GetBatchFromFileType(fileTypes []bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_type IN (?)", fileTypes).Find(&results).Error
	
	return
}
 
// GetFromSystemDel 通过system_del获取内容 系统删除(0为未删除 1为已删除) 
func (obj *_EeoFilesMgr) GetFromSystemDel(systemDel bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("system_del = ?", systemDel).Find(&results).Error
	
	return
}

// GetBatchFromSystemDel 批量唯一主键查找 系统删除(0为未删除 1为已删除)
func (obj *_EeoFilesMgr) GetBatchFromSystemDel(systemDels []bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("system_del IN (?)", systemDels).Find(&results).Error
	
	return
}
 
// GetFromSystemLock 通过system_lock获取内容 系统屏蔽(0为未屏蔽可以分享 1为屏蔽禁止分享) 
func (obj *_EeoFilesMgr) GetFromSystemLock(systemLock bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("system_lock = ?", systemLock).Find(&results).Error
	
	return
}

// GetBatchFromSystemLock 批量唯一主键查找 系统屏蔽(0为未屏蔽可以分享 1为屏蔽禁止分享)
func (obj *_EeoFilesMgr) GetBatchFromSystemLock(systemLocks []bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("system_lock IN (?)", systemLocks).Find(&results).Error
	
	return
}
 
// GetFromSaveNum 通过save_num获取内容 被保存次数 
func (obj *_EeoFilesMgr) GetFromSaveNum(saveNum int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("save_num = ?", saveNum).Find(&results).Error
	
	return
}

// GetBatchFromSaveNum 批量唯一主键查找 被保存次数
func (obj *_EeoFilesMgr) GetBatchFromSaveNum(saveNums []int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("save_num IN (?)", saveNums).Find(&results).Error
	
	return
}
 
// GetFromDownNum 通过down_num获取内容 被下载次数 
func (obj *_EeoFilesMgr) GetFromDownNum(downNum int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("down_num = ?", downNum).Find(&results).Error
	
	return
}

// GetBatchFromDownNum 批量唯一主键查找 被下载次数
func (obj *_EeoFilesMgr) GetBatchFromDownNum(downNums []int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("down_num IN (?)", downNums).Find(&results).Error
	
	return
}
 
// GetFromExp1 通过exp1获取内容 扩展字段 
func (obj *_EeoFilesMgr) GetFromExp1(exp1 bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 = ?", exp1).Find(&results).Error
	
	return
}

// GetBatchFromExp1 批量唯一主键查找 扩展字段
func (obj *_EeoFilesMgr) GetBatchFromExp1(exp1s []bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exp1 IN (?)", exp1s).Find(&results).Error
	
	return
}
 
// GetFromMigrateStatus 通过migrate_status获取内容 文件迁移cos状态(0为未迁移 1为已迁移 2迁移失败) 
func (obj *_EeoFilesMgr) GetFromMigrateStatus(migrateStatus bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("migrate_status = ?", migrateStatus).Find(&results).Error
	
	return
}

// GetBatchFromMigrateStatus 批量唯一主键查找 文件迁移cos状态(0为未迁移 1为已迁移 2迁移失败)
func (obj *_EeoFilesMgr) GetBatchFromMigrateStatus(migrateStatuss []bool) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("migrate_status IN (?)", migrateStatuss).Find(&results).Error
	
	return
}
 
// GetFromMigrateTime 通过migrate_time获取内容 文件迁移cos时间 
func (obj *_EeoFilesMgr) GetFromMigrateTime(migrateTime int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("migrate_time = ?", migrateTime).Find(&results).Error
	
	return
}

// GetBatchFromMigrateTime 批量唯一主键查找 文件迁移cos时间
func (obj *_EeoFilesMgr) GetBatchFromMigrateTime(migrateTimes []int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("migrate_time IN (?)", migrateTimes).Find(&results).Error
	
	return
}
 
// GetFromAddtime 通过addtime获取内容 上传时间 
func (obj *_EeoFilesMgr)  GetFromAddtime(addtime int) (result entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&result).Error
	
	return
}

// GetBatchFromAddtime 批量唯一主键查找 上传时间
func (obj *_EeoFilesMgr) GetBatchFromAddtime(addtimes []int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime IN (?)", addtimes).Find(&results).Error
	
	return
}
 
// GetFromUptime 通过uptime获取内容 状态更新时间 
func (obj *_EeoFilesMgr) GetFromUptime(uptime int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime = ?", uptime).Find(&results).Error
	
	return
}

// GetBatchFromUptime 批量唯一主键查找 状态更新时间
func (obj *_EeoFilesMgr) GetBatchFromUptime(uptimes []int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uptime IN (?)", uptimes).Find(&results).Error
	
	return
}
 
// GetFromTransitionType 通过transition_type获取内容  
func (obj *_EeoFilesMgr) GetFromTransitionType(transitionType string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_type = ?", transitionType).Find(&results).Error
	
	return
}

// GetBatchFromTransitionType 批量唯一主键查找 
func (obj *_EeoFilesMgr) GetBatchFromTransitionType(transitionTypes []string) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_type IN (?)", transitionTypes).Find(&results).Error
	
	return
}
 
// GetFromTransitionStartTime 通过transition_start_time获取内容 转换开始时间 
func (obj *_EeoFilesMgr) GetFromTransitionStartTime(transitionStartTime int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_start_time = ?", transitionStartTime).Find(&results).Error
	
	return
}

// GetBatchFromTransitionStartTime 批量唯一主键查找 转换开始时间
func (obj *_EeoFilesMgr) GetBatchFromTransitionStartTime(transitionStartTimes []int) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_start_time IN (?)", transitionStartTimes).Find(&results).Error
	
	return
}
 
// GetFromTransitionLongs 通过transition_longs获取内容 转换时长,单位（秒） 
func (obj *_EeoFilesMgr) GetFromTransitionLongs(transitionLongs int64) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_longs = ?", transitionLongs).Find(&results).Error
	
	return
}

// GetBatchFromTransitionLongs 批量唯一主键查找 转换时长,单位（秒）
func (obj *_EeoFilesMgr) GetBatchFromTransitionLongs(transitionLongss []int64) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("transition_longs IN (?)", transitionLongss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EeoFilesMgr) FetchByPrimaryKey(fileID int64 ) (result entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&result).Error
	
	return
}

 
 // FetchIndexByIEfFilesize  获取多个内容
 func (obj *_EeoFilesMgr) FetchIndexByIEfFilesize(fileSize int ) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_size = ?", fileSize).Find(&results).Error
	
	return
}
 
 // FetchIndexByEeoFilesUploadUserid  获取多个内容
 func (obj *_EeoFilesMgr) FetchIndexByEeoFilesUploadUserid(uploadUserid int64 ) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("upload_userid = ?", uploadUserid).Find(&results).Error
	
	return
}
 
 // FetchIndexByIEfAddt  获取多个内容
 func (obj *_EeoFilesMgr) FetchIndexByIEfAddt(addtime int ) (results []*entity.EeoFiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&results).Error
	
	return
}
 

	

