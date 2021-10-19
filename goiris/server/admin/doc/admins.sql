/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50714
Source Host           : localhost:3306
Source Database       : admins
Source Character Set  : utf8mb4
Source Collaction     : utf8mb4_general_ci

Target Server Type    : MYSQL
Target Server Version : 50714
File Encoding         : 65001
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `gender` tinyint(5) DEFAULT NULL COMMENT '性别',
  `enable` tinyint(1) DEFAULT '1' COMMENT '1：启用用户 0：禁用用户',
  `name` varchar(255) NOT NULL DEFAULT '',
  `age` int(5) DEFAULT '0',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `userface` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `label` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=147 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES ('1', 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '1', '超级管理员', '0', '13333333333', '', '', null, null, null);
INSERT INTO `admin_user` VALUES ('2', 'adminmaster', 'x04jpoIrc8/mvNRqAG59Wg==', '1', '1', '董事长', '0', '13888888888', '', '', null, null, null);

-- ----------------------------
-- Table structure for admin_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_rule`;
CREATE TABLE `admin_rule` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(100) DEFAULT '',
  `v0` varchar(100) DEFAULT '',
  `v1` varchar(100) DEFAULT '',
  `v2` varchar(100) DEFAULT '',
  `v3` varchar(100) DEFAULT '',
  `v4` varchar(100) DEFAULT '',
  `v5` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `IDX_admin_rule_p_type` (`p_type`) USING BTREE,
  KEY `IDX_admin_rule_v2` (`v2`) USING BTREE,
  KEY `IDX_admin_rule_v3` (`v3`) USING BTREE,
  KEY `IDX_admin_rule_v4` (`v4`) USING BTREE,
  KEY `IDX_admin_rule_v5` (`v5`) USING BTREE,
  KEY `IDX_admin_rule_v0` (`v0`) USING BTREE,
  KEY `IDX_admin_rule_v1` (`v1`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=185 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of admin_rule
-- ----------------------------
INSERT INTO `admin_rule` VALUES ('1', 'g', '1', 'admin', 'a', '', '', '系统管理员');
INSERT INTO `admin_rule` VALUES ('2', 'p', 'admin', 'a', '/api/admin/user*', 'GET|POST|PUT|DELETE', '.*', 'admin-用户权限');
INSERT INTO `admin_rule` VALUES ('3', 'p', 'admin', 'a', '/api/admin/menu*', 'GET|POST|PUT|DELETE', '.*', 'admin-菜单权限');
INSERT INTO `admin_rule` VALUES ('4', 'p', 'admin', 'a', '/api/admin/role*', 'GET|POST|PUT|DELETE', '.*', 'admin-角色权限');
INSERT INTO `admin_rule` VALUES ('5', 'p', 'admin', 'a', '/api/admin/policy*', 'GET|POST|PUT|DELETE', '.*', 'admin-权限模块');
INSERT INTO `admin_rule` VALUES ('6', 'p', 'demo', 'a', '/api/admin/demo*', 'GET|POST|PUT|DELETE', '.*', 'demo-测试权限');
INSERT INTO `admin_rule` VALUES ('7', 'g', '1', 'admin', 'a', '', '', '系统管理员');

-- ----------------------------
-- Table structure for dep
-- ----------------------------
DROP TABLE IF EXISTS `dep`;
CREATE TABLE `dep` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `dep_name` varchar(64) DEFAULT '' COMMENT '部门名称',
  `leader` varchar(64) NOT NULL COMMENT '部门领导人uid',
  `phone` varchar(64) DEFAULT '' COMMENT '部门电话',
  `email` varchar(64) DEFAULT '' COMMENT '部门邮箱',
  `disabled` tinyint(1) DEFAULT NULL COMMENT '0:可用 否则:不可用',
  `parent_id` int(20) DEFAULT NULL,
  `dep_desc` varchar(255) DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of dep
-- ----------------------------
INSERT INTO `dep` VALUES ('1', '集团公司', '董事长', '', '', '0', '0', '总公司', null, null);
INSERT INTO `dep` VALUES ('2', '成都分公司', '总经理', '', '', '0', '1', '成都分公司', null, null);
INSERT INTO `dep` VALUES ('3', '北京分公司', '总经理', '', '', '0', '1', '北京分公司', null, null);
INSERT INTO `dep` VALUES ('4', '人事部', '经理', '', '', '0', '2', '成都分公司-人事部', null, null);
INSERT INTO `dep` VALUES ('5', '公关部', '经理', '', '', '0', '2', '成都分公司-公关部', null, null);
INSERT INTO `dep` VALUES ('6', '财务部', '经理', '', '', '0', '2', '成都分公司-财务部', null, null);
INSERT INTO `dep` VALUES ('7', '技术部', '经理', '', '', '0', '2', '成都分公司-技术部', null, null);
INSERT INTO `dep` VALUES ('8', '市场部', '经理', '', '', '0', '2', '成都分公司-市场部', null, null);
INSERT INTO `dep` VALUES ('9', '人事部', '经理', '', '', '0', '3', '北京分公司-人事部', null, null);
INSERT INTO `dep` VALUES ('10', '公关部', '经理', '', '', '0', '3', '北京分公司-公关部', null, null);
INSERT INTO `dep` VALUES ('11', '财务部', '经理', '', '', '0', '3', '北京分公司-财务部', null, null);
INSERT INTO `dep` VALUES ('12', '技术部', '经理', '', '', '0', '3', '北京分公司-技术部', null, null);
INSERT INTO `dep` VALUES ('13', '市场部', '经理', '', '', '0', '3', '北京分公司-市场部', null, null);

-- ----------------------------
-- Table structure for dep_admin
-- ----------------------------
DROP TABLE IF EXISTS `dep_admin`;
CREATE TABLE `dep_admin` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `dep_ip` int(20) DEFAULT NULL COMMENT '部门id',
  `aid` int(20) DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of dep_admin
-- ----------------------------
INSERT INTO `dep_admin` VALUES ('1', '1', '2');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `path` varchar(64) DEFAULT '',
  `name` varchar(64) DEFAULT '' COMMENT '必须是唯一的',
  `modular` varchar(64) DEFAULT '',
  `component` varchar(64) DEFAULT '',
  `sort` int(3) DEFAULT '0' COMMENT '排序',
  `meta_id` int(20) DEFAULT NULL,
  `parent_id` int(20) DEFAULT NULL,
  `is_sub` tinyint(1) NOT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `key` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('1', 'sys_mgr', 'sys_mgr', '', '', '1', '1', null, '0', null, null);
INSERT INTO `menu` VALUES ('2', 'user', 'user', '/admin/user', '/user', '3', '2', '1', '1', null, null);
INSERT INTO `menu` VALUES ('3', 'role', 'role', '/admin/role', '/role', '2', '3', '1', '1', null, null);
INSERT INTO `menu` VALUES ('4', 'menu', 'menu', '/admin/menu', '/menu', '1', '4', '1', '1', null, null);
INSERT INTO `menu` VALUES ('5', 'dep', 'dep', '/admin/dep', '/dep', '0', '5', '1', '1', null, null);
INSERT INTO `menu` VALUES ('6', 'policy', 'policy', '/admin/policy', '/policy', '4', '26', '1', '1', null, null);
INSERT INTO `menu` VALUES ('10', 'doc', 'doc', '', '', '2', '6', '0', '0', null, null);
INSERT INTO `menu` VALUES ('15', 'components', 'components', '', '', '3', '7', null, '0', null, null);
INSERT INTO `menu` VALUES ('16', 'tree_select_page', 'tree_select_page', '/components/tree-select', '/index', '0', '8', '15', '1', null, null);
INSERT INTO `menu` VALUES ('17', 'count_to_page', 'count_to_page', '/components/count-to', '/count-to', '0', '9', '15', '1', null, null);
INSERT INTO `menu` VALUES ('18', 'drag_list_page', 'drag_list_page', '/components/drag-list', '/drag-list', '0', '10', '15', '1', null, null);
INSERT INTO `menu` VALUES ('19', 'drag_drawer_page', 'drag_drawer_page', '/components/drag-drawer', '/index', '0', '11', '15', '1', null, null);
INSERT INTO `menu` VALUES ('20', 'org_tree_page', 'org_tree_page', '/components/org-tree', '/index', '0', '12', '15', '1', null, null);
INSERT INTO `menu` VALUES ('21', 'tree_table_page', 'tree_table_page', '/components/tree-table', '/index', '0', '13', '15', '1', null, null);
INSERT INTO `menu` VALUES ('22', 'cropper_page', 'cropper_page', '/components/cropper', '/cropper', '0', '14', '15', '1', null, null);
INSERT INTO `menu` VALUES ('23', 'tables_page', 'tables_page', '/components/tables', '/tables', '0', '15', '15', '1', null, null);
INSERT INTO `menu` VALUES ('24', 'split_pane_page', 'split_pane_page', '/components/split-pane', '/split-pane', '0', '16', '15', '1', null, null);
INSERT INTO `menu` VALUES ('25', 'markdown_page', 'markdown_page', '/components/markdown', '/markdown', '0', '17', '15', '1', null, null);
INSERT INTO `menu` VALUES ('26', 'editor_page', 'editor_page', '/components/editor', '/editor', '0', '18', '15', '1', null, null);
INSERT INTO `menu` VALUES ('27', 'icons_page', 'icons_page', '/components/icons', '/icons', '0', '19', '15', '1', null, null);
INSERT INTO `menu` VALUES ('28', 'update', 'update', '', '', '3', '20', null, '0', null, null);
INSERT INTO `menu` VALUES ('29', 'update_table_page', 'update_table_page', '/update', '/update-table', '0', '21', '28', '1', null, null);
INSERT INTO `menu` VALUES ('30', 'update_paste_page', 'update_paste_page', '/update', '/update-paste', '0', '22', '28', '1', null, null);
INSERT INTO `menu` VALUES ('31', 'excel', 'excel', '', '', '4', '23', null, '0', null, null);
INSERT INTO `menu` VALUES ('32', 'upload-excel', 'upload-excel', '/excel', '/upload-excel', '0', '24', '31', '1', null, null);
INSERT INTO `menu` VALUES ('33', 'export-excel', 'export-excel', '/excel', '/export-excel', '0', '25', '31', '1', null, null);
INSERT INTO `menu` VALUES ('34', 'tools_methods', 'tools_methods', '', '', '5', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('35', 'tools_methods_page', 'tools_methods_page', '/tools-methods', '/tools-methods', '0', null, '34', '1', null, null);
INSERT INTO `menu` VALUES ('36', 'i18n', 'i18n', '', '', '6', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('37', 'i18n_page', 'i18n_page', '/i18n', '/i18n-page', '0', null, '36', '1', null, null);
INSERT INTO `menu` VALUES ('38', 'error_store', 'error_store', '', '', '7', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('39', 'error_store_page', 'error_store_page', '/error-store', '/error-store', '0', null, '38', '1', null, null);
INSERT INTO `menu` VALUES ('40', 'error_logger', 'error_logger', '', '', '8', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('41', 'error_logger_page', 'error_logger_page', '/single-page', '/error-logger', '0', null, '40', '1', null, null);
INSERT INTO `menu` VALUES ('42', 'directive', 'directive', '', '', '9', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('43', 'directive_page', 'directive_page', '/directive', '/directive', '0', null, '42', '1', null, null);
INSERT INTO `menu` VALUES ('47', '401', 'error_401', '/error-page', '/401', '10', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('48', '500', 'error_500', '/error-page', '/500', '11', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('49', '*', 'error_404', '/error-page', '/404', '12', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('50', 'ws-demo', 'ws测试', '/ws-demo', '/index', '8', '27', '1', '1', null, null);

-- ----------------------------
-- Table structure for menu_meta
-- ----------------------------
DROP TABLE IF EXISTS `menu_meta`;
CREATE TABLE `menu_meta` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '显示在侧边栏、面包屑和标签栏的文字',
  `hide_in_bread` tinyint(1) DEFAULT '0' COMMENT '设为true后此级路由将不会出现在面包屑中',
  `hide_in_menu` tinyint(1) DEFAULT '0' COMMENT '设为true后在左侧菜单不会显示该页面选项',
  `show_always` tinyint(1) DEFAULT '0',
  `not_cache` tinyint(1) DEFAULT '0' COMMENT '设为true后页面在切换标签后不会缓存，如果需要缓存，无需设置这个字段，而且需要设置页面组件name属性和路由配置的name一致',
  `access` varchar(255) DEFAULT NULL COMMENT '可访问该页面的权限数组，当前路由设置的权限会影响子路由',
  `icon` varchar(255) DEFAULT '' COMMENT '该页面在左侧菜单、面包屑和标签导航处显示的图标，如果是自定义图标，需要在图标名称前加下划线''_''',
  `before_close_name` varchar(255) DEFAULT '' COMMENT '设置该字段，则在关闭当前tab页时会去''@/router/before-close.js''里寻找该字段名对应的方法，作为关闭前的钩子函数',
  `href` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu_meta
-- ----------------------------
INSERT INTO `menu_meta` VALUES ('1', '系统管理', '0', '0', '0', '0', null, 'md-settings', '', '');
INSERT INTO `menu_meta` VALUES ('2', '用户管理', '0', '0', '0', '0', null, 'md-person', '', '');
INSERT INTO `menu_meta` VALUES ('3', '角色管理', '0', '0', '0', '0', null, 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('4', '菜单管理', '0', '0', '0', '0', null, 'md-flower', '', '');
INSERT INTO `menu_meta` VALUES ('5', '部门管理', '0', '0', '0', '0', null, 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('6', '文档', '0', '0', '0', '0', '', 'ios-book', '', 'https://pkg.go.dev/');
INSERT INTO `menu_meta` VALUES ('7', '组件', '0', '0', '0', '0', null, 'logo-buffer', '', '');
INSERT INTO `menu_meta` VALUES ('8', '树状下拉选择器', '0', '0', '0', '0', null, 'md-arrow-dropdown-circle', '', '');
INSERT INTO `menu_meta` VALUES ('9', '数字渐变', '0', '0', '0', '0', null, 'md-trending-up', '', '');
INSERT INTO `menu_meta` VALUES ('10', '拖拽列表', '0', '0', '0', '0', null, 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('11', '可拖拽抽屉', '0', '0', '0', '0', null, 'md-list', '', '');
INSERT INTO `menu_meta` VALUES ('12', '组织结构树', '0', '0', '0', '0', '', 'ios-appstore', '', '');
INSERT INTO `menu_meta` VALUES ('13', '树状表格', '0', '0', '0', '0', null, 'md-git-branch', '', '');
INSERT INTO `menu_meta` VALUES ('14', '图片裁剪', '0', '0', '0', '0', null, 'md-crop', '', '');
INSERT INTO `menu_meta` VALUES ('15', '多功能表格', '0', '0', '0', '0', null, 'md-grid', '', '');
INSERT INTO `menu_meta` VALUES ('16', '分割窗口', '0', '0', '0', '0', null, 'md-pause', '', '');
INSERT INTO `menu_meta` VALUES ('17', 'Markdown编辑器', '0', '0', '0', '0', null, 'logo-markdown', '', '');
INSERT INTO `menu_meta` VALUES ('18', '富文本编辑器', '0', '0', '0', '0', null, 'ios-create', '', '');
INSERT INTO `menu_meta` VALUES ('19', '自定义图标', '0', '0', '0', '0', null, 'ios-create', '', '');
INSERT INTO `menu_meta` VALUES ('20', '数据上传', '0', '0', '1', '0', null, 'md-cloud-upload', '', '');
INSERT INTO `menu_meta` VALUES ('21', '上传Csv', '0', '0', '0', '0', null, 'ios-document', '', '');
INSERT INTO `menu_meta` VALUES ('22', '粘贴表格数据', '0', '0', '0', '0', null, 'md-clipboard', '', '');
INSERT INTO `menu_meta` VALUES ('23', 'EXCEL导入导出', '0', '0', '0', '0', null, 'md-download', '', '');
INSERT INTO `menu_meta` VALUES ('24', '导入EXCEL', '0', '0', '0', '0', null, 'md-add', '', '');
INSERT INTO `menu_meta` VALUES ('25', '导出EXCEL', '0', '0', '0', '0', null, 'md-download', '', '');
INSERT INTO `menu_meta` VALUES ('26', '权限管理', '0', '0', '0', '0', null, 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('27', 'ws测试', '0', '0', '0', '0', null, 'logo-android', '', '');

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `role_key` varchar(100) NOT NULL DEFAULT '' COMMENT '角色key',
  `mid` int(20) NOT NULL COMMENT '菜单id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES ('admin', '2');
INSERT INTO `role_menu` VALUES ('admin', '3');
INSERT INTO `role_menu` VALUES ('admin', '6');
INSERT INTO `role_menu` VALUES ('demo', '16');
INSERT INTO `role_menu` VALUES ('demo', '22');
INSERT INTO `role_menu` VALUES ('admin', '4');
INSERT INTO `role_menu` VALUES ('admin', '10');
INSERT INTO `role_menu` VALUES ('admin', '50');
