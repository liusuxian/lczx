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

 Date: 12/04/2022 11:25:21
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
COMMIT;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父部门ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '部门名称',
  `status` tinyint unsigned DEFAULT NULL COMMENT '部门状态 0:停用 1:正常',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '修改人',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of dept
-- ----------------------------
BEGIN;
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
  `show` tinyint unsigned NOT NULL COMMENT '显示状态 0:隐藏 1:显示',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '路由地址',
  `jump_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转路由',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '组件路径',
  `is_frame` tinyint unsigned NOT NULL COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '所属模块',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `api_path_index` (`rule`),
  KEY `parent_id_index` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (1, 0, 'auth', '权限管理', NULL, 0, 1, 1, 'auth', NULL, 'auth', 0, 'sys_admin', NULL, '2022-04-08 09:10:35', '2022-04-08 09:10:35', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (2, 1, 'auth/menu/list', '菜单管理', NULL, 1, 1, 1, 'list', NULL, 'auth/menu/list', 0, 'sys_admin', NULL, '2022-04-08 09:15:06', '2022-04-08 09:15:06', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (3, 1, 'auth/menu/add', '添加菜单', NULL, 2, 1, 1, 'add', NULL, 'auth/menu/add', 0, 'sys_admin', NULL, '2022-04-08 09:15:31', '2022-04-08 09:15:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (4, 1, 'auth/menu/edit', '修改菜单', NULL, 2, 1, 1, 'edit', NULL, 'auth/menu/edit', 0, 'sys_admin', NULL, '2022-04-08 09:18:53', '2022-04-08 09:18:53', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (5, 1, 'auth/menu/delete', '删除菜单', NULL, 2, 1, 1, 'delete', NULL, 'auth/menu/delete', 0, 'sys_admin', NULL, '2022-04-08 09:19:59', '2022-04-08 09:19:59', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (6, 1, 'auth/role/list', '角色管理', NULL, 1, 1, 1, 'list', NULL, 'auth/role/list', 0, 'sys_admin', NULL, '2022-04-08 09:25:30', '2022-04-08 09:25:30', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (7, 1, 'auth/role/add', '添加角色', NULL, 2, 1, 1, 'add', NULL, 'auth/role/add', 0, 'sys_admin', NULL, '2022-04-08 09:27:55', '2022-04-08 09:27:55', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (8, 1, 'auth/role/edit', '修改角色', NULL, 2, 1, 1, 'edit', NULL, 'auth/role/edit', 0, 'sys_admin', NULL, '2022-04-08 09:28:52', '2022-04-08 09:28:52', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (9, 1, 'auth/role/setStatus', '设置角色状态', NULL, 2, 1, 1, 'setStatus', NULL, 'auth/role/setStatus', 0, 'sys_admin', NULL, '2022-04-08 09:36:14', '2022-04-08 09:36:14', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (10, 1, 'auth/role/setDataScope', '设置数据权限', NULL, 2, 1, 1, 'setDataScope', NULL, 'auth/role/setDataScope', 0, 'sys_admin', NULL, '2022-04-08 09:43:45', '2022-04-08 09:43:45', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (11, 1, 'auth/role/delete', '删除角色', NULL, 2, 1, 1, 'delete', NULL, 'auth/role/delete', 0, 'sys_admin', NULL, '2022-04-08 09:51:27', '2022-04-08 09:51:27', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (12, 1, 'auth/dept/list', '部门管理', NULL, 1, 1, 1, 'list', NULL, 'auth/dept/list', 0, 'sys_admin', NULL, '2022-04-08 09:55:37', '2022-04-08 09:55:37', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (13, 1, 'auth/dept/add', '添加部门', NULL, 2, 1, 1, 'add', NULL, 'auth/dept/add', 0, 'sys_admin', NULL, '2022-04-08 09:57:08', '2022-04-08 09:57:08', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (14, 1, 'auth/dept/edit', '修改部门', NULL, 2, 1, 1, 'edit', NULL, 'auth/dept/edit', 0, 'sys_admin', NULL, '2022-04-08 09:59:22', '2022-04-08 09:59:22', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (15, 1, 'auth/dept/delete', '删除部门', NULL, 2, 1, 1, 'delete', NULL, 'auth/dept/delete', 0, 'sys_admin', NULL, '2022-04-08 10:01:24', '2022-04-08 10:01:24', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (16, 1, 'auth/user/list', '用户管理', NULL, 1, 1, 1, 'list', NULL, 'auth/user/list', 0, 'sys_admin', NULL, '2022-04-08 10:04:07', '2022-04-08 10:04:07', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (17, 1, 'auth/user/add', '添加用户', NULL, 2, 1, 1, 'add', NULL, 'auth/user/add', 0, 'sys_admin', NULL, '2022-04-08 10:05:57', '2022-04-08 10:05:57', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (18, 1, 'auth/user/edit', '修改用户', NULL, 2, 1, 1, 'edit', NULL, 'auth/user/edit', 0, 'sys_admin', NULL, '2022-04-08 10:06:58', '2022-04-08 10:06:58', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (19, 1, 'auth/user/resetPwd', '密码重置', NULL, 2, 1, 1, 'resetPwd', NULL, 'auth/user/resetPwd', 0, 'sys_admin', NULL, '2022-04-08 10:08:54', '2022-04-08 10:08:54', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (20, 1, 'auth/user/setStatus', '设置用户状态', NULL, 2, 1, 1, 'setStatus', NULL, 'auth/user/setStatus', 0, 'sys_admin', NULL, '2022-04-08 10:10:31', '2022-04-08 10:10:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (21, 1, 'auth/user/delete', '删除用户', NULL, 2, 1, 1, 'delete', NULL, 'auth/user/delete', 0, 'sys_admin', NULL, '2022-04-08 10:12:03', '2022-04-08 10:12:03', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (22, 0, 'monitor', '系统监控', NULL, 0, 1, 1, 'monitor', NULL, 'monitor', 0, 'sys_admin', NULL, '2022-04-08 10:26:33', '2022-04-08 10:26:33', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (23, 22, 'monitor/userOnline/list', '在线用户', NULL, 1, 1, 1, 'list', NULL, 'monitor/userOnline/list', 0, 'sys_admin', NULL, '2022-04-08 10:29:20', '2022-04-08 10:29:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (24, 22, 'monitor/userOnline/forceLogout', '强制退出', NULL, 2, 1, 1, 'forceLogout', NULL, 'monitor/userOnline/forceLogout', 0, 'sys_admin', NULL, '2022-04-08 10:30:50', '2022-04-08 10:30:50', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (25, 22, 'monitor/loginLog/list', '登录日志', NULL, 1, 1, 1, 'list', NULL, 'monitor/loginLog/list', 0, 'sys_admin', NULL, '2022-04-08 10:34:26', '2022-04-08 10:34:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (26, 22, 'monitor/loginLog/delete', '删除登录日志', NULL, 2, 1, 1, 'delete', NULL, 'monitor/loginLog/delete', 0, 'sys_admin', NULL, '2022-04-08 10:35:26', '2022-04-08 10:35:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `show`, `path`, `jump_path`, `component`, `is_frame`, `module_type`, `remark`, `create_at`, `update_at`, `deleted_at`) VALUES (27, 22, 'monitor/loginLog/clear', '清除登录日志', NULL, 2, 1, 1, 'clear', NULL, 'monitor/loginLog/clear', 0, 'sys_admin', NULL, '2022-04-08 10:37:43', '2022-04-08 10:37:43', NULL);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `status` tinyint unsigned NOT NULL COMMENT '状态 0:停用 1:正常',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `data_scope` tinyint unsigned NOT NULL COMMENT '数据范围 1:全部数据权限 2:自定义数据权限 3:本部门数据权限 4:本部门及以下数据权限',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '备注',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of role
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `create_at`, `update_at`, `deleted_at`) VALUES (1, 'sadmin', 'b0167dd859b8a70478da36238b0b3e05', 'CqDQa4THP1', '超级管理员', '超级管理员', 0, 1, 1, NULL, NULL, NULL, NULL, '127.0.0.1', '2022-03-28 16:28:53', '2022-03-25 17:41:40', '2022-03-25 17:41:40', NULL);
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

SET FOREIGN_KEY_CHECKS = 1;
