package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"jupiter/domain/entity"
)

type _EeoUserSharesMgr struct {
	*_BaseMgr
}

// EeoUserSharesMgr open func
func EeoUserSharesMgr(db *gorm.DB) *_EeoUserSharesMgr {
	if db == nil {
		panic(fmt.Errorf("EeoUserSharesMgr need init by db"))
	}
	return &_EeoUserSharesMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EeoUserSharesMgr) GetTableName() string {
	return "eeo_user_shares"
}

// Get 获取
func (obj *_EeoUserSharesMgr) Get() (result entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EeoUserSharesMgr) Gets() (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithShareID share_id获取 分享ID
func (obj *_EeoUserSharesMgr) WithShareID(shareID int64) Option {
	return optionFunc(func(o *options) { o.query["share_id"] = shareID })
}

// WithUserID user_id获取 用户ID
func (obj *_EeoUserSharesMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserResourceType user_resource_type获取 分享类型（1为文件，2为文件夹,3为im聊天文件）
func (obj *_EeoUserSharesMgr) WithUserResourceType(userResourceType bool) Option {
	return optionFunc(func(o *options) { o.query["user_resource_type"] = userResourceType })
}

// WithUserResourceID user_resource_id获取 分享资源ID（用户文件关系ID或者文件夹ID或者im聊天文件ID）
func (obj *_EeoUserSharesMgr) WithUserResourceID(userResourceID int64) Option {
	return optionFunc(func(o *options) { o.query["user_resource_id"] = userResourceID })
}

// WithUserResourceName user_resource_name获取 分享资源名字
func (obj *_EeoUserSharesMgr) WithUserResourceName(userResourceName string) Option {
	return optionFunc(func(o *options) { o.query["user_resource_name"] = userResourceName })
}

// WithUserResourceExtension user_resource_extension获取 分享资源的类型（只有文件才有）
func (obj *_EeoUserSharesMgr) WithUserResourceExtension(userResourceExtension string) Option {
	return optionFunc(func(o *options) { o.query["user_resource_extension"] = userResourceExtension })
}

// WithUserResourceSize user_resource_size获取 分享资源的大小（只有文件才有）
func (obj *_EeoUserSharesMgr) WithUserResourceSize(userResourceSize int) Option {
	return optionFunc(func(o *options) { o.query["user_resource_size"] = userResourceSize })
}

// WithSaveNum save_num获取 保存次数
func (obj *_EeoUserSharesMgr) WithSaveNum(saveNum int) Option {
	return optionFunc(func(o *options) { o.query["save_num"] = saveNum })
}

// WithIsdel isdel获取 是否删除（0未删除，1删除）
func (obj *_EeoUserSharesMgr) WithIsdel(isdel bool) Option {
	return optionFunc(func(o *options) { o.query["isdel"] = isdel })
}

// WithAddtime addtime获取 分享时间
func (obj *_EeoUserSharesMgr) WithAddtime(addtime int) Option {
	return optionFunc(func(o *options) { o.query["addtime"] = addtime })
}

// WithFileID file_id获取 eeo_files中的file_id
func (obj *_EeoUserSharesMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithOldID old_id获取 试卷、考试迁移前的share_id
func (obj *_EeoUserSharesMgr) WithOldID(oldID int64) Option {
	return optionFunc(func(o *options) { o.query["old_id"] = oldID })
}

// WithType type获取 类型：0=课件、1=考试、2=作业
func (obj *_EeoUserSharesMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCourseID course_id获取 分享文件的群id
func (obj *_EeoUserSharesMgr) WithCourseID(courseID int64) Option {
	return optionFunc(func(o *options) { o.query["course_id"] = courseID })
}

// WithSid sid获取 分享文件的机构id
func (obj *_EeoUserSharesMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["sid"] = sid })
}

// GetByOption 功能选项模式获取
func (obj *_EeoUserSharesMgr) GetByOption(opts ...Option) (result entity.EeoUserShares, err error) {
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
func (obj *_EeoUserSharesMgr) GetByOptions(opts ...Option) (results []*entity.EeoUserShares, err error) {
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

// GetFromShareID 通过share_id获取内容 分享ID
func (obj *_EeoUserSharesMgr) GetFromShareID(shareID int64) (result entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("share_id = ?", shareID).Find(&result).Error

	return
}

// GetBatchFromShareID 批量唯一主键查找 分享ID
func (obj *_EeoUserSharesMgr) GetBatchFromShareID(shareIDs []int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("share_id IN (?)", shareIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_EeoUserSharesMgr) GetFromUserID(userID int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户ID
func (obj *_EeoUserSharesMgr) GetBatchFromUserID(userIDs []int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserResourceType 通过user_resource_type获取内容 分享类型（1为文件，2为文件夹,3为im聊天文件）
func (obj *_EeoUserSharesMgr) GetFromUserResourceType(userResourceType bool) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_type = ?", userResourceType).Find(&results).Error

	return
}

// GetBatchFromUserResourceType 批量唯一主键查找 分享类型（1为文件，2为文件夹,3为im聊天文件）
func (obj *_EeoUserSharesMgr) GetBatchFromUserResourceType(userResourceTypes []bool) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_type IN (?)", userResourceTypes).Find(&results).Error

	return
}

// GetFromUserResourceID 通过user_resource_id获取内容 分享资源ID（用户文件关系ID或者文件夹ID或者im聊天文件ID）
func (obj *_EeoUserSharesMgr) GetFromUserResourceID(userResourceID int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_id = ?", userResourceID).Find(&results).Error

	return
}

// GetBatchFromUserResourceID 批量唯一主键查找 分享资源ID（用户文件关系ID或者文件夹ID或者im聊天文件ID）
func (obj *_EeoUserSharesMgr) GetBatchFromUserResourceID(userResourceIDs []int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_id IN (?)", userResourceIDs).Find(&results).Error

	return
}

// GetFromUserResourceName 通过user_resource_name获取内容 分享资源名字
func (obj *_EeoUserSharesMgr) GetFromUserResourceName(userResourceName string) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_name = ?", userResourceName).Find(&results).Error

	return
}

// GetBatchFromUserResourceName 批量唯一主键查找 分享资源名字
func (obj *_EeoUserSharesMgr) GetBatchFromUserResourceName(userResourceNames []string) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_name IN (?)", userResourceNames).Find(&results).Error

	return
}

// GetFromUserResourceExtension 通过user_resource_extension获取内容 分享资源的类型（只有文件才有）
func (obj *_EeoUserSharesMgr) GetFromUserResourceExtension(userResourceExtension string) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_extension = ?", userResourceExtension).Find(&results).Error

	return
}

// GetBatchFromUserResourceExtension 批量唯一主键查找 分享资源的类型（只有文件才有）
func (obj *_EeoUserSharesMgr) GetBatchFromUserResourceExtension(userResourceExtensions []string) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_extension IN (?)", userResourceExtensions).Find(&results).Error

	return
}

// GetFromUserResourceSize 通过user_resource_size获取内容 分享资源的大小（只有文件才有）
func (obj *_EeoUserSharesMgr) GetFromUserResourceSize(userResourceSize int) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_size = ?", userResourceSize).Find(&results).Error

	return
}

