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

 Date: 01/07/2022 10:52:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `v0` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '5', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '6', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '7', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '8', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '9', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '16', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '17', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '18', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '35', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '36', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '38', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '39', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '40', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '41', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '42', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '43', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '44', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '45', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '46', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '47', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '48', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '49', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '50', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '51', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '52', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '54', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '55', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '56', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '57', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '58', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '59', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '60', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '61', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '62', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '63', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '64', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '1', '65', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '1', '1', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '2', '1', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for crontab
-- ----------------------------
DROP TABLE IF EXISTS `crontab`;
CREATE TABLE `crontab` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `group` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `invoke_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '?????????????????????',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'cron???????????????',
  `misfire_policy` tinyint unsigned NOT NULL COMMENT '?????????????????? 0:???????????? 1:????????????',
  `status` tinyint unsigned NOT NULL COMMENT '?????? 0:?????? 1:??????',
  `create_by` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `update_by` bigint unsigned DEFAULT NULL COMMENT '???????????????ID',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of crontab
-- ----------------------------
BEGIN;
INSERT INTO `crontab` (`id`, `name`, `group`, `params`, `invoke_target`, `cron_expression`, `misfire_policy`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES (1, '??????????????????', 'system', NULL, 'checkUserOnline', '0 0 */2 * * *', 1, 1, 1, NULL, NULL, '2022-06-29 11:52:59', '2022-06-29 11:52:59');
INSERT INTO `crontab` (`id`, `name`, `group`, `params`, `invoke_target`, `cron_expression`, `misfire_policy`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES (2, '??????????????????????????????', 'system', NULL, 'checkWdkProjectFileUploadStatus', '0 0 9 25 * *', 1, 1, 1, NULL, NULL, '2022-06-29 11:54:14', '2022-06-29 11:54:14');
COMMIT;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `parent_id` bigint unsigned NOT NULL COMMENT '?????????ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `status` tinyint unsigned NOT NULL COMMENT '???????????? 0:?????? 1:??????',
  `created_by` bigint unsigned NOT NULL COMMENT '?????????',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '?????????',
  `principal_uid` bigint unsigned DEFAULT NULL COMMENT '???????????????ID',
  `principal_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '???????????????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  `deleted_at` datetime DEFAULT NULL COMMENT '???????????????',
  PRIMARY KEY (`id`),
  KEY `parent_id_index` (`parent_id`) USING BTREE,
  KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of dept
-- ----------------------------
BEGIN;
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, '??????????????????', 1, 1, NULL, 0, '', '2022-04-20 12:18:04', '2022-04-20 12:18:04', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, '????????????', 1, 1, NULL, 0, '', '2022-04-20 12:18:31', '2022-04-20 12:18:31', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 2, '????????????', 1, 1, NULL, 0, '', '2022-04-20 12:18:52', '2022-04-20 12:18:52', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 3, '??????????????????', 1, 1, NULL, 0, '', '2022-04-20 12:19:14', '2022-04-20 12:19:14', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 2, '????????????', 1, 1, NULL, 0, '', '2022-04-20 12:20:53', '2022-04-20 12:20:53', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 5, '?????????????????????????????????????????????', 1, 1, NULL, 0, '', '2022-04-20 12:21:54', '2022-04-20 12:21:54', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 6, '????????????', 1, 1, NULL, 0, '', '2022-04-20 12:22:29', '2022-04-20 12:22:29', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 7, '????????????', 1, 1, NULL, 0, '', '2022-04-20 12:23:17', '2022-04-20 12:23:17', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 8, '???????????????', 1, 1, NULL, 0, '', '2022-04-20 12:23:45', '2022-04-20 12:23:45', NULL);
COMMIT;

