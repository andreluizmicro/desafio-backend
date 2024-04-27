-- transfer.user_types definition

CREATE TABLE `user_types` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO user_types (id, name, created_at) VALUES (1, "comum", NOW());
INSERT INTO user_types (id, name, created_at) VALUES (2, "lojista", NOW());