CREATE TABLE `note_user`
(
    `id`       integer      NOT NULL,
    `name`     varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,

    `nickname` varchar(255) NOT NULL,
    `identity` varchar(255) NOT NULL, -- 身份
    PRIMARY KEY (`id`),
    UNIQUE unique_name (`name`),
    INDEX      index_login_name_password (`name`,`password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;