-- ----------------------------
-- Table structure for login_log
-- ----------------------------
DROP TABLE IF EXISTS `login_log`;
CREATE TABLE `login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `passport` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????IP??????',
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `browser` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '???????????????',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `status` tinyint unsigned DEFAULT NULL COMMENT '???????????? 0:?????? 1:??????',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `time` datetime DEFAULT NULL COMMENT '????????????',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `parent_id` bigint unsigned NOT NULL COMMENT '?????????ID',
  `rule` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `menu_type` tinyint unsigned NOT NULL COMMENT '?????? 0:?????? 1:?????? 2:??????',
  `status` tinyint unsigned NOT NULL COMMENT '???????????? 0:?????? 1:??????',
  `jump_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `is_frame` tinyint unsigned NOT NULL COMMENT '???????????? 1??? 0???',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  `deleted_at` datetime DEFAULT NULL COMMENT '???????????????',
  PRIMARY KEY (`id`),
  UNIQUE KEY `rule_index` (`rule`),
  KEY `parent_id_index` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, 'auth', '????????????', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:10:35', '2022-04-08 09:10:35', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, 'auth/menu/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:15:06', '2022-04-08 09:15:06', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 2, 'auth/menu/add', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:15:31', '2022-04-08 09:15:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 2, 'auth/menu/edit', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:18:53', '2022-04-08 09:18:53', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 2, 'auth/menu/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:19:59', '2022-04-08 09:19:59', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 1, 'auth/role/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:25:30', '2022-04-08 09:25:30', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 6, 'auth/role/add', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:27:55', '2022-04-08 09:27:55', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 6, 'auth/role/edit', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:28:52', '2022-04-08 09:28:52', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 6, 'auth/role/setStatus', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:36:14', '2022-04-08 09:36:14', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 6, 'auth/role/setDataScope', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:43:45', '2022-04-08 09:43:45', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 6, 'auth/role/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:51:27', '2022-04-08 09:51:27', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 1, 'auth/dept/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:55:37', '2022-04-08 09:55:37', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 12, 'auth/dept/add', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:57:08', '2022-04-08 09:57:08', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 12, 'auth/dept/edit', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:59:22', '2022-04-08 09:59:22', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 12, 'auth/dept/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:01:24', '2022-04-08 10:01:24', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 1, 'auth/user/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:04:07', '2022-04-08 10:04:07', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, 16, 'auth/user/add', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:05:57', '2022-04-08 10:05:57', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 16, 'auth/user/edit', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:06:58', '2022-04-08 10:06:58', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 16, 'auth/user/resetPwd', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:08:54', '2022-04-08 10:08:54', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 16, 'auth/user/setStatus', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:10:31', '2022-04-08 10:10:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, 16, 'auth/user/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:12:03', '2022-04-08 10:12:03', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, 0, 'monitor', '????????????', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:26:33', '2022-04-08 10:26:33', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, 22, 'monitor/userOnline/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:29:20', '2022-04-08 10:29:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, 23, 'monitor/userOnline/forceLogout', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:30:50', '2022-04-08 10:30:50', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (25, 22, 'monitor/loginLog/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:34:26', '2022-04-08 10:34:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (26, 25, 'monitor/loginLog/delete', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:35:26', '2022-04-08 10:35:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (27, 25, 'monitor/loginLog/clear', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:37:43', '2022-04-08 10:37:43', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (28, 22, 'monitor/operLog/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:04:29', '2022-04-18 14:04:29', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (29, 28, 'monitor/operLog/delete', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:05:28', '2022-04-18 14:05:28', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (30, 28, 'monitor/operLog/clear', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:06:21', '2022-04-18 14:06:21', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (31, 22, 'monitor/server_monitor/info', '??????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (32, 22, 'monitor/crontab/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (33, 32, 'monitor/crontab/add', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (34, 32, 'monitor/crontab/edit', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (35, 32, 'monitor/crontab/start', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (36, 32, 'monitor/crontab/stop', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (37, 32, 'monitor/crontab/run', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (38, 32, 'monitor/crontab/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (39, 0, 'wdk', '?????????', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:17:10', '2022-04-24 10:17:10', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (40, 39, 'wdk/project/list', '????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:18:56', '2022-04-24 10:18:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (41, 40, 'wdk/project/add', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:20:38', '2022-04-24 10:20:38', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (42, 40, 'wdk/project/edit', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:22:56', '2022-04-24 10:22:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (43, 40, 'wdk/project/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:24:01', '2022-04-24 10:24:01', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (44, 40, 'wdk/project/export', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:24:01', '2022-04-24 10:24:01', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (45, 39, 'wdk/reportTypeCfg/list', '????????????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:14:16', '2022-04-25 10:14:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (46, 45, 'wdk/reportTypeCfg/add', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:20:31', '2022-04-25 10:20:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (47, 45, 'wdk/reportTypeCfg/edit', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:24:08', '2022-04-25 10:24:08', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (48, 45, 'wdk/reportTypeCfg/delete', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:26:56', '2022-04-25 10:26:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (49, 39, 'wdk/attachment/record', '??????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:15:17', '2022-04-26 16:15:17', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (50, 49, 'wdk/attachment/add', '????????????', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:16:16', '2022-04-26 16:16:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (51, 49, 'wdk/attachment/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:16:16', '2022-04-26 16:16:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (52, 39, 'wdk/service/record', '??????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:18:13', '2022-04-26 16:18:13', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (53, 52, 'wdk/service/add', '??????????????????', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:19:04', '2022-04-26 16:19:04', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (54, 52, 'wdk/service/delete', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:19:04', '2022-04-26 16:19:04', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (55, 39, 'wdk/file/record', '??????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:46:40', '2022-05-01 12:46:40', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (56, 55, 'wdk/file/add', '????????????', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:47:52', '2022-05-01 12:47:52', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (57, 39, 'wdk/report/record', '??????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:53:20', '2022-05-01 12:53:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (58, 57, 'wdk/report/add', '????????????', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:54:36', '2022-05-01 12:54:36', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (59, 57, 'wdk/report/delete', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:54:36', '2022-05-01 12:54:36', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (60, 39, 'wdk/reportAudit/uploadAuditList', '????????????????????????', NULL, 1, 1, NULL, 0, ' sys_admin', NULL, '2022-05-22 19:06:47', '2022-05-22 19:06:47', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (61, 60, 'wdk/reportAudit/rescindAudit', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-22 19:09:32', '2022-05-22 19:09:32', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (62, 39, 'wdk/reportAudit/list', '??????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 01:31:59', '2022-05-23 01:31:59', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (63, 62, 'wdk/reportAudit/audit', '????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 10:53:10', '2022-05-23 10:53:10', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (64, 39, 'wdk/report/excellenceList', '??????????????????', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 11:04:20', '2022-05-23 11:04:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (65, 64, 'wdk/report/chooseExcellence', '??????????????????', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 11:06:13', '2022-05-23 11:06:13', NULL);
COMMIT;

-- ----------------------------
-- Table structure for oper_log
-- ----------------------------
DROP TABLE IF EXISTS `oper_log`;
CREATE TABLE `oper_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `req_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `oper_type` tinyint unsigned DEFAULT NULL COMMENT '???????????? 0:?????? 1:???????????? 2:???????????????',
  `oper_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `dept_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `req_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '??????URL',
  `oper_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????IP??????',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `req_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '????????????',
  `json_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '????????????',
  `status` tinyint unsigned DEFAULT NULL COMMENT '???????????? 0:?????? 1:??????',
  `time` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of oper_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `status` tinyint unsigned NOT NULL COMMENT '???????????? 0:?????? 1:??????',
  `data_scope` tinyint unsigned NOT NULL COMMENT '???????????? 1:?????????????????? 2:????????????????????? 3:????????????????????? 4:?????????????????????????????? 5:?????????????????????',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `created_at`, `updated_at`) VALUES (1, '???????????????', 1, 3, '', '2022-04-19 14:30:16', '2022-04-19 14:30:16');
COMMIT;

-- ----------------------------
-- Table structure for role_dept
-- ----------------------------
DROP TABLE IF EXISTS `role_dept`;
CREATE TABLE `role_dept` (
  `role_id` bigint unsigned NOT NULL COMMENT '??????ID',
  `dept_id` bigint unsigned NOT NULL COMMENT '??????ID',
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of role_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `passport` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????',
  `salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '?????????',
  `realname` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `dept_id` bigint unsigned NOT NULL COMMENT '??????ID',
  `gender` tinyint unsigned NOT NULL COMMENT '?????? 1: ??? 2: ???',
  `status` tinyint unsigned NOT NULL COMMENT '?????? 0:?????? 1:??????',
  `is_admin` tinyint unsigned NOT NULL COMMENT '????????????????????? 0:??? 1:???',
  `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '????????????',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '?????????',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '??????????????????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  `deleted_at` datetime DEFAULT NULL COMMENT '???????????????',
  PRIMARY KEY (`id`),
  UNIQUE KEY `passport_index` (`passport`),
  UNIQUE KEY `mobile_index` (`mobile`),
  KEY `realname_index` (`realname`),
  KEY `nickname_index` (`nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'sadmin', 'b0167dd859b8a70478da36238b0b3e05', 'CqDQa4THP1', '???????????????', '???????????????', 4, 1, 1, 1, 'user/avatar/logo.png', '17364814710', '382185882@qq.com', NULL, '::1', '2022-05-23 00:11:46', '2022-03-25 17:41:40', '2022-03-25 17:41:40', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'liusuxian', 'b6d5d1dc1f39f8a10bfdb53bf78c895f', 'JKPT7WSeG0', '?????????', '', 4, 1, 1, 1, NULL, '15108274735', '', '', NULL, NULL, '2022-04-25 17:21:17', '2022-04-25 17:21:17', NULL);
COMMIT;

-- ----------------------------
-- Table structure for user_online
-- ----------------------------
DROP TABLE IF EXISTS `user_online`;
CREATE TABLE `user_online` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `uuid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????token',
  `passport` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `browser` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????ip',
  `time` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`),
  UNIQUE KEY `token_index` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user_online
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_attachment_file
-- ----------------------------
DROP TABLE IF EXISTS `wdk_attachment_file`;
CREATE TABLE `wdk_attachment_file` (
  `id` bigint unsigned NOT NULL COMMENT '??????????????????ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '?????????',
  `origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????url',
  `pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf??????url',
  KEY `id_index` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_attachment_file
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_attachment_record
-- ----------------------------
DROP TABLE IF EXISTS `wdk_attachment_record`;
CREATE TABLE `wdk_attachment_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????????????????ID',
  `project_id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`,`project_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_attachment_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_file
-- ----------------------------
DROP TABLE IF EXISTS `wdk_file`;
CREATE TABLE `wdk_file` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '????????????ID',
  `project_id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `type` tinyint unsigned NOT NULL COMMENT '???????????? 0:?????????????????? 1:????????????????????? 2:???????????? 3:???????????? 4:???????????? 5:??????????????? 6:??????????????????',
  `create_by` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `create_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????',
  `origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????url',
  `pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf??????url',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`,`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_file
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_project
-- ----------------------------
DROP TABLE IF EXISTS `wdk_project`;
CREATE TABLE `wdk_project` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '???????????????ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `type` tinyint unsigned NOT NULL COMMENT '???????????? 0:???????????? 1:??????',
  `origin` tinyint unsigned NOT NULL COMMENT '???????????? 0:?????? 1:???????????? 2:????????? 3:????????? 4:?????? 5:?????? 6:?????? 7:??????',
  `step` tinyint unsigned NOT NULL COMMENT '???????????? 0:????????? 1:???????????? 2:??????????????? 3:?????????-???????????? 4:?????????-????????????????????? 5:?????????-?????????????????? 6:?????????-?????????????????? 7:?????????-?????????????????? 8:?????????-?????????????????? 9:?????????-?????????????????? 30:???????????? 31:??????',
  `file_upload_status` tinyint unsigned NOT NULL COMMENT '???????????????????????? 0:?????? 1:?????? 2:?????????',
  `business_type` tinyint unsigned NOT NULL COMMENT '???????????? 0:?????? 1:?????? 2:?????????',
  `contract_status` tinyint unsigned NOT NULL COMMENT '???????????? 0:?????? 1:?????? 2:??????',
  `contract_sum` double unsigned NOT NULL COMMENT '????????????',
  `deep_culture` tinyint unsigned NOT NULL COMMENT '??????????????? 0:??? 1:???',
  `status` tinyint unsigned NOT NULL COMMENT '???????????? 0:????????? 1:?????? 2:???????????? 3:????????? 4:????????? 5:????????????',
  `entrust_company` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????',
  `sign_company` tinyint unsigned NOT NULL COMMENT '?????????????????? 0:??????????????????????????????????????? 1:????????????????????????????????????????????? 2:?????????????????????????????????????????????',
  `principal_uid` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `principal_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????',
  `dept_id` bigint unsigned NOT NULL COMMENT '??????????????????ID',
  `dept_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????????????????',
  `region` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????(???/???/??????)',
  `start_time` date NOT NULL COMMENT '??????????????????',
  `end_time` date NOT NULL COMMENT '??????????????????',
  `create_by` bigint unsigned NOT NULL COMMENT '?????????????????????ID',
  `create_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '?????????????????????',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '?????????????????????ID',
  `updated_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '?????????????????????',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `file_upload_last_time` datetime DEFAULT NULL COMMENT '????????????????????????????????????',
  `created_at` datetime DEFAULT NULL COMMENT '??????????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '??????????????????',
  `deleted_at` datetime DEFAULT NULL COMMENT '?????????????????????',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`) USING BTREE,
  KEY `type_index` (`type`) USING BTREE,
  KEY `origin_index` (`origin`) USING BTREE,
  KEY `step_index` (`step` DESC) USING BTREE,
  KEY `file_upload_status_index` (`file_upload_status` DESC) USING BTREE,
  KEY `business_type_index` (`business_type`) USING BTREE,
  KEY `contract_status_index` (`contract_status`) USING BTREE,
  KEY `deep_culture_index` (`deep_culture`) USING BTREE,
  KEY `status_index` (`status`) USING BTREE,
  KEY `sign_company_index` (`sign_company`) USING BTREE,
  KEY `dept_id_index` (`dept_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_project
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_project_businessforms
-- ----------------------------
DROP TABLE IF EXISTS `wdk_project_businessforms`;
CREATE TABLE `wdk_project_businessforms` (
  `project_id` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `business_forms` tinyint unsigned NOT NULL COMMENT '?????? 0:?????? 1:????????? 2:?????? 3:????????? 4:?????? 5:?????? 6:?????? 7:?????? 8:?????? 9:?????? 10:?????? 11:????????? 12:????????? 13:?????? 14:??????????????? 15:?????? 16:???????????? 17:???????????? 18:?????? 19:??????????????? 20:??????',
  PRIMARY KEY (`project_id`,`business_forms`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_project_businessforms
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_report
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report`;
CREATE TABLE `wdk_report` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '??????ID',
  `project_id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `project_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????????????????',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `create_by` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `create_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????',
  `rescind` tinyint unsigned NOT NULL COMMENT '??????????????? 0:??? 1:???',
  `audit_status` tinyint unsigned NOT NULL COMMENT '???????????? 0:????????? 1:????????? 2:????????? 3:???????????????????????????',
  `excellence` tinyint unsigned NOT NULL COMMENT '????????????????????? 0:??????????????????????????? 1:???????????????????????? 2:???????????????????????????',
  `audit_time` datetime DEFAULT NULL COMMENT '????????????',
  `origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????url',
  `pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf??????url',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`,`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_audit
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_audit`;
CREATE TABLE `wdk_report_audit` (
  `id` bigint unsigned NOT NULL COMMENT '??????ID',
  `audit_uid` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `auditor_type` tinyint unsigned NOT NULL COMMENT '??????????????? 0:????????? 1:?????????',
  `project_id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `audit_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????',
  `rescind` tinyint unsigned NOT NULL COMMENT '??????????????? 0:??? 1:???',
  `preaudit_status` tinyint unsigned NOT NULL COMMENT '??????????????????????????? 0:??? 1:???',
  `status` tinyint unsigned NOT NULL COMMENT '???????????? 0:????????? 1:????????? 2:?????????',
  `excellence` tinyint unsigned NOT NULL COMMENT '?????????????????????????????? 0:??????????????????????????? 1:???????????????????????????',
  `audit_time` datetime DEFAULT NULL COMMENT '????????????',
  `audit_opinion` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '????????????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`,`audit_uid`,`auditor_type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_audit
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_audit_cfg
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_audit_cfg`;
CREATE TABLE `wdk_report_audit_cfg` (
  `id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `audit_uid` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????????????????',
  `audit_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????',
  PRIMARY KEY (`id`,`audit_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_audit_cfg
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_audit_type
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_audit_type`;
CREATE TABLE `wdk_report_audit_type` (
  `id` bigint unsigned NOT NULL COMMENT '??????ID',
  `audit_uid` bigint unsigned NOT NULL COMMENT '???????????????ID',
  `auditor_type` tinyint unsigned NOT NULL COMMENT '??????????????? 0:????????? 1:?????????',
  `type_id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????????????????',
  PRIMARY KEY (`id`,`audit_uid`,`auditor_type`,`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_audit_type
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_cfg
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_cfg`;
CREATE TABLE `wdk_report_cfg` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '????????????ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????????????????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  `updated_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_cfg
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_type
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_type`;
CREATE TABLE `wdk_report_type` (
  `id` bigint unsigned NOT NULL COMMENT '??????ID',
  `type_id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????????????????',
  PRIMARY KEY (`id`,`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_type
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_service_photo
-- ----------------------------
DROP TABLE IF EXISTS `wdk_service_photo`;
CREATE TABLE `wdk_service_photo` (
  `id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????',
  `url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '????????????url',
  KEY `id_index` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_service_photo
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wdk_service_record
-- ----------------------------
DROP TABLE IF EXISTS `wdk_service_record`;
CREATE TABLE `wdk_service_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '????????????ID',
  `project_id` bigint unsigned NOT NULL COMMENT '????????????ID',
  `service_time` date NOT NULL COMMENT '????????????',
  `xch_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '??????????????????',
  `xch_origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '???????????????url',
  `xch_pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf?????????url',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '??????',
  `created_at` datetime DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`,`project_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_service_record
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
