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

 Date: 23/05/2022 00:13:48
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
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '1', '1', NULL, NULL, NULL, NULL);
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '2', '1', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '3', '2', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '41', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '42', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '43', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '44', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '45', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '46', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '47', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '48', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '49', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '50', 'All', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `parent_id` bigint unsigned NOT NULL COMMENT '父部门ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部门名称',
  `status` tinyint unsigned NOT NULL COMMENT '部门状态 0:停用 1:正常',
  `created_by` bigint unsigned NOT NULL COMMENT '创建人',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '修改人',
  `principal_uid` bigint unsigned DEFAULT NULL COMMENT '负责人用户ID',
  `principal_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人姓名',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `parent_id_index` (`parent_id`) USING BTREE,
  KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of dept
-- ----------------------------
BEGIN;
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (1, 0, '绿城服务集团', 1, 1, NULL, 0, '', '2022-04-20 12:18:04', '2022-04-20 12:18:04', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (2, 1, '咨询集团', 1, 1, NULL, 0, '', '2022-04-20 12:18:31', '2022-04-20 12:18:31', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (3, 2, '职能中心', 1, 1, NULL, 0, '', '2022-04-20 12:18:52', '2022-04-20 12:18:52', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (4, 3, '运营管理中心', 1, 1, NULL, 0, '', '2022-04-20 12:19:14', '2022-04-20 12:19:14', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (5, 2, '分子公司', 1, 1, NULL, 0, '', '2022-04-20 12:20:53', '2022-04-20 12:20:53', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (6, 5, '浙江幸福绿城房地产咨询有限公司', 1, 1, NULL, 0, '', '2022-04-20 12:21:54', '2022-04-20 12:21:54', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (7, 6, '业务部门', 1, 1, NULL, 0, '', '2022-04-20 12:22:29', '2022-04-20 12:22:29', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (8, 7, '物业咨询', 1, 1, NULL, 0, '', '2022-04-20 12:23:17', '2022-04-20 12:23:17', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (9, 8, '技术支持部', 1, 1, NULL, 0, '', '2022-04-20 12:23:45', '2022-04-20 12:23:45', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `create_at`, `update_at`, `deleted_at`) VALUES (10, 4, '测试', 1, 1, 1, 2, '刘苏贤', '2022-05-21 03:46:52', '2022-05-21 04:07:26', NULL);
COMMIT;

-- ----------------------------
-- Table structure for login_log
-- ----------------------------
DROP TABLE IF EXISTS `login_log`;
CREATE TABLE `login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `passport` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录账号',
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录IP地址',
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录地点',
  `browser` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '浏览器类型',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作系统',
  `status` tinyint unsigned DEFAULT NULL COMMENT '登录状态 0:失败 1:成功',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '提示消息',
  `time` datetime DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录模块',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '规则ID',
  `parent_id` bigint unsigned NOT NULL COMMENT '父规则ID',
  `rule` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限规则',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '条件',
  `menu_type` tinyint unsigned NOT NULL COMMENT '类型 0:目录 1:菜单 2:按钮',
  `status` tinyint unsigned NOT NULL COMMENT '菜单状态 0:停用 1:正常',
  `jump_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转路由',
  `is_frame` tinyint unsigned NOT NULL COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '所属模块',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `rule_index` (`rule`),
  KEY `parent_id_index` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (1, 0, 'auth', '权限管理', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:10:35', '2022-04-08 09:10:35', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (2, 1, 'auth/menu/list', '菜单管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:15:06', '2022-04-08 09:15:06', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (3, 2, 'auth/menu/add', '添加菜单', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:15:31', '2022-04-08 09:15:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (4, 2, 'auth/menu/edit', '修改菜单', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:18:53', '2022-04-08 09:18:53', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (5, 2, 'auth/menu/delete', '删除菜单', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:19:59', '2022-04-08 09:19:59', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (6, 1, 'auth/role/list', '角色管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:25:30', '2022-04-08 09:25:30', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (7, 6, 'auth/role/add', '添加角色', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:27:55', '2022-04-08 09:27:55', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (8, 6, 'auth/role/edit', '修改角色', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:28:52', '2022-04-08 09:28:52', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (9, 6, 'auth/role/setStatus', '设置角色状态', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:36:14', '2022-04-08 09:36:14', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (10, 6, 'auth/role/setDataScope', '设置数据权限', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:43:45', '2022-04-08 09:43:45', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (11, 6, 'auth/role/delete', '删除角色', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:51:27', '2022-04-08 09:51:27', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (12, 1, 'auth/dept/list', '部门管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:55:37', '2022-04-08 09:55:37', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (13, 12, 'auth/dept/add', '添加部门', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:57:08', '2022-04-08 09:57:08', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (14, 12, 'auth/dept/edit', '修改部门', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:59:22', '2022-04-08 09:59:22', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (15, 12, 'auth/dept/delete', '删除部门', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:01:24', '2022-04-08 10:01:24', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (16, 1, 'auth/user/list', '用户管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:04:07', '2022-04-08 10:04:07', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (17, 16, 'auth/user/add', '添加用户', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:05:57', '2022-04-08 10:05:57', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (18, 16, 'auth/user/edit', '修改用户', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:06:58', '2022-04-08 10:06:58', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (19, 16, 'auth/user/resetPwd', '密码重置', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:08:54', '2022-04-08 10:08:54', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (20, 16, 'auth/user/setStatus', '设置用户状态', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:10:31', '2022-04-08 10:10:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (21, 16, 'auth/user/delete', '删除用户', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:12:03', '2022-04-08 10:12:03', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (22, 0, 'monitor', '系统监控', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:26:33', '2022-04-08 10:26:33', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (23, 22, 'monitor/userOnline/list', '在线用户', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:29:20', '2022-04-08 10:29:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (24, 23, 'monitor/userOnline/forceLogout', '强制退出', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:30:50', '2022-04-08 10:30:50', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (25, 22, 'monitor/loginLog/list', '登录日志', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:34:26', '2022-04-08 10:34:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (26, 25, 'monitor/loginLog/delete', '删除登录日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:35:26', '2022-04-08 10:35:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (27, 25, 'monitor/loginLog/clear', '清除登录日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:37:43', '2022-04-08 10:37:43', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (28, 22, 'monitor/operLog/list', '操作日志', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:04:29', '2022-04-18 14:04:29', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (29, 28, 'monitor/operLog/delete', '删除操作日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:05:28', '2022-04-18 14:05:28', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (30, 28, 'monitor/operLog/clear', '清除操作日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:06:21', '2022-04-18 14:06:21', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (31, 22, 'monitor/server_monitor/info', '服务监控信息', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (32, 0, 'wdk', '文档库', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:17:10', '2022-04-24 10:17:10', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (33, 32, 'wdk/project/list', '项目管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:18:56', '2022-04-24 10:18:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (34, 33, 'wdk/project/add', '添加项目', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:20:38', '2022-04-24 10:20:38', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (35, 33, 'wdk/project/edit', '修改项目', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:22:56', '2022-04-24 10:22:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (36, 33, 'wdk/project/delete', '删除项目', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:24:01', '2022-04-24 10:24:01', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (37, 32, 'wdk/reportTypeCfg/list', '报告类型配置管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:14:16', '2022-04-25 10:14:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (38, 37, 'wdk/reportTypeCfg/add', '添加报告类型', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:20:31', '2022-04-25 10:20:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (39, 37, 'wdk/reportTypeCfg/edit', '修改报告类型', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:24:08', '2022-04-25 10:24:08', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (40, 37, 'wdk/reportTypeCfg/delete', '删除报告类型', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:26:56', '2022-04-25 10:26:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (41, 32, 'wdk/attachment/record', '附件记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:15:17', '2022-04-26 16:15:17', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (42, 41, 'wdk/attachment/add', '上传附件', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:16:16', '2022-04-26 16:16:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (43, 32, 'wdk/service/record', '服务记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:18:13', '2022-04-26 16:18:13', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (44, 43, 'wdk/service/add', '添加服务记录', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:19:04', '2022-04-26 16:19:04', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (45, 32, 'wdk/file/record', '文件记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:46:40', '2022-05-01 12:46:40', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (46, 45, 'wdk/file/add', '上传文件', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:47:52', '2022-05-01 12:47:52', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (47, 32, 'wdk/report/record', '报告记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:53:20', '2022-05-01 12:53:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (48, 47, 'wdk/report/add', '上传报告', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:54:36', '2022-05-01 12:54:36', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (49, 32, 'wdk/reportAudit/uploadAuditList', '报告上传审核管理', NULL, 1, 1, NULL, 0, ' sys_admin', NULL, '2022-05-22 19:06:47', '2022-05-22 19:06:47', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (50, 49, 'wdk/reportAudit/rescindAudit', '报告撤销审核', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-22 19:09:32', '2022-05-22 19:09:32', NULL);
COMMIT;

-- ----------------------------
-- Table structure for oper_log
-- ----------------------------
DROP TABLE IF EXISTS `oper_log`;
CREATE TABLE `oper_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模块标题',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '方法名称',
  `req_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求方式',
  `oper_type` tinyint unsigned DEFAULT NULL COMMENT '操作类别 0:其它 1:后台用户 2:手机端用户',
  `oper_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作人员',
  `dept_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '部门名称',
  `req_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '请求URL',
  `oper_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作IP地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作地点',
  `req_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '返回参数',
  `status` tinyint unsigned DEFAULT NULL COMMENT '操作状态 0:异常 1:正常',
  `time` datetime DEFAULT NULL COMMENT '操作时间',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `status` tinyint unsigned NOT NULL COMMENT '角色状态 0:停用 1:正常',
  `data_scope` tinyint unsigned NOT NULL COMMENT '数据范围 1:全部数据权限 2:自定义数据权限 3:本部门数据权限 4:本部门及以下数据权限 5:仅本人数据权限',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `create_at`, `update_at`) VALUES (1, '超级管理员', 1, 3, '', '2022-04-19 14:30:16', '2022-04-19 14:30:16');
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `create_at`, `update_at`) VALUES (2, '测试角色', 1, 3, '', '2022-05-21 04:01:56', '2022-05-23 00:12:22');
COMMIT;

