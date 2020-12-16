USE `test`;
CREATE TABLE IF NOT EXISTS `person` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `Name` text,
  `Age` int,
  `Location` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
