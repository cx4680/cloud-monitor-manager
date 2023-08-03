UPDATE t_monitor_product SET status = '1' WHERE abbreviation IN ('cbr', 'kafka', 'bms', 'ebms', 'redis', 'mongo', 'cgw', 'mysql', 'dm', 'postgresql');

INSERT INTO t_monitor_item (metric_name, product_metric_name, name, metric_name, labels, metrics_linux, metrics_windows, statistics, unit, frequency, type, is_display, status, description, create_user, create_time, show_expression, display) VALUES ('197', '1', '(基础)内存使用率', 'ecs_memory_base_usage', 'instance', '(1-sum by(instance,instanceType)(ecs_base_memory_unused_bytes{$INSTANCE})/sum by(instance,instanceType)(ecs_base_memory_available_bytes{$INSTANCE}))*100', null, null, '%', null, '1', '1', '1', null, null, NOW(), null, 'chart');

UPDATE t_monitor_item SET metrics_linux='100 - (100 * (sum by(instance,instanceType)(irate(ecs_cpu_seconds_total{mode="idle",$INSTANCE}[3m])) / sum by(instance,instanceType)(irate(ecs_cpu_seconds_total{$INSTANCE}[3m]))))' WHERE metric_name='ecs_cpu_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load1{$INSTANCE})' WHERE metric_name='ecs_load1';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load5{$INSTANCE})' WHERE metric_name='ecs_load5';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load15{$INSTANCE})' WHERE metric_name='ecs_load15';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_memory_MemTotal_bytes{$INSTANCE}) - sum by(instance,instanceType)(ecs_memory_MemFree_bytes{$INSTANCE})' WHERE metric_name='ecs_memory_used';
UPDATE t_monitor_item SET metrics_linux = '100 * ((sum by(instance)(ecs_memory_MemTotal_bytes{$INSTANCE}) - (sum by(instance)(ecs_memory_MemFree_bytes{$INSTANCE}) + sum by(instance)(ecs_memory_Cached_bytes{$INSTANCE}))) / sum by(instance)(ecs_memory_MemTotal_bytes{$INSTANCE}))' WHERE metric_name = 'ecs_memory_usage';
UPDATE t_monitor_item SET metrics_linux='100 * ((sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}) - sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})) / sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}))' WHERE metric_name='ecs_disk_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_read_bytes_total{$INSTANCE}[3m]))' WHERE metric_name='ecs_disk_read_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_written_bytes_total{$INSTANCE}[3m]))' WHERE metric_name='ecs_disk_write_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_reads_completed_total{$INSTANCE}[3m]))' WHERE metric_name='ecs_disk_read_iops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_writes_completed_total{$INSTANCE}[3m]))' WHERE metric_name='ecs_disk_write_iops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_receive_bytes_total{$INSTANCE}[3m])) / 1024 / 1024 * 8' WHERE metric_name='ecs_network_receive_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_transmit_bytes_total{$INSTANCE}[3m])) / 1024 / 1024 * 8' WHERE metric_name='ecs_network_transmit_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_receive_packets_total{$INSTANCE}[3m]))' WHERE metric_name='ecs_network_receive_packets_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_transmit_packets_total{$INSTANCE}[3m]))' WHERE metric_name='ecs_network_transmit_packets_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})' WHERE metric_name='ecs_filesystem_free_bytes';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}) - sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})' WHERE metric_name='ecs_disk_used';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE})' WHERE metric_name='ecs_filesystem_size_bytes';
UPDATE t_monitor_item SET metrics_linux='100 * avg by(instance,instanceType)(irate(ecs_base_vcpu_seconds{$INSTANCE}[6m]))' WHERE metric_name='ecs_cpu_base_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,drive)(irate(ecs_base_storage_traffic_bytes_total{type="read",$INSTANCE}[6m]))' WHERE metric_name='ecs_disk_base_read_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,drive)(irate(ecs_base_storage_traffic_bytes_total{type="write",$INSTANCE}[6m]))' WHERE metric_name='ecs_disk_base_write_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,interface)(irate(ecs_base_network_traffic_bytes_total{type="rx",$INSTANCE}[6m]))' WHERE metric_name='ecs_network_base_receive_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,interface)(irate(ecs_base_network_traffic_bytes_total{type="tx",$INSTANCE}[6m]))' WHERE metric_name='ecs_network_base_transmit_rate';
UPDATE t_monitor_item SET metrics_linux='sum(irate(ecs_base_storage_iops_total{type="read",$INSTANCE}[15m])) by(instance,instanceType,drive)' WHERE metric_name='ecs_disk_base_read_iops';
UPDATE t_monitor_item SET metrics_linux='sum(irate(ecs_base_storage_iops_total{type="write",$INSTANCE}[15m])) by(instance,instanceType,drive)' WHERE metric_name='ecs_disk_base_write_iops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,cmd_line)(ecs_processes_top5Cpus{cmd_line!="",$INSTANCE})' WHERE metric_name='ecs_processes_top5Cpus';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,cmd_line)(ecs_processes_top5Mems{cmd_line!="",$INSTANCE})' WHERE metric_name='ecs_processes_top5Mems';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_procs_running{$INSTANCE})' WHERE metric_name='ecs_procs_running';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,cmd_line)(ecs_processes_top5Fds{cmd_line!="",$INSTANCE})' WHERE metric_name='ecs_processes_top5Fds';