-- ----------------------------
-- Table structure for role_dept
-- ----------------------------
DROP TABLE IF EXISTS `role_dept`;
CREATE TABLE `role_dept` (
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `dept_id` bigint unsigned NOT NULL COMMENT '部门ID',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `passport` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '加密盐',
  `realname` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `dept_id` bigint unsigned NOT NULL COMMENT '部门ID',
  `gender` tinyint unsigned NOT NULL COMMENT '性别 1: 男 2: 女',
  `status` tinyint unsigned NOT NULL COMMENT '状态 0:禁用 1:启用',
  `is_admin` tinyint unsigned NOT NULL COMMENT '是否后台管理员 0:否 1:是',
  `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '头像地址',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '最后登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `passport_index` (`passport`),
  UNIQUE KEY `mobile_index` (`mobile`),
  KEY `realname_index` (`realname`),
  KEY `nickname_index` (`nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `create_at`, `update_at`, `deleted_at`) VALUES (1, 'sadmin', 'b0167dd859b8a70478da36238b0b3e05', 'CqDQa4THP1', '超级管理员', '超级管理员', 4, 1, 1, 1, 'https://lczx-sys.oss-cn-hangzhou.aliyuncs.com/user/avatar/logo.png', '17364814710', '382185882@qq.com', NULL, '::1', '2022-05-23 00:11:46', '2022-03-25 17:41:40', '2022-03-25 17:41:40', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `create_at`, `update_at`, `deleted_at`) VALUES (2, 'liusuxian', 'b6d5d1dc1f39f8a10bfdb53bf78c895f', 'JKPT7WSeG0', '刘苏贤', '', 4, 1, 1, 1, NULL, '15108274735', '', '', NULL, NULL, '2022-04-25 17:21:17', '2022-04-25 17:21:17', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `create_at`, `update_at`, `deleted_at`) VALUES (3, 'ceshi123', 'bed32e4ac7a83b854d6e1b279915352a', 'nw3VIpRP8y', '测试用户', '', 10, 1, 1, 0, NULL, '15110100101', '', '', '::1', '2022-05-21 04:07:49', '2022-05-21 04:02:36', '2022-05-21 04:02:36', NULL);
