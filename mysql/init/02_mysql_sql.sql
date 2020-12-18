CREATE DATABASE IF NOT EXISTS webgo_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use webgo_db;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_base
-- ----------------------------
CREATE TABLE IF NOT EXISTS `identity`  (
  `userid` varchar(32) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`userid`) USING BTREE,
  UNIQUE INDEX `userid`(`userid`) USING BTREE,
  UNIQUE INDEX `name`(`username`) USING BTREE
);

SET FOREIGN_KEY_CHECKS = 1;
