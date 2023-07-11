/*
 Navicat Premium Data Transfer

 Source Server         : 81.71.31.250
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 81.71.31.250:3306
 Source Schema         : giligili_video

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 11/07/2023 13:49:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '''标题''',
  `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '什么都没有' COMMENT '''视频简介''',
  `uid` bigint unsigned NOT NULL COMMENT '''用户ID''',
  `copyright` tinyint(1) NOT NULL COMMENT '''是否为原创''',
  `clicks` bigint DEFAULT '0' COMMENT '''点击量''',
  `status` bigint NOT NULL COMMENT '''审核状态''',
  `partition_id` bigint unsigned DEFAULT NULL COMMENT '''分区ID''',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_video_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_video_title` (`title`) USING BTREE,
  KEY `idx_video_uid` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
