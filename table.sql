CREATE TABLE IF NOT EXISTS `dirs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `loc` varchar(4096) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `size` bigint DEFAULT NULL,
  `name` varchar(1024) DEFAULT NULL,
  `ext` varchar(1024) DEFAULT NULL,
    PRIMARY KEY (`id`)
)