COMMIT;

-- ----------------------------
-- Table structure for user_online
-- ----------------------------
DROP TABLE IF EXISTS `user_online`;
CREATE TABLE `user_online` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `uuid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户token',
  `passport` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录账号',
  `browser` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '浏览器类型',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '操作系统',
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录ip',
  `time` datetime DEFAULT NULL COMMENT '登录时间',
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
  `id` bigint unsigned NOT NULL COMMENT '附件上传记录ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '附件名',
  `origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原始附件url',
  `pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf附件url',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '附件上传记录ID',
  `project_id` bigint unsigned NOT NULL COMMENT '所属项目ID',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '上传文件ID',
  `project_id` bigint unsigned NOT NULL COMMENT '所属项目ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件名称',
  `type` tinyint unsigned NOT NULL COMMENT '文件类型 0:合同扫描文件 1:年度服务计划书 2:总结报告 3:项目移交 4:复盘报告 5:文件签收单 6:满意度调查表',
  `create_by` bigint unsigned NOT NULL COMMENT '上传者用户ID',
  `create_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '上传者姓名',
  `origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原始文件url',
  `pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf文件url',
  `create_at` datetime DEFAULT NULL COMMENT '上传时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '文档库项目ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目名称',
  `type` tinyint unsigned NOT NULL COMMENT '项目性质 0:蓝绿体系 1:非绿',
  `origin` tinyint unsigned NOT NULL COMMENT '项目来源 0:物业公司 1:分子公司 2:老客户 3:自拓',
  `step` tinyint unsigned NOT NULL COMMENT '项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘',
  `file_upload_status` tinyint unsigned NOT NULL COMMENT '项目文件上传状态 0:未传完 1:已传完',
  `business_type` tinyint unsigned NOT NULL COMMENT '业务类型 0:物业 1:专项 2:全过程',
  `contract_status` tinyint unsigned NOT NULL COMMENT '签约状态 0:新签 1:续签',
  `contract_sum` double unsigned NOT NULL COMMENT '合同金额',
  `deep_culture` tinyint unsigned NOT NULL COMMENT '是否为深耕 0:否 1:是',
  `status` tinyint unsigned NOT NULL COMMENT '服务状态 0:服务中 1:暂停 2:提前终止 3:正常结束',
  `entrust_company` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '委托方公司',
  `sign_company` tinyint unsigned NOT NULL COMMENT '我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司',
  `principal_uid` bigint unsigned NOT NULL COMMENT '负责人用户ID',
  `principal_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '负责人姓名',
  `dept_id` bigint unsigned NOT NULL COMMENT '项目所属部门ID',
  `dept_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目所属部门名称',
  `region` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '地区(省/市/区县)',
  `start_time` datetime NOT NULL COMMENT '项目开始时间',
  `end_time` datetime NOT NULL COMMENT '项目结束时间',
  `create_by` bigint unsigned NOT NULL COMMENT '项目创建者用户ID',
  `create_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目创建者姓名',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '项目修改者用户ID',
  `updated_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '项目修改者姓名',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `create_at` datetime DEFAULT NULL COMMENT '项目创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '项目更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '项目软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_project
