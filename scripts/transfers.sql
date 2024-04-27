-- transfer.transfers definition

CREATE TABLE `transfers` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `payer_id` bigint(20) NOT NULL,
    `payee_id` bigint(20) NOT NULL,
    `value` float NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `transfers_accounts_FK` (`payer_id`),
    KEY `transfers_accounts_FK_1` (`payee_id`),
    CONSTRAINT `transfers_accounts_FK` FOREIGN KEY (`payer_id`) REFERENCES `accounts` (`id`),
    CONSTRAINT `transfers_accounts_FK_1` FOREIGN KEY (`payee_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;