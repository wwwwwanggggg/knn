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

 Date: 15/12/2023 13:21:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2023-12-14 23:13:57.284', '2023-12-14 23:13:57.284', NULL, '不灭的安黛因', 'e10adc3949ba59abbe56e057f20f883e');
INSERT INTO `users` VALUES (2, '2023-12-14 23:44:31.763', '2023-12-14 23:44:31.763', NULL, 'sans', 'e35cf7b66449df565f93c607d5a81d09');
INSERT INTO `users` VALUES (3, '2023-12-15 12:57:22.219', '2023-12-15 12:57:22.219', NULL, 'Wang', 'c4aab061a42fcbdb23b1aaafe05575f0');

SET FOREIGN_KEY_CHECKS = 1;
