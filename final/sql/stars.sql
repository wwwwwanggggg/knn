/*
 Navicat Premium Data Transfer

 Source Server         : finaltest
 Source Server Type    : MySQL
 Source Server Version : 80035 (8.0.35)
 Source Host           : localhost:3306
 Source Schema         : tenzor2023

 Target Server Type    : MySQL
 Target Server Version : 80035 (8.0.35)
 File Encoding         : 65001

 Date: 15/12/2023 13:21:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for stars
-- ----------------------------
DROP TABLE IF EXISTS `stars`;
CREATE TABLE `stars`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `intro` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_stars_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of stars
-- ----------------------------
INSERT INTO `stars` VALUES (1, '2023-12-14 14:03:30.728', '2023-12-14 14:03:30.728', NULL, '不灭的安黛因', '狂风呼啸，站在你面前的是真正的英雄');
INSERT INTO `stars` VALUES (2, '2023-12-14 14:04:57.009', '2023-12-15 13:12:33.136', NULL, 'sans', '随便改改');
INSERT INTO `stars` VALUES (7, '2023-12-14 14:21:05.420', '2023-12-14 14:21:05.420', NULL, '小花', '我叫小花，一朵名为小花的小花');
INSERT INTO `stars` VALUES (8, '2023-12-15 12:54:30.073', '2023-12-15 12:54:30.073', NULL, '马跃', '天网恢恢，肥而不腻');

SET FOREIGN_KEY_CHECKS = 1;
