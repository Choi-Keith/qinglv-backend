
CREATE TABLE `collection_group` (
    `id` bigint(20) unsigned NOT NULL,
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '收藏分组名称',
    `biz_type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1帖子，2文章，3视频',
    `count` int(8) unsigned NOT NULL DEFAULT '1' COMMENT '数量',
    `visibility` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否公开: 1是，2否',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_name` (`name`),
    KEY `idx_biz_type` (`biz_type`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '收藏分组表';


CREATE TABLE `collection` (
    `id` bigint(20) unsigned NOT NULL,
    `group_id` bigint(20) unsigned NOT NULL COMMENT '收藏分组id',
    `target_id` bigint(20) unsigned NOT NULL COMMENT '帖子，文章或视频id',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`target_id`),
    KEY `idx_group_id` (`group_id`),
    KEY `idx_target_id` (`target_id`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '收藏表';

CREATE TABLE `post_share` (
    `id` bigint(20) unsigned NOT NULL,
    `post_id` bigint(20) unsigned NOT NULL COMMENT '帖子id',
    `type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1轻语,2微信,3朋友圈,4微博',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_type` (`type`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '帖子分享表';

CREATE TABLE `article_share` (
    `id` bigint(20) unsigned NOT NULL,
    `article_id` bigint(20) unsigned NOT NULL COMMENT '文章id',
    `type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1轻语,2微信,3朋友圈,4微博',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_type` (`type`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '文章分享表';

CREATE TABLE `post_thumb` (
    `id` bigint(20) unsigned NOT NULL,
    `post_id` bigint(20) unsigned NOT NULL COMMENT '帖子id',
    `like` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1喜欢，0取消',
    `dislike` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1不喜欢，0取消',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`post_id`),
    KEY `idx_post_id` (`post_id`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '帖子点赞和点踩表';

CREATE TABLE `post_comment_thumb` (
    `id` bigint(20) unsigned NOT NULL,
    `post_id` bigint(20) unsigned NOT NULL COMMENT '帖子id',
    `comment_id` bigint(20) unsigned NOT NULL COMMENT '评论id',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '帖子评论点赞和点踩表';

CREATE TABLE `article_thumb` (
    `id` bigint(20) unsigned NOT NULL,
    `article_id` bigint(20) unsigned NOT NULL COMMENT '文章id',
    `like` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1喜欢，0取消',
    `dislike` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1不喜欢，0取消',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`post_id`),
    KEY `idx_post_id` (`post_id`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '文章点赞和点踩表';

CREATE TABLE `article_comment_thumb` (
    `id` bigint(20) unsigned NOT NULL,
    `post_id` bigint(20) unsigned NOT NULL COMMENT '帖子id',
    `comment_id` bigint(20) unsigned NOT NULL COMMENT '评论id',
    `creator_id` bigint(20) unsigned NOT NULL COMMENT '创建人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
    KEY `idx_creator_id` (`creator_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '文章评论点赞和点踩表';