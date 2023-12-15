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

 Date: 15/12/2023 13:20:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for show_star
-- ----------------------------
DROP TABLE IF EXISTS `show_star`;
CREATE TABLE `show_star`  (
  `show_id` bigint UNSIGNED NOT NULL COMMENT '主键',
  `star_id` bigint UNSIGNED NOT NULL COMMENT '主键',
  PRIMARY KEY (`show_id`, `star_id`) USING BTREE,
  INDEX `fk_show_star_star`(`star_id` ASC) USING BTREE,
  CONSTRAINT `fk_show_star_show` FOREIGN KEY (`show_id`) REFERENCES `shows` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_show_star_star` FOREIGN KEY (`star_id`) REFERENCES `stars` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of show_star
-- ----------------------------
INSERT INTO `show_star` VALUES (3, 1);
INSERT INTO `show_star` VALUES (4, 1);
INSERT INTO `show_star` VALUES (3, 2);
INSERT INTO `show_star` VALUES (4, 2);
INSERT INTO `show_star` VALUES (3, 7);
INSERT INTO `show_star` VALUES (5, 8);

SET FOREIGN_KEY_CHECKS = 1;
