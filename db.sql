-- 数据库表结构
-- CREATE TABLE IF NOT EXISTS `tb_user`;
CREATE TABLE IF NOT EXISTS `tb_user` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL,
    PRIMARY KEY(`id`),
    UNIQUE KEY `username` (`username`)
    -- KEY `idx_user_deleted_at`(`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

insert into `tb_user` (`id`, `username`, `password`) values(0, "admin", "");
