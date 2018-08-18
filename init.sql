
-- 关键字表
CREATE TABLE `keywords_reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key_word` varchar(255) NOT NULL COMMENT '关键字key',
  `msg_type` varchar(255) NOT NULL COMMENT 'text:文本,image:图片,voice:声音,video:视频,music:音乐,news:图文消息',
  `value` varchar(255) DEFAULT NULL COMMENT 'text:文本,image:图片,voice:声音的值',
  `status` int(11) NOT NULL COMMENT '0:normal,9:delete',
  `created_at` datetime DEFAULT NULL,
  `created_person` varchar(255) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_person` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `account_id` varchar(255) NOT NULL COMMENT '公众号账户Id',
  `raw_id` varchar(255) NOT NULL COMMENT '公众号原始Id（冗余）',
  `deleted_person` varchar(255) DEFAULT NULL COMMENT '删除人员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8

-- 测试数据
INSERT INTO `test`.`keywords_reply`(`id`, `key_word`, `msg_type`, `value`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `account_id`, `raw_id`, `deleted_person`) VALUES (1, '测试文字', '0', '测试文本！', 0, '2018-05-16 22:28:17', 'SYSTEM', '2018-08-05 16:45:11', 'SYSTEM', NULL, '123', '123', NULL);
INSERT INTO `test`.`keywords_reply`(`id`, `key_word`, `msg_type`, `value`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `account_id`, `raw_id`, `deleted_person`) VALUES (15, '测试音乐', '4', NULL, 0, '2018-06-22 10:57:00', 'SYSTEM', '2018-08-05 16:45:15', 'SYSTEM', NULL, 'default', 'default', '');
INSERT INTO `test`.`keywords_reply`(`id`, `key_word`, `msg_type`, `value`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `account_id`, `raw_id`, `deleted_person`) VALUES (16, '测试图文', '5', NULL, 0, NULL, NULL, NULL, NULL, NULL, '123', '123', NULL);
INSERT INTO `test`.`keywords_reply`(`id`, `key_word`, `msg_type`, `value`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `account_id`, `raw_id`, `deleted_person`) VALUES (17, '测试图片', '1', 'http://127.0.0.1:8023/api/file/qrcode_for_gh_275ff6ac308f_344.jpg', 0, NULL, NULL, '2018-07-11 23:14:38', 'SYSTEM', NULL, '123', '123', NULL);


-- 音乐
CREATE TABLE `keywords_reply_music_sub` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '音乐标题',
  `description` varchar(255) DEFAULT NULL COMMENT '音乐描述',
  `music_url` varchar(255) DEFAULT NULL COMMENT '音乐地址',
  `hq_music_url` varchar(255) DEFAULT NULL COMMENT '高清音乐地址，一般用不到',
  `thumb_media_id` varchar(255) DEFAULT NULL COMMENT 'thumb_media_id 微信媒体Id',
  `status` int(11) NOT NULL COMMENT '0:normal,9:delete',
  `created_at` datetime DEFAULT NULL,
  `created_person` varchar(255) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_person` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `reply_id` varchar(255) NOT NULL COMMENT '冗余',
  `deleted_person` varchar(255) DEFAULT NULL COMMENT '删除人员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8


INSERT INTO `test`.`keywords_reply_music_sub`(`id`, `title`, `description`, `music_url`, `hq_music_url`, `thumb_media_id`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `reply_id`, `deleted_person`) VALUES (17, '我是音乐1', '我是音乐的描述', 'www.baidu.com', 'www.baidui.com', 'www.baidu.com', 0, '2018-06-24 15:45:00', 'SYSTEM', '2018-08-05 16:45:15', 'SYSTEM', NULL, '15', '');
INSERT INTO `test`.`keywords_reply_music_sub`(`id`, `title`, `description`, `music_url`, `hq_music_url`, `thumb_media_id`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `reply_id`, `deleted_person`) VALUES (19, 'titile', 'titile', 'titile', 's d', '', 9, '2018-07-08 17:47:18', 'SYSTEM', '2018-07-08 18:05:44', 'SYSTEM', '2018-07-08 18:38:04', '1', 'SYSTEM');
INSERT INTO `test`.`keywords_reply_music_sub`(`id`, `title`, `description`, `music_url`, `hq_music_url`, `thumb_media_id`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `reply_id`, `deleted_person`) VALUES (20, '史蒂夫', '史蒂夫', '史蒂夫', '', '', 9, '2018-07-08 18:25:24', 'SYSTEM', '2018-07-08 18:25:24', '', '2018-07-08 18:38:04', '1', 'SYSTEM');
INSERT INTO `test`.`keywords_reply_music_sub`(`id`, `title`, `description`, `music_url`, `hq_music_url`, `thumb_media_id`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `reply_id`, `deleted_person`) VALUES (21, '手动方式', '手动方式', '手动方式', '', '', 9, '2018-07-08 18:37:49', 'SYSTEM', '2018-07-08 18:37:49', '', '2018-07-08 18:38:04', '1', 'SYSTEM');

-- 图文
CREATE TABLE `keywords_reply_news_sub` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '图文标题',
  `description` varchar(255) DEFAULT NULL COMMENT '图文描述',
  `pic_url` varchar(255) DEFAULT NULL COMMENT '图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200',
  `url` varchar(255) DEFAULT NULL COMMENT '点击图文消息跳转链接',
  `status` int(11) NOT NULL COMMENT '0:normal,9:delete',
  `created_at` datetime DEFAULT NULL,
  `created_person` varchar(255) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_person` varchar(255) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `reply_id` varchar(255) NOT NULL COMMENT '冗余',
  `deleted_person` varchar(255) DEFAULT NULL COMMENT '删除人员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8

INSERT INTO `test`.`keywords_reply_news_sub`(`id`, `title`, `description`, `pic_url`, `url`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `reply_id`, `deleted_person`) VALUES (1, '图文1', '我是测试图文1', 'http://gb.cri.cn/mmsource/images/2014/02/10/37/3886148822463260913.jpg', 'http://baidu.com', 0, NULL, NULL, NULL, NULL, NULL, '16', NULL);
INSERT INTO `test`.`keywords_reply_news_sub`(`id`, `title`, `description`, `pic_url`, `url`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `reply_id`, `deleted_person`) VALUES (2, '图文2', '我是测试图文2', 'http://uploads.5068.com/allimg/1801/85-1P119154919.jpg', 'http://baidu.com', 0, NULL, NULL, NULL, NULL, NULL, '16', NULL);
INSERT INTO `test`.`keywords_reply_news_sub`(`id`, `title`, `description`, `pic_url`, `url`, `status`, `created_at`, `created_person`, `updated_at`, `updated_person`, `deleted_at`, `reply_id`, `deleted_person`) VALUES (3, 'title', '我是测试图文3', 'https://i04picsos.sogoucdn.com/09cfc8705fd94ebb', 'http://baidu.com', 0, NULL, NULL, NULL, NULL, NULL, '16', NULL);