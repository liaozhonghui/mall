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
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
create table mall_goods (
    `id` int unsigned not null auto_increment COMMENT '商品自增id',
    `name` varchar(100) not null default '' COMMENT '商品名称',
    `description` varchar(255) not null default '' COMMENT '商品描述',
    `tags` varchar(255) not null default '' COMMENT '商品标签',
    `detail` text not null COMMENT '商品详情',
    `category_id` int not null default 0 COMMENT '商品分类id',
    `small_image` varchar(255) not null default '' COMMENT '小图片',
    `detail_image` varchar(255) not null default '' COMMENT '详情图片',
    `price` int not null default 0 COMMENT '价格',
    `status` tinyint unsigned not null default 1 COMMENT '状态：1-正常，2-下架',
    `created_at` timestamp not null default current_timestamp COMMENT '创建时间',
    `updated_at` timestamp not null default current_tim estamp on update current_timestamp COMMENT '更新时间',
    PRIMARY KEY(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
create table `mall_goods_sku` (
    `id` int unsigned not null auto_increment COMMENT 'SKU自增id',
    `goods_id` int unsigned not null default 0 COMMENT '商品id',
    `attribute_ids` varchar(255) not null default '' COMMENT '商品属性id列表',
    `spend_price` int unsigned not null default 0 COMMENT '商品划算价',
    `price` int unsigned not null default 0 COMMENT '商品价格',
    `discount_price` int unsigned not null default 0 COMMENT '折扣价格',
    `left_store` int unsigned not null default 0 COMMENT '剩余库存',
    `all_store` int unsigned not null default 0 COMMENT '总库存',
    `status` tinyint unsigned not null default 1 COMMENT '状态：1-正常，2-下架',
    `created_at` timestamp not null default current_timestamp COMMENT '创建时间',
    `updated_at` timestamp not null default current_timestamp on update current_timestamp COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_goods_id` (`goods_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;