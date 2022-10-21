/*
 Navicat Premium Data Transfer

 Source Server         : localhostTest
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : game

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 21/10/2022 17:12:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for room
-- ----------------------------
DROP TABLE IF EXISTS `room`;
CREATE TABLE `room`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '房间id',
  `userId1` int(0) NULL DEFAULT NULL COMMENT '用户id',
  `userId2` int(0) NULL DEFAULT NULL COMMENT '用户id',
  `size` int(0) NULL DEFAULT NULL COMMENT '地图总格子数',
  `user1Steps` int(0) NULL DEFAULT NULL COMMENT '用户1步数',
  `user2Steps` int(0) NULL DEFAULT NULL COMMENT '用户2步数',
  `delete_flag` int(0) NOT NULL COMMENT '0表示未删除，1表示已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of room
-- ----------------------------
INSERT INTO `room` VALUES (1, 1, 2, 30, 3, 3, 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', 'admin');
INSERT INTO `user` VALUES (2, 'test', 'test');

SET FOREIGN_KEY_CHECKS = 1;
