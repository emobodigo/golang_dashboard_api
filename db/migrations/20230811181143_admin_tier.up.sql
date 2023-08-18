CREATE TABLE IF NOT EXISTS `admin_tier` (
  `admin_tier_id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_level` int(11) NOT NULL,
  `division_id` int(11) NOT NULL,
  `level_title` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fulltime` tinyint(4) NOT NULL DEFAULT 1,
  PRIMARY KEY (`admin_tier_id`),
  KEY `division_id` (`division_id`),
  CONSTRAINT `admin_tier_ibfk_1` FOREIGN KEY (`division_id`) REFERENCES `admin_division` (`division_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;