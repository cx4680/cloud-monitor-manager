UPDATE t_monitor_product SET status = '0' WHERE abbreviation IN ('cbr', 'kafka', 'redis', 'mongo', 'cgw', 'mysql', 'dm', 'postgresql');
DELETE FROM t_monitor_item WHERE metric_name IN ('ecs_memory_base_usage','ecs_base_gpu_seconds','ecs_base_gpu_memory_total','ecs_base_gpu_memory_usage','kafka_logdirsusage_partition_usage');
