
CREATE TABLE `user` (
    `id` bigint(20) unsigned NOT NULL COMMENT '用户id',
    `role_id` bigint(20) unsigned NOT NULL COMMENT '角色id',
    `account` varchar(64) NOT NULL COMMENT '账号',
    `nickname` varchar(64) NOT NULL COMMENT '昵称',
    `motto` text COMMENT '个性签名',
    `email` varchar(64) NOT NULL COMMENT '邮箱',
    `we_chat` varchar(64) COMMENT '微信号',
    `auth_type` int(6) NOT NULL DEFAULT 1 COMMENT '注册方式:1邮件注册；2手机号注册',
    `phone` varchar(20) COMMENT '手机号',
    `password` varchar(128) COMMENT '密码',
    `avatar` varchar(128) COMMENT '头像',
    `profile_bg` varchar(128) COMMENT '个人主页背景图',
    `status` int(6) COMMENT '状态',
    `location` text COMMENT '位置',
    `age` int(10) COMMENT '年龄',
    `gender` int(6) COMMENT '性别',
    `level` int(6) DEFAULT 0 COMMENT '等级',
    `score` int(10) DEFAULT 0 COMMENT '积分',
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
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
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '角色表';

INSERT INTO TABLE role values()

CREATE TABLE `following` (
    `id` bigint(20) unsigned NOT NULL COMMENT '关注id',
    `user_id`  bigint(20) unsigned NOT NULL COMMENT '用户id',
    `following_id` bigint(20) unsigned NOT NULL COMMENT '关注用户id',
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '关注表'

CREATE TABLE `follower`(
    `id` bigint(20) unsigned NOT NULL COMMENT '关注id',
    `user_id`  bigint(20) unsigned NOT NULL COMMENT '用户id',
    `follower_id` bigint(20) unsigned NOT NULL COMMENT '粉丝id',
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '粉丝表'


CREATE TABLE `blacklist` (
    `id` bigint(20) unsigned NOT NULL COMMENT '黑名单id',
    `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '黑名单表';