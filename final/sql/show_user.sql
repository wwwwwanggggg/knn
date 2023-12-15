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

 Date: 15/12/2023 13:20:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for show_user
-- ----------------------------
DROP TABLE IF EXISTS `show_user`;
CREATE TABLE `show_user`  (
  `show_id` bigint UNSIGNED NOT NULL COMMENT '主键',
  `user_id` bigint UNSIGNED NOT NULL COMMENT '主键',
  PRIMARY KEY (`show_id`, `user_id`) USING BTREE,
  INDEX `fk_show_user_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_show_user_show` FOREIGN KEY (`show_id`) REFERENCES `shows` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_show_user_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of show_user
-- ----------------------------
INSERT INTO `show_user` VALUES (4, 1);

SET FOREIGN_KEY_CHECKS = 1;
