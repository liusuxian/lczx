/*
 Navicat Premium Data Transfer

 Source Server         : LCZX
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : 116.62.212.185:3306
 Source Schema         : lczx

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 14/07/2022 16:42:40
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
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '3', '2', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '14', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '21', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '21', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '48', '2', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '49', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '47', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '51', '2', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '6', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '16', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '41', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '43', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '45', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '47', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '6', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '50', '6', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '52', '6', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '53', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '54', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '25', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '25', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '26', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '32', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '7', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '7', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '8', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '24', '5', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '24', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '16', '5', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '16', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '38', '5', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '38', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '5', '5', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '5', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '4', '2', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '4', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '4', '5', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '6', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '6', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '12', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '15', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '9', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '10', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '13', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '11', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '18', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '18', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '20', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '17', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '17', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '29', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '41', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '41', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '42', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '44', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '46', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '39', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '40', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '40', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '43', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '45', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '22', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '22', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '23', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '23', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '19', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '30', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '34', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '36', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '36', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '37', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '31', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '31', '4', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '27', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '35', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '33', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', '28', '3', '', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '6', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '7', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '8', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '9', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '16', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '17', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '18', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '35', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '36', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '38', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '39', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '40', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '41', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '42', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '43', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '44', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '45', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '46', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '47', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '48', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '2', '54', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '41', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '42', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '43', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '44', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '45', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '46', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '47', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '48', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '49', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '50', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '3', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '4', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '4', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '4', '51', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '4', '52', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '4', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '41', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '42', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '43', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '44', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '45', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '46', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '47', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '48', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '49', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '50', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '51', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '52', 'All', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '5', '53', 'All', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for crontab
-- ----------------------------
DROP TABLE IF EXISTS `crontab`;
CREATE TABLE `crontab` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '任务名称',
  `group` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '任务组名',
  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '参数',
  `invoke_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'cron执行表达式',
  `misfire_policy` tinyint unsigned NOT NULL COMMENT '计划执行策略 0:执行一次 1:执行多次',
  `status` tinyint unsigned NOT NULL COMMENT '状态 0:暂停 1:正常',
  `create_by` bigint unsigned NOT NULL COMMENT '创建者用户ID',
  `update_by` bigint unsigned DEFAULT NULL COMMENT '更新者用户ID',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of crontab
-- ----------------------------
BEGIN;
INSERT INTO `crontab` (`id`, `name`, `group`, `params`, `invoke_target`, `cron_expression`, `misfire_policy`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES (1, '检查在线用户', 'system', NULL, 'checkUserOnline', '0 0 */2 * * *', 1, 1, 1, NULL, NULL, '2022-06-29 11:52:59', '2022-06-29 11:52:59');
INSERT INTO `crontab` (`id`, `name`, `group`, `params`, `invoke_target`, `cron_expression`, `misfire_policy`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`) VALUES (2, '检查项目文件上传状态', 'system', NULL, 'checkWdkProjectFileUploadStatus', '0 0 9 25 * *', 1, 1, 1, NULL, NULL, '2022-06-29 11:54:14', '2022-06-29 11:54:14');
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
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  KEY `parent_id_index` (`parent_id`) USING BTREE,
  KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of dept
-- ----------------------------
BEGIN;
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, '绿城服务集团', 1, 1, NULL, 0, '', '2022-04-20 12:18:04', '2022-04-20 12:18:04', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, '咨询集团', 1, 1, NULL, 0, '', '2022-04-20 12:18:31', '2022-04-20 12:18:31', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 2, '职能中心', 1, 1, NULL, 0, '', '2022-04-20 12:18:52', '2022-04-20 12:18:52', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 3, '运营管理中心', 1, 1, NULL, 0, '', '2022-04-20 12:19:14', '2022-04-20 12:19:14', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 2, '分子公司', 1, 1, NULL, 0, '', '2022-04-20 12:20:53', '2022-04-20 12:20:53', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 5, '浙江幸福绿城房地产咨询有限公司', 1, 1, NULL, 0, '', '2022-04-20 12:21:54', '2022-04-20 12:21:54', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 6, '业务部门', 1, 1, NULL, 0, '', '2022-04-20 12:22:29', '2022-04-20 12:22:29', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 7, '物业咨询', 1, 1, NULL, 0, '', '2022-04-20 12:23:17', '2022-04-20 12:23:17', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 8, '技术支持部', 1, 1, 1, 4, '周荣辉', '2022-04-20 12:23:45', '2022-06-06 15:57:43', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 8, '区域一部', 1, 1, 1, 24, '许虎', '2022-06-06 13:22:27', '2022-06-06 15:50:27', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 8, '区域二部', 1, 1, 1, 16, '张磊', '2022-06-06 13:22:48', '2022-06-06 15:50:33', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 8, '区域三部', 1, 1, 1, 4, '周荣辉', '2022-06-06 13:23:25', '2022-06-06 15:50:50', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 8, '区域四部', 1, 1, 1, 5, '陈锋', '2022-06-06 13:24:33', '2022-06-06 15:51:10', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 6, '职能部门', 1, 1, NULL, 0, '', '2022-06-06 13:28:22', '2022-06-06 13:28:22', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 14, '市场运营部', 1, 1, 1, 50, '孙瑞良', '2022-06-06 13:29:51', '2022-06-06 14:57:13', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 6, '管理层', 1, 1, NULL, 0, '', '2022-06-06 13:30:25', '2022-06-06 13:30:25', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, 7, '房产咨询', 1, 1, NULL, 0, '', '2022-06-06 13:31:18', '2022-06-06 13:31:18', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 17, '全过程咨询部', 1, 1, NULL, 0, '', '2022-06-06 13:31:56', '2022-06-06 13:31:56', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 8, '区域五部', 1, 1, 1, 16, '张磊', '2022-06-06 15:48:16', '2022-06-06 15:51:26', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 8, '区域六部', 1, 1, 1, 24, '许虎', '2022-06-06 15:48:41', '2022-06-06 15:52:14', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, 8, '区域七部', 1, 1, 1, 38, '孟龙飞', '2022-06-06 15:49:12', '2022-06-06 15:52:49', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, 8, '西南区域', 1, 1, 1, 16, '张磊', '2022-06-06 15:53:37', '2022-06-06 16:17:09', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, 8, '江苏区域', 1, 1, 1, 24, '许虎', '2022-06-06 15:53:58', '2022-06-06 16:17:22', NULL);
INSERT INTO `dept` (`id`, `parent_id`, `name`, `status`, `created_by`, `updated_by`, `principal_uid`, `principal_name`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, 8, '京津冀区域', 1, 1, 1, 24, '许虎', '2022-06-06 15:54:51', '2022-06-06 16:17:29', NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of login_log
-- ----------------------------
BEGIN;
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (1, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-07 09:03:06', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (2, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-07 09:10:10', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (3, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-07 10:20:08', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (4, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-07 11:12:28', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (5, 'yaoyisahn', '183.129.209.28', '浙江省 杭州市', 'Safari', 'Intel Mac OS X 10_15_7', 0, '账号不存在', '2022-06-07 14:18:31', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (6, 'yaoyishan', '183.129.209.28', '浙江省 杭州市', 'Safari', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-07 14:18:43', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (7, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-07 17:54:57', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (8, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-08 09:32:28', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (9, 'chenxin74', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 7', 0, '账号不存在', '2022-06-08 12:43:31', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (10, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-08 15:15:57', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (11, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-08 17:07:11', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (12, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-08 17:10:24', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (13, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-08 17:10:40', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (14, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-08 17:13:20', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (15, 'liusuxian', '183.129.209.28', '浙江省 杭州市', 'Safari', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-08 17:30:11', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (16, 'yaoyishan', '183.129.209.28', '浙江省 杭州市', 'Safari', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-08 17:31:38', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (17, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-09 18:06:58', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (18, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-06-10 09:20:16', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (19, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-10 10:31:12', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (20, 'Fanli1', '36.19.232.3', '浙江省 杭州市', 'Safari', 'CPU OS 15_4_1 like Mac OS X', 1, '登录成功', '2022-06-11 14:23:42', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (21, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-13 10:29:29', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (22, 'zhangxiaobo', '183.129.209.27', '浙江省 杭州市', 'Chrome', 'Windows 7', 1, '登录成功', '2022-06-13 15:03:01', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (23, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-14 13:57:17', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (24, 'hanfei', '220.189.47.96', '浙江省 宁波市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-14 14:35:17', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (25, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-15 10:32:08', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (26, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-16 13:18:08', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (27, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-20 13:44:25', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (28, 'chenxin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 7', 1, '登录成功', '2022-06-22 16:08:19', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (29, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-27 11:20:57', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (30, 'sadmin', '183.253.72.80', '福建省 莆田市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-06-29 14:22:59', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (31, 'sadmin', '42.3.19.118', '香港 ', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-07-01 15:56:17', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (32, 'sadmin', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-07-04 10:31:24', '系统登录');
INSERT INTO `login_log` (`id`, `passport`, `ip`, `location`, `browser`, `os`, `status`, `msg`, `time`, `module`) VALUES (33, 'wulijun', '183.129.209.28', '浙江省 杭州市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-07-12 11:04:44', '系统登录');
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
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `rule_index` (`rule`),
  KEY `parent_id_index` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, 'auth', '权限管理', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:10:35', '2022-04-08 09:10:35', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, 'auth/menu/list', '菜单管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:15:06', '2022-04-08 09:15:06', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 2, 'auth/menu/add', '添加菜单', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:15:31', '2022-04-08 09:15:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 2, 'auth/menu/edit', '修改菜单', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:18:53', '2022-04-08 09:18:53', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 2, 'auth/menu/delete', '删除菜单', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:19:59', '2022-04-08 09:19:59', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 1, 'auth/role/list', '角色管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:25:30', '2022-04-08 09:25:30', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 6, 'auth/role/add', '添加角色', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:27:55', '2022-04-08 09:27:55', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 6, 'auth/role/edit', '修改角色', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:28:52', '2022-04-08 09:28:52', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 6, 'auth/role/setStatus', '设置角色状态', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:36:14', '2022-04-08 09:36:14', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 6, 'auth/role/setDataScope', '设置数据权限', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:43:45', '2022-04-08 09:43:45', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 6, 'auth/role/delete', '删除角色', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:51:27', '2022-04-08 09:51:27', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 1, 'auth/dept/list', '部门管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:55:37', '2022-04-08 09:55:37', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 12, 'auth/dept/add', '添加部门', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:57:08', '2022-04-08 09:57:08', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 12, 'auth/dept/edit', '修改部门', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 09:59:22', '2022-04-08 09:59:22', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 12, 'auth/dept/delete', '删除部门', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:01:24', '2022-04-08 10:01:24', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 1, 'auth/user/list', '用户管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:04:07', '2022-04-08 10:04:07', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, 16, 'auth/user/add', '添加用户', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:05:57', '2022-04-08 10:05:57', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 16, 'auth/user/edit', '修改用户', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:06:58', '2022-04-08 10:06:58', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 16, 'auth/user/resetPwd', '密码重置', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:08:54', '2022-04-08 10:08:54', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 16, 'auth/user/setStatus', '设置用户状态', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:10:31', '2022-04-08 10:10:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, 16, 'auth/user/delete', '删除用户', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:12:03', '2022-04-08 10:12:03', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, 0, 'monitor', '系统监控', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:26:33', '2022-04-08 10:26:33', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, 22, 'monitor/userOnline/list', '在线用户', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:29:20', '2022-04-08 10:29:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, 23, 'monitor/userOnline/forceLogout', '强制退出', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:30:50', '2022-04-08 10:30:50', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (25, 22, 'monitor/loginLog/list', '登录日志', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:34:26', '2022-04-08 10:34:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (26, 25, 'monitor/loginLog/delete', '删除登录日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:35:26', '2022-04-08 10:35:26', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (27, 25, 'monitor/loginLog/clear', '清除登录日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-08 10:37:43', '2022-04-08 10:37:43', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (28, 22, 'monitor/operLog/list', '操作日志', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:04:29', '2022-04-18 14:04:29', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (29, 28, 'monitor/operLog/delete', '删除操作日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:05:28', '2022-04-18 14:05:28', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (30, 28, 'monitor/operLog/clear', '清除操作日志', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 14:06:21', '2022-04-18 14:06:21', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (31, 22, 'monitor/server_monitor/info', '服务监控信息', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (32, 22, 'monitor/crontab/list', '定时任务', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (33, 32, 'monitor/crontab/add', '添加任务', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (34, 32, 'monitor/crontab/edit', '修改任务', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (35, 32, 'monitor/crontab/start', '开启任务', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (36, 32, 'monitor/crontab/stop', '停止任务', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (37, 32, 'monitor/crontab/run', '执行任务', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (38, 32, 'monitor/crontab/delete', '删除任务', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-18 17:03:42', '2022-04-18 17:03:42', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (39, 0, 'wdk', '文档库', NULL, 0, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:17:10', '2022-04-24 10:17:10', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (40, 39, 'wdk/project/list', '项目管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:18:56', '2022-04-24 10:18:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (41, 40, 'wdk/project/add', '添加项目', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:20:38', '2022-04-24 10:20:38', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (42, 40, 'wdk/project/edit', '修改项目', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:22:56', '2022-04-24 10:22:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (43, 40, 'wdk/project/delete', '删除项目', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:24:01', '2022-04-24 10:24:01', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (44, 40, 'wdk/project/export', '导出项目', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-24 10:24:01', '2022-04-24 10:24:01', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (45, 39, 'wdk/reportTypeCfg/list', '报告类型配置管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:14:16', '2022-04-25 10:14:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (46, 45, 'wdk/reportTypeCfg/add', '添加报告类型', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:20:31', '2022-04-25 10:20:31', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (47, 45, 'wdk/reportTypeCfg/edit', '修改报告类型', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:24:08', '2022-04-25 10:24:08', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (48, 45, 'wdk/reportTypeCfg/delete', '删除报告类型', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-25 10:26:56', '2022-04-25 10:26:56', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (49, 39, 'wdk/attachment/record', '附件记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:15:17', '2022-04-26 16:15:17', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (50, 49, 'wdk/attachment/add', '上传附件', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:16:16', '2022-04-26 16:16:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (51, 49, 'wdk/attachment/delete', '删除附件', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:16:16', '2022-04-26 16:16:16', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (52, 39, 'wdk/service/record', '服务记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:18:13', '2022-04-26 16:18:13', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (53, 52, 'wdk/service/add', '添加服务记录', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:19:04', '2022-04-26 16:19:04', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (54, 52, 'wdk/service/delete', '删除服务记录', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-04-26 16:19:04', '2022-04-26 16:19:04', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (55, 39, 'wdk/file/record', '文件记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:46:40', '2022-05-01 12:46:40', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (56, 55, 'wdk/file/add', '上传文件', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:47:52', '2022-05-01 12:47:52', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (57, 39, 'wdk/report/record', '报告记录管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:53:20', '2022-05-01 12:53:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (58, 57, 'wdk/report/add', '上传报告', 'check_dept', 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:54:36', '2022-05-01 12:54:36', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (59, 57, 'wdk/report/delete', '删除报告', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-01 12:54:36', '2022-05-01 12:54:36', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (60, 39, 'wdk/reportAudit/uploadAuditList', '报告上传审核管理', NULL, 1, 1, NULL, 0, ' sys_admin', NULL, '2022-05-22 19:06:47', '2022-05-22 19:06:47', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (61, 60, 'wdk/reportAudit/rescindAudit', '报告撤销审核', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-22 19:09:32', '2022-05-22 19:09:32', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (62, 39, 'wdk/reportAudit/list', '报告审核管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 01:31:59', '2022-05-23 01:31:59', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (63, 62, 'wdk/reportAudit/audit', '报告审核', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 10:53:10', '2022-05-23 10:53:10', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (64, 39, 'wdk/report/excellenceList', '优秀报告管理', NULL, 1, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 11:04:20', '2022-05-23 11:04:20', NULL);
INSERT INTO `menu` (`id`, `parent_id`, `rule`, `name`, `condition`, `menu_type`, `status`, `jump_path`, `is_frame`, `module_type`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (65, 64, 'wdk/report/chooseExcellence', '评选优秀报告', NULL, 2, 1, NULL, 0, 'sys_admin', NULL, '2022-05-23 11:06:13', '2022-05-23 11:06:13', NULL);
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
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `created_at`, `updated_at`) VALUES (1, '超级管理员', 1, 3, '', '2022-04-19 14:30:16', '2022-04-19 14:30:16');
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `created_at`, `updated_at`) VALUES (2, '一般管理员', 1, 3, '', '2022-06-06 13:12:49', '2022-06-07 01:07:37');
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `created_at`, `updated_at`) VALUES (3, '普通用户', 1, 3, '', '2022-06-06 13:13:31', '2022-06-07 01:13:38');
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `created_at`, `updated_at`) VALUES (4, '专家组', 1, 3, '', '2022-06-06 13:16:55', '2022-06-07 01:13:48');
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `created_at`, `updated_at`) VALUES (5, '负责人', 1, 3, '', '2022-06-06 13:17:36', '2022-06-07 01:14:34');
INSERT INTO `role` (`id`, `name`, `status`, `data_scope`, `remark`, `created_at`, `updated_at`) VALUES (6, '普通管理员', 1, 3, '', '2022-06-06 14:44:21', '2022-06-06 14:44:21');
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
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `passport_index` (`passport`),
  UNIQUE KEY `mobile_index` (`mobile`),
  KEY `realname_index` (`realname`),
  KEY `nickname_index` (`nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'sadmin', 'b0167dd859b8a70478da36238b0b3e05', 'CqDQa4THP1', '超级管理员', '超级管理员', 4, 1, 1, 1, 'user/avatar/logo.png', '17364814710', '382185882@qq.com', NULL, '183.129.209.28', '2022-07-04 10:31:24', '2022-03-25 17:41:40', '2022-03-25 17:41:40', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'liusuxian', 'b6d5d1dc1f39f8a10bfdb53bf78c895f', 'JKPT7WSeG0', '刘苏贤', '', 4, 1, 1, 1, NULL, '15108274735', '', '', '183.129.209.28', '2022-06-08 17:30:11', '2022-04-25 17:21:17', '2022-04-25 17:21:17', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'fanli1', 'd4ef6db5cadae7717791bae071d6317d', 'qehBH2mMDI', '樊琍', '', 16, 2, 1, 1, NULL, '13758278788', '', '', '36.19.232.3', '2022-06-11 14:23:42', '2022-06-06 13:53:43', '2022-06-06 13:53:43', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'zhouronghui', 'd2209e0b9d67d3dd878aeda1af7472f8', 'T1hE8ugA8V', '周荣辉', '', 16, 1, 1, 1, NULL, '13396572227', '', '', NULL, NULL, '2022-06-06 13:55:25', '2022-06-06 16:05:56', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'chenfeng', 'fc849406c501e25fbfe7e8ed13167600', 'PMEgf76nLu', '陈锋', '', 13, 1, 1, 0, NULL, '15857110571', '', '', NULL, NULL, '2022-06-06 13:56:57', '2022-06-06 16:04:18', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'xueqing', 'a762dfb90a71d9e5942d71fd7d990d70', 'QyHgsF4L1B', '薛擎', '', 13, 1, 1, 0, NULL, '13616520155', '', '', NULL, NULL, '2022-06-06 13:58:30', '2022-06-06 16:06:47', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'fukangzheng', 'd6d2ac8d2727990c3e762d89ed2f20f9', 'uFZXdRtIsF', '傅康振', '', 12, 1, 1, 0, NULL, '18858173050', '', '', NULL, NULL, '2022-06-06 13:59:18', '2022-06-06 16:01:17', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 'daihaiying', '2e662b24648e7598e734c1c8d5e15a02', '7Ih8J6BlDU', '戴海印', '', 12, 1, 1, 0, NULL, '13065950051', '', '', '112.10.226.49', '2022-06-07 01:12:07', '2022-06-06 14:00:05', '2022-06-06 16:01:36', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'zhangyongqin', 'c4e935cea588232e7b141715a5f29d7d', 'enliTclVPg', '张永钦', '', 13, 1, 1, 0, NULL, '13757155216', '', '', NULL, NULL, '2022-06-06 14:00:54', '2022-06-06 16:07:27', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 'dingsongchao', '5b2a89349ea3094807944bd02c4b3576', 'JssFicMbP1', '丁松超', '', 13, 1, 1, 0, NULL, '13750857023', '', '', NULL, NULL, '2022-06-06 14:02:15', '2022-06-06 16:07:39', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 'wanghaiyang', 'a39b830daed59542a17b95595f60e529', 'EWCPnkfdrA', '王海洋', '', 13, 1, 1, 0, NULL, '18656131054', '', '', NULL, NULL, '2022-06-06 14:02:55', '2022-06-06 16:08:05', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 'chenxiang', 'b6ac33986a840238bae40bbed511206c', 'SssGFMebLJ', '陈祥', '', 13, 1, 1, 0, NULL, '18357234654', '', '', NULL, NULL, '2022-06-06 14:03:29', '2022-06-06 16:07:00', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 'chenqiwen', '11b3dd18d7ee5fd1eb7babd4a31e022b', '83Lv6rlzFO', '程起文', '', 13, 1, 1, 0, NULL, '18556511761', '', '', NULL, NULL, '2022-06-06 14:04:46', '2022-06-06 16:07:52', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 'liyaodong', '738f9270825e6e5441d356dbac778a2a', 'Itvf8n45xf', '李耀栋', '', 10, 1, 0, 0, NULL, '18351073783', '', '', NULL, NULL, '2022-06-06 14:05:20', '2022-06-06 16:21:36', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 'huangyougang', 'de20309949a3887cca18c76f27a11605', 'M0PfMG0Lwm', '黄友钢', '', 13, 1, 1, 0, NULL, '18906529803', '', '', NULL, NULL, '2022-06-06 14:05:51', '2022-06-06 16:07:14', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 'zhanglei', '140a3d97cb1473b0ab10e1e89e054597', 'Qef1yalbXE', '张磊', '', 19, 1, 1, 0, NULL, '18768486708', '', '', NULL, NULL, '2022-06-06 14:07:15', '2022-06-06 16:03:37', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, 'zhangyang5', '4c998588684788b40c5d9498f479b2da', 'mvJi3EwSsX', '张洋', '', 19, 1, 1, 0, NULL, '15868199430', '', '', NULL, NULL, '2022-06-06 14:09:36', '2022-06-06 16:09:28', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 'heyijiang', 'b245e1ae2a6e33cc4cc0f8d77ead859c', 'NYpJF8SNJ3', '何一江', '', 19, 1, 1, 0, NULL, '15858165163', '', '', NULL, NULL, '2022-06-06 14:10:14', '2022-06-06 16:08:22', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 'liaozhida', '9389885c6462f8d23bb6bfde1b878e31', 'rj6gzxtafu', '廖智达', '', 22, 1, 1, 0, NULL, '18628060812', '', '', NULL, NULL, '2022-06-06 14:10:51', '2022-06-06 16:16:56', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 'wubingbing', '7a15a01233d5ddddd7f77481cf2a6563', 'ETg4yHUTrH', '吴兵兵', '', 19, 1, 1, 0, NULL, '18706409646', '', '', NULL, NULL, '2022-06-06 14:11:26', '2022-06-06 16:08:37', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, 'zhangxiaobo', 'a08fbcb128b4f0004d5c261c50830a2a', 'taQzfKIhcd', '章晓波', '', 11, 1, 1, 0, NULL, '15167192621', '', '', '183.129.209.27', '2022-06-13 15:03:01', '2022-06-06 14:12:19', '2022-06-06 14:12:19', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, 'yuxiaowei', '2aad2ae9fcab0c7415a10b7accb5e406', '5DfsYjbIkk', '余晓伟', '', 22, 1, 1, 0, NULL, '18358583833', '', '', NULL, NULL, '2022-06-06 14:12:55', '2022-06-06 16:16:13', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, 'tongke', '17aafe7f4ff4e9460578a61e5a925df9', 'oQAdiUEYDc', '童柯', '', 22, 1, 1, 0, NULL, '18966008633', '', '', NULL, NULL, '2022-06-06 14:14:19', '2022-06-06 16:16:41', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, 'xuhu66', 'a4b82b82b2a6b74b21e02645740bc451', '0nRDA3aNjo', '许虎', '', 20, 1, 1, 0, NULL, '18667908989', '', '', NULL, NULL, '2022-06-06 14:15:03', '2022-06-06 16:03:19', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (25, 'zhouzhenghua', '67d071ed503911daa5a4a082db24a629', 'x8vx5sP6Vy', '周振华', '', 10, 1, 1, 0, NULL, '13588119134', '', '', NULL, NULL, '2022-06-06 14:15:39', '2022-06-06 15:59:05', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (26, 'hanfei', '32fda78a55a4a7f376b44dc3e3991563', 'JIosgSh1nD', '韩非', '', 10, 1, 1, 0, NULL, '13339202872', '', '', '220.189.47.96', '2022-06-14 14:35:17', '2022-06-06 14:16:10', '2022-06-06 15:59:27', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (27, 'huyuqi', '1212b3ed1d92d74d1dd09279cabf0126', '2hXzzeuQJu', '胡毓琦', '', 24, 1, 1, 0, NULL, '17519544757', '', '', NULL, NULL, '2022-06-06 14:16:42', '2022-06-06 16:19:38', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (28, 'hankaihao', '6b772f4c9b3b32f7e014836c2748a678', 'zZzlR3KTKr', '韩恺豪', '', 24, 1, 1, 0, NULL, '18137780329', '', '', NULL, NULL, '2022-06-06 14:17:16', '2022-06-06 16:20:23', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (29, 'zhujianlong', '38b4721fca505a3f0563e754787feadf', 'G8gwzNtVue', '朱建龙', '', 20, 1, 1, 0, NULL, '17361553932', '', '', NULL, NULL, '2022-06-06 14:17:43', '2022-06-06 16:11:03', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (30, 'xuzheng', '115c4e5ed7baf15b9af1747bf8eb2e95', '457KD2zMpZ', '徐征', '', 23, 1, 1, 0, NULL, '15557162333', '', '', NULL, NULL, '2022-06-06 14:18:44', '2022-06-06 16:18:32', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (31, 'zhaoqiang', 'f8619d2d5d4ff338bef1261daf967fa9', '8JfeZG0hnE', '赵强', '', 24, 1, 1, 0, NULL, '13591751761', '', '', NULL, NULL, '2022-06-06 14:19:22', '2022-06-06 16:19:27', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (32, 'wengyuefeng', 'a943d27f7caa25924b7a4df2d207705b', 'hMxD1zJSdT', '翁跃丰', '', 10, 1, 1, 0, NULL, '15967159926', '', '', NULL, NULL, '2022-06-06 14:19:57', '2022-06-06 15:59:45', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (33, 'zhangyang6', '2ba6655be28a391baef7876e27981193', 'Pyt3sDgZEI', '张洋', '', 24, 1, 1, 0, NULL, '15001893001', '', '', NULL, NULL, '2022-06-06 14:20:22', '2022-06-06 16:20:09', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (34, 'huwendi', '2178d47d37e7d7f1d24a83a24e873d70', 'Xxjybesxc0', '胡文迪', '', 23, 1, 1, 0, NULL, '13777451207', '', '', NULL, NULL, '2022-06-06 14:20:48', '2022-06-06 16:18:47', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (35, 'dangzejun', '7c1b0407dff23eabc04196ffbd3158db', '0zItllvHTi', '党泽君', '', 24, 1, 1, 0, NULL, '18802205313', '', '', NULL, NULL, '2022-06-06 14:21:20', '2022-06-06 16:19:50', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (36, 'yuanhengye', '660601167d0d7acb7025a3965ff4fae6', 'cC6KXRMSim', '袁恒业', '', 23, 1, 1, 0, NULL, '13906520763', '', '', NULL, NULL, '2022-06-06 14:22:00', '2022-06-06 16:19:00', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (37, 'zhangxu', 'f49a721d98ef72f4a7909245c686fdb9', 'ZA7CxVwYyw', '张旭', '', 23, 1, 1, 0, NULL, '18626016698', '', '', NULL, NULL, '2022-06-06 14:22:28', '2022-06-06 16:19:12', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (38, 'menglongfei', '3859324cdadac48bd526963328a63312', 'dgCn0rTkNB', '孟龙飞', '', 21, 1, 1, 0, NULL, '15925672995', '', '', NULL, NULL, '2022-06-06 14:23:10', '2022-06-06 16:03:57', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (39, 'zhuzhenyu', '0f56bd98ecad4df2d4a3035b8f785907', 'DO9x9VvT1t', '朱振宇', '', 21, 1, 1, 0, NULL, '18858122871', '', '', NULL, NULL, '2022-06-06 14:23:40', '2022-06-06 16:12:45', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (40, 'sunlei', '35e7bc87bcf82f4f218f9b5278538c44', 'DSbBfOo7J7', '孙雷', '', 21, 1, 1, 0, NULL, '15372566738', '', '', NULL, NULL, '2022-06-06 14:24:11', '2022-06-06 16:13:25', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (41, 'huangyufan', '634f855db1962fe240a1edab16e317c4', '9NxsXLVzWL', '黄裕凡', '', 21, 1, 1, 0, NULL, '18100176010', '', '', NULL, NULL, '2022-06-06 14:24:37', '2022-06-06 16:11:17', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (42, 'dongbinjie', '410da21ffbc1a6bf1d284a312d7846b0', 'Ma1gJyJPlN', '董滨杰', '', 21, 1, 1, 0, NULL, '18367131585', '', '', NULL, NULL, '2022-06-06 14:25:14', '2022-06-06 16:11:30', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (43, 'houli7', '395d9fa6b4a261de1982a8efa919354c', 'HbB7i5uP8v', '侯力', '', 21, 1, 1, 0, NULL, '15716274786', '', '', NULL, NULL, '2022-06-06 14:25:42', '2022-06-06 16:13:37', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (44, 'chenchangli', 'e3609ef1d26a127033d98f54ce68a8fa', 'NRf7u4Q0hd', '陈长利', '', 21, 1, 1, 0, NULL, '13262207916', '', '', NULL, NULL, '2022-06-06 14:26:18', '2022-06-06 16:11:43', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (45, 'xiaoningning', '95b41c7da90e60fb9f882b66da46fbc8', 'g9PuB9gYSQ', '肖宁宁', '', 21, 1, 1, 0, NULL, '18258234094', '', '', NULL, NULL, '2022-06-06 14:26:50', '2022-06-06 16:13:51', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (46, 'xieyingjie', '5f85085484b0635df24e980f9b7394fb', 'EkUsqQ18Hl', '谢英杰', '', 21, 1, 1, 0, NULL, '15669910099', '', '', NULL, NULL, '2022-06-06 14:27:28', '2022-06-06 16:12:24', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (47, 'yuehongliang', '5234f0729af3a4472cb411ccdbf0f805', 'NADTNAIdnO', '岳洪亮', '', 9, 1, 1, 0, NULL, '13958066519', '', '', NULL, NULL, '2022-06-06 14:28:11', '2022-06-06 14:34:16', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (48, 'wulijun', '003c53773fdc51ea401fa58069f5348a', 'UHqnmkhAtf', '吴利军', '', 9, 1, 1, 1, NULL, '18810023517', '', '', '183.129.209.28', '2022-07-12 11:04:44', '2022-06-06 14:28:57', '2022-06-06 14:28:57', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (49, 'chenxin', '4e48153e89a42bf965972db575f4b7f5', 'iPAsGlI5qQ', '陈欣', '', 9, 1, 1, 0, NULL, '15858173043', '', '', '183.129.209.28', '2022-06-22 16:08:19', '2022-06-06 14:32:35', '2022-06-06 14:32:35', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (50, 'sunruiliang', 'e3c96e3e2a2797341baa97f53dd85502', 'FbjqY2BRTv', '孙瑞良', '', 15, 1, 1, 1, NULL, '13634154818', '', '', NULL, NULL, '2022-06-06 14:35:15', '2022-06-06 14:46:18', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (51, 'yaoyishan', '46780e3b959666e81fe427e86ea16ef5', 'aO6C88WM9T', '姚艺珊', '', 15, 2, 1, 1, NULL, '18858153378', '', '', '183.129.209.28', '2022-06-08 17:31:38', '2022-06-06 14:36:00', '2022-06-06 14:36:00', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (52, 'wangjun', '5b642b18c7cd6d9c5e58b3a329b5e5a0', 'EzuSZd78Pl', '王军', '', 15, 1, 1, 1, NULL, '15702439287', '', '', NULL, NULL, '2022-06-06 14:37:17', '2022-06-06 14:47:03', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (53, 'chenlisheng', 'ef3d6dbe57e844921f47afbe4e544a9c', '6EIuW8e2w4', '陈立胜', '', 18, 1, 1, 0, NULL, '13819987129', '', '', NULL, NULL, '2022-06-06 14:48:06', '2022-06-06 14:48:06', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `salt`, `realname`, `nickname`, `dept_id`, `gender`, `status`, `is_admin`, `avatar`, `mobile`, `email`, `remark`, `last_login_ip`, `last_login_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (54, 'yukai1', '178507acfbe6c81e20bfa4b7baa76b5e', 'bFSJnxcJ8T', '俞凯', '', 18, 1, 1, 0, NULL, '18958065090', '', '', NULL, NULL, '2022-06-06 14:48:31', '2022-06-06 14:48:31', NULL);
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
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
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
  `created_at` datetime DEFAULT NULL COMMENT '上传时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_file
-- ----------------------------
BEGIN;
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (1, 1, '5、德清县舞阳街道2021-001号地块项目物业管理顾问咨询合同扫描件.pdf', 0, 48, '吴利军', 'wdk/file/ckj7oib4lq180ugbsfcdmo.pdf', 'wdk/file/ckj7oib4lq180ugbsfcdmo.pdf', '2022-06-07 01:02:02', '2022-06-07 01:02:02');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (2, 2, '1、凤悦印湖项目生活服务顾问咨询合同.pdf', 0, 48, '吴利军', 'wdk/file/ckkngykesm95zcbgwvatto.pdf', 'wdk/file/ckkngykesm95zcbgwvatto.pdf', '2022-06-08 17:37:08', '2022-06-08 17:37:08');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (3, 3, '1、温州市鹿城广场三、四期项目物业管理顾问咨询合同（扫描件）.pdf', 0, 48, '吴利军', 'wdk/file/ckook8bm19luyx2n9p7g2q.pdf', 'wdk/file/ckook8bm19luyx2n9p7g2q.pdf', '2022-06-13 11:19:06', '2022-06-13 11:19:06');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (4, 4, '5、温州绿城桂语江南项目前期物业管理顾问咨询合同2022续（扫描件）.pdf', 0, 48, '吴利军', 'wdk/file/ckor81kxnsp9nntm9plg9o.pdf', 'wdk/file/ckor81kxnsp9nntm9plg9o.pdf', '2022-06-13 13:24:14', '2022-06-13 13:24:14');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (5, 5, '4、温州绿城春月江澜项目物业管理顾问咨询合同 2022续（扫描件）.pdf', 0, 48, '吴利军', 'wdk/file/ckorhghw97pz91jdfbounp.pdf', 'wdk/file/ckorhghw97pz91jdfbounp.pdf', '2022-06-13 13:36:32', '2022-06-13 13:36:32');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (6, 6, '温州瑞安兰园项目前期物业管理顾问咨询合同（扫描件）(1).pdf', 0, 21, '章晓波', 'wdk/file/ckotgcxejmnoupaenf02gc.pdf', 'wdk/file/ckotgcxejmnoupaenf02gc.pdf', '2022-06-13 15:09:08', '2022-06-13 15:09:08');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (7, 7, '2、温州南湖未来社区工作联系单.pdf', 0, 48, '吴利军', 'wdk/file/ckotjgbt2fgpbbta31kx5h.pdf', 'wdk/file/ckotjgbt2fgpbbta31kx5h.pdf', '2022-06-13 15:13:10', '2022-06-13 15:13:10');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (8, 8, '3、温州市七都片区04-B-04，19地块项目前期顾问咨询合同.pdf', 0, 48, '吴利军', 'wdk/file/ckou2ze2dwdtkzgqgucsnd.pdf', 'wdk/file/ckou2ze2dwdtkzgqgucsnd.pdf', '2022-06-13 15:38:41', '2022-06-13 15:38:41');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (9, 9, '温岭豪成新河镇XH050802-04地块项目精装修工程、机电工程、景观工程管理顾咨询合同2021(扫描件）.pdf', 0, 48, '吴利军', 'wdk/file/ckouqhkozgsolyftpmks2f.pdf', 'wdk/file/ckouqhkozgsolyftpmks2f.pdf', '2022-06-13 16:09:23', '2022-06-13 16:09:23');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (10, 10, '椒江区现代大道北地块项目物业管理顾问咨询合同22.5.19.pdf', 0, 48, '吴利军', 'wdk/file/ckouzob9iv5zm8rzqbffht.pdf', 'wdk/file/ckouzob9iv5zm8rzqbffht.pdf', '2022-06-13 16:21:23', '2022-06-13 16:21:23');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (11, 12, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckpo0mcq8dklssjnxmkojq.pdf', 'wdk/file/ckpo0mcq8dklssjnxmkojq.pdf', '2022-06-14 15:06:09', '2022-06-14 15:06:09');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (12, 13, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckpob9lrzfihtctfgmvcgb.pdf', 'wdk/file/ckpob9lrzfihtctfgmvcgb.pdf', '2022-06-14 15:20:03', '2022-06-14 15:20:03');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (13, 14, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckqd3r7a34al6sj4ysngpg.pdf', 'wdk/file/ckqd3r7a34al6sj4ysngpg.pdf', '2022-06-15 10:45:42', '2022-06-15 10:45:42');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (14, 15, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckqgprqa233owibnq76jnb.pdf', 'wdk/file/ckqgprqa233owibnq76jnb.pdf', '2022-06-15 13:35:31', '2022-06-15 13:35:31');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (15, 16, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckqi5xjcthv41fdynuynua.pdf', 'wdk/file/ckqi5xjcthv41fdynuynua.pdf', '2022-06-15 14:43:38', '2022-06-15 14:43:38');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (16, 17, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckqifp2hknd70madz7jqhf.pdf', 'wdk/file/ckqifp2hknd70madz7jqhf.pdf', '2022-06-15 14:56:23', '2022-06-15 14:56:23');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (17, 18, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckqjzohpo2l7h3a1dmqc7c.pdf', 'wdk/file/ckqjzohpo2l7h3a1dmqc7c.pdf', '2022-06-15 16:09:31', '2022-06-15 16:09:31');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (18, 19, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckqkewj97ks67ohawt1c86.pdf', 'wdk/file/ckqkewj97ks67ohawt1c86.pdf', '2022-06-15 16:29:24', '2022-06-15 16:29:24');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (19, 20, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckqkmvn5injhvddgxtwha2.pdf', 'wdk/file/ckqkmvn5injhvddgxtwha2.pdf', '2022-06-15 16:39:48', '2022-06-15 16:39:48');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (20, 21, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckus3fozja4fw600yckmnw.pdf', 'wdk/file/ckus3fozja4fw600yckmnw.pdf', '2022-06-20 15:21:14', '2022-06-20 15:21:14');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (21, 22, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckus8bkd29x5y4jl5cdgma.pdf', 'wdk/file/ckus8bkd29x5y4jl5cdgma.pdf', '2022-06-20 15:27:36', '2022-06-20 15:27:36');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (22, 23, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckusfmscyqgnh5tkahcy01.pdf', 'wdk/file/ckusfmscyqgnh5tkahcy01.pdf', '2022-06-20 15:37:09', '2022-06-20 15:37:09');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (23, 24, '宁波区域开发项目前期物业咨询合作协议.pdf', 0, 48, '吴利军', 'wdk/file/ckusogl639422vrwxwneuw.pdf', 'wdk/file/ckusogl639422vrwxwneuw.pdf', '2022-06-20 15:48:41', '2022-06-20 15:48:41');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (24, 26, '宁波文创港核心区启动地块（二三期）项目(2).pdf', 0, 48, '吴利军', 'wdk/file/ckutmua657s5l0hishejna.pdf', 'wdk/file/ckutmua657s5l0hishejna.pdf', '2022-06-20 16:33:35', '2022-06-20 16:33:35');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (25, 27, '1、凤悦印湖项目生活服务顾问咨询合同.pdf', 0, 48, '吴利军', 'wdk/file/cl0oc3zr2vd0hm0rwl1tpn.pdf', 'wdk/file/cl0oc3zr2vd0hm0rwl1tpn.pdf', '2022-06-27 13:40:29', '2022-06-27 13:40:29');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (26, 28, '檀山府项目（绿城雲庐）(1).pdf', 0, 48, '吴利军', 'wdk/file/cl0oio03qligpw02ir5je5.pdf', 'wdk/file/cl0oio03qligpw02ir5je5.pdf', '2022-06-27 13:49:02', '2022-06-27 13:49:02');
INSERT INTO `wdk_file` (`id`, `project_id`, `name`, `type`, `create_by`, `create_name`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (27, 28, '20220112关于嵊州云庐项目年度顾问咨询服务计划.doc', 1, 48, '吴利军', 'wdk/file/cl0ooyqozu49jm7j4vf9nd.doc', 'wdk/file/cl0ooyqozu49jm7j4vf9nd.pdf', '2022-06-27 13:57:18', '2022-06-27 13:57:18');
COMMIT;

-- ----------------------------
-- Table structure for wdk_project
-- ----------------------------
DROP TABLE IF EXISTS `wdk_project`;
CREATE TABLE `wdk_project` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '文档库项目ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目名称',
  `type` tinyint unsigned NOT NULL COMMENT '项目性质 0:蓝绿体系 1:非绿',
  `origin` tinyint unsigned NOT NULL COMMENT '项目来源 0:绿中 1:分子公司 2:合伙人 3:老客户 4:中交 5:蓝城 6:自拓 7:其他',
  `step` tinyint unsigned NOT NULL COMMENT '项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中-规划设计 4:服务中-项目展示区施工 5:服务中-主体结构工程 6:服务中-主体安装工程 7:服务中-装饰装修工程 8:服务中-景观市政工程 9:服务中-项目交付验收 30:合同结束 31:复盘',
  `file_upload_status` tinyint unsigned NOT NULL COMMENT '项目文件上传状态 0:异常 1:正常 2:已完成',
  `business_type` tinyint unsigned NOT NULL COMMENT '业务类型 0:物业 1:专项 2:全过程',
  `contract_status` tinyint unsigned NOT NULL COMMENT '签约状态 0:新签 1:续签 2:未签',
  `contract_sum` double unsigned NOT NULL COMMENT '合同金额',
  `deep_culture` tinyint unsigned NOT NULL COMMENT '是否为深耕 0:否 1:是',
  `status` tinyint unsigned NOT NULL COMMENT '服务状态 0:服务中 1:暂停 2:提前终止 3:跟踪期 4:洽谈中 5:正常结束',
  `entrust_company` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '委托方公司',
  `sign_company` tinyint unsigned NOT NULL COMMENT '我方签订公司 0:绿城房地产咨询集团有限公司 1:浙江幸福绿城房地产咨询有限公司 2:浙江美好绿城房地产咨询有限公司',
  `principal_uid` bigint unsigned NOT NULL COMMENT '负责人用户ID',
  `principal_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '负责人姓名',
  `dept_id` bigint unsigned NOT NULL COMMENT '项目所属部门ID',
  `dept_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目所属部门名称',
  `region` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '地区(省/市/区县)',
  `start_time` date NOT NULL COMMENT '项目开始时间',
  `end_time` date NOT NULL COMMENT '项目结束时间',
  `create_by` bigint unsigned NOT NULL COMMENT '项目创建者用户ID',
  `create_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目创建者姓名',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '项目修改者用户ID',
  `updated_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '项目修改者姓名',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `file_upload_last_time` datetime DEFAULT NULL COMMENT '项目文件最近一次上传时间',
  `created_at` datetime DEFAULT NULL COMMENT '项目创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '项目更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '项目软删除时间',
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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_project
-- ----------------------------
BEGIN;
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '德清县舞阳街道2021-001号地块项目', 0, 0, 3, 0, 0, 0, 18, 0, 0, '德清绿城浙然置业有限公司', 2, 6, '薛擎', 13, '区域四部', '浙江省/湖州市/德清县', '2022-04-01', '2023-03-31', 48, '吴利军', 48, '吴利军', '', NULL, '2022-06-07 01:01:45', '2022-06-08 17:26:00', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '凤悦印印湖府项目', 0, 0, 3, 0, 0, 0, 12, 0, 0, '宁波奉化甬轩房地产开发有限公司', 2, 25, '周振华', 10, '区域一部', '浙江省/宁波市/奉化区', '2022-01-01', '2022-12-31', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-08 17:34:38', '2022-06-08 17:38:33', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, '温州鹿城广场', 0, 0, 3, 0, 0, 0, 18, 1, 0, '温州绿景置业有限公司', 0, 21, '章晓波', 11, '区域二部', '浙江省/温州市/鹿城区', '2022-01-01', '2022-12-31', 48, '吴利军', 48, '吴利军', '合同联系人：何一江', NULL, '2022-06-13 11:18:32', '2022-06-13 13:23:35', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, '温州桂语江南', 0, 0, 3, 0, 0, 1, 9, 1, 0, '温州浙扬房地产开发有限公司', 0, 21, '章晓波', 11, '区域二部', '浙江省/温州市/鹿城区', '2022-07-01', '2022-12-31', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-13 13:22:24', '2022-06-13 13:25:47', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '温州春月江澜项目', 0, 0, 3, 0, 0, 0, 13.5, 1, 0, '温州绿城浙琴房地产开发有限公司', 0, 21, '章晓波', 11, '区域二部', '浙江省/温州市/鹿城区', '2022-04-01', '2022-12-31', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-13 13:35:46', '2022-06-13 13:39:03', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '瑞安兰园', 0, 0, 3, 0, 0, 1, 13.5, 1, 4, '温州绿城浙冠房地产开发有限公司', 2, 21, '章晓波', 11, '区域二部', '浙江省/温州市/瑞安市', '2022-04-01', '2022-12-31', 21, '章晓波', NULL, NULL, '待签', NULL, '2022-06-13 15:07:32', '2022-06-13 15:26:26', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, '南湖未来社区', 1, 7, 3, 0, 0, 0, 18, 0, 0, '未来社区公司', 2, 21, '章晓波', 11, '区域二部', '浙江省/温州市/瓯海区', '2022-01-01', '2022-12-31', 21, '章晓波', NULL, NULL, '', NULL, '2022-06-13 15:11:54', '2022-06-13 15:14:26', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, '温州七都科技岛项目', 1, 1, 3, 0, 1, 0, 99, 1, 0, '温州泽大产业发展运营管理有限公司', 0, 21, '章晓波', 11, '区域二部', '浙江省/温州市/鹿城区', '2022-04-01', '2023-03-31', 48, '吴利军', NULL, NULL, '合同联系人：何一江', NULL, '2022-06-13 15:37:22', '2022-06-13 15:39:45', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, '温岭豪成栖江悦项目（豪成新河镇XH050802-04地块）', 1, 6, 3, 0, 0, 0, 174, 1, 0, '台州豪成骏越房地产有限公司', 0, 21, '章晓波', 11, '区域二部', '浙江省/台州市/温岭市', '2021-11-01', '2022-11-30', 48, '吴利军', NULL, NULL, '合同联系人：何一江', NULL, '2022-06-13 16:08:54', '2022-06-13 16:11:06', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, '台州晓风印月（椒江区现代大道北地块）', 0, 0, 3, 0, 0, 0, 18, 1, 0, '台州绿珵置业有限公司', 0, 21, '章晓波', 11, '区域二部', '浙江省/台州市/椒江区', '2022-03-01', '2023-02-28', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-13 16:20:06', '2022-06-13 16:25:11', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, '嵊州绿城雲庐（檀山府）', 0, 0, 0, 0, 0, 0, 18, 1, 0, '浙江红贝投资有限公司', 2, 25, '周振华', 20, '区域六部', '浙江省/绍兴市/嵊州市', '2022-01-01', '2022-12-31', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-14 14:41:12', '2022-06-14 14:41:12', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, '宁波高新13号地块（春月云锦）', 0, 0, 3, 0, 0, 0, 1, 1, 0, '宁波绿城房地产投资有限公司', 2, 32, '翁跃丰', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', NULL, NULL, '打包合同中赠送项目', NULL, '2022-06-14 15:05:18', '2022-06-14 15:09:02', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, '宁波杭州湾5#地块（春熙月明）', 0, 0, 3, 0, 0, 0, 6, 1, 0, '宁波杭州湾新区海育开发建设有限公司', 2, 26, '韩非', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '打包合同项目之一', NULL, '2022-06-14 15:19:28', '2022-06-20 16:28:00', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, '宁波杭州湾10#地块（春来云潮）', 0, 0, 3, 0, 0, 0, 6, 1, 0, '宁波杭州湾新区海鹏置业有限公司', 2, 26, '韩非', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '', NULL, '2022-06-15 10:45:22', '2022-06-20 16:27:51', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, '宁波北仑未来社区', 0, 0, 3, 0, 0, 0, 12, 0, 0, '宁波北仑绿城未来社区置业有限公司', 2, 32, '翁跃丰', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-15 13:34:52', '2022-06-20 16:27:45', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, '宁波东钱湖地铁上盖地块（云栖桃花源）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波轨道交通绿城钱湖置业有限公司', 2, 32, '翁跃丰', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-15 14:43:12', '2022-06-20 16:27:36', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, '宁波宁丰路项目（滨河鸣翠）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波绿城浙凡置业有限公司', 2, 32, '翁跃丰', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-15 14:55:38', '2022-06-20 16:27:30', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, '宁波镇海植物园地块（春语云树）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波绿城浙砾置业有限公司', 2, 26, '韩非', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-15 16:09:05', '2022-06-20 16:27:23', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, '宁波江南路项目（春熙云境）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波绿城浙硕置业有限公司', 2, 32, '翁跃丰', 10, '区域一部', '浙江省/杭州市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-15 16:28:58', '2022-06-20 16:27:18', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, '宁波姚江8号地块（滨河沁月）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波绿城甬姚置业有限公司', 2, 32, '翁跃丰', 10, '区域一部', '浙江省/杭州市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-15 16:37:43', '2022-06-20 16:27:05', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, '宁波镇海骆驼项目（新桂沁澜）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波绿城甬勋置业有限公司', 2, 26, '韩非', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-20 15:20:34', '2022-06-20 16:26:58', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, '宁波蛟川项目（春语文澜）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波绿城甬蛟置业有限公司', 2, 26, '韩非', 10, '区域一部', '浙江省/宁波市/市辖区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-20 15:27:17', '2022-06-20 16:26:53', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, '宁波奉化江口三号地块(锦上月鸣）', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波奉化绿城甬承房地产开发有限公司', 2, 26, '韩非', 10, '区域一部', '浙江省/宁波市/奉化区', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同之一', NULL, '2022-06-20 15:36:50', '2022-06-20 16:26:46', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, '诸暨绿城广场', 0, 0, 1, 0, 0, 1, 1, 1, 0, '诸暨市越都置业有限公司', 2, 32, '翁跃丰', 20, '区域六部', '浙江省/绍兴市/诸暨市', '2021-08-15', '2022-08-14', 48, '吴利军', 48, '吴利军', '宁波打包合同赠送合同', '2022-06-20 15:48:23', NULL, '2022-06-20 16:26:41', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (25, '宁波璟丽时尚中心项目（中哲）', 1, 1, 0, 0, 0, 0, 30, 1, 3, '宁波中汇投资有限公司', 1, 26, '韩非', 10, '区域一部', '浙江省/宁波市/鄞州区', '2021-06-01', '2022-05-30', 48, '吴利军', 48, '吴利军', '已结束', '2022-06-20 16:04:29', NULL, '2022-06-20 16:26:04', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (26, '宁波文创港', 1, 1, 3, 0, 0, 0, 9.6, 1, 0, '宁波文创港环球产城发展有限公司', 2, 26, '韩非', 10, '区域一部', '浙江省/宁波市/市辖区', '2022-01-01', '2022-12-31', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-20 16:33:18', '2022-06-20 16:36:51', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (27, '奉化凤悦印湖', 0, 0, 3, 0, 0, 0, 12, 1, 0, '宁波奉化甬轩房地产开发有限公司', 0, 26, '韩非', 10, '区域一部', '浙江省/宁波市/奉化区', '2021-12-01', '2022-11-30', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-27 13:39:35', '2022-06-27 13:41:37', NULL);
INSERT INTO `wdk_project` (`id`, `name`, `type`, `origin`, `step`, `file_upload_status`, `business_type`, `contract_status`, `contract_sum`, `deep_culture`, `status`, `entrust_company`, `sign_company`, `principal_uid`, `principal_name`, `dept_id`, `dept_name`, `region`, `start_time`, `end_time`, `create_by`, `create_name`, `updated_by`, `updated_name`, `remark`, `file_upload_last_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (28, '嵊州·雲庐', 0, 0, 3, 0, 0, 0, 18, 0, 0, '浙江红贝投资有限公司', 1, 32, '翁跃丰', 20, '区域六部', '浙江省/绍兴市/嵊州市', '2021-01-01', '2021-12-31', 48, '吴利军', NULL, NULL, '', NULL, '2022-06-27 13:48:29', '2022-06-27 13:54:18', NULL);
COMMIT;

-- ----------------------------
-- Table structure for wdk_project_businessforms
-- ----------------------------
DROP TABLE IF EXISTS `wdk_project_businessforms`;
CREATE TABLE `wdk_project_businessforms` (
  `project_id` bigint unsigned NOT NULL COMMENT '文档库项目ID',
  `business_forms` tinyint unsigned NOT NULL COMMENT '业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房',
  PRIMARY KEY (`project_id`,`business_forms`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_project_businessforms
-- ----------------------------
BEGIN;
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (1, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (2, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (2, 16);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (3, 3);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (3, 4);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (3, 11);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (3, 13);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (3, 15);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (4, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (5, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (5, 2);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (6, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (6, 2);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (7, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (7, 12);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (7, 15);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (7, 18);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (8, 4);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (8, 15);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (8, 18);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (9, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (9, 2);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (9, 15);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (10, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (11, 5);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (11, 7);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (11, 9);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (12, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (13, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (14, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (15, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (16, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (17, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (18, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (19, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (20, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (21, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (22, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (23, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (24, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (25, 4);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (25, 15);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (25, 18);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (26, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (26, 4);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (26, 13);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (26, 15);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (26, 18);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (27, 0);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (27, 16);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (28, 5);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (28, 7);
INSERT INTO `wdk_project_businessforms` (`project_id`, `business_forms`) VALUES (28, 9);
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
  `created_at` datetime DEFAULT NULL COMMENT '上传时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`,`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=184 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (1, 1, '德清县舞阳街道2021-001号地块项目', '20220324德清晓月西地块地下室图纸的意见和建议.ppt', 48, '吴利军', 0, 3, 0, '2022-06-07 01:02:22', 'wdk/report/ckj7ond5is27nuxwnx874u.ppt', 'wdk/report/ckj7ond5is27nuxwnx874u.pdf', '2022-06-07 01:02:22', '2022-06-07 01:02:22');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (2, 1, '德清县舞阳街道2021-001号地块项目', '20220407德清晓月西地块总图及景观概念方案的意见和建议（修改稿）.ppt', 48, '吴利军', 0, 3, 0, '2022-06-07 01:02:51', 'wdk/report/ckj7ozdatqu6mxjfjgjhdx.ppt', 'wdk/report/ckj7ozdatqu6mxjfjgjhdx.pdf', '2022-06-07 01:02:51', '2022-06-07 01:02:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (3, 1, '德清县舞阳街道2021-001号地块项目', '20220420德清晓月西地块建筑图纸的意见和建议.ppt', 48, '吴利军', 0, 3, 0, '2022-06-07 01:04:32', 'wdk/report/ckj7qb3wddqxtseqkrmdkv.ppt', 'wdk/report/ckj7qb3wddqxtseqkrmdkv.pdf', '2022-06-07 01:04:32', '2022-06-07 01:04:32');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (4, 2, '凤悦印印湖府项目', '20211117关于宁波凤悦印湖项目不利因素梳理的意见和建议.docx', 48, '吴利军', 0, 3, 0, '2022-06-08 17:38:33', 'wdk/report/ckknhzxrk8xgghwdkmjasq.docx', 'wdk/report/ckknhzxrk8xgghwdkmjasq.pdf', '2022-06-08 17:38:33', '2022-06-08 17:38:33');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (5, 2, '凤悦印印湖府项目', '20210726关于奉化68号地块总平图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-08 17:40:27', 'wdk/report/ckknjfbke7fakrdgms4act.pptx', 'wdk/report/ckknjfbke7fakrdgms4act.pdf', '2022-06-08 17:40:27', '2022-06-08 17:40:27');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (6, 2, '凤悦印印湖府项目', '20210831关于奉化68号地块出入口管控的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-08 17:40:53', 'wdk/report/ckknjsl89zxhuy0vnop8ds.pptx', 'wdk/report/ckknjsl89zxhuy0vnop8ds.pdf', '2022-06-08 17:40:53', '2022-06-08 17:40:53');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (7, 3, '温州鹿城广场', '2022.1.3温州鹿城广场三、四期关于电梯维保条款及事项的意见及建议1.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 11:19:48', 'wdk/report/ckookqypbakwhazozecwhq.doc', 'wdk/report/ckookqypbakwhazozecwhq.pdf', '2022-06-13 11:19:48', '2022-06-13 11:19:48');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (8, 3, '温州鹿城广场', '2022.1.7关于温州鹿城广场三、四期关于业主访客流线的专题建议报告.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 11:21:15', 'wdk/report/ckoolu735ln15pmmge5sj2.pptx', 'wdk/report/ckoolu735ln15pmmge5sj2.pdf', '2022-06-13 11:21:15', '2022-06-13 11:21:15');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (9, 3, '温州鹿城广场', '2022.2.15温州鹿城广场关于四期塔楼景观方案的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-13 11:21:47', 'wdk/report/ckoomai42dqxijpfchwgmv.pdf', 'wdk/report/ckoomai42dqxijpfchwgmv.pdf', '2022-06-13 11:21:47', '2022-06-13 11:21:47');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (10, 3, '温州鹿城广场', '2022.1.3温州鹿城广场三、四期关于电梯招标合同注意事项的意见及建议 - 副本.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 11:24:00', 'wdk/report/ckoonz19hrrhj4m3gfmtxg.doc', 'wdk/report/ckoonz19hrrhj4m3gfmtxg.pdf', '2022-06-13 11:24:00', '2022-06-13 11:24:00');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (11, 4, '温州桂语江南', '20220118关于温州桂语江南项目施工现场巡视的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 13:25:47', 'wdk/report/ckor96urn7qk9eelojiwbx.doc', 'wdk/report/ckor96urn7qk9eelojiwbx.pdf', '2022-06-13 13:25:47', '2022-06-13 13:25:47');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (12, 4, '温州桂语江南', '2022.2.14温州桂语江南项目智能化施工图的意见及建议.ppt', 48, '吴利军', 0, 3, 0, '2022-06-13 13:26:54', 'wdk/report/ckora0otrnyxt1clnxmq3k.ppt', 'wdk/report/ckora0otrnyxt1clnxmq3k.pdf', '2022-06-13 13:26:54', '2022-06-13 13:26:54');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (13, 4, '温州桂语江南', '2022.2.25温州桂语江南项目物业用房平面布局的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 13:27:32', 'wdk/report/ckoraif49jc9rwhblpcvbr.pptx', 'wdk/report/ckoraif49jc9rwhblpcvbr.pdf', '2022-06-13 13:27:32', '2022-06-13 13:27:32');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (14, 4, '温州桂语江南', '2022.3.22温州桂语江南项目物业用房平面布局的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 13:28:28', 'wdk/report/ckorb7k3dwtw8tngf1bxbt.pptx', 'wdk/report/ckorb7k3dwtw8tngf1bxbt.pdf', '2022-06-13 13:28:28', '2022-06-13 13:28:28');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (15, 4, '温州桂语江南', '2022.3.23温州桂语江南各口部流线分析报告.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 13:29:27', 'wdk/report/ckorbxmx8ydx3hirasuzej.pptx', 'wdk/report/ckorbxmx8ydx3hirasuzej.pdf', '2022-06-13 13:29:27', '2022-06-13 13:29:27');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (16, 5, '温州春月江澜项目', '20220124关于温州春月江澜项目园区大堂平面图的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 13:39:03', 'wdk/report/ckorjd2biq873jchchunlk.doc', 'wdk/report/ckorjd2biq873jchchunlk.pdf', '2022-06-13 13:39:03', '2022-06-13 13:39:03');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (17, 5, '温州春月江澜项目', '2022.1.17春月江澜售楼处样板间现场巡视的意见和建议1.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 13:40:31', 'wdk/report/ckorkg3lmxo3xjfktfyrvg.pptx', 'wdk/report/ckorkg3lmxo3xjfktfyrvg.pdf', '2022-06-13 13:40:31', '2022-06-13 13:40:31');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (18, 5, '温州春月江澜项目', '2022.2.12州春月江澜项目物业用房功能布置反提.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 14:38:02', 'wdk/report/ckossiy9njmxgovibc1zp4.doc', 'wdk/report/ckossiy9njmxgovibc1zp4.pdf', '2022-06-13 14:38:02', '2022-06-13 14:38:02');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (19, 5, '温州春月江澜项目', '2022.2.23绿城春月江澜项目景观扩初施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 14:38:30', 'wdk/report/ckossuc3ek1gpk6vu9h5yg.pptx', 'wdk/report/ckossuc3ek1gpk6vu9h5yg.pdf', '2022-06-13 14:38:30', '2022-06-13 14:38:30');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (20, 7, '南湖未来社区', '2022.2.18温州南湖未来社区一期项目关于弱电施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 15:14:26', 'wdk/report/ckotkaj5dsa93dexfsauuf.pptx', 'wdk/report/ckotkaj5dsa93dexfsauuf.pdf', '2022-06-13 15:14:26', '2022-06-13 15:14:26');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (21, 7, '南湖未来社区', '2022.2.20温州南湖未来社区项目04地块景观施工图的意见及建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 15:15:55', 'wdk/report/ckotlfdxtq5l0clkwem62f.pptx', 'wdk/report/ckotlfdxtq5l0clkwem62f.pdf', '2022-06-13 15:15:55', '2022-06-13 15:15:55');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (22, 7, '南湖未来社区', '2022.3.21温州南湖未来社区05、08地块智能化调整方案的意见和建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 15:16:44', 'wdk/report/ckotm5u8lqownewwrarpny.doc', 'wdk/report/ckotm5u8lqownewwrarpny.pdf', '2022-06-13 15:16:44', '2022-06-13 15:16:44');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (23, 7, '南湖未来社区', '2022.4.21南湖未来社区05、08地块智能化调整施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 15:17:53', 'wdk/report/ckotmxbrytmint04djh1bz.pptx', 'wdk/report/ckotmxbrytmint04djh1bz.pdf', '2022-06-13 15:17:53', '2022-06-13 15:17:53');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (24, 7, '南湖未来社区', '南湖未来社区B08、D05地块景观施工图的意见及建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 15:19:56', 'wdk/report/ckotohbxx7166jrhocl3nz.pptx', 'wdk/report/ckotohbxx7166jrhocl3nz.pdf', '2022-06-13 15:19:56', '2022-06-13 15:19:56');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (25, 6, '瑞安兰园', '2022.4.18瑞安兰园地库地坪漆方案评审会的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 15:26:26', 'wdk/report/ckottkq7bk6hbzxupgpwvh.doc', 'wdk/report/ckottkq7bk6hbzxupgpwvh.pdf', '2022-06-13 15:26:26', '2022-06-13 15:26:26');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (26, 6, '瑞安兰园', '20220324温州瑞安兰园实体样板房现场巡视的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 15:27:08', 'wdk/report/ckotu4cx86mv9lewb3fynp.doc', 'wdk/report/ckotu4cx86mv9lewb3fynp.pdf', '2022-06-13 15:27:08', '2022-06-13 15:27:08');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (27, 6, '瑞安兰园', '20220601温州瑞安兰园标识标牌方案评审会的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 15:27:39', 'wdk/report/ckotuimqqlvzx69mifwair.doc', 'wdk/report/ckotuimqqlvzx69mifwair.pdf', '2022-06-13 15:27:39', '2022-06-13 15:27:39');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (28, 8, '温州七都科技岛项目', '2022.05.08温州七都国际科技岛项目关于景观示范区延展区概念方案的咨询建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 15:39:45', 'wdk/report/ckou3rx0hmkzp3xdlpyr2s.doc', 'wdk/report/ckou3rx0hmkzp3xdlpyr2s.pdf', '2022-06-13 15:39:45', '2022-06-13 15:39:45');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (29, 8, '温州七都科技岛项目', '温州七都科技岛项目关于项目建筑施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 15:44:08', 'wdk/report/ckou6i2jmata0bcpbowe3i.pptx', 'wdk/report/ckou6i2jmata0bcpbowe3i.pdf', '2022-06-13 15:44:08', '2022-06-13 15:44:08');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (30, 9, '温岭豪成栖江悦项目（豪成新河镇XH050802-04地块）', '2022.1.1温岭豪成项目关于景观总图方案的意见和建议(1).pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 16:11:06', 'wdk/report/ckourqd0f16dz3kre7mp96.pptx', 'wdk/report/ckourqd0f16dz3kre7mp96.pdf', '2022-06-13 16:11:06', '2022-06-13 16:11:06');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (31, 9, '温岭豪成栖江悦项目（豪成新河镇XH050802-04地块）', '2022.01.26豪成栖江悦项目关于售楼部平面的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 16:11:35', 'wdk/report/ckous5fty6hxfsnfehrb02.doc', 'wdk/report/ckous5fty6hxfsnfehrb02.pdf', '2022-06-13 16:11:35', '2022-06-13 16:11:35');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (32, 9, '温岭豪成栖江悦项目（豪成新河镇XH050802-04地块）', '2022.4.15温岭栖江悦项目关于施工场地大门及围挡设置的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 16:13:12', 'wdk/report/ckoutacz248f428edbizzl.pptx', 'wdk/report/ckoutacz248f428edbizzl.pdf', '2022-06-13 16:13:12', '2022-06-13 16:13:12');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (33, 9, '温岭豪成栖江悦项目（豪成新河镇XH050802-04地块）', '2022.5.26温岭栖江悦项目关于智能化施工图的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-13 16:14:28', 'wdk/report/ckouudpl0v39hbdpld09dy.pdf', 'wdk/report/ckouudpl0v39hbdpld09dy.pdf', '2022-06-13 16:14:28', '2022-06-13 16:14:28');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (34, 10, '台州晓风印月（椒江区现代大道北地块）', '2022.3.16台州绿城现代大道北侧、台州大道东侧出让地块项目规划方案现场会的意见及建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-13 16:25:11', 'wdk/report/ckov28loqagydzt7u6xcqe.pptx', 'wdk/report/ckov28loqagydzt7u6xcqe.pdf', '2022-06-13 16:25:11', '2022-06-13 16:25:11');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (35, 10, '台州晓风印月（椒江区现代大道北地块）', '2022.3.30台州绿城新项目景观概念方案评审会的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 16:26:04', 'wdk/report/ckov38hm5w7mzikaqevsbt.doc', 'wdk/report/ckov38hm5w7mzikaqevsbt.pdf', '2022-06-13 16:26:04', '2022-06-13 16:26:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (36, 10, '台州晓风印月（椒江区现代大道北地块）', '2022.4.18台州绿城新项目规划方案终稿会评审会的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 16:26:42', 'wdk/report/ckov3pwussl99cb3gvc8mw.doc', 'wdk/report/ckov3pwussl99cb3gvc8mw.pdf', '2022-06-13 16:26:42', '2022-06-13 16:26:42');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (37, 10, '台州晓风印月（椒江区现代大道北地块）', '2022.5.5台州晓风印月项目大区景观方案评审会的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 16:28:31', 'wdk/report/ckov53vgv8ph1lw44lwof6.doc', 'wdk/report/ckov53vgv8ph1lw44lwof6.pdf', '2022-06-13 16:28:31', '2022-06-13 16:28:31');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (38, 10, '台州晓风印月（椒江区现代大道北地块）', '2022.5.18台州晓风印月项目199户型样板房概念方案及售楼部平面概念方案的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 16:29:52', 'wdk/report/ckov659uanui7nvz4dxgj4.doc', 'wdk/report/ckov659uanui7nvz4dxgj4.pdf', '2022-06-13 16:29:52', '2022-06-13 16:29:52');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (39, 10, '台州晓风印月（椒江区现代大道北地块）', '2022.6.6台州晓风印月项目售楼部精装修设计评审会的意见及建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-13 16:30:16', 'wdk/report/ckov6gcxxqwjxjrxwzxqip.doc', 'wdk/report/ckov6gcxxqwjxjrxwzxqip.pdf', '2022-06-13 16:30:16', '2022-06-13 16:30:16');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (40, 12, '宁波高新13号地块（春月云锦）', '20200723关于宁波高新13号地块启动会报告的意见和建议.ppt', 48, '吴利军', 0, 3, 0, '2022-06-14 15:09:02', 'wdk/report/ckpo2s8ks4mfepxp91v7j5.ppt', 'wdk/report/ckpo2s8ks4mfepxp91v7j5.pdf', '2022-06-14 15:09:02', '2022-06-14 15:09:02');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (41, 12, '宁波高新13号地块（春月云锦）', '20200819关于宁波高新13号地块项目景观概念方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:09:40', 'wdk/report/ckpo3a2imgpli8aeqrmrw3.pptx', 'wdk/report/ckpo3a2imgpli8aeqrmrw3.pdf', '2022-06-14 15:09:40', '2022-06-14 15:09:40');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (42, 12, '宁波高新13号地块（春月云锦）', '20200917关于宁波高新13号地块项目景观深化方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:10:05', 'wdk/report/ckpo3li31pxa83rmo2nns0.pptx', 'wdk/report/ckpo3li31pxa83rmo2nns0.pdf', '2022-06-14 15:10:05', '2022-06-14 15:10:05');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (43, 12, '宁波高新13号地块（春月云锦）', '20201118关于宁波高新区13号地块项目精装修图纸评审的意见.docx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:10:18', 'wdk/report/ckpo3seg3wq5hemcsn840w.docx', 'wdk/report/ckpo3seg3wq5hemcsn840w.pdf', '2022-06-14 15:10:18', '2022-06-14 15:10:18');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (44, 12, '宁波高新13号地块（春月云锦）', '20201203宁波高新13#地块项目公区精装扩初施工图评审会会议纪要-1.doc', 48, '吴利军', 0, 3, 0, '2022-06-14 15:10:34', 'wdk/report/ckpo3zl1orxip90sna5ftu.doc', 'wdk/report/ckpo3zl1orxip90sna5ftu.pdf', '2022-06-14 15:10:34', '2022-06-14 15:10:34');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (45, 12, '宁波高新13号地块（春月云锦）', '20201231关于宁波春月云锦项目建筑方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:10:59', 'wdk/report/ckpo49xcecvybjrusscy8k.pptx', 'wdk/report/ckpo49xcecvybjrusscy8k.pdf', '2022-06-14 15:10:59', '2022-06-14 15:10:59');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (46, 12, '宁波高新13号地块（春月云锦）', '20210106关于宁波春月云锦项目智能化方案评审的意见.docx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:11:13', 'wdk/report/ckpo4haew428oi6cmjarns.docx', 'wdk/report/ckpo4haew428oi6cmjarns.pdf', '2022-06-14 15:11:13', '2022-06-14 15:11:13');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (47, 12, '宁波高新13号地块（春月云锦）', '20210802关于宁波绿城高新GX05-01-13地块物业用房配置的建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:11:51', 'wdk/report/ckpo4y7wetzjuzkolcbgah.pptx', 'wdk/report/ckpo4y7wetzjuzkolcbgah.pdf', '2022-06-14 15:11:51', '2022-06-14 15:11:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (48, 12, '宁波高新13号地块（春月云锦）', '20210909鄞州区GX05-01-13地块项目智能化设计的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-14 15:12:26', 'wdk/report/ckpo5fgnmmw6ian5u4besb.pdf', 'wdk/report/ckpo5fgnmmw6ian5u4besb.pdf', '2022-06-14 15:12:26', '2022-06-14 15:12:26');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (49, 12, '宁波高新13号地块（春月云锦）', '20220116春月云锦项目关于地下室除湿机设置点位的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:13:15', 'wdk/report/ckpo60z2hlqayupnw6he4s.pptx', 'wdk/report/ckpo60z2hlqayupnw6he4s.pdf', '2022-06-14 15:13:15', '2022-06-14 15:13:15');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (50, 12, '宁波高新13号地块（春月云锦）', '20220307宁波春月云锦项目关于地下室保洁取水点的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:14:08', 'wdk/report/ckpo6p5n9xs3hmc9rd6umo.pptx', 'wdk/report/ckpo6p5n9xs3hmc9rd6umo.pdf', '2022-06-14 15:14:08', '2022-06-14 15:14:08');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (51, 12, '宁波高新13号地块（春月云锦）', '20220415宁波春月云锦关于交通流线管控的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:14:54', 'wdk/report/ckpo77k4srcb4l8fycif7x.pptx', 'wdk/report/ckpo77k4srcb4l8fycif7x.pdf', '2022-06-14 15:14:54', '2022-06-14 15:14:54');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (52, 12, '宁波高新13号地块（春月云锦）', '20220420宁波春月云锦关于实体样板房现场巡视的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:15:47', 'wdk/report/ckpo7w4eq63hgpd8uxep7p.pptx', 'wdk/report/ckpo7w4eq63hgpd8uxep7p.pdf', '2022-06-14 15:15:47', '2022-06-14 15:15:47');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (53, 13, '宁波杭州湾5#地块（春熙月明）', '20210720关于宁波杭州湾5#地块项目弱电图纸的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:21:16', 'wdk/report/ckpoc4x4mk1z45xhgbc1z3.pptx', 'wdk/report/ckpoc4x4mk1z45xhgbc1z3.pdf', '2022-06-14 15:21:16', '2022-06-14 15:21:16');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (54, 13, '宁波杭州湾5#地块（春熙月明）', '20210917杭州湾5#地块景观施工图评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-14 15:21:27', 'wdk/report/ckpocc6v3g77lytdcu1lfz.pdf', 'wdk/report/ckpocc6v3g77lytdcu1lfz.pdf', '2022-06-14 15:21:27', '2022-06-14 15:21:27');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (55, 13, '宁波杭州湾5#地块（春熙月明）', '20210926杭州湾5#地块泛光设计方案评审会会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-14 15:21:40', 'wdk/report/ckpoci7rcxkto372hb5rul.pdf', 'wdk/report/ckpoci7rcxkto372hb5rul.pdf', '2022-06-14 15:21:40', '2022-06-14 15:21:40');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (56, 13, '宁波杭州湾5#地块（春熙月明）', '20211202关于杭州湾春熙月明标识系统方案的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-14 15:22:04', 'wdk/report/ckpoctbibyi6ty4q1ua7jf.pdf', 'wdk/report/ckpoctbibyi6ty4q1ua7jf.pdf', '2022-06-14 15:22:04', '2022-06-14 15:22:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (57, 13, '宁波杭州湾5#地块（春熙月明）', '20211203杭州湾春熙月明项目关于地下室保洁取水点及除湿机设置的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-14 15:22:42', 'wdk/report/ckpodagqagmnlmhvxvaz2d.pdf', 'wdk/report/ckpodagqagmnlmhvxvaz2d.pdf', '2022-06-14 15:22:42', '2022-06-14 15:22:42');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (58, 13, '宁波杭州湾5#地块（春熙月明）', '20211203杭州湾春熙月明项目关于地下室保洁取水点及除湿机设置的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:23:33', 'wdk/report/ckpodvns13pgbs2xytcbny.pptx', 'wdk/report/ckpodvns13pgbs2xytcbny.pdf', '2022-06-14 15:23:33', '2022-06-14 15:23:33');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (59, 13, '宁波杭州湾5#地块（春熙月明）', '20211220杭州湾春熙月明项目物业交通管理方案.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:24:02', 'wdk/report/ckpoe9qe1bc5lb5xrhi97h.pptx', 'wdk/report/ckpoe9qe1bc5lb5xrhi97h.pdf', '2022-06-14 15:24:02', '2022-06-14 15:24:02');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (60, 13, '宁波杭州湾5#地块（春熙月明）', '20211222杭州湾5#、10#地块项目导示系统方案评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-14 15:24:32', 'wdk/report/ckpoeolyq8v4fmmr7i03qu.doc', 'wdk/report/ckpoeolyq8v4fmmr7i03qu.pdf', '2022-06-14 15:24:32', '2022-06-14 15:24:32');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (61, 13, '宁波杭州湾5#地块（春熙月明）', '20220124杭州湾春熙月明项目物业交通管理方案.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:25:09', 'wdk/report/ckpof4kapqxfi7w1qyehoz.pptx', 'wdk/report/ckpof4kapqxfi7w1qyehoz.pdf', '2022-06-14 15:25:09', '2022-06-14 15:25:09');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (62, 13, '宁波杭州湾5#地块（春熙月明）', '20220214杭州湾春熙月明项目室外管综评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-14 15:25:34', 'wdk/report/ckpofgu0s5ntuefsibjh7z.doc', 'wdk/report/ckpofgu0s5ntuefsibjh7z.pdf', '2022-06-14 15:25:34', '2022-06-14 15:25:34');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (63, 13, '宁波杭州湾5#地块（春熙月明）', '20220525宁波春熙月明项目现场巡视报告.pptx', 48, '吴利军', 0, 3, 0, '2022-06-14 15:26:34', 'wdk/report/ckpog5n7il5wkdikkusgpo.pptx', 'wdk/report/ckpog5n7il5wkdikkusgpo.pdf', '2022-06-14 15:26:34', '2022-06-14 15:26:34');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (64, 13, '宁波杭州湾5#地块（春熙月明）', '202107022杭州湾春熙月明（5#地块）项目智能化扩初图评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-14 15:27:04', 'wdk/report/ckpogmxd2hhtj6aj2fjw39.pdf', 'wdk/report/ckpogmxd2hhtj6aj2fjw39.pdf', '2022-06-14 15:27:04', '2022-06-14 15:27:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (65, 13, '宁波杭州湾5#地块（春熙月明）', '202107021杭州湾春熙月明（5#地块）项目景观扩初图评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-14 15:27:19', 'wdk/report/ckpogu0wzv34fmilygbas5.pdf', 'wdk/report/ckpogu0wzv34fmilygbas5.pdf', '2022-06-14 15:27:19', '2022-06-14 15:27:19');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (66, 14, '宁波杭州湾10#地块（春来云潮）', '20200205杭州湾新区10#地块项目景概念方案评审会会议意见(1).doc', 48, '吴利军', 0, 3, 0, '2022-06-15 10:46:44', 'wdk/report/ckqd4ifop3ffedjotfolr5.doc', 'wdk/report/ckqd4ifop3ffedjotfolr5.pdf', '2022-06-15 10:46:44', '2022-06-15 10:46:44');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (67, 14, '宁波杭州湾10#地块（春来云潮）', '20210304宁波杭州湾新区10#地块大区景观概念方案及示范区方案评审会纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 10:48:21', 'wdk/report/ckqd5rhr5nlq0bmyngw78m.doc', 'wdk/report/ckqd5rhr5nlq0bmyngw78m.pdf', '2022-06-15 10:48:21', '2022-06-15 10:48:21');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (68, 14, '宁波杭州湾10#地块（春来云潮）', '20210903杭州湾10#地块关于智能化图纸的意见与建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 10:50:58', 'wdk/report/ckqd7s3eooermj4kzfptlt.pdf', 'wdk/report/ckqd7s3eooermj4kzfptlt.pdf', '2022-06-15 10:50:58', '2022-06-15 10:50:58');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (69, 14, '宁波杭州湾10#地块（春来云潮）', '20210906杭州湾10#地块十号地块智能化扩初评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 10:51:42', 'wdk/report/ckqd8cb6vr4qcvwpyagam2.pdf', 'wdk/report/ckqd8cb6vr4qcvwpyagam2.pdf', '2022-06-15 10:51:42', '2022-06-15 10:51:42');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (70, 14, '宁波杭州湾10#地块（春来云潮）', '20210916杭州湾10#地块景观施工图评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 10:52:01', 'wdk/report/ckqd8lafhkmabkg9ugkxjk.pdf', 'wdk/report/ckqd8lafhkmabkg9ugkxjk.pdf', '2022-06-15 10:52:01', '2022-06-15 10:52:01');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (71, 14, '宁波杭州湾10#地块（春来云潮）', '20210926杭州湾10#地块泛光设计方案评审会会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 10:54:34', 'wdk/report/ckqdajlc3279e3vf2jjfot.pdf', 'wdk/report/ckqdajlc3279e3vf2jjfot.pdf', '2022-06-15 10:54:34', '2022-06-15 10:54:34');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (72, 14, '宁波杭州湾10#地块（春来云潮）', '20211203杭州湾春来云潮项目关于地下室保洁取水点及除湿机设置的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 13:06:56', 'wdk/report/ckqg3w5k1q7frqwbg5nlcb.pdf', 'wdk/report/ckqg3w5k1q7frqwbg5nlcb.pdf', '2022-06-15 13:06:56', '2022-06-15 13:06:56');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (73, 14, '宁波杭州湾10#地块（春来云潮）', '20211216杭州湾春来云潮项目物业交通管理方案.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 13:12:01', 'wdk/report/ckqg7qbpuv3j5ja72vi4hf.pptx', 'wdk/report/ckqg7qbpuv3j5ja72vi4hf.pdf', '2022-06-15 13:12:01', '2022-06-15 13:12:01');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (74, 14, '宁波杭州湾10#地块（春来云潮）', '20211220关于杭州湾春来云潮标识系统方案的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 13:12:35', 'wdk/report/ckqg87zqree9bd2vsl6ivd.pdf', 'wdk/report/ckqg87zqree9bd2vsl6ivd.pdf', '2022-06-15 13:12:35', '2022-06-15 13:12:35');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (75, 14, '宁波杭州湾10#地块（春来云潮）', '20211222杭州湾5#、10#地块项目导示系统方案评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 13:13:03', 'wdk/report/ckqg8k1aa41ufwoiasrd3e.doc', 'wdk/report/ckqg8k1aa41ufwoiasrd3e.pdf', '2022-06-15 13:13:03', '2022-06-15 13:13:03');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (76, 14, '宁波杭州湾10#地块（春来云潮）', '20220331杭州湾5#、10#地块关于架空层方案及景观软装方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 13:14:19', 'wdk/report/ckqg9fy3rulatgs4xzwkkb.pptx', 'wdk/report/ckqg9fy3rulatgs4xzwkkb.pdf', '2022-06-15 13:14:19', '2022-06-15 13:14:19');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (77, 15, '宁波北仑未来社区', '20210721宁波北仑未来社区景观概念设计评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 13:36:06', 'wdk/report/ckqgq801c0yy6heovcbfs3.pdf', 'wdk/report/ckqgq801c0yy6heovcbfs3.pdf', '2022-06-15 13:36:06', '2022-06-15 13:36:06');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (78, 15, '宁波北仑未来社区', '20210813宁波北仑未来社区景观方案评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 13:36:19', 'wdk/report/ckqgqdtqbkt76fklgx42ij.pdf', 'wdk/report/ckqgqdtqbkt76fklgx42ij.pdf', '2022-06-15 13:36:19', '2022-06-15 13:36:19');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (79, 15, '宁波北仑未来社区', '20210908北仑未来社区方案终稿会（住宅区）评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 13:36:43', 'wdk/report/ckqgqou6dwbl1tweloikd1.pdf', 'wdk/report/ckqgqou6dwbl1tweloikd1.pdf', '2022-06-15 13:36:43', '2022-06-15 13:36:43');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (80, 15, '宁波北仑未来社区', '20210920北仑未来社区项目关于门窗防渗漏专项的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 13:37:37', 'wdk/report/ckqgrbpb3b6yaohbudljz6.pptx', 'wdk/report/ckqgrbpb3b6yaohbudljz6.pdf', '2022-06-15 13:37:37', '2022-06-15 13:37:37');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (81, 15, '宁波北仑未来社区', '20211014关于北仑未来社区智慧园区硬件设备配置标准要求的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 13:37:56', 'wdk/report/ckqgrllv8p53xflzcdikqi.doc', 'wdk/report/ckqgrllv8p53xflzcdikqi.pdf', '2022-06-15 13:37:56', '2022-06-15 13:37:56');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (82, 15, '宁波北仑未来社区', '20211023关于北仑未来社区电梯功能定位的意见和建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 13:38:09', 'wdk/report/ckqgrrko9k07tcpilyiay8.doc', 'wdk/report/ckqgrrko9k07tcpilyiay8.pdf', '2022-06-15 13:38:09', '2022-06-15 13:38:09');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (83, 15, '宁波北仑未来社区', '20211116宁波北仑未来社区2#、5#地块景观扩初图评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 13:38:23', 'wdk/report/ckqgrygtqe2lmk5s01jfty.doc', 'wdk/report/ckqgrygtqe2lmk5s01jfty.pdf', '2022-06-15 13:38:23', '2022-06-15 13:38:23');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (84, 15, '宁波北仑未来社区', '20211220北仑未来社区项目物业交通管理方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 13:39:04', 'wdk/report/ckqgsg3jd2u3hpmdbhoapw.pptx', 'wdk/report/ckqgsg3jd2u3hpmdbhoapw.pdf', '2022-06-15 13:39:04', '2022-06-15 13:39:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (85, 15, '宁波北仑未来社区', '20211220关于通山未来社区公区深化方案物业提资的报告.docx', 48, '吴利军', 0, 3, 0, '2022-06-15 13:41:51', 'wdk/report/ckqguluyqny5xzkxvk9ykh.docx', 'wdk/report/ckqguluyqny5xzkxvk9ykh.pdf', '2022-06-15 13:41:51', '2022-06-15 13:41:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (86, 15, '宁波北仑未来社区', '20220118关于通山未来社区项目公区深化方案的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 13:42:59', 'wdk/report/ckqgvhwey1z5zakfl8by2r.pdf', 'wdk/report/ckqgvhwey1z5zakfl8by2r.pdf', '2022-06-15 13:42:59', '2022-06-15 13:42:59');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (87, 15, '宁波北仑未来社区', '20220227关于北仑未来社区5#地块项目智能化施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 13:43:37', 'wdk/report/ckqgvwik2wnqdjsmfqqvoz.pptx', 'wdk/report/ckqgvwik2wnqdjsmfqqvoz.pdf', '2022-06-15 13:43:37', '2022-06-15 13:43:37');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (88, 15, '宁波北仑未来社区', '20220602关于北仑未来社区样板房巡视的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 13:46:26', 'wdk/report/ckqgxz5wfehlcbag2io1ww.pptx', 'wdk/report/ckqgxz5wfehlcbag2io1ww.pdf', '2022-06-15 13:46:26', '2022-06-15 13:46:26');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (89, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '2020年9月25日关于东钱湖地铁上盖景观概念方案的意见和建议(1).doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:44:50', 'wdk/report/ckqi6ttlloutdpbiudkpas.doc', 'wdk/report/ckqi6ttlloutdpbiudkpas.pdf', '2022-06-15 14:44:50', '2022-06-15 14:44:50');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (90, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '2020年9月28日关于东钱湖项目园区垃圾收集配置系统的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 14:45:40', 'wdk/report/ckqi7fhc0ue8pjhxcngj89.pptx', 'wdk/report/ckqi7fhc0ue8pjhxcngj89.pdf', '2022-06-15 14:45:40', '2022-06-15 14:45:40');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (91, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '2020年10月29日关于东钱湖项目示范区景观设计方案的意见与建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 14:46:06', 'wdk/report/ckqi7s7hyvi9flmcd2ujyp.pptx', 'wdk/report/ckqi7s7hyvi9flmcd2ujyp.pdf', '2022-06-15 14:46:06', '2022-06-15 14:46:06');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (92, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20200812关于东钱湖地铁上盖项目方案设计的意见与建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 14:46:37', 'wdk/report/ckqi86qnqa7f3d91fvxolm.pptx', 'wdk/report/ckqi86qnqa7f3d91fvxolm.pdf', '2022-06-15 14:46:37', '2022-06-15 14:46:37');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (93, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20201207关于宁波云栖桃花源项目景观方案设计的意见与建议-2.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 14:47:08', 'wdk/report/ckqi8kvk8lpjqcbk3vzays.pptx', 'wdk/report/ckqi8kvk8lpjqcbk3vzays.pdf', '2022-06-15 14:47:08', '2022-06-15 14:47:08');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (94, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20210705宁波绿城东钱湖地铁上盖项目样板房巡查的建议和意见.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 14:47:56', 'wdk/report/ckqi97zdrzntl3jtk8xfqy.pdf', 'wdk/report/ckqi97zdrzntl3jtk8xfqy.pdf', '2022-06-15 14:47:56', '2022-06-15 14:47:56');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (95, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20210914关于宁波云栖桃花源项目智慧园区硬件设备配置标准要求的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:48:25', 'wdk/report/ckqi9khoo00a6fwjl9kzgk.doc', 'wdk/report/ckqi9khoo00a6fwjl9kzgk.pdf', '2022-06-15 14:48:25', '2022-06-15 14:48:25');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (96, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20211016宁波云栖桃花源关于水电能耗计量设置的意见.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:48:50', 'wdk/report/ckqi9w5fha2fbrghdjzxhy.doc', 'wdk/report/ckqi9w5fha2fbrghdjzxhy.pdf', '2022-06-15 14:48:50', '2022-06-15 14:48:50');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (97, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20211123关于云栖桃花源项目电梯功能定位的意见和建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:49:02', 'wdk/report/ckqia1t4r03qwytswmjbni.doc', 'wdk/report/ckqia1t4r03qwytswmjbni.pdf', '2022-06-15 14:49:02', '2022-06-15 14:49:02');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (98, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20220109关于宁波云栖桃花源项目综合管网施工现场常见问题及注意事项的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:50:24', 'wdk/report/ckqib2zog02d0nu0twyhqg.doc', 'wdk/report/ckqib2zog02d0nu0twyhqg.pdf', '2022-06-15 14:50:24', '2022-06-15 14:50:24');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (99, 16, '宁波东钱湖地铁上盖地块（云栖桃花源）', '20220224云栖桃花源项目关于门窗防渗漏专项的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 14:54:03', 'wdk/report/ckqiduggl6x9yjgubypzvb.pptx', 'wdk/report/ckqiduggl6x9yjgubypzvb.pdf', '2022-06-15 14:54:03', '2022-06-15 14:54:03');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (100, 17, '宁波宁丰路项目（滨河鸣翠）', '20210524关于宁波宁丰路项目景观概念方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 14:57:19', 'wdk/report/ckqigd5vo2jx5rstvgggvk.pptx', 'wdk/report/ckqigd5vo2jx5rstvgggvk.pdf', '2022-06-15 14:57:19', '2022-06-15 14:57:19');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (101, 17, '宁波宁丰路项目（滨河鸣翠）', '20210525宁波绿城宁丰路项目景观概念设计评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:57:38', 'wdk/report/ckqigmg0vs1crcd5crp1a4.doc', 'wdk/report/ckqigmg0vs1crcd5crp1a4.pdf', '2022-06-15 14:57:38', '2022-06-15 14:57:38');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (102, 17, '宁波宁丰路项目（滨河鸣翠）', '20210630宁波绿城宁丰路项目景观概念设计评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:57:50', 'wdk/report/ckqigs9kc507aar3mwebkn.doc', 'wdk/report/ckqigs9kc507aar3mwebkn.pdf', '2022-06-15 14:57:50', '2022-06-15 14:57:50');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (103, 17, '宁波宁丰路项目（滨河鸣翠）', '20211215滨河鸣翠架空层、园区大堂方案评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 14:58:36', 'wdk/report/ckqihcpzqcnwva5o6by4eo.doc', 'wdk/report/ckqihcpzqcnwva5o6by4eo.pdf', '2022-06-15 14:58:36', '2022-06-15 14:58:36');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (104, 17, '宁波宁丰路项目（滨河鸣翠）', '20220428宁波滨河鸣翠项目建筑施工图的意见和建议(1).pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 14:59:31', 'wdk/report/ckqihwvvs7i2qxll2p2ndi.pptx', 'wdk/report/ckqihwvvs7i2qxll2p2ndi.pdf', '2022-06-15 14:59:31', '2022-06-15 14:59:31');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (105, 17, '宁波宁丰路项目（滨河鸣翠）', '20220507关于宁波滨河鸣翠项目公区照明控制方式的意见与建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 15:00:12', 'wdk/report/ckqiihs1thlzjmdcuykqnl.pptx', 'wdk/report/ckqiihs1thlzjmdcuykqnl.pdf', '2022-06-15 15:00:12', '2022-06-15 15:00:12');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (106, 17, '宁波宁丰路项目（滨河鸣翠）', '20220509关于宁波滨河鸣翠项目弱电施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 15:01:17', 'wdk/report/ckqij9r15a0ywbhbp9escy.pptx', 'wdk/report/ckqij9r15a0ywbhbp9escy.pdf', '2022-06-15 15:01:17', '2022-06-15 15:01:17');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (107, 17, '宁波宁丰路项目（滨河鸣翠）', '20220526滨河鸣翠项目物业交通管控方案.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 15:02:04', 'wdk/report/ckqijxwx5q75urecf6sxxl.pptx', 'wdk/report/ckqijxwx5q75urecf6sxxl.pdf', '2022-06-15 15:02:04', '2022-06-15 15:02:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (108, 17, '宁波宁丰路项目（滨河鸣翠）', '20220526宁波滨河鸣翠关于垃圾清运及外卖快递配送方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 15:05:05', 'wdk/report/ckqim3po4ecdznpf3zqzov.pptx', 'wdk/report/ckqim3po4ecdznpf3zqzov.pdf', '2022-06-15 15:05:05', '2022-06-15 15:05:05');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (109, 17, '宁波宁丰路项目（滨河鸣翠）', '20220527宁波滨河鸣翠项目景观施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:04:09', 'wdk/report/ckqjvdg4kc393uo86ocznp.pptx', 'wdk/report/ckqjvdg4kc393uo86ocznp.pdf', '2022-06-15 16:04:09', '2022-06-15 16:04:09');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (110, 17, '宁波宁丰路项目（滨河鸣翠）', '20220531宁波滨河鸣翠项目周例会会议纪要（物业部分）.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:04:27', 'wdk/report/ckqjvs8xyuq5twanygbodd.doc', 'wdk/report/ckqjvs8xyuq5twanygbodd.pdf', '2022-06-15 16:04:27', '2022-06-15 16:04:27');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (111, 18, '宁波镇海植物园地块（春语云树）', '20201112关于镇海植物园地块项目总图的意见与建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:10:19', 'wdk/report/ckqk09erq4ustmdfemibvz.pptx', 'wdk/report/ckqk09erq4ustmdfemibvz.pdf', '2022-06-15 16:10:19', '2022-06-15 16:10:19');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (112, 18, '宁波镇海植物园地块（春语云树）', '20201113关于镇海植物园地块项目景观概念方案评审的意见 - 副本.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:10:46', 'wdk/report/ckqk0m56se9ys4sz4bkchg.pptx', 'wdk/report/ckqk0m56se9ys4sz4bkchg.pdf', '2022-06-15 16:10:46', '2022-06-15 16:10:46');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (113, 18, '宁波镇海植物园地块（春语云树）', '20201229宁波春语云树项目景观方案评审的意见.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:11:02', 'wdk/report/ckqk0ttzv5rgjz2w8c2g8c.doc', 'wdk/report/ckqk0ttzv5rgjz2w8c2g8c.pdf', '2022-06-15 16:11:02', '2022-06-15 16:11:02');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (114, 18, '宁波镇海植物园地块（春语云树）', '20210803宁波镇海春语云树项目关于水电能耗计量设置的意见.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:11:15', 'wdk/report/ckqk0zq33a5zphdt7iwcjr.doc', 'wdk/report/ckqk0zq33a5zphdt7iwcjr.pdf', '2022-06-15 16:11:15', '2022-06-15 16:11:15');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (115, 18, '宁波镇海植物园地块（春语云树）', '20211014关于春语云树项目智慧园区硬件设备配置标准要求的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:11:30', 'wdk/report/ckqk16dcqwd6wlzpohbwje.doc', 'wdk/report/ckqk16dcqwd6wlzpohbwje.pdf', '2022-06-15 16:11:30', '2022-06-15 16:11:30');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (116, 18, '宁波镇海植物园地块（春语云树）', '20211023关于春语云树项目电梯功能定位的意见和建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:11:53', 'wdk/report/ckqk1hfauy33dhdxgf4dnv.doc', 'wdk/report/ckqk1hfauy33dhdxgf4dnv.pdf', '2022-06-15 16:11:53', '2022-06-15 16:11:53');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (117, 18, '宁波镇海植物园地块（春语云树）', '20211109关于镇海春语云树项目综合管网施工现场常见问题及注意事项的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:12:48', 'wdk/report/ckqk26eg8r12lym0kygwxa.doc', 'wdk/report/ckqk26eg8r12lym0kygwxa.pdf', '2022-06-15 16:12:48', '2022-06-15 16:12:48');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (118, 18, '宁波镇海植物园地块（春语云树）', '20220120春语云树项目关于门窗防渗漏专项的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:13:30', 'wdk/report/ckqk2okwvno3spgrndvbcf.pptx', 'wdk/report/ckqk2okwvno3spgrndvbcf.pdf', '2022-06-15 16:13:30', '2022-06-15 16:13:30');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (119, 18, '宁波镇海植物园地块（春语云树）', '20220127宁波镇海春语云树北项目物业交通管理方案.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:13:53', 'wdk/report/ckqk2zwzzkoqef6ccohl80.pptx', 'wdk/report/ckqk2zwzzkoqef6ccohl80.pdf', '2022-06-15 16:13:53', '2022-06-15 16:13:53');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (120, 18, '宁波镇海植物园地块（春语云树）', '20220218宁波春语云树水电施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:14:29', 'wdk/report/ckqk3fybezl62pwbt7c7f1.pptx', 'wdk/report/ckqk3fybezl62pwbt7c7f1.pdf', '2022-06-15 16:14:29', '2022-06-15 16:14:29');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (121, 18, '宁波镇海植物园地块（春语云树）', '20220301镇海春语云树北项目景观概念方案评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:15:03', 'wdk/report/ckqk3wbvbvgnhyya6jtflk.doc', 'wdk/report/ckqk3wbvbvgnhyya6jtflk.pdf', '2022-06-15 16:15:03', '2022-06-15 16:15:03');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (122, 18, '宁波镇海植物园地块（春语云树）', '20220303宁波镇海植物园北项目物业交通管理方案.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:15:32', 'wdk/report/ckqk49ot8a669xhyajyjck.pptx', 'wdk/report/ckqk49ot8a669xhyajyjck.pdf', '2022-06-15 16:15:32', '2022-06-15 16:15:32');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (123, 19, '宁波江南路项目（春熙云境）', '20210331绿城宁波江南路项目和宁波北仑未来社区项目景观概念方案.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:29:45', 'wdk/report/ckqkf5r7fwppbdkr9heace.doc', 'wdk/report/ckqkf5r7fwppbdkr9heace.pdf', '2022-06-15 16:29:45', '2022-06-15 16:29:45');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (124, 19, '宁波江南路项目（春熙云境）', '20210728宁波绿城春熙云境项目样板房巡查的建议和意见(1).pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 16:30:51', 'wdk/report/ckqkg0hizrdofgmr4cgua4.pdf', 'wdk/report/ckqkg0hizrdofgmr4cgua4.pdf', '2022-06-15 16:30:51', '2022-06-15 16:30:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (125, 19, '宁波江南路项目（春熙云境）', '20210923关于宁波江南路项目公区水电点位设置的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:31:24', 'wdk/report/ckqkge6ihqn1mytpbj3zvq.pptx', 'wdk/report/ckqkge6ihqn1mytpbj3zvq.pdf', '2022-06-15 16:31:24', '2022-06-15 16:31:24');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (126, 19, '宁波江南路项目（春熙云境）', '20211027宁波春熙云境项目大区景观施工图评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:32:20', 'wdk/report/ckqkh51n62c7c2hkzagdah.doc', 'wdk/report/ckqkh51n62c7c2hkzagdah.pdf', '2022-06-15 16:32:20', '2022-06-15 16:32:20');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (127, 19, '宁波江南路项目（春熙云境）', '20211028关于宁波春熙云境项目物业交通管控流线设置的建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:32:46', 'wdk/report/ckqkhg0im19ch4fztyiiq3.pptx', 'wdk/report/ckqkhg0im19ch4fztyiiq3.pdf', '2022-06-15 16:32:46', '2022-06-15 16:32:46');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (128, 19, '宁波江南路项目（春熙云境）', '20211112关于宁波春熙云镜项目物业用房功能定位的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:34:21', 'wdk/report/ckqkiln79t9ogxncd8etlk.pptx', 'wdk/report/ckqkiln79t9ogxncd8etlk.pdf', '2022-06-15 16:34:21', '2022-06-15 16:34:21');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (129, 19, '宁波江南路项目（春熙云境）', '20211215春熙云境项目公区架空层、归家门厅、地下门厅概念方案评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-15 16:34:39', 'wdk/report/ckqkiwkdyi21tzsvkpnhgu.doc', 'wdk/report/ckqkiwkdyi21tzsvkpnhgu.pdf', '2022-06-15 16:34:39', '2022-06-15 16:34:39');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (130, 19, '宁波江南路项目（春熙云境）', '20220415关于宁波春熙云镜项目弱电施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:35:31', 'wdk/report/ckqkjfs7lg1fvphgzspdr2.pptx', 'wdk/report/ckqkjfs7lg1fvphgzspdr2.pdf', '2022-06-15 16:35:31', '2022-06-15 16:35:31');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (131, 20, '宁波姚江8号地块（滨河沁月）', '20210701宁波绿城8#-1、8#-2地块景观概念设计方案会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 16:40:15', 'wdk/report/ckqkn7z0tg4vjxpa3l1pif.pdf', 'wdk/report/ckqkn7z0tg4vjxpa3l1pif.pdf', '2022-06-15 16:40:15', '2022-06-15 16:40:15');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (132, 20, '宁波姚江8号地块（滨河沁月）', '20210827宁波绿城姚江8号地块项目景观方案评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 16:40:41', 'wdk/report/ckqknjxzxy1ljtdxeboevd.pdf', 'wdk/report/ckqknjxzxy1ljtdxeboevd.pdf', '2022-06-15 16:40:41', '2022-06-15 16:40:41');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (133, 20, '宁波姚江8号地块（滨河沁月）', '20210923宁波姚江8号地块终审会方案评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-15 16:41:05', 'wdk/report/ckqknuxsclzu7w0jt8l8jt.pdf', 'wdk/report/ckqknuxsclzu7w0jt8l8jt.pdf', '2022-06-15 16:41:05', '2022-06-15 16:41:05');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (134, 20, '宁波姚江8号地块（滨河沁月）', '20211213滨河沁月项目关于垃圾分类点位的专题报告.pptx', 48, '吴利军', 0, 3, 0, '2022-06-15 16:45:16', 'wdk/report/ckqkqyo4xuh1lj0e1395wg.pptx', 'wdk/report/ckqkqyo4xuh1lj0e1395wg.pdf', '2022-06-15 16:45:16', '2022-06-15 16:45:16');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (135, 20, '宁波姚江8号地块（滨河沁月）', '20211223滨河沁月项目智能化初步设计评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 13:58:55', 'wdk/report/ckuqce6cm972890l9lyabg.doc', 'wdk/report/ckuqce6cm972890l9lyabg.pdf', '2022-06-20 13:58:55', '2022-06-20 13:58:55');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (136, 20, '宁波姚江8号地块（滨河沁月）', '20211227关于滨河沁月项目物业用房功能定位的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 13:59:38', 'wdk/report/ckuqcws2sdn9liigb6mzs8.pptx', 'wdk/report/ckuqcws2sdn9liigb6mzs8.pdf', '2022-06-20 13:59:38', '2022-06-20 13:59:38');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (137, 20, '宁波姚江8号地块（滨河沁月）', '20220228关于滨河沁月项目公区照明控制方式的意见与建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 14:00:17', 'wdk/report/ckuqdf3l4fwenke4vzbhhd.pptx', 'wdk/report/ckuqdf3l4fwenke4vzbhhd.pdf', '2022-06-20 14:00:17', '2022-06-20 14:00:17');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (138, 20, '宁波姚江8号地块（滨河沁月）', '20220509宁波滨河沁月关于交通流线管控的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 14:01:05', 'wdk/report/ckuqe2s1o304bhpnbtuodv.pdf', 'wdk/report/ckuqe2s1o304bhpnbtuodv.pdf', '2022-06-20 14:01:05', '2022-06-20 14:01:05');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (139, 21, '宁波镇海骆驼项目（新桂沁澜）', '20210701宁波绿城骆驼项目景观概念文本会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:21:38', 'wdk/report/ckus3qrddn9nfdqo14bbeu.pdf', 'wdk/report/ckus3qrddn9nfdqo14bbeu.pdf', '2022-06-20 15:21:38', '2022-06-20 15:21:38');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (140, 21, '宁波镇海骆驼项目（新桂沁澜）', '20210810宁波绿城骆驼项目景观方案评审会会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:21:49', 'wdk/report/ckus3vvmogwcgpbpo17psj.pdf', 'wdk/report/ckus3vvmogwcgpbpo17psj.pdf', '2022-06-20 15:21:49', '2022-06-20 15:21:49');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (141, 21, '宁波镇海骆驼项目（新桂沁澜）', '20210827镇海新桂沁澜（骆驼）项目（泳池取消后）景观方案评审会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:22:04', 'wdk/report/ckus42ygyo4ixzzovtyw7x.pdf', 'wdk/report/ckus42ygyo4ixzzovtyw7x.pdf', '2022-06-20 15:22:04', '2022-06-20 15:22:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (142, 21, '宁波镇海骆驼项目（新桂沁澜）', '20211102关于宁波镇海骆驼项目（新桂沁澜）弱电施工图设计的意见和建议.docx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:22:19', 'wdk/report/ckus48yezcp2l1wtyjcfv1.docx', 'wdk/report/ckus48yezcp2l1wtyjcfv1.pdf', '2022-06-20 15:22:19', '2022-06-20 15:22:19');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (143, 21, '宁波镇海骆驼项目（新桂沁澜）', '20211123宁波骆驼项目差异户型、合同附图等评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:23:38', 'wdk/report/ckus59c83yaefmlgeyzvxb.doc', 'wdk/report/ckus59c83yaefmlgeyzvxb.pdf', '2022-06-20 15:23:38', '2022-06-20 15:23:38');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (144, 21, '宁波镇海骆驼项目（新桂沁澜）', '20220301新桂沁澜项目关于车行道闸设置位置的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:24:18', 'wdk/report/ckus5qrcm1w0d7iaxeabbn.pptx', 'wdk/report/ckus5qrcm1w0d7iaxeabbn.pdf', '2022-06-20 15:24:18', '2022-06-20 15:24:18');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (145, 21, '宁波镇海骆驼项目（新桂沁澜）', '20220310镇海新桂沁澜项目物业交通管理方案.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:24:56', 'wdk/report/ckus685sjvi1nef6qam2p6.pptx', 'wdk/report/ckus685sjvi1nef6qam2p6.pdf', '2022-06-20 15:24:56', '2022-06-20 15:24:56');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (146, 21, '宁波镇海骆驼项目（新桂沁澜）', '20220415宁波新桂沁澜关于公区精装修方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:25:38', 'wdk/report/ckus6oo5j6nkzerad5vqzy.pptx', 'wdk/report/ckus6oo5j6nkzerad5vqzy.pdf', '2022-06-20 15:25:38', '2022-06-20 15:25:38');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (147, 22, '宁波蛟川项目（春语文澜）', '20210701宁波绿城镇海蛟川项目景观概念方案文本会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:27:59', 'wdk/report/ckus8lww8mn3rdnn3x5nri.pdf', 'wdk/report/ckus8lww8mn3rdnn3x5nri.pdf', '2022-06-20 15:27:59', '2022-06-20 15:27:59');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (148, 22, '宁波蛟川项目（春语文澜）', '20210810宁波绿城镇海蛟川项目景观方案评审会会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:28:09', 'wdk/report/ckus8qgagyqvhb26f34irv.pdf', 'wdk/report/ckus8qgagyqvhb26f34irv.pdf', '2022-06-20 15:28:09', '2022-06-20 15:28:09');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (149, 22, '宁波蛟川项目（春语文澜）', '20210926宁波镇海蛟川项目智能化设计的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:28:21', 'wdk/report/ckus8vvgxn4czadfgotppe.pdf', 'wdk/report/ckus8vvgxn4czadfgotppe.pdf', '2022-06-20 15:28:21', '2022-06-20 15:28:21');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (150, 22, '宁波蛟川项目（春语文澜）', '20211014关于春语文澜项目智慧园区硬件设备配置标准要求的建议报告 (1).doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:28:39', 'wdk/report/ckus93okumyg9fbwfmlf7x.doc', 'wdk/report/ckus93okumyg9fbwfmlf7x.pdf', '2022-06-20 15:28:39', '2022-06-20 15:28:39');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (151, 22, '宁波蛟川项目（春语文澜）', '20211023关于春语文澜项目电梯功能定位的意见和建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:28:54', 'wdk/report/ckus9aeivfccxcxo6d9umi.doc', 'wdk/report/ckus9aeivfccxcxo6d9umi.pdf', '2022-06-20 15:28:54', '2022-06-20 15:28:54');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (152, 22, '宁波蛟川项目（春语文澜）', '20211122宁波蛟川项目智能化评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:29:26', 'wdk/report/ckus9p09gd1u4q9nm248vz.doc', 'wdk/report/ckus9p09gd1u4q9nm248vz.pdf', '2022-06-20 15:29:26', '2022-06-20 15:29:26');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (153, 22, '宁波蛟川项目（春语文澜）', '20211126镇海春语文澜（蛟川）景观扩初评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:29:39', 'wdk/report/ckus9v1xiyljnbqxllfetp.doc', 'wdk/report/ckus9v1xiyljnbqxllfetp.pdf', '2022-06-20 15:29:39', '2022-06-20 15:29:39');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (154, 22, '宁波蛟川项目（春语文澜）', '20211213镇海春语文澜建筑施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:30:04', 'wdk/report/ckusa5pqb6bccuyrbrwjav.pptx', 'wdk/report/ckusa5pqb6bccuyrbrwjav.pdf', '2022-06-20 15:30:04', '2022-06-20 15:30:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (155, 22, '宁波蛟川项目（春语文澜）', '20220108关于奉化锦上月明综合管网施工现场常见问题及注意事项的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:30:53', 'wdk/report/ckusasrjp4o1atebgtwxls.doc', 'wdk/report/ckusasrjp4o1atebgtwxls.pdf', '2022-06-20 15:30:53', '2022-06-20 15:30:53');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (156, 23, '宁波奉化江口三号地块(锦上月鸣）', '20210706宁波奉化江口三号地块景观概念方案评审会会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:37:30', 'wdk/report/ckusfwb686nv7mqirzyanc.pdf', 'wdk/report/ckusfwb686nv7mqirzyanc.pdf', '2022-06-20 15:37:30', '2022-06-20 15:37:30');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (157, 23, '宁波奉化江口三号地块(锦上月鸣）', '20210803宁波奉化锦上月明关于水电能耗计量设置的意见.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:37:51', 'wdk/report/ckusg52j6bh2fnnvucmcba.doc', 'wdk/report/ckusg52j6bh2fnnvucmcba.pdf', '2022-06-20 15:37:51', '2022-06-20 15:37:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (158, 23, '宁波奉化江口三号地块(锦上月鸣）', '20210813宁波奉化江口三号地块景观概念方案评审会会议纪要.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 15:38:03', 'wdk/report/ckusgblxcn84cfnwhssyvv.pdf', 'wdk/report/ckusgblxcn84cfnwhssyvv.pdf', '2022-06-20 15:38:03', '2022-06-20 15:38:03');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (159, 23, '宁波奉化江口三号地块(锦上月鸣）', '20210911关于奉化锦上月明综合管网施工现场常见问题及注意事项的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:38:50', 'wdk/report/ckusgw39hdplxtam3kkkw3.doc', 'wdk/report/ckusgw39hdplxtam3kkkw3.pdf', '2022-06-20 15:38:50', '2022-06-20 15:38:50');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (160, 23, '宁波奉化江口三号地块(锦上月鸣）', '20211014关于奉化锦上月明项目智慧园区硬件设备配置标准要求的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:39:11', 'wdk/report/ckush5meellhhitcmsey9r.doc', 'wdk/report/ckush5meellhhitcmsey9r.pdf', '2022-06-20 15:39:11', '2022-06-20 15:39:11');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (161, 23, '宁波奉化江口三号地块(锦上月鸣）', '20211220奉化锦上月明项目关于门窗防渗漏专项的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:39:49', 'wdk/report/ckushm5md8tbph4disnwln.pptx', 'wdk/report/ckushm5md8tbph4disnwln.pdf', '2022-06-20 15:39:49', '2022-06-20 15:39:49');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (162, 23, '宁波奉化江口三号地块(锦上月鸣）', '20220119关于奉化锦上月明项目电梯功能定位的意见和建议.doc', 48, '吴利军', 0, 3, 0, '2022-06-20 15:40:02', 'wdk/report/ckushtbu5b6zq4zozlkrxn.doc', 'wdk/report/ckushtbu5b6zq4zozlkrxn.pdf', '2022-06-20 15:40:02', '2022-06-20 15:40:02');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (163, 23, '宁波奉化江口三号地块(锦上月鸣）', '20220423关于奉化锦上月明项目智能化方案和图纸的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:40:44', 'wdk/report/ckusi8r9gz0e2v1owb0krg.pptx', 'wdk/report/ckusi8r9gz0e2v1owb0krg.pdf', '2022-06-20 15:40:44', '2022-06-20 15:40:44');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (164, 23, '宁波奉化江口三号地块(锦上月鸣）', '20220425宁波奉化锦上月明项目物业交通管理方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 15:41:30', 'wdk/report/ckusiuiy80q1gdxajpdbhe.pptx', 'wdk/report/ckusiuiy80q1gdxajpdbhe.pdf', '2022-06-20 15:41:30', '2022-06-20 15:41:30');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (165, 26, '宁波文创港', '20220123关于宁波环球文创港项目施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 16:36:51', 'wdk/report/ckutp7fstemxcjlzwo56lp.pptx', 'wdk/report/ckutp7fstemxcjlzwo56lp.pdf', '2022-06-20 16:36:51', '2022-06-20 16:36:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (166, 26, '宁波文创港', '20210823关于环球集团文创港项目总平图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 17:20:34', 'wdk/report/ckuumpwxhcftgdp1cicwea.pptx', 'wdk/report/ckuumpwxhcftgdp1cicwea.pdf', '2022-06-20 17:20:34', '2022-06-20 17:20:34');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (167, 26, '宁波文创港', '20220123关于宁波环球文创港项目施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-20 17:23:34', 'wdk/report/ckuuoytjkyoevroibl0n3l.pptx', 'wdk/report/ckuuoytjkyoevroibl0n3l.pdf', '2022-06-20 17:23:34', '2022-06-20 17:23:34');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (168, 26, '宁波文创港', '20220318关于宁波文创港项目办公电梯系统分析方案的意见和建议.pdf', 48, '吴利军', 0, 3, 0, '2022-06-20 17:23:57', 'wdk/report/ckuupe9h56xdf5p0tn72nh.pdf', 'wdk/report/ckuupe9h56xdf5p0tn72nh.pdf', '2022-06-20 17:23:57', '2022-06-20 17:23:57');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (169, 26, '宁波文创港', '20220402关于文创港项目消防专业需甲方确认事项.docx', 48, '吴利军', 0, 3, 0, '2022-06-20 17:25:45', 'wdk/report/ckuuqrbkopmivia9ev2de0.docx', 'wdk/report/ckuuqrbkopmivia9ev2de0.pdf', '2022-06-20 17:25:45', '2022-06-20 17:25:45');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (170, 26, '宁波文创港', '20220406关于文创港项目消防专业需甲方确认事项--甲方回复版.docx', 48, '吴利军', 0, 3, 0, '2022-06-20 17:26:21', 'wdk/report/ckuur7mescfzb3sdrro4dc.docx', 'wdk/report/ckuur7mescfzb3sdrro4dc.pdf', '2022-06-20 17:26:21', '2022-06-20 17:26:21');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (171, 27, '奉化凤悦印湖', '20210726关于奉化68号地块总平图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-27 13:41:37', 'wdk/report/cl0ocwlvjch8f1svznpohf.pptx', 'wdk/report/cl0ocwlvjch8f1svznpohf.pdf', '2022-06-27 13:41:37', '2022-06-27 13:41:37');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (172, 27, '奉化凤悦印湖', '20210831关于奉化68号地块出入口管控的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-27 13:41:51', 'wdk/report/cl0od4nzw43qofh05wp1di.pptx', 'wdk/report/cl0od4nzw43qofh05wp1di.pdf', '2022-06-27 13:41:51', '2022-06-27 13:41:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (173, 27, '奉化凤悦印湖', '20211117关于宁波凤悦印湖项目不利因素梳理的意见和建议.docx', 48, '吴利军', 0, 3, 0, '2022-06-27 13:42:25', 'wdk/report/cl0odkewzlvtjef4otdp42.docx', 'wdk/report/cl0odkewzlvtjef4otdp42.pdf', '2022-06-27 13:42:25', '2022-06-27 13:42:25');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (174, 27, '奉化凤悦印湖', '20220414凤悦印湖项目说辞（物业咨询意见）.pdf', 48, '吴利军', 0, 3, 0, '2022-06-27 13:42:51', 'wdk/report/cl0odxfvyygcut9zrjpedo.pdf', 'wdk/report/cl0odxfvyygcut9zrjpedo.pdf', '2022-06-27 13:42:51', '2022-06-27 13:42:51');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (175, 27, '奉化凤悦印湖', '20220414宁波凤悦印湖项目景观扩初图评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-27 13:43:07', 'wdk/report/cl0oe3lkz97n38crvsmbbu.doc', 'wdk/report/cl0oe3lkz97n38crvsmbbu.pdf', '2022-06-27 13:43:07', '2022-06-27 13:43:07');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (176, 27, '奉化凤悦印湖', '20220426宁波凤悦印湖项目智能化施工图的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-27 13:43:49', 'wdk/report/cl0oej18jgdfzk391zxcgf.pptx', 'wdk/report/cl0oej18jgdfzk391zxcgf.pdf', '2022-06-27 13:43:49', '2022-06-27 13:43:49');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (177, 27, '奉化凤悦印湖', '20220510宁波凤悦印湖项目智能化评审会会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-27 13:44:04', 'wdk/report/cl0oeu2j4tqj8xtfpwsqih.doc', 'wdk/report/cl0oeu2j4tqj8xtfpwsqih.pdf', '2022-06-27 13:44:04', '2022-06-27 13:44:04');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (178, 28, '嵊州·雲庐', '20211013关于中式项目交付缺陷分享.pptx', 48, '吴利军', 0, 3, 0, '2022-06-27 13:54:18', 'wdk/report/cl0omh80d715boqgdlxija.pptx', 'wdk/report/cl0omh80d715boqgdlxija.pdf', '2022-06-27 13:54:18', '2022-06-27 13:54:18');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (179, 28, '嵊州·雲庐', '20211030嵊州云庐项目关于生活馆智能化设计方案的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-27 13:55:27', 'wdk/report/cl0onj8susygzvsuttmhii.pptx', 'wdk/report/cl0onj8susygzvsuttmhii.pdf', '2022-06-27 13:55:27', '2022-06-27 13:55:27');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (180, 28, '嵊州·雲庐', '20211114关于嵊州云庐项目智慧园区硬件设备配置标准要求的建议报告.doc', 48, '吴利军', 0, 3, 0, '2022-06-27 13:55:45', 'wdk/report/cl0onsa2cnb3sbiatahday.doc', 'wdk/report/cl0onsa2cnb3sbiatahday.pdf', '2022-06-27 13:55:45', '2022-06-27 13:55:45');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (181, 28, '嵊州·雲庐', '20211221嵊州檀山府项目示范区及大区景观方案评审会议纪要.doc', 48, '吴利军', 0, 3, 0, '2022-06-27 13:55:59', 'wdk/report/cl0onypvk4atb79ukbb1mb.doc', 'wdk/report/cl0onypvk4atb79ukbb1mb.pdf', '2022-06-27 13:55:59', '2022-06-27 13:55:59');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (182, 28, '嵊州·雲庐', '20220120嵊州云庐项目关于门窗防渗漏专项的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-27 13:58:00', 'wdk/report/cl0oph41s2p3z7hosiobjn.pptx', 'wdk/report/cl0oph41s2p3z7hosiobjn.pdf', '2022-06-27 13:58:00', '2022-06-27 13:58:00');
INSERT INTO `wdk_report` (`id`, `project_id`, `project_name`, `name`, `create_by`, `create_name`, `rescind`, `audit_status`, `excellence`, `audit_time`, `origin_url`, `pdf_url`, `created_at`, `updated_at`) VALUES (183, 28, '嵊州·雲庐', '20220414嵊州云庐项目关于生活美学馆现场巡视的意见和建议.pptx', 48, '吴利军', 0, 3, 0, '2022-06-27 14:00:00', 'wdk/report/cl0oqum6ukb9wbzg4vjahl.pptx', 'wdk/report/cl0oqum6ukb9wbzg4vjahl.pdf', '2022-06-27 14:00:00', '2022-06-27 14:00:00');
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
  `preaudit_status` tinyint unsigned NOT NULL COMMENT '前置审核是否已通过 0:否 1:是',
  `status` tinyint unsigned NOT NULL COMMENT '审核状态 0:未通过 1:审核中 2:已通过',
  `excellence` tinyint unsigned NOT NULL COMMENT '是否被推荐为优秀报告 0:未被推荐为优秀报告 1:已被推荐为优秀报告',
  `audit_time` datetime DEFAULT NULL COMMENT '审核时间',
  `audit_opinion` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '审核意见',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
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
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (1, 7, '建筑类报告', '傅康振');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (1, 18, '建筑类报告', '何一江');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (1, 38, '建筑类报告', '孟龙飞');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (1, 40, '建筑类报告', '孙雷');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (1, 53, '建筑类报告', '陈立胜');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (2, 5, '景观类报告', '陈锋');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (2, 23, '景观类报告', '童柯');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (2, 25, '景观类报告', '周振华');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (3, 6, '精装类报告', '薛擎');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (3, 24, '精装类报告', '许虎');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (3, 31, '精装类报告', '赵强');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (3, 54, '精装类报告', '俞凯');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (4, 4, '机电类报告', '周荣辉');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (4, 36, '机电类报告', '袁恒业');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (4, 47, '机电类报告', '岳洪亮');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (4, 49, '机电类报告', '陈欣');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (5, 16, '物业管理类报告', '张磊');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (5, 17, '物业管理类报告', '张洋');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (5, 21, '物业管理类报告', '章晓波');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (5, 22, '物业管理类报告', '余晓伟');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (5, 41, '物业管理类报告', '黄裕凡');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (6, 5, '现场巡视类报告', '陈锋');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (6, 16, '现场巡视类报告', '张磊');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (6, 24, '现场巡视类报告', '许虎');
INSERT INTO `wdk_report_audit_cfg` (`id`, `audit_uid`, `type_name`, `audit_name`) VALUES (6, 38, '现场巡视类报告', '孟龙飞');
COMMIT;

-- ----------------------------
-- Table structure for wdk_report_audit_type
-- ----------------------------
DROP TABLE IF EXISTS `wdk_report_audit_type`;
CREATE TABLE `wdk_report_audit_type` (
  `id` bigint unsigned NOT NULL COMMENT '报告ID',
  `audit_uid` bigint unsigned NOT NULL COMMENT '审核员用户ID',
  `auditor_type` tinyint unsigned NOT NULL COMMENT '审核员类型 0:负责人 1:专家组',
  `type_id` bigint unsigned NOT NULL COMMENT '报告类型ID',
  `type_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报告类型名称',
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
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '报告类型ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '报告类型名称',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_report_cfg
-- ----------------------------
BEGIN;
INSERT INTO `wdk_report_cfg` (`id`, `name`, `created_at`, `updated_at`) VALUES (1, '建筑类报告', '2022-06-06 15:22:22', '2022-06-06 15:22:22');
INSERT INTO `wdk_report_cfg` (`id`, `name`, `created_at`, `updated_at`) VALUES (2, '景观类报告', '2022-06-06 15:23:00', '2022-06-06 15:23:00');
INSERT INTO `wdk_report_cfg` (`id`, `name`, `created_at`, `updated_at`) VALUES (3, '精装类报告', '2022-06-06 15:23:41', '2022-06-06 15:23:41');
INSERT INTO `wdk_report_cfg` (`id`, `name`, `created_at`, `updated_at`) VALUES (4, '机电类报告', '2022-06-06 15:24:24', '2022-06-06 15:24:24');
INSERT INTO `wdk_report_cfg` (`id`, `name`, `created_at`, `updated_at`) VALUES (5, '物业管理类报告', '2022-06-06 15:25:23', '2022-06-06 15:25:23');
INSERT INTO `wdk_report_cfg` (`id`, `name`, `created_at`, `updated_at`) VALUES (6, '现场巡视类报告', '2022-06-06 15:26:09', '2022-06-06 15:26:09');
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
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (1, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (2, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (3, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (4, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (5, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (6, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (7, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (8, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (9, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (10, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (11, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (12, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (13, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (14, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (15, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (16, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (17, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (18, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (19, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (20, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (21, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (22, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (23, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (24, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (25, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (26, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (27, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (28, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (29, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (30, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (31, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (32, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (33, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (34, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (35, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (36, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (37, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (38, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (39, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (40, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (41, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (42, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (43, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (44, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (45, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (46, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (47, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (48, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (49, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (50, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (51, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (52, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (53, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (54, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (55, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (56, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (57, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (58, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (59, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (60, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (61, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (62, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (63, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (64, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (65, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (66, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (67, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (68, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (69, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (70, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (71, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (72, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (73, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (74, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (75, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (76, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (77, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (78, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (79, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (80, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (81, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (82, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (83, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (84, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (85, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (86, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (87, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (88, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (89, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (90, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (91, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (92, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (93, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (94, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (95, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (96, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (97, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (98, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (99, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (100, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (101, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (102, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (103, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (104, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (105, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (106, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (107, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (108, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (109, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (110, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (111, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (112, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (113, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (114, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (115, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (116, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (117, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (118, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (119, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (120, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (121, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (122, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (123, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (124, 6, '现场巡视类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (125, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (126, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (127, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (128, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (129, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (130, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (131, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (132, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (133, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (134, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (135, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (136, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (137, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (138, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (139, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (140, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (141, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (142, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (143, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (144, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (145, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (146, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (147, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (148, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (149, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (150, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (151, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (152, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (153, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (154, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (155, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (156, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (157, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (158, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (159, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (160, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (161, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (162, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (163, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (164, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (165, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (166, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (167, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (168, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (169, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (170, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (171, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (172, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (173, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (174, 5, '物业管理类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (175, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (176, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (177, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (178, 1, '建筑类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (179, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (180, 4, '机电类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (181, 2, '景观类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (182, 3, '精装类报告');
INSERT INTO `wdk_report_type` (`id`, `type_id`, `type_name`) VALUES (183, 6, '现场巡视类报告');
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
  `service_time` date NOT NULL COMMENT '服务时间',
  `xch_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '行程涵文件名',
  `xch_origin_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原始行程涵url',
  `xch_pdf_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'pdf行程涵url',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`,`project_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of wdk_service_record
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
