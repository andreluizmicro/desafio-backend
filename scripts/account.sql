-- transfer.accounts definition

CREATE TABLE `accounts` (
    `id` varchar(100) NOT NULL,
    `user_id` varchar(100) NOT NULL,
    `balance` float NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `active` tinyint(4) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `accounts_users_FK` (`user_id`),
    CONSTRAINT `accounts_users_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;