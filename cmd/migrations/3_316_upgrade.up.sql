UPDATE t_monitor_item SET metrics_linux = 'ecs_processes_top5Cpus{$INSTANCE}', labels = 'instance,pid' WHERE biz_id = '191';
UPDATE t_monitor_item SET metrics_linux = 'ecs_processes_top5Mems{$INSTANCE}', labels = 'instance,pid' WHERE biz_id = '192';
UPDATE t_monitor_item SET metrics_linux = 'ecs_procs_running{$INSTANCE}' WHERE biz_id = '193';
UPDATE t_monitor_item SET metrics_linux = 'ecs_processes_top5Fds{$INSTANCE}', labels = 'instance,pid' WHERE biz_id = '194';
