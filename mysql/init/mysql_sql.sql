use webgo_db;
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
-- ----------------------------
-- Table structure for wedding_images
-- ----------------------------
CREATE TABLE IF NOT EXISTS `wedding_images`  (
  `id` varchar(32) NOT NULL,
  `name` varchar(255) NOT NULL,
  `uri` varchar(255) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
);
SET FOREIGN_KEY_CHECKS = 1;
