package model

// EeoFileMiddleTransate 上传文件转换中间表
type EeoFileMiddleTransate struct {
	ID    uint   `gorm:"primary_key;column:id;type:int unsigned;not null"` // ID
	UId   int64  `gorm:"column:uid;type:bigint;not null"`                  // 上传者id
	Type  string `gorm:"column:type;type:varchar(50);not null"`            // 文件格式
	Count int    `gorm:"column:count;type:int;not null"`                   // 待转换文件个数
}

// EeoFiles 基础文件表
type EeoFiles struct {
	FileID          uint64 `gorm:"primary_key;column:file_id;type:bigint unsigned;not null"`                                // 文件ID
	FileSavePath    string `gorm:"column:file_save_path;type:varchar(100)"`                                                 // 文件保存路径
	FileSaveName    string `gorm:"column:file_save_name;type:varchar(300)"`                                                 // 文件保存名字
	FileExtension   string `gorm:"index:eeo_files#file_extension#transition_state;column:file_extension;type:varchar(10)"`  // 文件格式
	FileSize        int    `gorm:"column:file_size;type:int"`                                                               // 文件大小(单位是KB)
	TransitionFiles string `gorm:"column:transition_files;type:varchar(100)"`                                               // 转换后保存文件地址
	TransitionData  string `gorm:"column:transition_data;type:mediumtext"`                                                  // json格式数据
	TransitionState bool   `gorm:"index:eeo_files#file_extension#transition_state;column:transition_state;type:tinyint(1)"` // 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的）
	ModerationState uint8  `gorm:"column:moderation_state;type:tinyint unsigned;not null"`                                  // 内容审核状态（0.正常, 1.违规）
	UploadUserid    int64  `gorm:"index:eeo_files_upload_userid;column:upload_userid;type:bigint"`                          // 上传者id
	FileType        bool   `gorm:"column:file_type;type:tinyint(1)"`                                                        // 文件状态(0为个人上传，1为平台上传，2为第三方资料,3作业上传 ,4 im聊天图片)
	SystemDel       bool   `gorm:"column:system_del;type:tinyint(1)"`                                                       // 系统删除(0为未删除 1为已删除)
	SystemLock      bool   `gorm:"column:system_lock;type:tinyint(1)"`                                                      // 系统屏蔽(0为未屏蔽可以分享 1为屏蔽禁止分享)
	SaveNum         int    `gorm:"column:save_num;type:int"`                                                                // 被保存次数
	DownNum         int    `gorm:"column:down_num;type:int"`                                                                // 被下载次数
	Exp1            bool   `gorm:"column:exp1;type:tinyint(1)"`                                                             // 扩展字段
	MigrateStatus   bool   `gorm:"column:migrate_status;type:tinyint(1)"`                                                   // 文件迁移cos状态(0为未迁移 1为已迁移 2迁移失败)
	MigrateTime     int    `gorm:"column:migrate_time;type:int"`                                                            // 文件迁移cos时间
	Addtime         int    `gorm:"index:i_ef_addt;column:addtime;type:int"`                                                 // 上传时间
	Uptime          int    `gorm:"column:uptime;type:int;not null"`                                                         // 状态更新时间
}

// EeoFilesLack 文件缺失记录表
type EeoFilesLack struct {
	ID           uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null"` // ID
	FileID       int64  `gorm:"index:file_id;column:file_id;type:bigint"`            // 文件ID
	FileSavePath string `gorm:"column:file_save_path;type:varchar(100)"`             // 文件保存路径
	Addtime      int    `gorm:"column:addtime;type:int"`                             // 添加时间
}

