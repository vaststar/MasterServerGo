PRAGMA foreign_keys = OFF;

-- ----------------------------
-- Table structure for identity
-- ----------------------------
CREATE TABLE IF NOT EXISTS `identity` (
  `identity_id` VARCHAR (32) NOT NULL,
  `lastname` VARCHAR (255) NOT NULL,
  `firstname` VARCHAR (255) NOT NULL,
  PRIMARY KEY (`identity_id`),
  UNIQUE (`identity_id` ASC)
);

PRAGMA foreign_keys = ON;