UPDATE t_monitor_item SET metrics_linux='sum(eip_upstream_bits_rate{$INSTANCE}) by(instance,instanceType,eip)' WHERE metric_name='eip_upstream_bandwidth';
UPDATE t_monitor_item SET metrics_linux='sum(eip_downstream_bits_rate{$INSTANCE}) by(instance,instanceType,eip)' WHERE metric_name='eip_downstream_bandwidth';
UPDATE t_monitor_item SET metrics_linux='((sum(eip_upstream_bits_rate{$INSTANCE}) by(instance,instanceType,eip))/8)*60' WHERE metric_name='eip_upstream';
UPDATE t_monitor_item SET metrics_linux='((sum(eip_downstream_bits_rate{$INSTANCE}) by(instance,instanceType,eip))/8)*60' WHERE metric_name='eip_downstream';
UPDATE t_monitor_item SET metrics_linux='(sum(eip_upstream_bits_rate{$INSTANCE}) by(instance,instanceType,eip) / avg(eip_config_upstream_bandwidth{$INSTANCE}) by(instance,instanceType,eip)) * 100' WHERE metric_name='eip_upstream_bandwidth_usage';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_http_bps_out_rate{$INSTANCE})' WHERE metric_name='slb_out_bandwidth';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_http_bps_in_rate{$INSTANCE})' WHERE metric_name='slb_in_bandwidth';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_all_connection_count{$INSTANCE})' WHERE metric_name='slb_max_connection';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_all_est_connection_count{$INSTANCE})' WHERE metric_name='slb_active_connection';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_all_none_est_connection_count{$INSTANCE})' WHERE metric_name='slb_inactive_connection';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_new_connection_rate{$INSTANCE})' WHERE metric_name='slb_new_connection';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_drop_connection_rate{$INSTANCE})' WHERE metric_name='slb_drop_connection';
UPDATE t_monitor_item SET metrics_linux='avg by(instance,instanceType)(Slb_unhealthy_server_count{$INSTANCE})' WHERE metric_name='3slb_unhealthyserver';
UPDATE t_monitor_item SET metrics_linux='avg by(instance,instanceType)(Slb_healthy_server_count{$INSTANCE})' WHERE metric_name='slb_healthyserver';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_request_rate{$INSTANCE})' WHERE metric_name='slb_qps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_http_2xx_rate{$INSTANCE})' WHERE metric_name='slb_statuscode2xx';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_http_3xx_rate{$INSTANCE})' WHERE metric_name='slb_statuscode3xx';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_http_4xx_rate{$INSTANCE})' WHERE metric_name='slb_statuscode4xx';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Slb_http_5xx_rate{$INSTANCE})' WHERE metric_name='slb_statuscode5xx';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(cbr_vault_size{$INSTANCE})' WHERE metric_name='cbr_vault_size';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(cbr_vault_used{$INSTANCE})' WHERE metric_name='cbr_vault_used';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(cbr_vault_used{$INSTANCE}) / sum by(instance,instanceType)(cbr_vault_size{$INSTANCE}) * 100' WHERE metric_name='cbr_vault_usage_rate';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Nat_snat_total_connection_count{$INSTANCE})' WHERE metric_name='snat_connection';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(Nat_recv_bytes_total_count{$INSTANCE}[1m])*8)' WHERE metric_name='inbound_bandwidth';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(Nat_send_bytes_total_count{$INSTANCE}[1m])*8)' WHERE metric_name='outbound_bandwidth';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Nat_recv_bytes_total_count{$INSTANCE})' WHERE metric_name='inbound_traffic';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Nat_send_bytes_total_count{$INSTANCE})' WHERE metric_name='outbound_traffic';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(Nat_recv_packets_total_count{$INSTANCE}[1m]))' WHERE metric_name='inbound_pps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(Nat_send_packets_total_count{$INSTANCE}[1m]))' WHERE metric_name='outbound_pps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(Nat_snat_total_connection_count{$INSTANCE}) / avg by (instance,instanceType)(Nat_nat_max_connection_count{$INSTANCE}) *100' WHERE metric_name='snat_connection_ratio';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_slave_io{$INSTANCE})' WHERE metric_name='mysql_slave_io';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_slave_sql{$INSTANCE})' WHERE metric_name='mysql_slave_sql';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_slave_seconds_behind_master{$INSTANCE})' WHERE metric_name='mysql_slave_seconds_behind_master';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_active_connections{$INSTANCE})' WHERE metric_name='mysql_active_connections';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_current_connection_percent{$INSTANCE})' WHERE metric_name='mysql_current_connection_percent';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_qps{$INSTANCE})' WHERE metric_name='mysql_qps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_tps{$INSTANCE})' WHERE metric_name='mysql_tps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_select_ps{$INSTANCE})' WHERE metric_name='mysql_select_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_update_ps{$INSTANCE})' WHERE metric_name='mysql_update_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_insert_ps{$INSTANCE})' WHERE metric_name='mysql_insert_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_delete_ps{$INSTANCE})' WHERE metric_name='mysql_delete_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_cpu_usage{$INSTANCE})' WHERE metric_name='mysql_cpu_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_mem_usage{$INSTANCE})' WHERE metric_name='mysql_mem_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_disk_usage{$INSTANCE})' WHERE metric_name='mysql_disk_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_select_ps{$INSTANCE})' WHERE metric_name='mysql_innodb_select_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_update_ps{$INSTANCE})' WHERE metric_name='mysql_innodb_update_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_insert_ps{$INSTANCE})' WHERE metric_name='mysql_innodb_insert_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_delete_ps{$INSTANCE})' WHERE metric_name='mysql_innodb_delete_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_cache_hit_rate{$INSTANCE})' WHERE metric_name='mysql_innodb_cache_hit_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_reads_ps{$INSTANCE})' WHERE metric_name='mysql_innodb_reads_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_writes_ps{$INSTANCE})' WHERE metric_name='mysql_innodb_writes_ps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_buffer_pool_pages_dirty{$INSTANCE})' WHERE metric_name='mysql_innodb_buffer_pool_pages_dirty';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_buffer_pool_bytes_dirty{$INSTANCE})' WHERE metric_name='mysql_innodb_buffer_pool_bytes_dirty';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_innodb_log_waits{$INSTANCE})' WHERE metric_name='mysql_innodb_log_waits';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_binlog_cache_disk_use{$INSTANCE})' WHERE metric_name='mysql_binlog_cache_disk_use';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_slow_queries_per_min{$INSTANCE})' WHERE metric_name='mysql_slow_queries_per_min';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_long_query_count{$INSTANCE})' WHERE metric_name='mysql_long_query_count';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_long_query_alert_count{$INSTANCE})' WHERE metric_name='mysql_long_query_alert_count';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_exec_statememt_frequency{$INSTANCE})' WHERE metric_name='mysql_exec_statememt_frequency';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_read_frequency{$INSTANCE})' WHERE metric_name='mysql_read_frequency';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_write_frequency{$INSTANCE})' WHERE metric_name='mysql_write_frequency';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_top_statememt_avg_exec_time{$INSTANCE})' WHERE metric_name='mysql_top_statememt_avg_exec_time';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_top_statememt_exec_err_rate{$INSTANCE})' WHERE metric_name='mysql_top_statememt_exec_err_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mysql_current_cons_num{$INSTANCE})' WHERE metric_name='mysql_current_cons_num';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_tps{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_tps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_qps{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_qps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_ips{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_ips';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_dps{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_dps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_ups{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_ups';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_ddlps{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_ddlps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_nioips{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_nioips';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_nio_ops{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_nio_ops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_fio_ips{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_fio_ips';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(rate(dm_global_status_fio_ops{$INSTANCE}[1m]))' WHERE metric_name='dm_global_status_fio_ops';
UPDATE t_monitor_item SET metrics_linux='avg by(instance,instanceType)(dm_global_status_mem_used{$INSTANCE})' WHERE metric_name='dm_global_status_mem_used';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_cpu_use_rate{$INSTANCE})' WHERE metric_name='dm_global_status_cpu_use_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_mem_use_rate{$INSTANCE})' WHERE metric_name='dm_global_status_mem_use_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_sessions{$INSTANCE})' WHERE metric_name='dm_global_status_sessions';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_active_sessions{$INSTANCE})' WHERE metric_name='dm_global_status_active_sessions';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_task_waiting{$INSTANCE})' WHERE metric_name='dm_global_status_task_waiting';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_task_ready{$INSTANCE})' WHERE metric_name='dm_global_status_task_ready';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_task_total_wait_time{$INSTANCE})' WHERE metric_name='dm_global_status_task_total_wait_time';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_avg_wait_time{$INSTANCE})' WHERE metric_name='dm_global_status_avg_wait_time';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(dm_global_status_threads{$INSTANCE})' WHERE metric_name='dm_global_status_threads';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_cpu_usage{$INSTANCE})' WHERE metric_name='pg_cpu_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_mem_usage{$INSTANCE})' WHERE metric_name='pg_mem_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_disk_usage{$INSTANCE})' WHERE metric_name='pg_disk_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_qps{$INSTANCE})' WHERE metric_name='pg_qps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_rqps{$INSTANCE})' WHERE metric_name='pg_rqps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_wqps{$INSTANCE})' WHERE metric_name='pg_wqps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_tps{$INSTANCE})' WHERE metric_name='pg_tps';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_mean_exec_time{$INSTANCE})' WHERE metric_name='pg_mean_exec_time';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_open_ct_num{$INSTANCE})' WHERE metric_name='pg_open_ct_num';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(pg_active_ct_num{$INSTANCE})' WHERE metric_name='pg_active_ct_num';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(kafka_brokers{$INSTANCE})' WHERE metric_name='kafka_brokers';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(kafka_server_replicamanager_partitioncount{$INSTANCE})' WHERE metric_name='kafka_server_replicamanager_partitioncount';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(kafka_server_brokertopicmetrics_bytesinpersec{$INSTANCE})' WHERE metric_name='kafka_server_brokertopicmetrics_bytesinpersec';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(kafka_server_brokertopicmetrics_bytesoutpersec{$INSTANCE})' WHERE metric_name='kafka_server_brokertopicmetrics_bytesoutpersec';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(kafka_server_brokertopicmetrics_messagesinpersec{$INSTANCE})' WHERE metric_name='kafka_server_brokertopicmetrics_messagesinpersec';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(kafka_consumergroup_lag{$INSTANCE})' WHERE metric_name='kafka_consumergroup_lag';
INSERT INTO t_monitor_item (biz_id, product_biz_id, name, metric_name, labels, metrics_linux, metrics_windows, statistics, unit, frequency, type, is_display, status, description, create_user, create_time, show_expression, display) VALUES ('201', '9', '集群磁盘使用量', 'kafka_logdirsusage_partition_usage', 'instance', 'sum by(instance,instanceType)(kafka_logdirsusage_partition_usage{$INSTANCE})', null, null, 'Byte', null, null, '1', '1', null, null, null, null, 'chart');

