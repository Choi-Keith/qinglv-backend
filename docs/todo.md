

- [X] 优化帖子， 话题和文章推荐算法

- [X] 删除帖子和删除文章时，需要删除评论，收藏，分享和点赞信息，并且更新用户的经验值和等级 

- [] 文章, 消息，用户关注

- [] 文章浏览记录

- [] 帖子和文章生成短链，分享

- [] 使用es优化搜索

- [] 评论踩一下，待进行处理

- [] 优化数据库查询


## 缺失的接口
1. 关注和粉丝的总数量
2. 忘记密码接口
3. 点赞与收藏总数量

4. 新增用户数据表
```sql
CREATE table `user_article_count`(
    `id` bigint(20) unsigned NOT NULL COMMENT '用户数据表id',
    `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
    `level` int(6) DEFAULT 0 COMMENT '等级',
    `score` int(10) DEFAULT 0 COMMENT '积分',
    `like_count` int(10) DEFAULT 0 COMMENT '点赞数',
    `collect_count` int(10) DEFAULT 0 COMMENT '收藏数',
    `share_count` int(10) DEFAULT 0 COMMENT '转发数',
    `view_count` int(10) DEFAULT 0 COMMENT '观看数',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
    `is_del` tinyint(1) NOT NULL DEFAULT '0',
    `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
    PRIMARY KEY (`id`),
)
```




