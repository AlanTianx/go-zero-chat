CREATE TABLE `user_tokens`
(
    `id`          bigint unsigned                         NOT NULL AUTO_INCREMENT,
    `user_id`     bigint unsigned                         NOT NULL,
    `token`       varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `token_key`   varchar(64) COLLATE utf8mb4_general_ci  NOT NULL COMMENT '// token 的 MD5值',
    `status`      tinyint(1)                              NOT NULL DEFAULT '1' COMMENT '0失效-1有效',
    `expire_time` int                                     NOT NULL DEFAULT '0',
    `created_at`  datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime                                         DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

