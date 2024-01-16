CREATE TABLE `topic` (
    `id` bigint(20) unsigned NOT NULL COMMENT '主键id',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '话题名称',
    `bg` varchar(64) NOT NULL DEFAULT '' COMMENT '话题背景图',
    `description` text COMMENT '话题描述',
    `quote_count` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '话题引用数量',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建者',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`),
    KEY `idx_name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '帖子话题';


CREATE TABLE `category` (
    `id` bigint(20) unsigned NOT NULL COMMENT '主键id',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '分类名称',
    `description` text  COMMENT '分类描述',
    `image` varchar(64) NOT NULL DEFAULT '' COMMENT '分类图片',
    `quote_count` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '分类引用数量',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建者',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`),
    KEY `idx_name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '分类';


CREATE TABLE `post` (
    `id` bigint(20) unsigned NOT NULL COMMENT '主键id',
    `status` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '审核状态：1已通过,2正在审核中，3不通过',
    `visibility` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '可见性：1公开,私密,仅好友可见',
    `is_top` tinyint(3) unsigned NOT NULL DEFAULT 2 COMMENT '是否置顶:1是，2否',
    `score` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '积分',
    `location` varchar(64) NOT NULL DEFAULT '' COMMENT '位置',
    `comment_count`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '评论数',
    `collection_count`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '收藏数',
    `like_count`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '点赞数',
    `dislike_count`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '点踩数',
    `share_count`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '分享',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '发布者',
    `latest_replied_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '最近回复时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_creator_id` (`creator_id`),
    KEY `idx_is_top` (`is_top`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '帖子表';

CREATE TABLE `post_content` (
    `id` bigint(20) unsigned NOT NULL COMMENT '主键id',
    `post_id` bigint(20) unsigned NOT NULL COMMENT '帖子id',
    `category_id` bigint(20) unsigned COMMENT '分类id',
    `topic` varchar(128) COMMENT '话题',
    `content` text NOT NULL DEFAULT '' COMMENT '内容',
    `images` json NOT NULL DEFAULT '' COMMENT '图片',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_creator_id` (`creator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '帖子表';

