CREATE TABLE `eeo_webcast_user_record2` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `class_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '课节ID',
  `cookie` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'cookie值',
  `account` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '帐号',
  `account_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '帐号类型',
  `nickname` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `source` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '来源',
  `client` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '客户端来源',
  `in_time` int(11) NOT NULL DEFAULT '0' COMMENT '进入时间',
  `out_time` int(11) NOT NULL DEFAULT '0' COMMENT '离开时间',
  `forbidden_speak` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁言 0不禁言 1禁言 (废弃)',
  `client_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客户端IP',
  PRIMARY KEY (`id`),
  KEY `account` (`account`),
  KEY `class_id` (`class_id`),
  KEY `out_in_time` (`out_time`,`in_time`),
  KEY `class_cookie_time` (`class_id`,`cookie`,`out_time`)
) ENGINE=InnoDB AUTO_INCREMENT=17729 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='直播聊天室成员进出表';

CREATE TABLE `eeo_webcast_like_info2` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `class_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '课节ID',
  `account` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '帐号',
  `like_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '点赞状态',
  PRIMARY KEY (`id`),
  KEY `eeo_webcast_like_info#indexA` (`class_id`,`like_status`),
  KEY `eeo_webcast_like_info#indexB` (`class_id`,`account`)
) ENGINE=InnoDB AUTO_INCREMENT=1802 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='直播聊天室点赞信息表';

CREATE TABLE `eeo_webcast_cast_info2` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `class_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '课节ID',
  `client_class_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '客户端课节ID',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '聊天室状态',
  `play_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '播放次数',
  `like_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点赞次数',
  `end_timestamp` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '课节结束时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `eeo_webcast_cast_info#class_id` (`class_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3226 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='直播聊天室相关信息表';