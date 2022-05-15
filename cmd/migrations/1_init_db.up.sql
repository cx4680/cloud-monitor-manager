DROP TABLE IF EXISTS `t_monitor_product`;
CREATE TABLE `t_monitor_product`  (
                                      `biz_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '-1' COMMENT '业务Id',
                                      `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `status` tinyint UNSIGNED NULL DEFAULT NULL,
                                      `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `create_user` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `create_time` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `route` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `cron` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `host` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `page_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `abbreviation` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                                      `sort` int NULL DEFAULT NULL COMMENT '排序',
                                      `monitor_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `t_monitor_product` VALUES ('1', '云服务器ECS', 1, 'ecs', NULL, NULL, '/productmonitoring/ecs', '0 0 0/1 * * ?', 'http://cus-resource-center-svc.product-cec-cbc', '/rc/resource/page/outer', 'ecs', NULL, NULL, '云产品监控');
INSERT INTO `t_monitor_product` VALUES ('2', '弹性公网IP', 1, 'eip', NULL, NULL, '/productmonitoring/eip', '0 0 0/1 * * ?', 'http://cus-resource-center-svc.product-cec-cbc', '/rc/resource/page/outer', 'eip', NULL, NULL, '云产品监控');
INSERT INTO `t_monitor_product` VALUES ('3', '负载均衡SLB', 1, 'slb', NULL, NULL, '/productmonitoring/slb', '0 0 0/1 * * ?', 'http://cus-resource-center-svc.product-cec-cbc', '/rc/resource/page/outer', 'slb', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('4', '云备份CBR', 1, 'cbr', NULL, NULL, '/productmonitoring/cbr', '0 0 0/1 * * ?', 'http://product-backup-backup-manage.product-backup', '/noauth/backup/vault/pageList', 'cbr', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('5', 'NAT网关', 1, 'nat', NULL, NULL, '/productmonitoring/nat', '0 0 0/1 * * ?', 'http://product-nat-controller-nat-manage.product-nat-gw', '/nat-gw/inner/nat/page', 'nat', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('6', '云数据库MySQL', 1, 'mysql', NULL, NULL, '/productmonitoring/mysql', '0 0 0/1 * * ?', 'http://product-mysql-rds-mysql-manage.product-mysql.svc.cluster.local:8888', '/v1/mysql/instance', 'mysql', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('7', '云数据库达梦RdbDM', 1, 'dm', NULL, NULL, '/productmonitoring/dm', '0 0 0/1 * * ?', 'http://product-dm-rds-dm-manage.product-dm.svc.cluster.local:8888', '/v1/dm/instance', 'dm', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('8', '云数据库PostgreSQL', 1, 'pg', NULL, NULL, '/productmonitoring/pg', '0 0 0/1 * * ?', 'http://product-postgresql-rdb-pg-manage.product-pg.svc.cluster.local:8888', '/v1/pg/instance/', 'pg', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('9', '消息队列Kafka', 1, 'kafka', NULL, NULL, '/productmonitoring/kafka', '0 0 0/1 * * ?', 'http://cmq-kafka.product-cmq-kafka:8080', '/kafka/v1/cluster', 'kafka', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('10', '传统裸金属', 1, 'bms', NULL, NULL, '/productmonitoring/bms', '0 0 0/1 * * ?', 'http://bms-manage-bms-union.product-bms-union:8082', '/compute/bms/ops/v1/tenants/{tenantId}/servers', 'bms', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('11', '弹性裸金属', 1, 'ebms', NULL, NULL, '/productmonitoring/ebms', '0 0 0/1 * * ?', 'http://bms-manage-bms-union.product-bms-union:8081', '/compute/ebms/ops/v1/tenants/{tenantId}/servers', 'ebms', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('12', '缓存数据库Redis', 1, 'redis', NULL, NULL, '/productmonitoring/redis', '0 0 0/1 * * ?', 'http://product-redis-ndb-redis-manage.product-redis.svc.cluster.local:8888', '/v1/redis/instance', 'redis', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('13', '非关系型数据库NoSQL', 1, 'mongo', NULL, NULL, '/productmonitoring/mongo', '0 0 0/1 * * ?', 'http://dbaas-manage.product-dbaas.svc.cluster.local', '/v2/mongo/instance', 'mongo', NULL, NULL, '云产品监控');
-- INSERT INTO `t_monitor_product` VALUES ('14', 'API网关CGW', 1, 'cgw', NULL, NULL, '/productmonitoring/cgw', '0 0 0/1 * * ?', 'http://cgw-cgw-manage-admin.product-cgw.svc.cluster.local:8080', '/gateway/instance/page', 'cgw', NULL, NULL, '云产品监控');

DROP TABLE IF EXISTS `t_monitor_item`;
CREATE TABLE `t_monitor_item`  (
                                   `biz_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '-1' COMMENT '业务Id',
                                   `product_biz_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '-1' COMMENT '产品业务Id',
                                   `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `metric_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `labels` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `metrics_linux` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `metrics_windows` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `statistics` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `unit` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `frequency` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `type` tinyint UNSIGNED NULL DEFAULT NULL,
                                   `is_display` tinyint UNSIGNED NULL DEFAULT NULL,
                                   `status` tinyint UNSIGNED NULL DEFAULT NULL,
                                   `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `create_user` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `create_time` datetime NULL DEFAULT NULL,
                                   `show_expression` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
                                   `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `display` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'chart,rule,scaling' COMMENT '展示位置',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 207 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `t_monitor_item` VALUES ('1', '1', 'CPU使用率', 'ecs_cpu_usage', 'instance', '100 - (100 * (sum by(instance) (irate(ecs_cpu_seconds_total{mode=\"idle\",$INSTANCE}[3m])) / sum by(instance) (irate(ecs_cpu_seconds_total{$INSTANCE}[3m]))))', NULL, NULL, '%', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('2', '1', 'CPU1分钟平均负载', 'ecs_load1', 'instance', 'ecs_load1{$INSTANCE}', NULL, NULL, NULL, NULL, 2, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('3', '1', 'CPU5分钟平均负载', 'ecs_load5', 'instance', 'ecs_load5{$INSTANCE}', NULL, NULL, NULL, NULL, 2, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('4', '1', 'CPU15分钟平均负载', 'ecs_load15', 'instance', 'ecs_load15{$INSTANCE}', NULL, NULL, NULL, NULL, 2, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('5', '1', '内存使用量', 'ecs_memory_used', 'instance', 'ecs_memory_MemTotal_bytes{$INSTANCE} - ecs_memory_MemFree_bytes', NULL, NULL, 'Byte', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('6', '1', '内存使用率', 'ecs_memory_usage', 'instance', '100 * ((ecs_memory_MemTotal_bytes{$INSTANCE} - ecs_memory_MemFree_bytes) / ecs_memory_MemTotal_bytes)', NULL, NULL, '%', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('7', '1', '磁盘使用率', 'ecs_disk_usage', 'instance,device', '100 * ((ecs_filesystem_size_bytes{$INSTANCE} - ecs_filesystem_free_bytes) / ecs_filesystem_size_bytes)', NULL, NULL, '%', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('8', '1', '磁盘读速率', 'ecs_disk_read_rate', 'instance,device', 'irate(ecs_disk_read_bytes_total{$INSTANCE}[3m])', NULL, NULL, 'Byte/s', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('9', '1', '磁盘写速率', 'ecs_disk_write_rate', 'instance,device', 'irate(ecs_disk_written_bytes_total{$INSTANCE}[3m])', NULL, NULL, 'Byte/s', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('10', '1', '磁盘读IOPS', 'ecs_disk_read_iops', 'instance,device', 'irate(ecs_disk_reads_completed_total{$INSTANCE}[3m])', NULL, NULL, '次', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('11', '1', '磁盘写IOPS', 'ecs_disk_write_iops', 'instance,device', 'irate(ecs_disk_writes_completed_total{$INSTANCE}[3m])', NULL, NULL, '次', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('12', '1', '流入带宽', 'ecs_network_receive_rate', 'instance,device', 'irate(ecs_network_receive_bytes_total{$INSTANCE}[3m]) / 1024 / 1024 * 8', NULL, NULL, 'Mbps', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('13', '1', '流出带宽', 'ecs_network_transmit_rate', 'instance,device', 'irate(ecs_network_transmit_bytes_total{$INSTANCE}[3m]) / 1024 / 1024 * 8', NULL, NULL, 'Mbps', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('14', '1', '包接收速率', 'ecs_network_receive_packets_rate', 'instance,device', 'irate(ecs_network_receive_packets_total{$INSTANCE}[3m])', NULL, NULL, '个/s', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('15', '1', '包发送速率', 'ecs_network_transmit_packets_rate', 'instance,device', 'irate(ecs_network_transmit_packets_total{$INSTANCE}[3m])', NULL, NULL, '个/s', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('16', '1', '磁盘剩余存储量', 'ecs_filesystem_free_bytes', 'instance,device', 'ecs_filesystem_free_bytes{$INSTANCE}', NULL, NULL, 'GB', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('17', '1', '磁盘已用存储量', 'ecs_disk_used', 'instance,device', 'ecs_filesystem_size_bytes{$INSTANCE} - ecs_filesystem_free_bytes', NULL, NULL, 'GB', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('18', '1', '磁盘存储总量', 'ecs_filesystem_size_bytes', 'instance,device', 'ecs_filesystem_size_bytes{$INSTANCE}', NULL, NULL, 'GB', NULL, 2, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('19', '1', '(基础)CPU的平均使用率', 'ecs_cpu_base_usage', 'instance', '100 * avg by(instance)(irate(ecs_base_vcpu_seconds{$INSTANCE}[6m]))', NULL, NULL, '%', NULL, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('20', '1', '(基础)磁盘读速率', 'ecs_disk_base_read_rate', 'instance,drive', 'irate(ecs_base_storage_traffic_bytes_total{type=\"read\",$INSTANCE}[6m])', NULL, NULL, 'Byte/s', NULL, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('21', '1', '(基础)磁盘写速率', 'ecs_disk_base_write_rate', 'instance,drive', 'irate(ecs_base_storage_traffic_bytes_total{type=\"write\",$INSTANCE}[6m])', NULL, NULL, 'Byte/s', NULL, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('22', '1', '(基础)网卡下行带宽', 'ecs_network_base_receive_rate', 'instance,interface', 'irate(ecs_base_network_traffic_bytes_total{type=\"rx\",$INSTANCE}[6m])', NULL, NULL, 'Byte/s', NULL, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('23', '1', '(基础)网卡上行带宽', 'ecs_network_base_transmit_rate', 'instance,interface', 'irate(ecs_base_network_traffic_bytes_total{type=\"tx\",$INSTANCE}[6m])', NULL, NULL, 'Byte/s', NULL, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('24', '1', '(基础)磁盘读IOPS', 'ecs_disk_base_read_iops', 'instance,drive', 'sum(irate(ecs_base_storage_iops_total{type=\"read\",$INSTANCE}[15m])) by (instance,drive)', NULL, NULL, '次', NULL, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');
INSERT INTO `t_monitor_item` VALUES ('25', '1', '(基础)磁盘写IOPS', 'ecs_disk_base_write_iops', 'instance,drive', 'sum(irate(ecs_base_storage_iops_total{type=\"write\",$INSTANCE}[15m])) by (instance,drive)', NULL, NULL, '次', NULL, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule,scaling');

INSERT INTO `t_monitor_item` VALUES ('26', '2', '出网带宽', 'eip_upstream_bandwidth', 'instance', 'sum(eip_upstream_bits_rate{$INSTANCE}) by (instance,eip)', NULL, NULL, 'bps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('27', '2', '入网带宽', 'eip_downstream_bandwidth', 'instance', 'sum(eip_downstream_bits_rate{$INSTANCE}) by (instance,eip)', NULL, NULL, 'bps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('28', '2', '出网流量', 'eip_upstream', 'instance', '((sum(eip_upstream_bits_rate{$INSTANCE}) by (instance,eip))/8)*60', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('29', '2', '入网流量', 'eip_downstream', 'instance', '((sum(eip_downstream_bits_rate{$INSTANCE}) by (instance,eip))/8)*60', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('30', '2', '出网带宽使用率', 'eip_upstream_bandwidth_usage', 'instance', '(sum(eip_upstream_bits_rate{$INSTANCE}) by (instance,eip) / avg(eip_config_upstream_bandwidth{$INSTANCE}) by (instance,eip)) * 100', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');

INSERT INTO `t_monitor_item` VALUES ('31', '3', '出网带宽', 'slb_out_bandwidth', 'instance,slb_listener_id', 'sum by(instance) (Slb_http_bps_out_rate{$INSTANCE})', NULL, NULL, 'bps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('32', '3', '入网带宽', 'slb_in_bandwidth', 'instance,slb_listener_id', 'sum by(instance) (Slb_http_bps_in_rate{$INSTANCE})', NULL, NULL, 'bps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('33', '3', '并发连接数', 'slb_max_connection', 'instance,slb_listener_id', 'sum by(instance) (Slb_all_connection_count{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('34', '3', '活跃连接数', 'slb_active_connection', 'instance,slb_listener_id', 'sum by (instance)(Slb_all_est_connection_count{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('35', '3', '非活跃连接数', 'slb_inactive_connection', 'instance,slb_listener_id', 'sum by (instance) (Slb_all_none_est_connection_count{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('36', '3', '新建连接数', 'slb_new_connection', 'instance,slb_listener_id', 'sum by(instance) (Slb_new_connection_rate{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('37', '3', '丢弃连接数', 'slb_drop_connection', 'instance,slb_listener_id', 'sum by(instance)(Slb_drop_connection_rate{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('38', '3', '异常后端云服务器数', 'slb_unhealthyserver', 'instance', 'avg by(instance) (Slb_unhealthy_server_count{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('39', '3', '正常后端云服务器数', 'slb_healthyserver', 'instance', 'avg by(instance) (Slb_healthy_server_count{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('40', '3', '7层协议查询速率', 'slb_qps', 'instance,slb_listener_id', 'sum by(instance)(Slb_request_rate{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('41', '3', '7层协议返回客户端2xx状态码数', 'slb_statuscode2xx', 'instance', 'sum by(instance) (Slb_http_2xx_rate{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('42', '3', '7层协议返回客户端3xx状态码数', 'slb_statuscode3xx', 'instance', 'sum by(instance) (Slb_http_3xx_rate{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('43', '3', '7层协议返回客户端4xx状态码数', 'slb_statuscode4xx', 'instance', 'sum by(instance) (Slb_http_4xx_rate{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
INSERT INTO `t_monitor_item` VALUES ('44', '3', '7层协议返回客户端5xx状态码数', 'slb_statuscode5xx', 'instance', 'sum by(instance) (Slb_http_5xx_rate{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');

-- INSERT INTO `t_monitor_item` VALUES ('45', '4', '存储库总量', 'cbr_vault_size', 'instance', 'cbr_vault_size{$INSTANCE}', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('46', '4', '存储库使用量', 'cbr_vault_used', 'instance', 'cbr_vault_used{$INSTANCE}', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('47', '4', '存储库使用率', 'cbr_vault_usage_rate', 'instance', 'cbr_vault_used{$INSTANCE} / cbr_vault_size{$INSTANCE} * 100', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('48', '5', 'SNAT连接数', 'snat_connection', 'instance', 'sum by (instance)(Nat_snat_total_connection_count{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('49', '5', '入方向带宽', 'inbound_bandwidth', 'instance', 'sum by (instance)(rate(Nat_recv_bytes_total_count{$INSTANCE}[1m])*8)', NULL, NULL, 'bps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('50', '5', '出方向带宽', 'outbound_bandwidth', 'instance', 'sum by (instance)(rate(Nat_send_bytes_total_count{$INSTANCE}[1m])*8)', NULL, NULL, 'bps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('51', '5', '入方向流量', 'inbound_traffic', 'instance', 'sum by (instance)(Nat_recv_bytes_total_count{$INSTANCE})', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('52', '5', '出方向流量', 'outbound_traffic', 'instance', 'sum by (instance)(Nat_send_bytes_total_count{$INSTANCE})', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('53', '5', '入方向PPS', 'inbound_pps', 'instance', 'sum by (instance)(rate(Nat_recv_packets_total_count{$INSTANCE}[1m]))', NULL, NULL, 'pps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('54', '5', '出方向PPS', 'outbound_pps', 'instance', 'sum by (instance)(rate(Nat_send_packets_total_count{$INSTANCE}[1m]))', NULL, NULL, 'pps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('55', '5', 'SNAT连接数使用率', 'snat_connection_ratio', 'instance', 'sum by (instance)(Nat_snat_total_connection_count{$INSTANCE}) / avg by (instance)(Nat_nat_max_connection_count{$INSTANCE}) *100', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('56', '6', '主从复制IO线程状态', 'mysql_slave_io', 'instance', 'mysql_slave_io{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"Basic\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('57', '6', '主从复制SQL线程状态', 'mysql_slave_sql', 'instance', 'mysql_slave_sql{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"Basic\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('58', '6', '复制延迟', 'mysql_slave_seconds_behind_master', 'instance', 'mysql_slave_seconds_behind_master{$INSTANCE}', NULL, NULL, 's', NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"Basic\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('59', '6', '活跃连接数', 'mysql_active_connections', 'instance', 'mysql_active_connections{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('60', '6', '当前连接占比', 'mysql_current_connection_percent', 'instance', 'mysql_current_connection_percent{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('61', '6', 'QPS', 'mysql_qps', 'instance', 'mysql_qps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('62', '6', 'TPS', 'mysql_tps', 'instance', 'mysql_tps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('63', '6', '每秒查询数量', 'mysql_select_ps', 'instance', 'mysql_select_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('64', '6', '每秒更新数量', 'mysql_update_ps', 'instance', 'mysql_update_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('65', '6', '每秒插入数量', 'mysql_insert_ps', 'instance', 'mysql_insert_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('66', '6', '每秒删除数量', 'mysql_delete_ps', 'instance', 'mysql_delete_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('67', '6', 'CPU使用率', 'mysql_cpu_usage', 'instance', 'mysql_cpu_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('68', '6', '内存使用率', 'mysql_mem_usage', 'instance', 'mysql_mem_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('69', '6', '磁盘使用率', 'mysql_disk_usage', 'instance', 'mysql_disk_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('70', '6', 'InnoDB每秒查询行数', 'mysql_innodb_select_ps', 'instance', 'mysql_innodb_select_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('71', '6', 'InnoDB每秒更新行数', 'mysql_innodb_update_ps', 'instance', 'mysql_innodb_update_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('72', '6', 'InnoDB每秒插入行数', 'mysql_innodb_insert_ps', 'instance', 'mysql_innodb_insert_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('73', '6', 'InnoDB每秒删除行数', 'mysql_innodb_delete_ps', 'instance', 'mysql_innodb_delete_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('74', '6', 'InnoDB缓存命中率', 'mysql_innodb_cache_hit_rate', 'instance', 'mysql_innodb_cache_hit_rate{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('75', '6', 'InnoDB每秒读次数', 'mysql_innodb_reads_ps', 'instance', 'mysql_innodb_reads_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('76', '6', 'InnoDB每秒写次数', 'mysql_innodb_writes_ps', 'instance', 'mysql_innodb_writes_ps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('77', '6', 'InnoDB脏页数量', 'mysql_innodb_buffer_pool_pages_dirty', 'instance', 'mysql_innodb_buffer_pool_pages_dirty{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('78', '6', 'InnoDB脏页大小', 'mysql_innodb_buffer_pool_bytes_dirty', 'instance', 'mysql_innodb_buffer_pool_bytes_dirty{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('79', '6', 'InnoDB日志写等待', 'mysql_innodb_log_waits', 'instance', 'mysql_innodb_log_waits{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('80', '6', '大事务', 'mysql_binlog_cache_disk_use', 'instance', 'mysql_binlog_cache_disk_use{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('81', '6', '每分钟慢查询数量', 'mysql_slow_queries_per_min', 'instance', 'mysql_slow_queries_per_min{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('82', '6', '长时间执行SQL(执行时间超过600秒)', 'mysql_long_query_count', 'instance', 'mysql_long_query_count{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('83', '6', '长时间执行SQL报警(执行时间超过1800秒)', 'mysql_long_query_alert_count', 'instance', 'mysql_long_query_alert_count{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('84', '6', '执行语句频率', 'mysql_exec_statememt_frequency', 'instance', 'mysql_exec_statememt_frequency{$INSTANCE}', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('85', '6', '读频率', 'mysql_read_frequency', 'instance', 'mysql_exec_statememt_frequency{$INSTANCE}', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('86', '6', '写频率', 'mysql_write_frequency', 'instance', 'mysql_write_frequency{$INSTANCE}', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('87', '6', 'TOP语句平均执行时间', 'mysql_top_statememt_avg_exec_time', 'instance', 'mysql_top_statememt_avg_exec_time{$INSTANCE}', NULL, NULL, 'us', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('88', '6', 'TOP语句执行错误率', 'mysql_top_statememt_exec_err_rate', 'instance', 'mysql_top_statememt_exec_err_rate{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('89', '6', '当前打开连接数', 'mysql_current_cons_num', 'instance', 'mysql_current_cons_num{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('90', '7', '每秒事务数', 'dm_global_status_tps', 'instance', 'rate(dm_global_status_tps{$INSTANCE}[1m])', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('91', '7', '每秒执行select SQL语句数', 'dm_global_status_qps', 'instance', 'rate(dm_global_status_qps{$INSTANCE}[1m])', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('92', '7', '每秒执行insert SQL语句数', 'dm_global_status_ips', 'instance', 'rate(dm_global_status_ips{$INSTANCE}[1m])', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('93', '7', '每秒执行delete SQL语句数', 'dm_global_status_dps', 'instance', 'rate(dm_global_status_dps{$INSTANCE}[1m])', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('94', '7', '每秒执行update SQL语句数', 'dm_global_status_ups', 'instance', 'rate(dm_global_status_ups{$INSTANCE}[1m])', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('95', '7', '每秒执行DDL SQL语句数', 'dm_global_status_ddlps', 'instance', 'rate(dm_global_status_ddlps{$INSTANCE}[1m])', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('96', '7', '每秒从客户端接收字节数', 'dm_global_status_nioips', 'instance', 'rate(dm_global_status_nioips{$INSTANCE}[1m])', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('97', '7', '每秒往客户端发送字节数', 'dm_global_status_nio_ops', 'instance', 'rate(dm_global_status_nio_ops{$INSTANCE}[1m])', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('98', '7', '每秒读取字节数', 'dm_global_status_fio_ips', 'instance', 'rate(dm_global_status_fio_ips{$INSTANCE}[1m])', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('99', '7', '每秒写入字节数', 'dm_global_status_fio_ops', 'instance', 'rate(dm_global_status_fio_ops{$INSTANCE}[1m])', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('100', '7', '内存占用字节数', 'dm_global_status_mem_used', 'instance', 'dm_global_status_mem_used{$INSTANCE}', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('101', '7', 'CPU使用率', 'dm_global_status_cpu_use_rate', 'instance', 'dm_global_status_cpu_use_rate{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('102', '7', '内存使用率', 'dm_global_status_mem_use_rate', 'instance', 'dm_global_status_mem_use_rate{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('103', '7', '总会话数', 'dm_global_status_sessions', 'instance', 'dm_global_status_sessions{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('104', '7', '活动会话数', 'dm_global_status_active_sessions', 'instance', 'dm_global_status_active_sessions{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('105', '7', '等待处理任务数', 'dm_global_status_task_waiting', 'instance', 'dm_global_status_task_waiting{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('106', '7', '已处理任务数', 'dm_global_status_task_ready', 'instance', 'dm_global_status_task_ready{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('107', '7', '已处理任务的总等待时间', 'dm_global_status_task_total_wait_time', 'instance', 'dm_global_status_task_total_wait_time{$INSTANCE}', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('108', '7', '已处理任务的平均等待时间', 'dm_global_status_avg_wait_time', 'instance', 'dm_global_status_avg_wait_time{$INSTANCE}', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('109', '7', '线程数', 'dm_global_status_threads', 'instance', 'dm_global_status_threads{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('110', '8', 'CPU使用率', 'pg_cpu_usage', 'instance', 'pg_cpu_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('111', '8', '内存使用率', 'pg_mem_usage', 'instance', 'pg_mem_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('112', '8', '磁盘使用率', 'pg_disk_usage', 'instance', 'pg_disk_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('113', '8', 'QPS', 'pg_qps', 'instance', 'pg_qps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('114', '8', '读QPS', 'pg_rqps', 'instance', 'pg_rqps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('115', '8', '写QPS', 'pg_wqps', 'instance', 'pg_wqps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('116', '8', 'TPS', 'pg_tps', 'instance', 'pg_tps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('117', '8', '最长平均执行时间', 'pg_mean_exec_time', 'instance', 'pg_mean_exec_time{$INSTANCE}', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('118', '8', '当前打开连接数', 'pg_open_ct_num', 'instance', 'pg_open_ct_num{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('119', '8', '当前活跃连接数', 'pg_active_ct_num', 'instance', 'pg_active_ct_num{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('120', '9', '在线Broker数', 'kafka_brokers', 'instance', 'sum by (instance) (kafka_brokers{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('121', '9', '分区的数量', 'kafka_server_replicamanager_partitioncount', 'instance', 'sum by (instance) (kafka_server_replicamanager_partitioncount{$INSTANCE})', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('122', '9', '生产速率', 'kafka_server_brokertopicmetrics_bytesinpersec', 'instance', 'sum by (instance) (kafka_server_brokertopicmetrics_bytesinpersec{$INSTANCE})', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('123', '9', '消费速率', 'kafka_server_brokertopicmetrics_bytesoutpersec', 'instance', 'sum by (instance) (kafka_server_brokertopicmetrics_bytesoutpersec{$INSTANCE})', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('124', '9', '消息生产速率', 'kafka_server_brokertopicmetrics_messagesinpersec', 'instance', 'sum by (instance) (kafka_server_brokertopicmetrics_messagesinpersec{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('125', '9', '落后的消费量', 'kafka_consumergroup_lag', 'instance', 'sum by (instance) (kafka_consumergroup_lag{$INSTANCE})', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
--
-- INSERT INTO `t_monitor_item` VALUES ('126', '10', 'CPU使用率', 'bms_cpu_usage', 'instance', '100 - (100 * (sum by(instance) (irate(ecs_cpu_seconds_total{mode=\"idle\",$INSTANCE}[3m])) / sum by(instance) (irate(ecs_cpu_seconds_total{$INSTANCE}[3m]))))', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('127', '10', 'CPU1分钟平均负载', 'bms_load1', 'instance', 'ecs_load1{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('128', '10', 'CPU5分钟平均负载', 'bms_load5', 'instance', 'ecs_load5{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('129', '10', 'CPU15分钟平均负载', 'bms_load15', 'instance', 'ecs_load15{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('130', '10', '内存使用量', 'bms_memory_used', 'instance', 'ecs_memory_MemTotal_bytes{$INSTANCE} - ecs_memory_MemFree_bytes', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('131', '10', '内存使用率', 'bms_memory_usage', 'instance', '100 * ((ecs_memory_MemTotal_bytes{$INSTANCE} - ecs_memory_MemFree_bytes) / ecs_memory_MemTotal_bytes)', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('132', '10', '磁盘使用率', 'bms_disk_usage', 'instance,device', '100 * ((ecs_filesystem_size_bytes{$INSTANCE} - ecs_filesystem_free_bytes) / ecs_filesystem_size_bytes)', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('133', '10', '磁盘读速率', 'bms_disk_read_rate', 'instance,device', 'irate(ecs_disk_read_bytes_total{$INSTANCE}[3m])', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('134', '10', '磁盘写速率', 'bms_disk_write_rate', 'instance,device', 'irate(ecs_disk_written_bytes_total{$INSTANCE}[3m])', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('135', '10', '磁盘读IOPS', 'bms_disk_read_iops', 'instance,device', 'irate(ecs_disk_reads_completed_total{$INSTANCE}[3m])', NULL, NULL, '次', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('136', '10', '磁盘写IOPS', 'bms_disk_write_iops', 'instance,device', 'irate(ecs_disk_writes_completed_total{$INSTANCE}[3m])', NULL, NULL, '次', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('137', '10', '流入带宽', 'bms_network_transmit_rate', 'instance,device', 'irate(ecs_network_receive_bytes_total{$INSTANCE}[3m]) / 1024 / 1024 * 8', NULL, NULL, 'Mbps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('138', '10', '流出带宽', 'bms_network_receive_rate', 'instance,device', 'irate(ecs_network_transmit_bytes_total{$INSTANCE}[3m]) / 1024 / 1024 * 8', NULL, NULL, 'Mbps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('139', '10', '包接收速率', 'bms_network_receive_packets_rate', 'instance,device', 'irate(ecs_network_receive_packets_total{$INSTANCE}[3m])', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('140', '10', '包发送速率', 'bms_network_transmit_packets_rate', 'instance,device', 'irate(ecs_network_transmit_packets_total{$INSTANCE}[3m])', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('141', '10', '磁盘剩余存储量', 'bms_filesystem_free_bytes', 'instance,device', 'ecs_filesystem_free_bytes{$INSTANCE}', NULL, NULL, 'GB', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('142', '10', '磁盘已用存储量', 'bms_disk_used', 'instance,device', 'ecs_filesystem_size_bytes{$INSTANCE} - ecs_filesystem_free_bytes', NULL, NULL, 'GB', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('143', '10', '磁盘存储总量', 'bms_filesystem_size_bytes', 'instance,device', 'ecs_filesystem_size_bytes{$INSTANCE}', NULL, NULL, 'GB', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('144', '11', 'CPU使用率', 'bms_cpu_usage', 'instance', '100 - (100 * (sum by(instance) (irate(ecs_cpu_seconds_total{mode=\"idle\",$INSTANCE}[3m])) / sum by(instance) (irate(ecs_cpu_seconds_total{$INSTANCE}[3m]))))', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('145', '11', 'CPU1分钟平均负载', 'bms_load1', 'instance', 'ecs_load1{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('146', '11', 'CPU5分钟平均负载', 'bms_load5', 'instance', 'ecs_load5{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('147', '11', 'CPU15分钟平均负载', 'bms_load15', 'instance', 'ecs_load15{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, '{{ne .OSTYPE \"windows\"}}', NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('148', '11', '内存使用量', 'bms_memory_used', 'instance', 'ecs_memory_MemTotal_bytes{$INSTANCE} - ecs_memory_MemFree_bytes', NULL, NULL, 'Byte', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('149', '11', '内存使用率', 'bms_memory_usage', 'instance', '100 * ((ecs_memory_MemTotal_bytes{$INSTANCE} - ecs_memory_MemFree_bytes) / ecs_memory_MemTotal_bytes)', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('150', '11', '磁盘使用率', 'bms_disk_usage', 'instance,device', '100 * ((ecs_filesystem_size_bytes{$INSTANCE} - ecs_filesystem_free_bytes) / ecs_filesystem_size_bytes)', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('151', '11', '磁盘读速率', 'bms_disk_read_rate', 'instance,device', 'irate(ecs_disk_read_bytes_total{$INSTANCE}[3m])', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('152', '11', '磁盘写速率', 'bms_disk_write_rate', 'instance,device', 'irate(ecs_disk_written_bytes_total{$INSTANCE}[3m])', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('153', '11', '磁盘读IOPS', 'bms_disk_read_iops', 'instance,device', 'irate(ecs_disk_reads_completed_total{$INSTANCE}[3m])', NULL, NULL, '次', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('154', '11', '磁盘写IOPS', 'bms_disk_write_iops', 'instance,device', 'irate(ecs_disk_writes_completed_total{$INSTANCE}[3m])', NULL, NULL, '次', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('155', '11', '流入带宽', 'bms_network_transmit_rate', 'instance,device', 'irate(ecs_network_receive_bytes_total{$INSTANCE}[3m]) / 1024 / 1024 * 8', NULL, NULL, 'Mbps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('156', '11', '流出带宽', 'bms_network_receive_rate', 'instance,device', 'irate(ecs_network_transmit_bytes_total{$INSTANCE}[3m]) / 1024 / 1024 * 8', NULL, NULL, 'Mbps', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('157', '11', '包接收速率', 'bms_network_receive_packets_rate', 'instance,device', 'irate(ecs_network_receive_packets_total{$INSTANCE}[3m])', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('158', '11', '包发送速率', 'bms_network_transmit_packets_rate', 'instance,device', 'irate(ecs_network_transmit_packets_total{$INSTANCE}[3m])', NULL, NULL, '个/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('159', '11', '磁盘剩余存储量', 'bms_filesystem_free_bytes', 'instance,device', 'ecs_filesystem_free_bytes{$INSTANCE}', NULL, NULL, 'GB', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('160', '11', '磁盘已用存储量', 'bms_disk_used', 'instance,device', 'ecs_filesystem_size_bytes{$INSTANCE} - ecs_filesystem_free_bytes', NULL, NULL, 'GB', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('161', '11', '磁盘存储总量', 'bms_filesystem_size_bytes', 'instance,device', 'ecs_filesystem_size_bytes{$INSTANCE}', NULL, NULL, 'GB', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('162', '12', 'CPU使用率', 'redis_cpu_usage', 'instance', 'redis_cpu_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('163', '12', '内存使用率', 'redis_mem_usage', 'instance', 'redis_mem_usage{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('164', '12', '客户端连接数', 'redis_connected_clients', 'instance', 'redis_connected_clients{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('165', '12', 'TPS', 'redis_tps', 'instance', 'redis_tps{$INSTANCE}', NULL, NULL, NULL, NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
--
-- INSERT INTO `t_monitor_item` VALUES ('166', '13', 'mongos客户端当前连接数', 'mongo_mongos_current_connections', 'instance,pod', 'mongo_mongos_current_connections{$INSTANCE}', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('167', '13', 'shard客户端当前连接数', 'mongo_shard_current_connections', 'instance,pod', 'mongo_shard_current_connections{$INSTANCE}', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('168', '13', 'configServer客户端当前连接数', 'mongo_config_current_connections', 'instance,pod', 'mongo_config_current_connections{$INSTANCE}', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('169', '13', 'mongo各角色总当前连接数', 'mongo_total_current_connections', 'instance', 'mongo_total_current_connections{$INSTANCE}', NULL, NULL, '个', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('170', '13', '每个mongos的内存使用率', 'mongo_mongos_memory_ratio', 'instance,pod', 'mongo_mongos_memory_ratio{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('171', '13', 'configServer的内存使用率', 'mongo_config_memory_ratio', 'instance', 'mongo_config_memory_ratio{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('172', '13', '每个分片的内存使用率', 'mongo_shard_memory_ratio', 'instance,pod', 'mongo_shard_memory_ratio{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('173', '13', '每个mongos的CPU使用率', 'mongo_mongos_cpu_ratio', 'instance,pod', 'mongo_mongos_cpu_ratio{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('174', '13', '每个shard的CPU使用率', 'mongo_shard_cpu_ratio', 'instance,pod', 'mongo_shard_cpu_ratio{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('175', '13', 'configServer的CPU使用率', 'mongo_config_cpu_ratio', 'instance', 'mongo_config_cpu_ratio{$INSTANCE}', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
--
-- INSERT INTO `t_monitor_item` VALUES ('176', '14', 'QPS', 'guard_nginx_http_current_reqs', 'instance,route', 'sum by(instance,service,route)(rate(guard_nginx_http_current_reqs{$INSTANCE}[3m]))', NULL, NULL, '次/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('177', '14', 'P90接口响应延时', 'guard_http_latency_bucket_api_p90', 'instance,route', 'histogram_quantile(0.90, sum by(instance,service,route,le)(rate(guard_http_latency_bucket{type=\"request\",$INSTANCE}[3m])))', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('178', '14', 'P95接口响应延时', 'guard_http_latency_bucket_api_p95', 'instance,route', 'histogram_quantile(0.95, sum by(instance,service,route,le)(rate(guard_http_latency_bucket{type=\"request\",$INSTANCE}[3m])))', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('179', '14', 'P99接口响应延时', 'guard_http_latency_bucket_api_p99', 'instance,route', 'histogram_quantile(0.99, sum by(instance,service,route,le)(rate(guard_http_latency_bucket{type=\"request\",$INSTANCE}[3m])))', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('180', '14', 'P90服务响应延时', 'guard_http_latency_bucket_service_p90', 'instance,service', 'histogram_quantile(0.90, sum by(instance,service,le)(rate(guard_http_latency_bucket{type=\"request\",$INSTANCE}[3m])))', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('181', '14', 'P95服务响应延时', 'guard_http_latency_bucket_service_p95', 'instance,service', 'histogram_quantile(0.95, sum by(instance,service,le)(rate(guard_http_latency_bucket{type=\"request\",$INSTANCE}[3m])))', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('182', '14', 'P99服务响应延时', 'guard_http_latency_bucket_service_p99', 'instance,service', 'histogram_quantile(0.99, sum by(instance,service,le)(rate(guard_http_latency_bucket{type=\"request\",$INSTANCE}[3m])))', NULL, NULL, 'ms', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('183', '14', '接口成功率', 'guard_nginx_url_request_succ_api', 'instance,route', 'sum by(instance,service,route)(rate(guard_nginx_url_request_succ{code=\"200\",$INSTANCE}[3m]))/sum by(instance,service,route)(rate(guard_nginx_url_request_succ{code=\"total\",$INSTANCE}[3m]))', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('184', '14', '服务成功率', 'guard_nginx_url_request_succ_service', 'instance,service', 'sum by(instance,service)(rate(guard_nginx_url_request_succ{code=\"200\",$INSTANCE}[3m]))/sum by(instance,service)(rate(guard_nginx_url_request_succ{code=\"total\",$INSTANCE}[3m]))', NULL, NULL, '%', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart,rule');
-- INSERT INTO `t_monitor_item` VALUES ('185', '14', '接口入口带宽监控', 'guard_bandwidth_api_ingress', 'instance,route', 'sum by(instance,service,route)(rate(guard_bandwidth{type=\"ingress\",$INSTANCE}[3m]))', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('186', '14', '接口出口带宽监控', 'guard_bandwidth_api_egress', 'instance,route', 'sum by(instance,service,route)(rate(guard_bandwidth{type=\"egress\",$INSTANCE}[3m]))', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('187', '14', '服务入口带宽监控', 'guard_bandwidth_service_ingress', 'instance,service', 'sum by(instance,service)(irate(guard_bandwidth{type=\"ingress\",$INSTANCE}[3m])) ', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('188', '14', '服务出口带宽监控', 'guard_bandwidth_service_egress', 'instance,service', 'sum by(instance,service)(rate(guard_bandwidth{type=\"egress\",$INSTANCE}[3m]))', NULL, NULL, 'Byte/s', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('189', '14', '接口调用量', 'guard_http_apirequests', 'instance,service,route', 'guard_http_apirequests{$INSTANCE}', NULL, NULL, '次', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');
-- INSERT INTO `t_monitor_item` VALUES ('190', '14', '服务调用量', 'guard_http_servicerequests', 'instance,service', 'guard_http_apirequests{$INSTANCE}', NULL, NULL, '次', NULL, NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, 'chart');

DROP TABLE IF EXISTS `t_config_item`;
CREATE TABLE `t_config_item`
(
    `biz_id`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '-1' COMMENT '业务Id',
    `p_biz_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '-1' COMMENT '上级业务Id',
    `name`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '配置名称',
    `code`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '配置编码',
    `data`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '配置值',
    `sort_id`  int NULL DEFAULT 0 COMMENT '排序',
    `remark`   varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
    `id`       bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `t_config_item` VALUES ('1', '-1', '统计周期', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('10', '2', '持续3个周期', '3', '3', 1, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('11', '2', '持续5个周期', '5', '5', 2, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('12', '3', '平均值', 'Average', 'avg', 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('13', '3', '最大值', 'Maximum', 'max', 1, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('14', '3', '最小值', 'Minimum', 'min', 2, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('15', '4', '大于', 'greater', '>', 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('16', '4', '大于等于', 'greaterOrEqual', '>=', 1, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('17', '4', '小于', 'less', '<', 2, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('18', '4', '小于等于', 'lessOrEqual', '<=', 3, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('19', '4', '等于', 'equal', '==', 4, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('2', '-1', '持续周期', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('20', '4', '不等于', 'notEqual', '!=', 5, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('21', '-1', '概览监控项', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('22', '21', 'CPU使用率（操作系统）', NULL, 'ecs_cpu_usage', 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('23', '21', '内存使用率（操作系统）', NULL, 'ecs_memory_usage', 1, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('28', '-1', '监控周期', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('29', '28', '紧急', '1', 'MAIN', 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('3', '-1', '统计方式', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('30', '28', '重要', '2', 'MARJOR', 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('31', '28', '次要', '3', 'MINOR', 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('32', '28', '提醒', '4', 'WARN', 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('4', '-1', '对比方式', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('5', '-1', '监控数据', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('51', '5', '0-3H', '0,3', '60', 1, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('52', '5', '3H-12H', '3,12', '180', 2, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('53', '5', '12H-3D', '12,72', '900', 3, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('54', '5', '3D-10D', '72,240', '2700', 4, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('55', '-1', '监控周期', NULL, NULL, 0, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('6', '1', '5分钟', '300', '5m', 1, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('7', '1', '15分钟', '900', '15m', 2, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('8', '1', '30分钟', '1800', '30m', 3, NULL, NULL);
INSERT INTO `t_config_item` VALUES ('9', '2', '持续1个周期', '1', '1', 0, NULL, NULL);
