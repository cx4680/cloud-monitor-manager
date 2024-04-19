CREATE TABLE `large_screen_remote_database`
(
    `region` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '区域',
    `host`   varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '主机',
    `port`   varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '端口',
    `db`     varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '数据库',
    `user`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '账号',
    `pass`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码'
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

INSERT INTO `large_screen_remote_database`
VALUES ('cc-wuhan-0', '10.109.62.11', '33331', 'unified_billing', 'unified_billing', 'R3Vkg_f!YKq9QGOu'),
       ('cc-wuhan-2', '10.109.62.11', '33331', 'unified_billing', 'unified_billing', 'R3Vkg_f!YKq9QGOu'),
       ('nc-beijing-2', '10.109.62.11', '33331', 'unified_billing', 'unified_billing', 'R3Vkg_f!YKq9QGOu'),
       ('nc-cec-gts', '10.127.142.8', '33331', 'unified_billing', 'unified_billing', 'R3Vkg_f!YKq9QGOu'),
       ('nc-cec-cts', '10.127.142.8', '33331', 'unified_billing', 'unified_billing', 'R3Vkg_f!YKq9QGOu');


CREATE TABLE `large_screen_storage_login`
(
    `region`     varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '区域',
    `vendor`     varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '厂商',
    `type`       varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '类型',
    `username`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '账号',
    `password`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
    `manage_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '管理地址'
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

INSERT INTO `large_screen_storage_login`
VALUES ('cc-wuhan-0', 'inspur', 'ceph', 'superuser', 'passw0rd', 'https://10.127.135.5'),
       ('cc-wuhan-2', 'inspur', 'ceph', 'superuser', 'passw0rd', 'https://10.127.207.132'),
       ('cc-wuhan-2', 'inspur', 'ceph', 'superuser', 'passw0rd', 'https://10.127.207.5'),
       ('cc-wuhan-2', 'inspur', 'ceph', 'cecloud01', 'Cecloud@12345#$', 'https://10.127.208.67'),
       ('cc-wuhan-2', 'inspur', 'ceph', 'superuser', 'Passw0rd!', 'https://10.127.208.65'),
       ('cc-wuhan-2', 'inspur', 'ceph', 'cecloud', 'Passw0rd@#', 'https://100.65.9.193'),
       ('nc-beijing-2', 'inspur', 'ceph', 'superuser', 'passw0rd', 'https://10.110.10.1'),
       ('nc-beijing-2', 'inspur', 'ceph', 'superuser', 'passw0rd', 'https://10.110.10.10'),
       ('nc-beijing-2', 'inspur', 'ceph', 'cecloud', 'Passw0rd@#', 'https://100.65.10.193'),
       ('nc-beijing-2', 'inspur', 'ceph', 'cecloud01', 'Cecloud@12345#$', 'https://10.110.9.80'),
       ('nc-beijing-2', 'inspur', 'ceph', 'superuser', 'Passw0rd!', 'https://10.110.9.70'),
       ('nc-cec-gts', 'inspur', 'ceph', 'superuser', 'passw0rd', 'https://10.109.3.193'),
       ('nc-cec-gts', 'inspur', 'ceph', 'superuser', 'passw0rd', 'https://100.65.11.193'),
       ('nc-cec-gts', 'inspur', 'ceph', 'cecloud01', 'Cecloud@1234#$', 'https://10.109.3.140'),
       ('nc-cec-gts', 'inspur', 'ceph', 'cecloud01', 'Cecloud@1234#$', 'https://10.109.3.130'),
       ('nc-cec-cts', 'inspur', 'ceph', 'superuser', 'Cecloud@1234!@', 'https://10.109.67.193'),
       ('nc-cec-cts', 'inspur', 'ceph', 'superuser', 'Cecloud@1234!@', 'http://10.109.75.2:8056');


CREATE TABLE `large_screen_resource_storage`
(
    `resource_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '资源id',
    `time`        varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '时间',
    `type`        varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '资源类型',
    `value`       bigint                                                 NULL DEFAULT NULL COMMENT '数值',
    `create_time` datetime                                               NULL DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`resource_id`, `time`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = DYNAMIC;