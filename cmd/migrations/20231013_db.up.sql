UPDATE t_monitor_item SET metrics_linux = 'sum by(instance,instanceType)(ecs_memory_MemTotal_bytes{$INSTANCE}) - sum by(instance,instanceType)(ecs_memory_MemAvailable_bytes{$INSTANCE})' WHERE metric_name = 'ecs_memory_used';
UPDATE t_monitor_item SET metrics_linux = '100 * ((sum by(instance,instanceType)(ecs_memory_MemTotal_bytes{$INSTANCE}) - (sum by(instance,instanceType)(ecs_memory_MemAvailable_bytes{$INSTANCE}))) / sum by(instance,instanceType)(ecs_memory_MemTotal_bytes{$INSTANCE}))' WHERE metric_name = 'ecs_memory_usage';
