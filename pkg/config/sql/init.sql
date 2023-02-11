CREATE TABLE `user`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_name`      varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`       varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `follow_count`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Follow Count',
    `follower_count` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Follower Count',
    `avatar`         TEXT NULL COMMENT 'Avatar URL',
    `created_at`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at`     timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY              `idx_user_name` (`user_name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`        bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User ID',
    `play_url`       TEXT NULL COMMENT 'Video url',
    `cover_url`      TEXT NULL COMMENT 'Cover url',
    `favorite_count` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Favorite count',
    `comment_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Comment count',
    `title`          TEXT NULL  COMMENT 'Title',
    `created_at`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Video create time',
    `updated_at`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Video update time',
    `deleted_at`     timestamp NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY              `idx_user_id` (`user_id`) COMMENT 'UserID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `favorite`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `video_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video ID',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favorite create time',
    PRIMARY KEY (`id`),
    KEY          `idx_video_user_id` (`video_id`, `user_id`) COMMENT 'Video User ID index',
    KEY          `idx_user_id` (`user_id`) COMMENT 'User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite table';

CREATE TABLE `comment`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `video_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video ID',
    `user_id`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'User ID',
    `content`    TEXT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favorite create time',
    PRIMARY KEY (`id`),
    KEY          `idx_video_id` (`video_id`) COMMENT 'Video ID index',
    KEY          `idx_user_id` (`user_id`) COMMENT 'User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';

CREATE TABLE `follow`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'From User ID',
    `to_user_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'To User ID',
    `created_at`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favorite create time',
    PRIMARY KEY (`id`),
    KEY            `idx_from_user_id` (`from_user_id`) COMMENT 'From User ID index',
    KEY            `idx_to_user_id` (`to_user_id`) COMMENT 'To User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Follow table';

CREATE TABLE `message`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'From User ID',
    `to_user_id`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'To User ID',
    `content`      TEXT NULL COMMENT 'Content',
    `created_at`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favorite create time',
    PRIMARY KEY (`id`),
    KEY            `idx_user_id` (`from_user_id`,`to_user_id`) COMMENT 'User ID index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';