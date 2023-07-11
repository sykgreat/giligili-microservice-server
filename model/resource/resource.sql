/*
 Navicat Premium Data Transfer

 Source Server         : 81.71.31.250
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 81.71.31.250:3306
 Source Schema         : giligili_resource

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 11/07/2023 13:22:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `vid` bigint unsigned DEFAULT NULL COMMENT '''所属视频''',
  `uid` bigint unsigned DEFAULT NULL COMMENT '''所属用户''',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''分P使用的标题''',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''视频链接''',
  `duration` double DEFAULT '0' COMMENT '''视频时长''',
  `status` bigint NOT NULL COMMENT '''审核状态''',
  `quality` bigint DEFAULT NULL COMMENT '''视频最大质量''',
  `original_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''原始mp4链接''',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_resource_status` (`status`) USING BTREE,
  KEY `idx_resource_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_resource_vid` (`vid`) USING BTREE,
  KEY `idx_resource_uid` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