-- ----------------------------
BEGIN;
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (1, '测试项目1', 0, 0, 0, 0, 0, 0, 20.11, 0, 0, '委托公司123', 0, 3, '测试用户', 10, '测试', '浙江省/杭州市/西湖区', '2022-05-01 00:00:00', '2022-05-31 23:59:59', 1, '超级管理员', NULL, NULL, '', '2022-05-21 04:03:31', '2022-05-21 04:03:31', NULL);
COMMIT;

-- ----------------------------
-- Table structure for wdk_project_businessforms
-- ----------------------------
DROP TABLE IF EXISTS `wdk_project_businessforms`;
CREATE TABLE `wdk_project_businessforms` (
  `project_id` bigint unsigned NOT NULL COMMENT '文档库项目ID',
  `business_forms` tinyint unsigned NOT NULL COMMENT '业态 0:住宅-综合性住宅 1:住宅-高层住宅 2:住宅-多层住宅 3:住宅-联排别墅 4:住宅-独立式住宅 5:非住宅-办公用房 6:非住宅-场馆（体育馆、博物馆、机场） 7:非住宅-综合类 8:非住宅-学校 9:非住宅-工业园 10:非住宅-医院 11:非住宅-商场 12:非住宅-商铺 13:非住宅-车位 14:非住宅-酒店式公寓 15:非住宅-仓储 16:非住宅-经营用房 17:非住宅-物业用房 18:非住宅-配套用房 19:非住宅-车库 20:非住宅-会所 21:非住宅-临停车位 22:非住宅-非机动车位 23:非住宅-社会福利院 24:非住宅-旅游景区',
  PRIMARY KEY (`project_id`,`business_forms`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_project_businessforms
-- ----------------------------
BEGIN;
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (1, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (1, 2);
COMMIT;

-- ----------------------------
-- Table structure for wdk_report
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report`;
CREATE TABLE `wdk_report` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '报告ID',
  `project_id` bigint unsigned NOT NULL COMMENT '所属项目ID',
  `project_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '所属项目名称',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报告名称',
  `create_by` bigint unsigned NOT NULL COMMENT '上传者用户ID',
  `create_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '上传者姓名',
  `rescind` tinyint unsigned NOT NULL COMMENT '是否已撤销 0:否 1:是',
  `audit_status` tinyint unsigned NOT NULL COMMENT '审核状态 0:未通过 1:审核中 2:已通过 3:后台管理员自动通过',
  `excellence` tinyint unsigned NOT NULL COMMENT '是否是优秀报告 0:未被评选为优秀报告 1:被推荐为优秀报告 2:已被评选为优秀报告',
  `audit_time` datetime DEFAULT NULL COMMENT '审核时间',
  `origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原始文件url',
  `pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf文件url',
  `create_at` datetime DEFAULT NULL COMMENT '上传时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `create_at`, `update_at`) VALUES (1, 1, '测试项目1', '一季度市场拓展数据分析-A.pptx', 3, '测试用户', 0, 1, 0, NULL, 'https://lczx-sys.oss-accelerate.aliyuncs.com/wdk/report/ck4uzzsy45bcaeglfvchci.pptx', 'https://lczx-sys.oss-accelerate.aliyuncs.com/wdk/report/ck4uzzsy45bcaeglfvchci.pdf', '2022-05-21 04:08:35', '2022-05-21 04:08:35');
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_audit
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_audit`;
CREATE TABLE `wdk_report_audit` (
  `id` bigint unsigned NOT NULL COMMENT '报告ID',
  `audit_uid` bigint unsigned NOT NULL COMMENT '审核员用户ID',
  `auditor_type` tinyint unsigned NOT NULL COMMENT '审核员类型 0:负责人 1:专家组',
  `project_id` bigint unsigned NOT NULL COMMENT '所属项目ID',
  `audit_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '审核员姓名',
  `rescind` tinyint unsigned NOT NULL COMMENT '是否已撤销 0:否 1:是',
  `status` tinyint unsigned NOT NULL COMMENT '审核状态 0:未通过 1:审核中 2:已通过',
  `excellence` tinyint unsigned NOT NULL COMMENT '是否被推荐为优秀报告 0:未被推荐为优秀报告 1:已被推荐为优秀报告',
  `audit_time` datetime DEFAULT NULL COMMENT '审核时间',
  `audit_opinion` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '审核意见',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`audit_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_audit
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report_audit` (`id`, `audit_uid`, `auditor_type`, `project_id`, `audit_name`, `rescind`, `status`, `excellence`, `audit_time`, `audit_opinion`, `create_at`, `update_at`) VALUES (1, 1, 1, 1, '超级管理员', 0, 1, 0, NULL, NULL, '2022-05-21 04:08:35', '2022-05-21 04:08:35');
INSERT INTO `wdk_report_audit` (`id`, `audit_uid`, `auditor_type`, `project_id`, `audit_name`, `rescind`, `status`, `excellence`, `audit_time`, `audit_opinion`, `create_at`, `update_at`) VALUES (1, 2, 0, 1, '刘苏贤', 0, 1, 0, NULL, NULL, '2022-05-21 04:08:35', '2022-05-21 04:08:35');
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_audit_cfg
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_audit_cfg`;
CREATE TABLE `wdk_report_audit_cfg` (
  `id` bigint unsigned NOT NULL COMMENT '报告类型ID',
  `audit_uid` bigint unsigned NOT NULL COMMENT '审核员用户ID',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报告类型名称',
  `audit_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '审核员姓名',
  PRIMARY KEY (`id`,`audit_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_audit_cfg
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (1, 1, '水电报告', '超级管理员');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (1, 2, '水电报告', '刘苏贤');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (2, 1, '建筑规划报告', '超级管理员');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (2, 2, '建筑规划报告', '刘苏贤');
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_audit_type
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_audit_type`;
CREATE TABLE `wdk_report_audit_type` (
  `id` bigint unsigned NOT NULL COMMENT '报告ID',
  `audit_uid` bigint unsigned NOT NULL COMMENT '审核员用户ID',
  `type_id` bigint unsigned NOT NULL COMMENT '报告类型ID',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报告类型名称',
  PRIMARY KEY (`id`,`audit_uid`,`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_audit_type
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report_audit_type` (`id`, `audit_uid`, `type_id`, `type_name`) VALUES (1, 1, 2, '建筑规划报告');
INSERT INTO `wdk_report_audit_type` (`id`, `audit_uid`, `type_id`, `type_name`) VALUES (1, 2, 1, '水电报告');
INSERT INTO `wdk_report_audit_type` (`id`, `audit_uid`, `type_id`, `type_name`) VALUES (1, 2, 2, '建筑规划报告');
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_cfg
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_cfg`;
CREATE TABLE `wdk_report_cfg` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '报告类型ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报告类型名称',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_cfg
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report_cfg` (`id`, `name`, `create_at`, `update_at`) VALUES (1, '水电报告', '2022-05-21 04:03:55', '2022-05-21 04:03:55');
INSERT INTO `wdk_report_cfg` (`id`, `name`, `create_at`, `update_at`) VALUES (2, '建筑规划报告', '2022-05-21 04:04:10', '2022-05-21 04:04:10');
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_type
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_type`;
CREATE TABLE `wdk_report_type` (
  `id` bigint unsigned NOT NULL COMMENT '报告ID',
  `type_id` bigint unsigned NOT NULL COMMENT '报告类型ID',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报告类型名称',
  PRIMARY KEY (`id`,`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_type
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (1, 1, '水电报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (1, 2, '建筑规划报告');
COMMIT;

-- ----------------------------
-- Table structure for wdk_service_photo
-- ----------------------------
DROP TABLE IF EXISTS `wdk_service_photo`;
CREATE TABLE `wdk_service_photo` (
  `id` bigint unsigned NOT NULL COMMENT '服务记录ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '照片名称',
  `url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '照片文件url',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '服务记录ID',
  `project_id` bigint unsigned NOT NULL COMMENT '所属项目ID',
  `service_time` datetime NOT NULL COMMENT '服务时间',
  `xch_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '行程涵文件名',
  `xch_origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原始行程涵url',
  `xch_pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf行程涵url',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`,`project_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_service_record
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
