/*
 Navicat Premium Data Transfer

 Source Server         : 81.71.31.250
 Source Server Type    : MySQL
 Source Server Version : 80028 (8.0.28)
 Source Host           : 81.71.31.250:3306
 Source Schema         : giligili_user

 Target Server Type    : MySQL
 Target Server Version : 80028 (8.0.28)
 File Encoding         : 65001

 Date: 31/05/2023 12:39:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL,
  `created_time` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3),
  `updated_time` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_time` datetime(3) DEFAULT NULL,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '''用户名''',
  `email` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '''邮箱''',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '''密码''',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://i1.hdslb.com/bfs/face/2c72223afa74b0036daee60cd99c069760b653df.jpg@240w_240h_1c_1s_!web-avatar-space-header.avif' COMMENT '''头像''',
  `space_cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://i0.hdslb.com/bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png@2560w_400h_100q_1o.webp' COMMENT '''空间封面''',
  `gender` tinyint NOT NULL DEFAULT '0' COMMENT '''性别:0未知、1男、3女''',
  `birthday` datetime(3) NOT NULL DEFAULT '1970-01-01 00:00:00.000' COMMENT '''生日''',
  `sign` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '这个人很懒，什么都没有留下' COMMENT '''个性签名''',
  `client_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '''客户端IP''',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0',
  `role` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_deleted_at` (`deleted_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
