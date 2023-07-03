/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50520
 Source Host           : localhost:3306
 Source Schema         : eframe.dev.main

 Target Server Type    : MySQL
 Target Server Version : 50520
 File Encoding         : 65001

 Date: 02/07/2023 16:24:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for player
-- ----------------------------
DROP TABLE IF EXISTS `player`;
CREATE TABLE `player`  (
  `id` int(11) NOT NULL,
  `account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `online` int(11) UNSIGNED ZEROFILL NOT NULL,
  `conn_url` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `conn_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of player
-- ----------------------------
INSERT INTO `player` VALUES (0, 'test01', '123456', 00000000000, '', NULL);

SET FOREIGN_KEY_CHECKS = 1;
