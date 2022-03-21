/*
 Navicat Premium Data Transfer

 Source Server         : LsxMysql
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : lczx

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 21/03/2022 13:54:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for branch
-- ----------------------------
DROP TABLE IF EXISTS `branch`;
CREATE TABLE `branch` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部门名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of branch
-- ----------------------------
BEGIN;
INSERT INTO `branch` (`id`, `name`) VALUES (1, '市场运营部');
INSERT INTO `branch` (`id`, `name`) VALUES (2, '技术支持部');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `passport` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `realname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `gender` int unsigned NOT NULL COMMENT '性别 0: 未设置 1: 男 2: 女',
  `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '头像地址',
  `telno` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `status` int unsigned NOT NULL COMMENT '状态 0:启用 1:禁用',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `update_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `passport`, `password`, `realname`, `nickname`, `gender`, `avatar`, `telno`, `status`, `create_at`, `update_at`) VALUES (1, 'liusuxian', 'lsx19890329', '刘苏贤', NULL, 1, NULL, NULL, 0, '2022-03-18 13:20:00', '2022-03-18 13:20:00');
COMMIT;

-- ----------------------------
-- Table structure for user_ext
-- ----------------------------
DROP TABLE IF EXISTS `user_ext`;
CREATE TABLE `user_ext` (
  `id` int unsigned NOT NULL COMMENT '用户ID',
  `branch_id` int unsigned NOT NULL COMMENT '部门ID',
  `zsk_role_id` int unsigned DEFAULT NULL COMMENT '知识库角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员',
  `wdk_role_id` int unsigned DEFAULT NULL COMMENT '文档库角色ID 0: 默认普通用户 1000: 超级管理员 900: 普通管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_ext
-- ----------------------------
BEGIN;
INSERT INTO `user_ext` (`id`, `branch_id`, `zsk_role_id`, `wdk_role_id`) VALUES (1, 1, NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
