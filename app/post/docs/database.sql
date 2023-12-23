CREATE TABLE `post`(
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT "动态id",
    `user_id` bigint(20) unsigned COMMENT "用户id",
    `content` text COMMENT "内容",
    `is_top` tinyint(3) unsigned DEFAULT 0 COMMENT "是否置顶:0否, 1是",
    `visibility` tinyint(3) unsigned DEFAULT 0 COMMENT "可见性:0公开,1私密,2朋友可见",
    `status` tinyint(3) unsigned DEFAULT 1 COMMENT "动态状态:0正在审核,1审核通过,12审核不通过",
    `sort` int(10) unsigned COMMENT "排序",
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(6),
    PRIMARY KEY (`id`)
    KEY `idx_user_id` (`user_id`)
    KEY `idx_content` (`content`) 
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '动态表';

CREATE TABLE `media` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT "媒体id",
    `post_id` bigint(20) unsigned COMMENT "动态id",
    `user_id` bigint(20) unsigned COMMENT "用户id",
    `type` tinyint(3) unsigned NOT NULL COMMENT "文件类型:0图片,1视频",
    `file_url` varchar(20) NOT NULL COMMENT "媒体文件地址",
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(6),
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '媒体表';