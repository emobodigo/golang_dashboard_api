CREATE TABLE IF NOT EXISTS `admin_status` (
  `admin_status_id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_status` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`admin_status_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;