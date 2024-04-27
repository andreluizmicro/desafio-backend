-- transfer.users definition

CREATE TABLE `users` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_type_id` bigint(20) NOT NULL,
    `name` varchar(200) NOT NULL,
    `email` varchar(200) NOT NULL,
    `cpf` varchar(14) NOT NULL,
    `cnpj` varchar(18) DEFAULT NULL,
    `password` varchar(100) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted` tinyint(4) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_unique` (`cpf`),
    UNIQUE KEY `users_unique_1` (`email`),
    UNIQUE KEY `users_unique_2` (`cnpj`),
    KEY `users_user_types_FK` (`user_type_id`),
    CONSTRAINT `users_user_types_FK` FOREIGN KEY (`user_type_id`) REFERENCES `user_types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;