// EeoUserFiles 用户网盘文件对应表
type EeoUserFiles struct {
	ID              uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null"`     // ID
	UserID          int64  `gorm:"index:eeo_user_files#user_id;column:user_id;type:bigint"` // 用户id
	FileID          int64  `gorm:"column:file_id;type:bigint"`                              // 文件id
	FolderID        int64  `gorm:"index:folder_id;column:folder_id;type:bigint"`            // 所属文件夹id
	FileName        string `gorm:"column:file_name;type:varchar(300)"`                      // 文件名称
	FileSize        int    `gorm:"column:file_size;type:int"`                               // 文件大小(单位是KB)
	FileExtension   string `gorm:"column:file_extension;type:varchar(10)"`                  // 文件格式
	TransitionState bool   `gorm:"column:transition_state;type:tinyint(1)"`                 // 转换状态（0.等待转换, 1.转换中 2.转换成功 3.转换失败 4需要重新转换的）
	FileNamePy      string `gorm:"column:file_name_py;type:varchar(155)"`                   // 名字拼音
	Isdel           bool   `gorm:"column:isdel;type:tinyint(1)"`                            // 是否删除(在我的网盘里不显示)
	Exp1            bool   `gorm:"column:exp1;type:tinyint(1)"`                             // 扩展字段
	Addtime         int    `gorm:"column:addtime;type:int"`                                 // 存入我的网盘时间
	Uptime          int    `gorm:"column:uptime;type:int"`                                  // 最后修改时间
	AllowOperate    uint8  `gorm:"column:allow_operate;type:tinyint unsigned;not null"`     // 授权文件是否允许操作
}

// EeoUserFolders [...]
type EeoUserFolders struct {
	FolderID     uint64 `gorm:"primary_key;column:folder_id;type:bigint unsigned;not null"` // ID
	UserID       int64  `gorm:"index:user_id;column:user_id;type:bigint"`                   // 用户id
	FolderName   string `gorm:"column:folder_name;type:varchar(255)"`
	FolderNamePy string `gorm:"column:folder_name_py;type:varchar(155)"`                // 文件夹名称拼音
	FolderNode   bool   `gorm:"column:folder_node;type:tinyint(1)"`                     // 文件层级
	ParentID     int64  `gorm:"index:parenter_id;column:parent_id;type:bigint"`         // 父级id
	FolderPath   string `gorm:"index:folder_path;column:folder_path;type:varchar(200)"` // 目录结构
	SystemLock   bool   `gorm:"column:system_lock;type:tinyint(1)"`                     // 是否可分享(根据文件夹下用户存放的文件决定的 0为可分享 1为不可分享)
	Exp1         bool   `gorm:"column:exp1;type:tinyint(1)"`                            // 备用字段
	OperateLevel bool   `gorm:"column:operate_level;type:tinyint(1)"`                   // 操作级别（0 自己创建可以随意删除复制移动，1系统创建不可删除、移动、复制、重命名、分享）
	Isdel        bool   `gorm:"column:isdel;type:tinyint(1)"`                           // 是否删除(0没删除 1删除)
	Addtime      int    `gorm:"column:addtime;type:int"`                                // 添加时间
	Uptime       int    `gorm:"column:uptime;type:int"`                                 // 最后修改时间
	FromID       int64  `gorm:"column:from_id;type:bigint"`                             // 来自文件夹id
	FromParent   int64  `gorm:"column:from_parent;type:bigint"`                         // 原来的父级id
}

// EeoUserShares 用户云盘分享表
type EeoUserShares struct {
	ShareID               uint64 `gorm:"primary_key;column:share_id;type:bigint unsigned;not null"` // 分享ID
	UserID                uint64 `gorm:"column:user_id;type:bigint unsigned;not null"`              // 用户ID
	UserResourceType      bool   `gorm:"column:user_resource_type;type:tinyint(1);not null"`        // 分享类型（1为文件，2为文件夹,3为im聊天文件）
	UserResourceID        uint64 `gorm:"column:user_resource_id;type:bigint unsigned;not null"`     // 分享资源ID（用户文件关系ID或者文件夹ID或者im聊天文件ID）
	UserResourceName      string `gorm:"column:user_resource_name;type:varchar(300);not null"`      // 分享资源名字
	UserResourceExtension string `gorm:"column:user_resource_extension;type:varchar(20);not null"`  // 分享资源的类型（只有文件才有）
	UserResourceSize      int    `gorm:"column:user_resource_size;type:int;not null"`               // 分享资源的大小（只有文件才有）
	SaveNum               int    `gorm:"column:save_num;type:int;not null"`                         // 保存次数
	Isdel                 bool   `gorm:"column:isdel;type:tinyint(1);not null"`                     // 是否删除（0未删除，1删除）
	Addtime               int    `gorm:"column:addtime;type:int;not null"`                          // 分享时间
	FileID                uint64 `gorm:"column:file_id;type:bigint unsigned;not null"`              // eeo_files中的file_id
}
