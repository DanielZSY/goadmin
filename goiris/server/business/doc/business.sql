/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50714
Source Host           : localhost:3306
Source Database       : business
Source Character Set  : utf8mb4
Source Collaction     : utf8mb4_general_ci

Target Server Type    : MYSQL
Target Server Version : 50714
File Encoding         : 65001
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for b_user
-- ----------------------------
DROP TABLE IF EXISTS `b_user`;
CREATE TABLE `b_user` (
    `id` int(64) NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL DEFAULT '',
    `password` varchar(255) NOT NULL DEFAULT '',
    `gender` tinyint(5) DEFAULT NULL COMMENT '性别 1 男 2 女  0 未选',
    `enable` tinyint(1) DEFAULT '1' COMMENT '1：启用用户 0：禁用用户',
    `name` varchar(255) NOT NULL DEFAULT '',
    `age` int(5) DEFAULT '0',
    `phone` varchar(255) NOT NULL DEFAULT '',
    `email` varchar(255) NOT NULL DEFAULT '',
    `avatar` varchar(255) NOT NULL DEFAULT '',
    `create_time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `username` (`username`) USING BTREE,
    UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=147 DEFAULT CHARSET=utf8_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of b_user
-- ----------------------------
INSERT INTO `b_user` VALUES ('1', 'daniel', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '1', '轩诚', '18', '13333333333', 'daniel@163.com', '', null);