UPDATE t_monitor_item SET metrics_linux='100 - (100 * (sum by(instance,instanceType)(irate(ecs_cpu_seconds_total{mode="idle",$INSTANCE}[3m])) / sum by(instance,instanceType)(irate(ecs_cpu_seconds_total{$INSTANCE}[3m]))))' WHERE metric_name='bms_cpu_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load1{$INSTANCE})' WHERE metric_name='bms_load1';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load5{$INSTANCE})' WHERE metric_name='bms_load5';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load15{$INSTANCE})' WHERE metric_name='bms_load15';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_memory_MemTotal_bytes{$INSTANCE}) - sum by(instance,instanceType)(ecs_memory_MemFree_bytes{$INSTANCE})' WHERE metric_name='bms_memory_used';
UPDATE t_monitor_item SET metrics_linux = '100 * ((sum by(instance)(ecs_memory_MemTotal_bytes{$INSTANCE}) - (sum by(instance)(ecs_memory_MemFree_bytes{$INSTANCE}) + sum by(instance)(ecs_memory_Cached_bytes{$INSTANCE}))) / sum by(instance)(ecs_memory_MemTotal_bytes{$INSTANCE}))' WHERE metric_name = 'bms_memory_usage';
UPDATE t_monitor_item SET metrics_linux='100 * ((sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}) - sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})) / sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}))' WHERE metric_name='bms_disk_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_read_bytes_total{$INSTANCE}[3m]))' WHERE metric_name='bms_disk_read_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_written_bytes_total{$INSTANCE}[3m]))' WHERE metric_name='bms_disk_write_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_reads_completed_total{$INSTANCE}[3m]))' WHERE metric_name='bms_disk_read_iops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_writes_completed_total{$INSTANCE}[3m]))' WHERE metric_name='bms_disk_write_iops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_receive_bytes_total{$INSTANCE}[3m])) / 1024 / 1024 * 8' WHERE metric_name='bms_network_receive_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_transmit_bytes_total{$INSTANCE}[3m])) / 1024 / 1024 * 8' WHERE metric_name='bms_network_transmit_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_receive_packets_total{$INSTANCE}[3m]))' WHERE metric_name='bms_network_receive_packets_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_transmit_packets_total{$INSTANCE}[3m]))' WHERE metric_name='bms_network_transmit_packets_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})' WHERE metric_name='bms_filesystem_free_bytes';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}) - sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})' WHERE metric_name='bms_disk_used';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE})' WHERE metric_name='bms_filesystem_size_bytes';