// GetBatchFromUserResourceSize 批量唯一主键查找 分享资源的大小（只有文件才有）
func (obj *_EeoUserSharesMgr) GetBatchFromUserResourceSize(userResourceSizes []int) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_resource_size IN (?)", userResourceSizes).Find(&results).Error

	return
}

// GetFromSaveNum 通过save_num获取内容 保存次数
func (obj *_EeoUserSharesMgr) GetFromSaveNum(saveNum int) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("save_num = ?", saveNum).Find(&results).Error

	return
}

// GetBatchFromSaveNum 批量唯一主键查找 保存次数
func (obj *_EeoUserSharesMgr) GetBatchFromSaveNum(saveNums []int) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("save_num IN (?)", saveNums).Find(&results).Error

	return
}

// GetFromIsdel 通过isdel获取内容 是否删除（0未删除，1删除）
func (obj *_EeoUserSharesMgr) GetFromIsdel(isdel bool) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel = ?", isdel).Find(&results).Error

	return
}

// GetBatchFromIsdel 批量唯一主键查找 是否删除（0未删除，1删除）
func (obj *_EeoUserSharesMgr) GetBatchFromIsdel(isdels []bool) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("isdel IN (?)", isdels).Find(&results).Error

	return
}

// GetFromAddtime 通过addtime获取内容 分享时间
func (obj *_EeoUserSharesMgr) GetFromAddtime(addtime int) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime = ?", addtime).Find(&results).Error

	return
}

// GetBatchFromAddtime 批量唯一主键查找 分享时间
func (obj *_EeoUserSharesMgr) GetBatchFromAddtime(addtimes []int) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addtime IN (?)", addtimes).Find(&results).Error

	return
}

// GetFromFileID 通过file_id获取内容 eeo_files中的file_id
func (obj *_EeoUserSharesMgr) GetFromFileID(fileID int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id = ?", fileID).Find(&results).Error

	return
}

// GetBatchFromFileID 批量唯一主键查找 eeo_files中的file_id
func (obj *_EeoUserSharesMgr) GetBatchFromFileID(fileIDs []int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("file_id IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromOldID 通过old_id获取内容 试卷、考试迁移前的share_id
func (obj *_EeoUserSharesMgr) GetFromOldID(oldID int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("old_id = ?", oldID).Find(&results).Error

	return
}

// GetBatchFromOldID 批量唯一主键查找 试卷、考试迁移前的share_id
func (obj *_EeoUserSharesMgr) GetBatchFromOldID(oldIDs []int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("old_id IN (?)", oldIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型：0=课件、1=考试、2=作业
func (obj *_EeoUserSharesMgr) GetFromType(_type int8) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型：0=课件、1=考试、2=作业
func (obj *_EeoUserSharesMgr) GetBatchFromType(_types []int8) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromCourseID 通过course_id获取内容 分享文件的群id
func (obj *_EeoUserSharesMgr) GetFromCourseID(courseID int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id = ?", courseID).Find(&results).Error

	return
}

// GetBatchFromCourseID 批量唯一主键查找 分享文件的群id
func (obj *_EeoUserSharesMgr) GetBatchFromCourseID(courseIDs []int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("course_id IN (?)", courseIDs).Find(&results).Error

	return
}

// GetFromSid 通过sid获取内容 分享文件的机构id
func (obj *_EeoUserSharesMgr) GetFromSid(sid int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sid = ?", sid).Find(&results).Error

	return
}

// GetBatchFromSid 批量唯一主键查找 分享文件的机构id
func (obj *_EeoUserSharesMgr) GetBatchFromSid(sids []int64) (results []*entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sid IN (?)", sids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EeoUserSharesMgr) FetchByPrimaryKey(shareID int64) (result entity.EeoUserShares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("share_id = ?", shareID).Find(&result).Error

	return
}
