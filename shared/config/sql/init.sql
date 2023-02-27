CREATE TABLE `user`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `name`              varchar(32) NOT NULL DEFAULT '' COMMENT 'Name',
    `password`          varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
    `avatar`            TEXT NULL COMMENT 'Avatar URL',
    `back_ground_image` TEXT NULL COMMENT 'BackGroundImage',
    `signature`         TEXT NULL COMMENT 'Signature',
    PRIMARY KEY (`id`),
    KEY                 `idx_name` (`name`) COMMENT 'Name index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User ID',
    `play_url`   TEXT NULL COMMENT 'Video url',
    `cover_url`  TEXT NULL COMMENT 'Cover url',
    `title`      TEXT NULL COMMENT 'Title',
    `created_at` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video create time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`) COMMENT 'UserID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `favorite`
(
    `video_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video ID',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User ID',
    `created_at` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Favorite create time',
    PRIMARY KEY (`video_id`, `user_id`),
    KEY          `idx_user_id` (`user_id`) COMMENT 'User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite table';

CREATE TABLE `comment`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `video_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video ID',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User ID',
    `content`    TEXT NULL COMMENT 'Content',
    `created_at` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Favorite create time',
    PRIMARY KEY (`id`),
    KEY          `idx_video_id` (`video_id`) COMMENT 'Video ID index',
    KEY          `idx_user_id` (`user_id`) COMMENT 'User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';

CREATE TABLE `follow`
(
    `from_user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'From User ID',
    `to_user_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'To User ID',
    PRIMARY KEY (`from_user_id`, `to_user_id`),
    KEY            `idx_to_user_id` (`to_user_id`) COMMENT 'To User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Follow table';

CREATE TABLE `message`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'From User ID',
    `to_user_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'To User ID',
    `content`      TEXT NULL COMMENT 'Content',
    `created_at`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Favorite create time',
    PRIMARY KEY (`id`),
    KEY            `idx_user_id` (`from_user_id`,`to_user_id`) COMMENT 'User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';