UPDATE t_monitor_item SET metrics_linux='100 - (100 * (sum by(instance,instanceType)(irate(ecs_cpu_seconds_total{mode="idle",$INSTANCE}[3m])) / sum by(instance,instanceType)(irate(ecs_cpu_seconds_total{$INSTANCE}[3m]))))' WHERE metric_name='ebms_cpu_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load1{$INSTANCE})' WHERE metric_name='ebms_load1';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load5{$INSTANCE})' WHERE metric_name='ebms_load5';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_load15{$INSTANCE})' WHERE metric_name='ebms_load15';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(ecs_memory_MemTotal_bytes{$INSTANCE}) - sum by(instance,instanceType)(ecs_memory_MemFree_bytes{$INSTANCE})' WHERE metric_name='ebms_memory_used';
UPDATE t_monitor_item SET metrics_linux = '100 * ((sum by(instance)(ecs_memory_MemTotal_bytes{$INSTANCE}) - (sum by(instance)(ecs_memory_MemFree_bytes{$INSTANCE}) + sum by(instance)(ecs_memory_Cached_bytes{$INSTANCE}))) / sum by(instance)(ecs_memory_MemTotal_bytes{$INSTANCE}))' WHERE metric_name = 'ebms_memory_usage';
UPDATE t_monitor_item SET metrics_linux='100 * ((sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}) - sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})) / sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}))' WHERE metric_name='ebms_disk_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_read_bytes_total{$INSTANCE}[3m]))' WHERE metric_name='ebms_disk_read_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_written_bytes_total{$INSTANCE}[3m]))' WHERE metric_name='ebms_disk_write_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_reads_completed_total{$INSTANCE}[3m]))' WHERE metric_name='ebms_disk_read_iops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_disk_writes_completed_total{$INSTANCE}[3m]))' WHERE metric_name='ebms_disk_write_iops';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_receive_bytes_total{$INSTANCE}[3m])) / 1024 / 1024 * 8' WHERE metric_name='ebms_network_receive_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_transmit_bytes_total{$INSTANCE}[3m])) / 1024 / 1024 * 8' WHERE metric_name='ebms_network_transmit_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_receive_packets_total{$INSTANCE}[3m]))' WHERE metric_name='ebms_network_receive_packets_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(irate(ecs_network_transmit_packets_total{$INSTANCE}[3m]))' WHERE metric_name='ebms_network_transmit_packets_rate';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})' WHERE metric_name='ebms_filesystem_free_bytes';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE}) - sum by(instance,instanceType,device)(ecs_filesystem_free_bytes{$INSTANCE})' WHERE metric_name='ebms_disk_used';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,device)(ecs_filesystem_size_bytes{$INSTANCE})' WHERE metric_name='ebms_filesystem_size_bytes';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(redis_cpu_usage{$INSTANCE})' WHERE metric_name='redis_cpu_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(redis_mem_usage{$INSTANCE})' WHERE metric_name='redis_mem_usage';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(redis_connected_clients{$INSTANCE})' WHERE metric_name = 'redis_connected_clients';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(redis_tps{$INSTANCE})' WHERE metric_name = 'redis_tps';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,pod)(mongo_mongos_current_connections{$INSTANCE})' WHERE metric_name='mongo_mongos_current_connections';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,pod)(mongo_shard_current_connections{$INSTANCE})' WHERE metric_name='mongo_shard_current_connections';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,pod)(mongo_config_current_connections{$INSTANCE})' WHERE metric_name='mongo_config_current_connections';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mongo_total_current_connections{$INSTANCE})' WHERE metric_name='mongo_total_current_connections';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,pod)(mongo_mongos_memory_ratio{$INSTANCE})' WHERE metric_name='mongo_mongos_memory_ratio';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mongo_config_memory_ratio{$INSTANCE})' WHERE metric_name='mongo_config_memory_ratio';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,pod)(mongo_shard_memory_ratio{$INSTANCE})' WHERE metric_name='mongo_shard_memory_ratio';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,pod)(mongo_mongos_cpu_ratio{$INSTANCE})' WHERE metric_name='mongo_mongos_cpu_ratio';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,pod)(mongo_shard_cpu_ratio{$INSTANCE})' WHERE metric_name='mongo_shard_cpu_ratio';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType)(mongo_config_cpu_ratio{$INSTANCE})' WHERE metric_name='mongo_config_cpu_ratio';

UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service,route)(rate(guard_nginx_http_current_reqs{$INSTANCE}[3m]))' WHERE metric_name='guard_nginx_http_current_reqs';
UPDATE t_monitor_item SET metrics_linux='histogram_quantile(0.90, sum by(instance,instanceType,service,route,le)(rate(guard_http_latency_bucket{type="request",$INSTANCE}[3m])))' WHERE metric_name='guard_http_latency_bucket_api_p90';
UPDATE t_monitor_item SET metrics_linux='histogram_quantile(0.95, sum by(instance,instanceType,service,route,le)(rate(guard_http_latency_bucket{type="request",$INSTANCE}[3m])))' WHERE metric_name='guard_http_latency_bucket_api_p95';
UPDATE t_monitor_item SET metrics_linux='histogram_quantile(0.99, sum by(instance,instanceType,service,route,le)(rate(guard_http_latency_bucket{type="request",$INSTANCE}[3m])))' WHERE metric_name='guard_http_latency_bucket_api_p99';
UPDATE t_monitor_item SET metrics_linux='histogram_quantile(0.90, sum by(instance,instanceType,service,le)(rate(guard_http_latency_bucket{type="request",$INSTANCE}[3m])))' WHERE metric_name='guard_http_latency_bucket_service_p90';
UPDATE t_monitor_item SET metrics_linux='histogram_quantile(0.95, sum by(instance,instanceType,service,le)(rate(guard_http_latency_bucket{type="request",$INSTANCE}[3m])))' WHERE metric_name='guard_http_latency_bucket_service_p95';
UPDATE t_monitor_item SET metrics_linux='histogram_quantile(0.99, sum by(instance,instanceType,service,le)(rate(guard_http_latency_bucket{type="request",$INSTANCE}[3m])))' WHERE metric_name='guard_http_latency_bucket_service_p99';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service,route)(rate(guard_nginx_url_request_succ{code="200",$INSTANCE}[3m]))/sum by(instance,instanceType,service,route)(rate(guard_nginx_url_request_succ{code="total",$INSTANCE}[3m]))*100' WHERE metric_name='guard_nginx_url_request_succ_api';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service)(rate(guard_nginx_url_request_succ{code="200",$INSTANCE}[3m]))/sum by(instance,instanceType,service)(rate(guard_nginx_url_request_succ{code="total",$INSTANCE}[3m]))*100' WHERE metric_name='guard_nginx_url_request_succ_service';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service,route)(rate(guard_bandwidth{type="ingress",$INSTANCE}[3m]))' WHERE metric_name='guard_bandwidth_api_ingress';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service,route)(rate(guard_bandwidth{type="egress",$INSTANCE}[3m]))' WHERE metric_name='guard_bandwidth_api_egress';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service)(irate(guard_bandwidth{type="ingress",$INSTANCE}[3m])) ' WHERE metric_name='guard_bandwidth_service_ingress';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service)(rate(guard_bandwidth{type="egress",$INSTANCE}[3m]))' WHERE metric_name='guard_bandwidth_service_egress';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service,route)(guard_http_apirequests{$INSTANCE})' WHERE metric_name='guard_http_apirequests';
UPDATE t_monitor_item SET metrics_linux='sum by(instance,instanceType,service)(guard_http_apirequests{$INSTANCE})' WHERE metric_name='guard_http_servicerequests';
