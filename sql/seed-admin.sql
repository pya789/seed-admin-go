/*
 Navicat MySQL Data Transfer

 Source Server         : seed-admin
 Source Server Type    : MySQL
 Source Server Version : 50739
 Source Host           : 119.8.52.66:3306
 Source Schema         : seed-admin

 Target Server Type    : MySQL
 Target Server Version : 50739
 File Encoding         : 65001

 Date: 12/08/2022 08:54:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dept`;
CREATE TABLE `admin_sys_dept`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `parent_id` int(11) NULL DEFAULT NULL COMMENT '父级ID',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_dept
-- ----------------------------
INSERT INTO `admin_sys_dept` VALUES (1, '心脏跳动科技', 0, 0, '2022-07-02 14:48:55', '2022-07-21 17:34:35');
INSERT INTO `admin_sys_dept` VALUES (2, '北京总公司', 1, 0, '2022-07-02 14:51:09', '2022-07-08 18:52:48');
INSERT INTO `admin_sys_dept` VALUES (5, '后勤部', 2, 0, '2022-07-02 14:52:48', '2022-07-08 19:31:37');
INSERT INTO `admin_sys_dept` VALUES (6, '西藏分公司', 1, 0, '2022-07-02 14:52:26', '2022-07-08 19:35:38');
INSERT INTO `admin_sys_dept` VALUES (21, '财务部', 6, 0, '2022-07-08 17:46:46', '2022-07-08 19:35:49');
INSERT INTO `admin_sys_dept` VALUES (23, '工程部', 6, 0, '2022-07-08 19:53:15', '2022-07-08 19:53:15');
INSERT INTO `admin_sys_dept` VALUES (25, '江浙沪分公司', 1, 0, '2022-07-10 10:50:02', '2022-07-10 10:50:02');
INSERT INTO `admin_sys_dept` VALUES (26, '开发部', 25, 0, '2022-07-10 10:50:08', '2022-07-10 10:50:08');
INSERT INTO `admin_sys_dept` VALUES (27, '财务部', 25, 0, '2022-07-10 10:50:18', '2022-07-10 10:50:18');
INSERT INTO `admin_sys_dept` VALUES (28, '工程部', 25, 0, '2022-07-10 10:50:23', '2022-07-10 10:50:23');
INSERT INTO `admin_sys_dept` VALUES (29, '东北分公司', 1, 0, '2022-07-10 10:50:32', '2022-07-10 10:50:32');
INSERT INTO `admin_sys_dept` VALUES (30, '销售部', 29, 0, '2022-07-10 10:50:39', '2022-07-10 10:50:39');
INSERT INTO `admin_sys_dept` VALUES (31, '后勤2部', 2, 0, '2022-07-19 18:30:04', '2022-07-22 23:04:24');

-- ----------------------------
-- Table structure for admin_sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dict_data`;
CREATE TABLE `admin_sys_dict_data`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(11) NOT NULL COMMENT '主键',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典标签',
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态0.正常 1.禁用',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_dict_data
-- ----------------------------
INSERT INTO `admin_sys_dict_data` VALUES (18, 7, '正常', '0', 0, '2022-07-16 15:49:08', '2022-07-16 19:09:14');
INSERT INTO `admin_sys_dict_data` VALUES (19, 7, '禁用', '1', 0, '2022-07-16 15:49:14', '2022-07-17 14:22:40');
INSERT INTO `admin_sys_dict_data` VALUES (20, 8, '目录', '0', 0, '2022-07-17 17:04:45', '2022-07-17 17:04:45');
INSERT INTO `admin_sys_dict_data` VALUES (21, 8, '菜单', '1', 0, '2022-07-17 17:04:51', '2022-07-17 17:04:51');
INSERT INTO `admin_sys_dict_data` VALUES (22, 8, '按钮', '2', 0, '2022-07-17 17:04:59', '2022-07-17 17:04:59');
INSERT INTO `admin_sys_dict_data` VALUES (23, 9, '否', '0', 0, '2022-07-17 17:17:26', '2022-07-17 17:17:26');
INSERT INTO `admin_sys_dict_data` VALUES (24, 9, '是', '1', 0, '2022-07-17 17:17:31', '2022-07-17 17:17:31');
INSERT INTO `admin_sys_dict_data` VALUES (26, 19, '显示', '0', 0, '2022-07-27 15:10:11', '2022-07-27 15:10:11');
INSERT INTO `admin_sys_dict_data` VALUES (27, 19, '隐藏', '1', 0, '2022-07-27 15:10:16', '2022-07-27 15:10:16');

-- ----------------------------
-- Table structure for admin_sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dict_type`;
CREATE TABLE `admin_sys_dict_type`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名称',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典类型',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态0.正常 1.禁用',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_dict_type
-- ----------------------------
INSERT INTO `admin_sys_dict_type` VALUES (7, '状态', 'sys_common_status', 0, '2022-07-16 15:48:41', '2022-07-16 19:14:34');
INSERT INTO `admin_sys_dict_type` VALUES (8, '菜单类型', 'sys_menu_type', 0, '2022-07-17 17:04:25', '2022-07-17 17:04:31');
INSERT INTO `admin_sys_dict_type` VALUES (9, '页面缓存', 'sys_page_keepAlive', 0, '2022-07-17 17:17:00', '2022-07-17 17:17:00');
INSERT INTO `admin_sys_dict_type` VALUES (19, '菜单显示', 'sys_menu_visible', 0, '2022-07-27 15:09:48', '2022-07-27 15:10:25');

-- ----------------------------
-- Table structure for admin_sys_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_log`;
CREATE TABLE `admin_sys_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL COMMENT '角色ID',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方式',
  `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '行为',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'IP',
  `status_code` int(11) NULL DEFAULT NULL COMMENT '响应状态',
  `params` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '参数',
  `results` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1587 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_log
-- ----------------------------

-- ----------------------------
-- Table structure for admin_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_menu`;
CREATE TABLE `admin_sys_menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `router_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由名称',
  `router_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由地址',
  `page_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '页面路径',
  `perms` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型0.目录 1.菜单 2.按钮',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `visible` tinyint(1) NOT NULL DEFAULT 0 COMMENT '隐藏0.显示 1.隐藏',
  `keep_alive` tinyint(1) NOT NULL DEFAULT 0 COMMENT '页面缓存0.否 1.是',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态0.正常 1.禁用',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_menu
-- ----------------------------
INSERT INTO `admin_sys_menu` VALUES (1, 0, '仪表盘', 'dashboard', '/dashboard/index', '/dashboard/index.vue', 'sys:dashboard', 1, 'dashboardOutlined', 0, 0, 0, 0, '2022-06-18 04:46:24', '2022-07-12 17:12:33');
INSERT INTO `admin_sys_menu` VALUES (2, 0, '系统管理', 'system', '/system', '', '', 0, 'settingOutlined', 1, 0, 0, 0, '2022-06-18 10:28:30', '2022-06-28 10:17:08');
INSERT INTO `admin_sys_menu` VALUES (3, 2, '用户管理', 'user', '/system/user', '/system/user/index.vue', 'sys:user:list', 1, 'userOutlined', 0, 0, 0, 0, '2022-06-18 10:34:04', '2022-06-29 12:36:52');
INSERT INTO `admin_sys_menu` VALUES (4, 2, '菜单管理', 'menu', '/system/menu', '/system/menu/index.vue', 'sys:menu:list', 1, 'menuOutlined', 2, 0, 0, 0, '2022-06-18 10:37:23', '2022-07-24 14:19:08');
INSERT INTO `admin_sys_menu` VALUES (5, 2, '角色管理', 'role', '/system/role', '/system/role/index.vue', 'sys:role:list', 1, 'robotOutlined', 1, 0, 0, 0, '2022-06-18 10:37:26', '2022-07-24 14:19:05');
INSERT INTO `admin_sys_menu` VALUES (6, 2, '操作日志', 'log', '/system/log', '/system/log/index.vue', 'sys:log:list', 1, 'minusSquareOutlined', 3, 0, 0, 0, '2022-06-18 10:38:48', '2022-07-29 08:13:25');
INSERT INTO `admin_sys_menu` VALUES (7, 2, '系统监控', 'server', '/system/server', '/system/server/index.vue', 'sys:server:info', 1, 'fundProjectionScreenOutlined', 4, 0, 0, 0, '2022-06-20 13:13:36', '2022-08-01 16:38:20');
INSERT INTO `admin_sys_menu` VALUES (8, 2, '字典管理', 'dict', '/system/dict', '/system/dict/index.vue', 'sys:dict:list', 1, 'bookOutlined', 5, 0, 0, 0, '2022-06-21 09:02:15', '2022-07-23 18:32:51');
INSERT INTO `admin_sys_menu` VALUES (46, 5, '新增', NULL, NULL, NULL, 'sys:role:add', 2, '', 0, 0, 0, 0, '2022-06-26 19:27:46', '2022-06-26 19:27:46');
INSERT INTO `admin_sys_menu` VALUES (50, 2, '字典数据', 'dictDetails', '/system/dict/details/:id', '/system/dict/details/index.vue', 'sys:dict:details:list', 1, '', 10, 1, 0, 0, '2022-07-11 17:20:33', '2022-07-22 22:12:06');
INSERT INTO `admin_sys_menu` VALUES (51, 8, '信息', '', '', '', 'sys:dict:info', 2, '', 0, 0, 0, 0, '2022-07-22 18:44:20', '2022-07-22 18:44:20');
INSERT INTO `admin_sys_menu` VALUES (52, 8, '新增', '', '', '', 'sys:dict:add', 2, '', 0, 0, 0, 0, '2022-07-22 22:04:53', '2022-07-22 22:04:53');
INSERT INTO `admin_sys_menu` VALUES (53, 8, '编辑', '', '', '', 'sys:dict:update', 2, '', 0, 0, 0, 0, '2022-07-22 22:05:12', '2022-07-22 22:06:24');
INSERT INTO `admin_sys_menu` VALUES (54, 8, '删除', '', '', '', 'sys:dict:del', 2, '', 0, 0, 0, 0, '2022-07-22 22:06:18', '2022-07-22 22:06:18');
INSERT INTO `admin_sys_menu` VALUES (55, 50, '信息', '', '', '', 'sys:dict:details:info', 2, '', 0, 0, 0, 0, '2022-07-22 22:13:23', '2022-07-22 22:13:23');
INSERT INTO `admin_sys_menu` VALUES (56, 50, '新增', '', '', '', 'sys:dict:details:add', 2, '', 0, 0, 0, 0, '2022-07-22 22:14:40', '2022-07-22 22:14:40');
INSERT INTO `admin_sys_menu` VALUES (57, 50, '编辑', '', '', '', 'sys:dict:details:update', 2, '', 0, 0, 0, 0, '2022-07-22 22:15:00', '2022-07-22 22:15:00');
INSERT INTO `admin_sys_menu` VALUES (58, 50, '删除', '', '', '', 'sys:dict:details:del', 2, '', 0, 0, 0, 0, '2022-07-22 22:15:12', '2022-07-22 22:15:12');
INSERT INTO `admin_sys_menu` VALUES (59, 3, '新增', '', '', '', 'sys:user:add', 2, '', 0, 0, 0, 0, '2022-07-22 22:29:04', '2022-07-22 22:29:04');
INSERT INTO `admin_sys_menu` VALUES (60, 3, '删除', '', '', '', 'sys:user:del', 2, '', 0, 0, 0, 0, '2022-07-22 22:29:15', '2022-07-22 22:29:15');
INSERT INTO `admin_sys_menu` VALUES (61, 3, '信息', '', '', '', 'sys:user:info', 2, '', 0, 0, 0, 0, '2022-07-22 22:29:27', '2022-07-22 22:29:27');
INSERT INTO `admin_sys_menu` VALUES (62, 3, '编辑', '', '', '', 'sys:user:update', 2, '', 0, 0, 0, 0, '2022-07-22 22:29:38', '2022-07-22 22:29:38');
INSERT INTO `admin_sys_menu` VALUES (63, 3, '移动用户部门', '', '', '', 'sys:user:move', 2, '', 0, 0, 0, 0, '2022-07-22 22:29:58', '2022-07-22 22:29:58');
INSERT INTO `admin_sys_menu` VALUES (64, 3, '更新用户角色', '', '', '', 'sys:user:updateUserRole', 2, '', 0, 0, 0, 0, '2022-07-22 22:30:45', '2022-07-22 22:30:45');
INSERT INTO `admin_sys_menu` VALUES (65, 5, '信息', '', '', '', 'sys:role:info', 2, '', 0, 0, 0, 0, '2022-07-22 22:35:42', '2022-07-22 22:35:42');
INSERT INTO `admin_sys_menu` VALUES (66, 5, '编辑', '', '', '', 'sys:role:update', 2, '', 0, 0, 0, 0, '2022-07-22 22:35:53', '2022-07-22 22:35:53');
INSERT INTO `admin_sys_menu` VALUES (67, 5, '删除', '', '', '', 'sys:role:del', 2, '', 0, 0, 0, 0, '2022-07-22 22:36:03', '2022-07-22 22:36:03');
INSERT INTO `admin_sys_menu` VALUES (68, 4, '信息', '', '', '', 'sys:menu:info', 2, '', 0, 0, 0, 0, '2022-07-22 22:42:14', '2022-07-22 22:42:14');
INSERT INTO `admin_sys_menu` VALUES (69, 4, '新增', '', '', '', 'sys:menu:add', 2, '', 0, 0, 0, 0, '2022-07-22 22:42:33', '2022-07-22 22:42:33');
INSERT INTO `admin_sys_menu` VALUES (70, 4, '编辑', '', '', '', 'sys:menu:update', 2, '', 0, 0, 0, 0, '2022-07-22 22:42:43', '2022-07-22 22:42:43');
INSERT INTO `admin_sys_menu` VALUES (71, 4, '删除', '', '', '', 'sys:menu:del', 2, '', 0, 0, 0, 0, '2022-07-22 22:42:55', '2022-07-22 22:42:55');
INSERT INTO `admin_sys_menu` VALUES (72, 3, '部门列表', '', '', '', 'sys:dept:list', 2, '', 0, 0, 0, 0, '2022-07-22 22:52:04', '2022-07-22 22:52:04');
INSERT INTO `admin_sys_menu` VALUES (73, 3, '部门信息', '', '', '', 'sys:dept:info', 2, '', 0, 0, 0, 0, '2022-07-22 22:53:14', '2022-07-22 22:53:14');
INSERT INTO `admin_sys_menu` VALUES (74, 3, '部门新增', '', '', '', 'sys:dept:add', 2, '', 0, 0, 0, 0, '2022-07-22 22:53:31', '2022-07-22 22:53:31');
INSERT INTO `admin_sys_menu` VALUES (75, 3, '部门编辑', '', '', '', 'sys:dept:update', 2, '', 0, 0, 0, 0, '2022-07-22 22:53:44', '2022-07-22 22:53:44');
INSERT INTO `admin_sys_menu` VALUES (76, 3, '部门删除', '', '', '', 'sys:dept:del', 2, '', 0, 0, 0, 0, '2022-07-22 22:53:56', '2022-07-22 22:53:56');
INSERT INTO `admin_sys_menu` VALUES (79, 6, '删除', '', '', '', 'sys:log:del', 2, '', 0, 0, 0, 0, '2022-07-29 11:59:19', '2022-07-29 11:59:19');

-- ----------------------------
-- Table structure for admin_sys_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role`;
CREATE TABLE `admin_sys_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色标签',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `relevance` tinyint(1) NOT NULL DEFAULT 1 COMMENT '上下级数据权限是否关联0.是 1.否',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_role
-- ----------------------------
INSERT INTO `admin_sys_role` VALUES (1, '超级管理员', 'admin', '超级管理员', 0, '2022-06-18 10:49:48', '2022-08-09 18:31:37');
INSERT INTO `admin_sys_role` VALUES (35, '普通角色', 'test', '我只是一个普通的角色', 0, '2022-08-09 18:04:04', '2022-08-09 18:04:04');

-- ----------------------------
-- Table structure for admin_sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role_dept`;
CREATE TABLE `admin_sys_role_dept`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `dept_id` int(11) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 119 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_role_dept
-- ----------------------------
INSERT INTO `admin_sys_role_dept` VALUES (92, 35, 1);
INSERT INTO `admin_sys_role_dept` VALUES (93, 35, 2);
INSERT INTO `admin_sys_role_dept` VALUES (94, 35, 5);
INSERT INTO `admin_sys_role_dept` VALUES (95, 35, 31);
INSERT INTO `admin_sys_role_dept` VALUES (96, 35, 6);
INSERT INTO `admin_sys_role_dept` VALUES (97, 35, 21);
INSERT INTO `admin_sys_role_dept` VALUES (98, 35, 23);
INSERT INTO `admin_sys_role_dept` VALUES (99, 35, 25);
INSERT INTO `admin_sys_role_dept` VALUES (100, 35, 26);
INSERT INTO `admin_sys_role_dept` VALUES (101, 35, 27);
INSERT INTO `admin_sys_role_dept` VALUES (102, 35, 28);
INSERT INTO `admin_sys_role_dept` VALUES (103, 35, 29);
INSERT INTO `admin_sys_role_dept` VALUES (104, 35, 30);
INSERT INTO `admin_sys_role_dept` VALUES (105, 1, 24);
INSERT INTO `admin_sys_role_dept` VALUES (106, 1, 5);
INSERT INTO `admin_sys_role_dept` VALUES (107, 1, 31);
INSERT INTO `admin_sys_role_dept` VALUES (108, 1, 21);
INSERT INTO `admin_sys_role_dept` VALUES (109, 1, 23);
INSERT INTO `admin_sys_role_dept` VALUES (110, 1, 26);
INSERT INTO `admin_sys_role_dept` VALUES (111, 1, 27);
INSERT INTO `admin_sys_role_dept` VALUES (112, 1, 28);
INSERT INTO `admin_sys_role_dept` VALUES (113, 1, 30);
INSERT INTO `admin_sys_role_dept` VALUES (114, 1, 1);
INSERT INTO `admin_sys_role_dept` VALUES (115, 1, 2);
INSERT INTO `admin_sys_role_dept` VALUES (116, 1, 6);
INSERT INTO `admin_sys_role_dept` VALUES (117, 1, 25);
INSERT INTO `admin_sys_role_dept` VALUES (118, 1, 29);

-- ----------------------------
-- Table structure for admin_sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role_menu`;
CREATE TABLE `admin_sys_role_menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3198 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_role_menu
-- ----------------------------
INSERT INTO `admin_sys_role_menu` VALUES (3152, 35, 2);
INSERT INTO `admin_sys_role_menu` VALUES (3153, 35, 7);
INSERT INTO `admin_sys_role_menu` VALUES (3154, 35, 1);
INSERT INTO `admin_sys_role_menu` VALUES (3155, 35, 2);
INSERT INTO `admin_sys_role_menu` VALUES (3156, 35, 7);
INSERT INTO `admin_sys_role_menu` VALUES (3157, 35, 1);
INSERT INTO `admin_sys_role_menu` VALUES (3158, 35, 2);
INSERT INTO `admin_sys_role_menu` VALUES (3159, 35, 7);
INSERT INTO `admin_sys_role_menu` VALUES (3160, 35, 1);
INSERT INTO `admin_sys_role_menu` VALUES (3161, 1, 1);
INSERT INTO `admin_sys_role_menu` VALUES (3162, 1, 59);
INSERT INTO `admin_sys_role_menu` VALUES (3163, 1, 60);
INSERT INTO `admin_sys_role_menu` VALUES (3164, 1, 61);
INSERT INTO `admin_sys_role_menu` VALUES (3165, 1, 62);
INSERT INTO `admin_sys_role_menu` VALUES (3166, 1, 63);
INSERT INTO `admin_sys_role_menu` VALUES (3167, 1, 64);
INSERT INTO `admin_sys_role_menu` VALUES (3168, 1, 72);
INSERT INTO `admin_sys_role_menu` VALUES (3169, 1, 73);
INSERT INTO `admin_sys_role_menu` VALUES (3170, 1, 74);
INSERT INTO `admin_sys_role_menu` VALUES (3171, 1, 75);
INSERT INTO `admin_sys_role_menu` VALUES (3172, 1, 76);
INSERT INTO `admin_sys_role_menu` VALUES (3173, 1, 46);
INSERT INTO `admin_sys_role_menu` VALUES (3174, 1, 65);
INSERT INTO `admin_sys_role_menu` VALUES (3175, 1, 66);
INSERT INTO `admin_sys_role_menu` VALUES (3176, 1, 67);
INSERT INTO `admin_sys_role_menu` VALUES (3177, 1, 68);
INSERT INTO `admin_sys_role_menu` VALUES (3178, 1, 69);
INSERT INTO `admin_sys_role_menu` VALUES (3179, 1, 70);
INSERT INTO `admin_sys_role_menu` VALUES (3180, 1, 71);
INSERT INTO `admin_sys_role_menu` VALUES (3181, 1, 79);
INSERT INTO `admin_sys_role_menu` VALUES (3182, 1, 7);
INSERT INTO `admin_sys_role_menu` VALUES (3183, 1, 51);
INSERT INTO `admin_sys_role_menu` VALUES (3184, 1, 52);
INSERT INTO `admin_sys_role_menu` VALUES (3185, 1, 53);
INSERT INTO `admin_sys_role_menu` VALUES (3186, 1, 54);
INSERT INTO `admin_sys_role_menu` VALUES (3187, 1, 55);
INSERT INTO `admin_sys_role_menu` VALUES (3188, 1, 56);
INSERT INTO `admin_sys_role_menu` VALUES (3189, 1, 57);
INSERT INTO `admin_sys_role_menu` VALUES (3190, 1, 58);
INSERT INTO `admin_sys_role_menu` VALUES (3191, 1, 2);
INSERT INTO `admin_sys_role_menu` VALUES (3192, 1, 3);
INSERT INTO `admin_sys_role_menu` VALUES (3193, 1, 5);
INSERT INTO `admin_sys_role_menu` VALUES (3194, 1, 4);
INSERT INTO `admin_sys_role_menu` VALUES (3195, 1, 6);
INSERT INTO `admin_sys_role_menu` VALUES (3196, 1, 8);
INSERT INTO `admin_sys_role_menu` VALUES (3197, 1, 50);

-- ----------------------------
-- Table structure for admin_sys_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_user`;
CREATE TABLE `admin_sys_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dept_id` int(11) NULL DEFAULT NULL COMMENT '部门ID',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '未命名' COMMENT '用户名称',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态0.正常 1.禁用',
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '软删除0.正常 1.删除',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_user
-- ----------------------------
INSERT INTO `admin_sys_user` VALUES (1, 1, 'admin', '73d77a08df7a0026706cd4be8317e3cd', '超级管理员', '13938749240', '120547546@qq.com', '', 0, 0, '2022-06-18 10:04:52', '2022-08-09 18:29:21');
INSERT INTO `admin_sys_user` VALUES (2, 2, 'test', '73D77A08DF7A0026706CD4BE8317E3CD', '普通用户', '13503635596', '', '', 0, 0, '2022-06-18 10:40:02', '2022-07-17 14:39:13');

-- ----------------------------
-- Table structure for admin_sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_user_role`;
CREATE TABLE `admin_sys_user_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 747 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_sys_user_role
-- ----------------------------
INSERT INTO `admin_sys_user_role` VALUES (743, 1, 1);
INSERT INTO `admin_sys_user_role` VALUES (746, 2, 35);

-- ----------------------------
-- Table structure for admin_uploads
-- ----------------------------
DROP TABLE IF EXISTS `admin_uploads`;
CREATE TABLE `admin_uploads`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `type` int(11) NULL DEFAULT NULL COMMENT '分类',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 93 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_uploads
-- ----------------------------
INSERT INTO `admin_uploads` VALUES (91, 'logo.png', 'https://admin.seed-app.com/wwwroot/uploads/2022-08-09/db41cb1e-47d0-481c-b928-17165ef9ea6d.png', 2, '2022-08-09 18:29:00', '2022-08-09 18:29:00');
INSERT INTO `admin_uploads` VALUES (92, '微信截图_20220702182931.png', 'https://admin.seed-app.com/wwwroot/uploads/2022-08-09/5e963439-eb00-4250-8251-a029fb6eb2ea.png', 2, '2022-08-09 18:29:21', '2022-08-09 18:29:21');

-- ----------------------------
-- Table structure for admin_uploads_type
-- ----------------------------
DROP TABLE IF EXISTS `admin_uploads_type`;
CREATE TABLE `admin_uploads_type`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类名称',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分类标识',
  `created_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_uploads_type
-- ----------------------------
INSERT INTO `admin_uploads_type` VALUES (1, 'admin端通用上传', 'admin', '2022-07-04 15:17:13', '2022-07-04 15:17:16');
INSERT INTO `admin_uploads_type` VALUES (2, 'admin头像上传', 'admin_avatar', '2022-08-03 17:09:46', '2022-08-03 17:09:48');

SET FOREIGN_KEY_CHECKS = 1;
