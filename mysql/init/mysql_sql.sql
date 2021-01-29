use webgo_db;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_base
-- ----------------------------
CREATE TABLE IF NOT EXISTS `identity`  (
  `userid` VARCHAR(32) NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`userid`) USING BTREE,
  UNIQUE INDEX `userid`(`userid`) USING BTREE,
  UNIQUE INDEX `name`(`username`) USING BTREE
);
-- ----------------------------
-- Table structure for secret_key
-- ----------------------------
CREATE TABLE IF NOT EXISTS `secret_key`  (
  `userid` VARCHAR(32) NOT NULL,
  `keySalt` VARCHAR(32) NOT NULL,
  `accessExpireTime` INT UNSIGNED NOT NULL,
  `refreshExpireTime` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`userid`) USING BTREE,
  UNIQUE INDEX `userid`(`userid`) USING BTREE,
  CONSTRAINT `secret_key` FOREIGN KEY (`userid`) REFERENCES `identity` (`userid`) ON DELETE CASCADE ON UPDATE CASCADE
);
SET FOREIGN_KEY_CHECKS = 1;
