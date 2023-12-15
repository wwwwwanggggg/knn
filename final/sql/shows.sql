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

 Date: 15/12/2023 13:20:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for shows
-- ----------------------------
DROP TABLE IF EXISTS `shows`;
CREATE TABLE `shows`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `starttime` datetime(3) NOT NULL,
  `endtime` datetime(3) NOT NULL,
  `location` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `max_capacity` bigint NOT NULL,
  `show_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `promo` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `sold` bigint NOT NULL,
  `curr_capacity` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_shows_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of shows
-- ----------------------------
INSERT INTO `shows` VALUES (1, '2023-12-14 14:06:33.541', '2023-12-14 23:50:44.811', '2023-12-15 10:53:40.053', '打打屠杀线', '2023-12-15 02:00:00.000', '2023-12-15 03:00:00.000', 'ut', 100, '', '', 12, 88);
INSERT INTO `shows` VALUES (3, '2023-12-15 10:06:41.382', '2023-12-15 10:06:41.382', NULL, '测试用', '2323-12-15 20:00:00.000', '2023-12-15 21:00:00.000', 'PC', 100, '没有内容', '不想宣传', 0, 100);
INSERT INTO `shows` VALUES (4, '2023-12-15 12:39:27.927', '2023-12-15 12:40:42.208', NULL, '新的演出', '2023-12-17 02:00:00.000', '2023-12-17 02:30:00.000', 'C-106', 50, '瞎编的', '随便宣传', 1, 49);
INSERT INTO `shows` VALUES (5, '2023-12-15 13:10:03.431', '2023-12-15 13:10:03.431', NULL, '数学分析', '2023-12-19 03:10:00.000', '2023-12-19 06:00:00.000', 'C-406', 130, '就你还想听懂晚上的课，想多了', '少一分太瘦，多一分太肥', 0, 130);

SET FOREIGN_KEY_CHECKS = 1;
