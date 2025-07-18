-- 初始数据
-- 插入默认部门
INSERT INTO `depts` (`dept_id`, `parent_id`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, '总公司', 1, 'admin', '12345678901', 'admin@example.com', '1', '2021-01-01 00:00:00', '2021-01-01 00:00:00', NULL);

-- 插入默认角色
INSERT INTO `roles` (`role_id`, `role_name`, `role_key`, `status`, `remark`) VALUES
(1, '超级管理员', 'admin','1', '超级管理员'),
(2, '普通用户', 'user', '1', '普通用户');

-- 插入默认菜单
INSERT INTO `menus` (`menu_id`, `menu_name`, `parent_id`, `perms`, `menu_type`, `menu_visible`, `menu_is_link`, `icon`, `path`, `sort`, `status`, `remark`) VALUES
-- 系统管理
(1, '系统管理', 0, '', '1', '1', '0', 'system', '/system', 1, '1', '系统管理目录'),
-- 用户管理
(2, '用户管理', 1, 'system:user:list', '2', '1', '0', 'user', '/system/user', 1, '1', '用户管理菜单'),
(3, '用户查询', 2, 'system:user:query', '3', '1', '0', '', '', 1, '1', '用户查询按钮'),
(4, '用户新增', 2, 'system:user:add', '3', '1', '0', '', '', 2, '1', '用户新增按钮'),
(5, '用户修改', 2, 'system:user:edit', '3', '1', '0', '', '', 3, '1', '用户修改按钮'),
(6, '用户删除', 2, 'system:user:remove', '3', '1', '0', '', '', 4, '1', '用户删除按钮'),
-- 角色管理
(7, '角色管理', 1, 'system:role:list', '2', '1', '0', 'role', '/system/role', 2, '1', '角色管理菜单'),
(8, '角色查询', 7, 'system:role:query', '3', '1', '0', '', '', 1, '1', '角色查询按钮'),
(9, '角色新增', 7, 'system:role:add', '3', '1', '0', '', '', 2, '1', '角色新增按钮'),
(10, '角色修改', 7, 'system:role:edit', '3', '1', '0', '', '', 3, '1', '角色修改按钮'),
(11, '角色删除', 7, 'system:role:remove', '3', '1', '0', '', '', 4, '1', '角色删除按钮'),
-- 菜单管理
(12, '菜单管理', 1, 'system:menu:list', '2', '1', '0', 'menu', '/system/menu', 3, '1', '菜单管理菜单'),
(13, '菜单查询', 12, 'system:menu:query', '3', '1', '0', '', '', 1, '1', '菜单查询按钮'),
(14, '菜单新增', 12, 'system:menu:add', '3', '1', '0', '', '', 2, '1', '菜单新增按钮'),
(15, '菜单修改', 12, 'system:menu:edit', '3', '1', '0', '', '', 3, '1', '菜单修改按钮'),
(16, '菜单删除', 12, 'system:menu:remove', '3', '1', '0', '', '', 4, '1', '菜单删除按钮');

-- 为超级管理员分配所有菜单权限
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6),
(1, 7), (1, 8), (1, 9), (1, 10), (1, 11),
(1, 12), (1, 13), (1, 14), (1, 15), (1, 16);

-- 为普通用户分配基础查询权限
INSERT INTO `role_menus` (`role_id`, `menu_id`) VALUES
(2, 1), (2, 2), (2, 3), (2, 7), (2, 8), (2, 12), (2, 13);

-- 插入默认管理员用户（密码：admin123）
INSERT INTO `users` (`user_id`, `username`, `nickname`, `email`, `password`, `status`, `ack_code`) VALUES
(1, 'admin', '系统管理员', 'admin@example.com', '$2a$10$F5pYyOtY.eGS1/AFLEE/r.QGa/FyTzEsL7Ps5uANOKlZAlFiCIIou', '1', '123456');

-- 为管理员分配超级管理员角色
INSERT INTO `user_roles` (`user_id`, `role_id`) VALUES (1, 1); 