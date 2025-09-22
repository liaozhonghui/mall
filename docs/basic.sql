create table mall_user (
    `id` int unsigned not null auto_increment COMMENT '用户自增id',
    `nick_name` varchar(64) not null default '' COMMENT '用户昵称',
    `account` varchar(32) not null default '' COMMENT '用户账号',
    `password` varchar(64) not null default '' COMMENT '用户密码',
    `icon` varchar(256) not null default '' COMMENT '用户头像',
    `gender` tinyint unsigned not null default 0 COMMENT '用户性别 0:女 1:男',
    `status` tinyint unsigned not null default 1 COMMENT '用户状态 0:禁用 1:启用',
    `created_at` timestamp not null default current_timestamp COMMENT '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
