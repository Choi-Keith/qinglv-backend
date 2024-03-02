
CREATE TABLE `user` (
    `id` bigint(20) unsigned NOT NULL COMMENT '用户id',
    `role_id` bigint(20) unsigned NOT NULL COMMENT '角色id',
    `account` varchar(64) NOT NULL COMMENT '账号',
    `nickname` varchar(64) NOT NULL COMMENT '昵称',
    `motto` text COMMENT '个性签名',
    `profession` varchar(65) COMMENT '职业',
    `email` varchar(64) NOT NULL COMMENT '邮箱',
    `we_chat` varchar(64) COMMENT '微信号',
    `auth_type` int(6) NOT NULL DEFAULT 1 COMMENT '注册方式:1邮件注册；2手机号注册',
    `phone` varchar(20) COMMENT '手机号',
    `password` varchar(128) NOT NULL COMMENT '密码',
    `avatar` varchar(128) NOT NULL COMMENT '头像',
    `profile_bg` varchar(128) COMMENT '个人主页背景图',
    `status` int(6) NOT NULL DEFAULT 2 COMMENT '状态:1已注销，2正常',
    `mail_status` int(6) NOT NULL DEFAULT 1 COMMENT '状态:1未激活，2正常',
    `location` text COMMENT '位置',
    `age` int(10) COMMENT '年龄',
    `gender` int(6) COMMENT '性别',
    `level` int(6) DEFAULT 0 COMMENT '等级',
    `score` int(10) DEFAULT 0 COMMENT '积分',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `last_login_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最近一次登录时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`nickname`),
    UNIQUE KEY (`email`),
    UNIQUE KEY (`phone`),
    UNIQUE KEy (`we_chat`),
    KEY `idx_nickname` (`nickname`),
    KEY `idx_email` (`email`),
    KEY `idx_phone` (`phone`),
    KEY `idx_we_chat` (`we_chat`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';

CREATE TABLE `role` (
    `id` bigint(20) unsigned NOT NULL COMMENT '角色id',
    `name` varchar(64) NOT NULL COMMENT '角色名称',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '角色表';


CREATE TABLE `following` (
    `id` bigint(20) unsigned NOT NULL COMMENT '关注id',
    `user_id`  bigint(20) unsigned NOT NULL COMMENT '用户id',
    `following_id` bigint(20) unsigned NOT NULL COMMENT '关注用户id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_following_user` (`user_id`,`following_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '关注表'

CREATE TABLE `follower`(
    `id` bigint(20) unsigned NOT NULL COMMENT '关注id',
    `user_id`  bigint(20) unsigned NOT NULL COMMENT '用户id',
    `follower_id` bigint(20) unsigned NOT NULL COMMENT '粉丝id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_follower_user` (`user_id`,`follower_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '粉丝表'


CREATE TABLE `blacklist` (
    `id` bigint(20) unsigned NOT NULL COMMENT '黑名单id',
    `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
    `black_id` bigint(20) unsigned NOT NULL COMMENT '黑名单用户id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_black_user` (`user_id`,`black_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '黑名单表';

CREATE TABLE `comment_notify` (
    `id` bigint(20) unsigned NOT NULL COMMENT '消息通知Id',
    `sender_user_id` bigint(20) unsigned NOT NULL COMMENT '发送方用户id',
    `receiver_user_id` bigint(20) unsigned NOT NULL COMMENT '接收方用户id',
    `comment_id` bigint(20) unsigned NOT NULL COMMENT '评论id',
    `comment_content`  varchar(512) NOT NULL DEFAULT '' COMMENT '评论内容',
    `reply_id` bigint(20) unsigned NOT NULL COMMENT '回复id',
    `reply_content` varchar(512) NOT NULL DEFAULT '' COMMENT '回复内容',
    `target_id` bigint(20) unsigned NOT NULL COMMENT '内容id,可能是文章或其它',
    `target_title` varchar(60) NOT NULL DEFAULT '' COMMENT '内容标题,可能是文章或其它',
    `type` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '内容类型:1动态，2文章，3其它',
    `is_read` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已读：0否，1是',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_comment_notify_is_read` (`is_read`) USING BTREE,
    KEY `idx_comment_notify_receiver_user_id` (`receiver_user_id`) USING BTREE,
    KEY `idx_comment_notify_target_id` (`target_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '评论通知';


CREATE TABLE `like_notify` (
    `id` bigint(20) unsigned NOT NULL COMMENT '消息通知Id',
    `sender_user_id` bigint(20) unsigned NOT NULL COMMENT '发送方用户id',
    `receiver_user_id` bigint(20) unsigned NOT NULL COMMENT '接收方用户id',
    `target_id` bigint(20) unsigned NOT NULL COMMENT '内容id,可能是文章或其它',
    `type` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '内容类型:1文章，2其它',
    `is_read` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已读：0否，1是',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_like_notify_is_read` (`is_read`) USING BTREE,
    KEY `idx_like_notify_receiver_user_id` (`receiver_user_id`) USING BTREE,
    KEY `idx_like_notify_target_id` (`target_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '点赞和收藏通知';


CREATE TABLE `follow_notify` (
    `id` bigint(20) unsigned NOT NULL COMMENT '消息通知Id',
    `sender_user_id` bigint(20) unsigned NOT NULL COMMENT '发送方用户id',
    `receiver_user_id` bigint(20) unsigned NOT NULL COMMENT '接收方用户id',
    `is_read` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已读：0否，1是',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_follow_notify_is_read` (`is_read`) USING BTREE,
    KEY `idx_follow_notify_receiver_user_id` (`receiver_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '新增粉丝通知';


CREATE TABLE `os_notify` (
    `id` bigint(20) unsigned NOT NULL COMMENT '消息通知Id',
    `sender_user_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '发送方用户id',
    `receiver_user_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '接收方用户id',
    `message` varchar(512) NOT NULL DEFAULT '' COMMENT '消息内容', 
    `is_read` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已读：0否，1是',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_os_notify_is_read` (`is_read`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '系统通知';

