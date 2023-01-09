CREATE TABLE IF NOT EXISTS `users` (
`id` bigint NOT NULL AUTO_INCREMENT,
`first_name`varchar(100) NOT NULL,
`last_name`varchar(100) NOT NULL,
`email` varchar(100) NOT NULL,
`password` varchar(100) NULL,
`is_admin` tinyint(1) NOT NULL DEFAULT 0,
`is_active` tinyint NOT NULL DEFAULT 0,
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `idx_user_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;