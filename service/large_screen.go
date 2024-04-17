package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type LargeScreenService struct{}

func NewLargeScreenService() *LargeScreenService {
	return &LargeScreenService{}
}

func (s *LargeScreenService) Tags() ([]*form.LargeScreenResourceTag, error) {
	var response *form.LargeScreenResourceTagResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.ResourceTagApi + "/list-all")
	if err != nil || response == nil || response.Module == nil {
		return nil, err
	}
	return response.Module.List, nil
}

func (s *LargeScreenService) ResourceOverview(tag string) (*form.LargeScreenResourceOverview, error) {
	var resourceOverview = &form.LargeScreenResourceOverview{}
	resources, _, err := s.getResourcesByTag(tag, "", "")
	if err != nil {
		return nil, err
	}
	for _, v := range resources {
		if v.CloudProductCode == constant.CloudProductCodeEcs && v.ResourceTypeCode == constant.ResourceTypeCodeInstance {
			resourceOverview.Ecs.Total++
			if v.StatusDesc == constant.StatusDescActive {
				resourceOverview.Ecs.Normal++
			}
		}
		if v.CloudProductCode == constant.CloudProductCodeEip && v.ResourceTypeCode == constant.ResourceTypeCodeInstance {
			resourceOverview.Eip.Total++
			var additional = &form.EipAdditional{}
			jsonutil.ToObject(v.Additional, additional)
			resourceOverview.Eip.Bandwidth += additional.BandWidth.BandWidthSize
		}
		if v.CloudProductCode == constant.CloudProductCodeRdb {
			switch v.ResourceTypeCode {
			case constant.ResourceTypeCodeMysql:
				resourceOverview.Rdb.Mysql++
			case constant.ResourceTypeCodeDm:
				resourceOverview.Rdb.Dm++
			case constant.ResourceTypeCodePg:
				resourceOverview.Rdb.Pg++
			}
		}
		if v.CloudProductCode == constant.CloudProductCodeSlb && v.ResourceTypeCode == constant.ResourceTypeCodeInstance {
			resourceOverview.Slb.Total++
			if v.StatusDesc != constant.StatusDescFailed {
				resourceOverview.Slb.Normal++
			}
		}
		if v.CloudProductCode == constant.CloudProductCodeNat && v.ResourceTypeCode == constant.ResourceTypeCodeInstance {
			resourceOverview.Nat.Total++
			if v.StatusDesc == constant.StatusDescRunning {
				resourceOverview.Nat.Normal++
			}
		}
	}
	return resourceOverview, nil
}

func (s *LargeScreenService) ResourceAlert(tag string) (*form.YyLargeScreen, error) {
	var response = &form.YyLargeScreen{}
	_, resources, err := s.getResourcesByTag(tag, "", "")
	if err != nil {
		return nil, err
	}
	var param = map[string][]string{"resources": resources}
	_, err = httputil.GetHttpClient().R().SetBody(param).SetResult(&response).Post(config.Cfg.Common.CcosHawkeyeApi + "/yyLargeScreen/alerts")
	if err != nil {
		logger.Logger().Error("ResourceAlert error: ", err)
		return nil, err
	}
	return response, nil
}

func (s *LargeScreenService) ResourceEcs(tag string) (*form.LargeScreenEcs, error) {
	var largeScreenEcs = &form.LargeScreenEcs{}
	largeScreenEcs.Cpu.Unit = "core"
	largeScreenEcs.Memory.Unit = "GiB"
	largeScreenEcs.Disk.Unit = "GiB"

	_, resourceList, err := s.getResourcesByTag(tag, constant.CloudProductCodeEcs, constant.ResourceTypeCodeInstance)
	if err != nil {
		return nil, err
	}
	resources := strings.Join(resourceList, "|")

	//ECS vCpu
	vCpuPql := fmt.Sprintf("count(ecs_base_vcpu_seconds{%v})", constant.INSTANCE+"=~'"+resources+"'")
	vCpuUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(vCpuPql)
	vCpuResponse := &form.PrometheusResponse{}
	_, err = httputil.GetHttpClient().R().SetResult(&vCpuResponse).Get(vCpuUrl)
	if err != nil {
		logger.Logger().Error("ECS vCpu error: ", err)
	}
	if vCpuResponse != nil && vCpuResponse.Data != nil && vCpuResponse.Data.Result != nil && len(vCpuResponse.Data.Result) != 0 {
		largeScreenEcs.Cpu.Value = vCpuResponse.Data.Result[0].Value[1].(string)
	}
	//ECS 内存
	memoryPql := fmt.Sprintf("sum(ecs_base_memory_available_bytes{%v})/1024/1024", constant.INSTANCE+"=~'"+resources+"'")
	memoryUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(memoryPql)
	memoryResponse := &form.PrometheusResponse{}
	_, err = httputil.GetHttpClient().R().SetResult(&memoryResponse).Get(memoryUrl)
	if err != nil {
		logger.Logger().Error("ECS memory error: ", err)
	}
	if memoryResponse != nil && memoryResponse.Data != nil && memoryResponse.Data.Result != nil && len(memoryResponse.Data.Result) != 0 {
		float, _ := strconv.ParseFloat(memoryResponse.Data.Result[0].Value[1].(string), 64)
		if float < 10000 {
			largeScreenEcs.Memory.Unit = "MiB"
			largeScreenEcs.Memory.Value = fmt.Sprintf("%.2f", float)
		} else if float/1024 < 10000 {
			largeScreenEcs.Memory.Value = fmt.Sprintf("%.2f", float/1024)
			largeScreenEcs.Memory.Unit = "GiB"
		} else {
			largeScreenEcs.Memory.Value = fmt.Sprintf("%.2f", float/1024/1024)
			largeScreenEcs.Memory.Unit = "TiB"
		}
	}
	//ECS 磁盘
	diskPql := fmt.Sprintf("sum(ecs_filesystem_size_bytes{%v})/1024/1024", constant.INSTANCE+"=~'"+resources+"'")
	diskUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(diskPql)
	diskResponse := &form.PrometheusResponse{}
	_, err = httputil.GetHttpClient().R().SetResult(&diskResponse).Get(diskUrl)
	if err != nil {
		logger.Logger().Error("ECS disk error: ", err)
	}
	if diskResponse != nil && diskResponse.Data != nil && diskResponse.Data.Result != nil && len(diskResponse.Data.Result) != 0 {
		float, _ := strconv.ParseFloat(diskResponse.Data.Result[0].Value[1].(string), 64)
		if float < 10000 {
			largeScreenEcs.Disk.Unit = "MiB"
			largeScreenEcs.Disk.Value = fmt.Sprintf("%.2f", float)
		} else if float/1024 < 10000 {
			largeScreenEcs.Disk.Value = fmt.Sprintf("%.2f", float/1024)
			largeScreenEcs.Disk.Unit = "GiB"
		} else {
			largeScreenEcs.Disk.Value = fmt.Sprintf("%.2f", float/1024/1024)
			largeScreenEcs.Disk.Unit = "TiB"
		}
	}
	return largeScreenEcs, nil
}

var ResourceEcsTopCpuPqlMap = map[string]string{
	"cpu1":  "sum by(instance,instanceType)(ecs_load1{$INSTANCE})",  //CPU1分钟平均负载
	"cpu5":  "sum by(instance,instanceType)(ecs_load5{$INSTANCE})",  //CPU5分钟平均负载
	"cpu15": "sum by(instance,instanceType)(ecs_load15{$INSTANCE})", //CPU15分钟平均负载
}

var ResourceEcsTopMemoryPqlMap = map[string]string{
	"memory1":  "avg_over_time(((1-sum by(instance,instanceType)(ecs_base_memory_unused_bytes{$INSTANCE})/sum by(instance,instanceType)(ecs_base_memory_available_bytes{$INSTANCE}))*100)[1m:1m])",  //(基础)内存1分钟使用率
	"memory5":  "avg_over_time(((1-sum by(instance,instanceType)(ecs_base_memory_unused_bytes{$INSTANCE})/sum by(instance,instanceType)(ecs_base_memory_available_bytes{$INSTANCE}))*100)[5m:1m])",  //(基础)内存5分钟使用率
	"memory15": "avg_over_time(((1-sum by(instance,instanceType)(ecs_base_memory_unused_bytes{$INSTANCE})/sum by(instance,instanceType)(ecs_base_memory_available_bytes{$INSTANCE}))*100)[15m:1m])", //(基础)内存15分钟使用率
}

var ResourceEcsTopDiskPqlMap = map[string]string{
	"disk": "100 * ((sum by(instance,instanceType)(ecs_filesystem_size_bytes{$INSTANCE}) - sum by(instance,instanceType)(ecs_filesystem_free_bytes{$INSTANCE})) / sum by(instance,instanceType)(ecs_filesystem_size_bytes{$INSTANCE}))", //磁盘使用率
}

func (s *LargeScreenService) ResourceEcsTop(tag, item string) ([]*form.LargeScreenEcsTop, error) {
	_, resourceList, err := s.getResourcesByTag(tag, constant.CloudProductCodeEcs, constant.ResourceTypeCodeInstance)
	if err != nil {
		return nil, err
	}
	resources := strings.Join(resourceList, "|")

	var resourceEcsPqlMap = make(map[string]string)
	switch item {
	case "cpu":
		resourceEcsPqlMap = ResourceEcsTopCpuPqlMap
	case "memory":
		resourceEcsPqlMap = ResourceEcsTopMemoryPqlMap
	case "disk":
		resourceEcsPqlMap = ResourceEcsTopDiskPqlMap
	}

	var list []*form.LargeScreenEcsTop
	for name, pql := range resourceEcsPqlMap {
		var ecsTop = &form.LargeScreenEcsTop{Name: name}
		list = append(list, ecsTop)
		pql = fmt.Sprintf(constant.TopExpr, "5", strings.ReplaceAll(pql, constant.MetricLabel, constant.INSTANCE+"=~'"+resources+"'"))
		prometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(pql)
		prometheusResponse := &form.PrometheusResponse{}
		_, err = httputil.GetHttpClient().R().SetResult(&prometheusResponse).Get(prometheusUrl)
		if err != nil {
			logger.Logger().Error("getUsageData error: ", err)
			continue
		}
		if prometheusResponse == nil || prometheusResponse.Data == nil || prometheusResponse.Data.Result == nil {
			logger.Logger().Infof("query prometheus empty, pql: %v, origin response:%v", pql, prometheusResponse)
			continue
		}
		for _, v := range prometheusResponse.Data.Result {
			ecsTop.Data = append(ecsTop.Data, &form.LargeScreenEcsTopValue{
				ResourceId: v.Metric[constant.INSTANCE],
				Value:      changeDecimal(v.Value[1].(string)),
			})
		}
	}
	return list, nil
}

func (s *LargeScreenService) ResourceEip(tag, interval string) ([]*form.LargeScreenMonitorEip, error) {
	_, resourceList, err := s.getResourcesByTag(tag, constant.CloudProductCodeEip, constant.ResourceTypeCodeInstance)
	if err != nil {
		return nil, err
	}
	resources := strings.Join(resourceList, "|")

	end := time.Now().Unix()
	start := end - intervalMap[interval]
	step := intervalMap[interval] / 10

	//出网带宽
	pql := fmt.Sprintf("sum(eip_upstream_bits_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	upstreamBandwidthMap := querySinglePrometheusRange(pql, start, end, step)
	//入网带宽
	pql = fmt.Sprintf("sum(eip_downstream_bits_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	downstreamBandwidthMap := querySinglePrometheusRange(pql, start, end, step)
	//出网流量
	pql = fmt.Sprintf("sum(eip_upstream_bits_rate{%v})/8*60", constant.INSTANCE+"=~'"+resources+"'")
	upstreamMap := querySinglePrometheusRange(pql, start, end, step)
	//入网流量
	pql = fmt.Sprintf("sum(eip_downstream_bits_rate{%v})/8*60", constant.INSTANCE+"=~'"+resources+"'")
	downstreamMap := querySinglePrometheusRange(pql, start, end, step)

	var list []*form.LargeScreenMonitorEip
	timeList := handleTimeInterval(start, end, step, start)
	for _, v := range timeList {
		list = append(list, &form.LargeScreenMonitorEip{
			Time:                time.Unix(v, 0).Format(util.HourMinuteTimeFmt),
			UpstreamBandwidth:   upstreamBandwidthMap[v],
			DownstreamBandwidth: downstreamBandwidthMap[v],
			Upstream:            upstreamMap[v],
			Downstream:          downstreamMap[v],
		})
	}
	return list, err
}

func (s *LargeScreenService) ResourceNat(tag, interval string) ([]*form.LargeScreenMonitorNat, error) {
	_, resourceList, err := s.getResourcesByTag(tag, constant.CloudProductCodeNat, constant.ResourceTypeCodeInstance)
	if err != nil {
		return nil, err
	}
	resources := strings.Join(resourceList, "|")

	end := time.Now().Unix()
	start := end - intervalMap[interval]
	step := intervalMap[interval] / 10

	//出方向带宽
	pql := fmt.Sprintf("sum(rate(Nat_send_bytes_total_count{%v}[3m])*8)", constant.INSTANCE+"=~'"+resources+"'")
	upstreamBandwidthMap := querySinglePrometheusRange(pql, start, end, step)
	//入方向带宽
	pql = fmt.Sprintf("sum(rate(Nat_recv_bytes_total_count{%v}[3m])*8)", constant.INSTANCE+"=~'"+resources+"'")
	downstreamBandwidthMap := querySinglePrometheusRange(pql, start, end, step)
	//出方向流量
	pql = fmt.Sprintf("sum(Nat_send_bytes_total_count{%v})", constant.INSTANCE+"=~'"+resources+"'")
	upstreamMap := querySinglePrometheusRange(pql, start, end, step)
	//入方向流量
	pql = fmt.Sprintf("sum(Nat_recv_bytes_total_count{%v})", constant.INSTANCE+"=~'"+resources+"'")
	downstreamMap := querySinglePrometheusRange(pql, start, end, step)

	var list []*form.LargeScreenMonitorNat
	timeList := handleTimeInterval(start, end, step, start)
	for _, v := range timeList {
		list = append(list, &form.LargeScreenMonitorNat{
			Time:                time.Unix(v, 0).Format(util.HourMinuteTimeFmt),
			UpstreamBandwidth:   upstreamBandwidthMap[v],
			DownstreamBandwidth: downstreamBandwidthMap[v],
			Upstream:            upstreamMap[v],
			Downstream:          downstreamMap[v],
		})
	}
	return list, err
}

func (s *LargeScreenService) ResourceSlb(tag, interval string) ([]*form.LargeScreenMonitorSlb, error) {
	_, resourceList, err := s.getResourcesByTag(tag, constant.CloudProductCodeSlb, constant.ResourceTypeCodeInstance)
	if err != nil {
		return nil, err
	}
	resources := strings.Join(resourceList, "|")

	end := time.Now().Unix()
	start := end - intervalMap[interval]
	step := intervalMap[interval] / 10

	//并发连接数
	pql := fmt.Sprintf("sum(Slb_all_connection_count{%v})", constant.INSTANCE+"=~'"+resources+"'")
	allMap := querySinglePrometheusRange(pql, start, end, step)
	//活跃连接数
	pql = fmt.Sprintf("sum(Slb_all_est_connection_count{%v})", constant.INSTANCE+"=~'"+resources+"'")
	allEstMap := querySinglePrometheusRange(pql, start, end, step)
	//非活跃连接数
	pql = fmt.Sprintf("sum(Slb_all_none_est_connection_count{%v})", constant.INSTANCE+"=~'"+resources+"'")
	allNoneEstMap := querySinglePrometheusRange(pql, start, end, step)
	//新建连接数
	pql = fmt.Sprintf("sum(Slb_new_connection_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	newMap := querySinglePrometheusRange(pql, start, end, step)
	//丢弃连接数
	pql = fmt.Sprintf("sum(Slb_drop_connection_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	dropMap := querySinglePrometheusRange(pql, start, end, step)
	//7层协议查询速率
	pql = fmt.Sprintf("sum(Slb_request_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	requestMap := querySinglePrometheusRange(pql, start, end, step)
	//7层协议返回客户端2xx状态码数
	pql = fmt.Sprintf("sum(Slb_http_2xx_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	http2xxMap := querySinglePrometheusRange(pql, start, end, step)
	//7层协议返回客户端3xx状态码数
	pql = fmt.Sprintf("sum(Slb_http_3xx_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	http3xxMap := querySinglePrometheusRange(pql, start, end, step)
	//7层协议返回客户端4xx状态码数
	pql = fmt.Sprintf("sum(Slb_http_4xx_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	http4xxMap := querySinglePrometheusRange(pql, start, end, step)
	//7层协议返回客户端5xx状态码数
	pql = fmt.Sprintf("sum(Slb_http_5xx_rate{%v})", constant.INSTANCE+"=~'"+resources+"'")
	http5xxMap := querySinglePrometheusRange(pql, start, end, step)

	var list []*form.LargeScreenMonitorSlb
	timeList := handleTimeInterval(start, end, step, start)
	for _, v := range timeList {
		list = append(list, &form.LargeScreenMonitorSlb{
			Time:                      time.Unix(v, 0).Format(util.HourMinuteTimeFmt),
			AllConnectionCount:        allMap[v],
			AllEstConnection:          allEstMap[v],
			AllNoneEstConnectionCount: allNoneEstMap[v],
			NewConnectionRate:         newMap[v],
			DropConnectionRate:        dropMap[v],
			RequestRate:               requestMap[v],
			Http2xxRate:               http2xxMap[v],
			Http3xxRate:               http3xxMap[v],
			Http4xxRate:               http4xxMap[v],
			Http5xxRate:               http5xxMap[v],
		})
	}
	return list, err
}

func querySinglePrometheusRange(pql string, start, end, step int64) map[int64]float64 {
	prometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.QueryRange + url.QueryEscape(pql)
	prometheusResponse := &form.PrometheusResponse{}
	_, err := httputil.GetHttpClient().R().SetResult(&prometheusResponse).Get(fmt.Sprintf("%v&start=%v&end=%v&step=%v", prometheusUrl, start, end, step))
	if err != nil {
		logger.Logger().Error("querySinglePrometheusRange error: ", err)
	}
	var prometheusMap = make(map[int64]float64)
	if prometheusResponse != nil || prometheusResponse.Data != nil || prometheusResponse.Data.Result != nil {
		for _, result := range prometheusResponse.Data.Result {
			for _, v := range result.Values {
				prometheusMap[int64(v[0].(float64))] = changeDecimalFloat(v[1].(string))
			}
		}
	}
	return prometheusMap
}

var resourceMysqlCodeList = []string{
	"mysql_cpu_usage",          //CPU使用率
	"mysql_mem_usage",          //内存使用率
	"mysql_disk_usage",         //磁盘使用率
	"mysql_current_cons_num",   //当前打开连接数
	"mysql_active_connections", //活跃连接数
	"mysql_qps",                //QPS
	"mysql_tps",                //TPS
}

var resourceDmCodeList = []string{
	"dm_global_status_cpu_use_rate",    //CPU使用率
	"dm_global_status_mem_use_rate",    //内存使用率
	"dm_global_status_sessions",        //总会话数
	"dm_global_status_active_sessions", //活动会话数
	"dm_global_status_qps",             //每秒执行select SQL语句数
	"dm_global_status_tps",             //每秒事务数
}

var resourcePgCodeList = []string{
	"pg_cpu_usage",     //CPU使用率
	"pg_mem_usage",     //内存使用率
	"pg_disk_usage",    //磁盘使用率
	"pg_open_ct_num",   //当前打开连接数
	"pg_active_ct_num", //当前活跃连接数
	"pg_qps",           //QPS
	"pg_tps",           //TPS
}

func (s *LargeScreenService) ResourceRdb(tag, interval, productCode string) ([]*form.LargeScreenMonitor, error) {
	var monitorItemCodeList []string
	var cloudProductCode, ResourceTypeCode string
	switch productCode {
	case "mysql":
		monitorItemCodeList = resourceMysqlCodeList
		cloudProductCode = constant.CloudProductCodeRdb
		ResourceTypeCode = constant.ResourceTypeCodeMysql
	case "dm":
		monitorItemCodeList = resourceDmCodeList
		cloudProductCode = constant.CloudProductCodeRdb
		ResourceTypeCode = constant.ResourceTypeCodeDm
	case "pg":
		monitorItemCodeList = resourcePgCodeList
		cloudProductCode = constant.CloudProductCodeRdb
		ResourceTypeCode = constant.ResourceTypeCodePg
	}

	_, resourceList, err := s.getResourcesByTag(tag, cloudProductCode, ResourceTypeCode)
	if err != nil {
		return nil, err
	}
	resources := strings.Join(resourceList, "|")
	end := time.Now().Unix()
	start := end - intervalMap[interval]
	step := intervalMap[interval] / 10

	var list []*form.LargeScreenMonitor
	for _, code := range monitorItemCodeList {
		var monitorItem *model.MonitorItem
		if err = global.DB.Where("metric_name = ?", code).Find(&monitorItem).Error; err != nil {
			logger.Logger().Error("ResourceRdb error: ", err)
			continue
		}
		var monitor = &form.LargeScreenMonitor{Name: monitorItem.Name, Code: monitorItem.MetricName, Unit: monitorItem.Unit}
		if strutil.IsNotBlank(resources) {
			pql := strings.ReplaceAll(monitorItem.MetricsLinux, constant.MetricLabel, constant.INSTANCE+"=~'"+resources+"'")
			prometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.QueryRange + url.QueryEscape(pql)
			prometheusResponse := &form.PrometheusResponse{}
			_, err = httputil.GetHttpClient().R().SetResult(&prometheusResponse).Get(fmt.Sprintf("%v&start=%v&end=%v&step=%v", prometheusUrl, start, end, step))
			if err != nil {
				logger.Logger().Error("getUsageData error: ", err)
				continue
			}
			if prometheusResponse == nil || prometheusResponse.Data == nil || prometheusResponse.Data.Result == nil {
				logger.Logger().Infof("query prometheus empty, pql: %v, origin response:%v", pql, prometheusResponse)
				continue
			}
			result := prometheusResponse.Data.Result
			var timeList []int64
			if len(result) == 0 {
				timeList = handleTimeInterval(start, end, step, start)
			} else {
				timeList = handleTimeInterval(start, end, step, int64(result[0].Values[0][0].(float64)))
			}

			for _, v := range prometheusResponse.Data.Result {
				var monitorChart = &form.LargeScreenMonitorChart{ResourceId: v.Metric[constant.INSTANCE]}
				var timeMap = make(map[int64]string)
				for _, value := range v.Values {
					timeMap[int64(value[0].(float64))] = changeDecimal(value[1].(string))
				}
				for _, timestamp := range timeList {
					monitorChart.Data = append(monitorChart.Data, &form.LargeScreenMonitorData{
						Time:  time.Unix(timestamp, 0).Format(util.HourMinuteTimeFmt),
						Value: timeMap[timestamp],
					})
				}
				monitor.Chart = append(monitor.Chart, monitorChart)
			}
		}
		list = append(list, monitor)
	}
	return list, nil
}

func (s *LargeScreenService) ResourceStorage(tag, cloudProductCode, ResourceTypeCode string) ([]*form.LargeScreenStorageTrend, error) {
	_, resourceList, err := s.getResourcesByTag(tag, cloudProductCode, ResourceTypeCode)
	if err != nil {
		return nil, err
	}
	var resourceStorage []*model.LargeScreenResourceStorage
	if err = global.DB.Where("type = ? AND resource_id IN (?)", cloudProductCode, resourceList).Order("create_time ASC").Find(&resourceStorage).Error; err != nil {
		logger.Logger().Error(err)
		return nil, err
	}
	now := util.GetNow()
	var timeList []string
	for i := 7; i >= 0; i-- {
		timeList = append(timeList, now.AddDate(0, 0, -i).Format("01-02"))
	}
	var valueMap = make(map[string]float64)
	var conversionMap = make(map[string]string)
	var unit string
	var min, max float64
	for _, v := range resourceStorage {
		t := util.StrToTime(util.DayTimeFmt, v.Time).Format("01-02")
		valueMap[t] += float64(v.Value)
	}
	for _, v := range valueMap {
		if v == 0 {
			continue
		}
		if min > v || min == 0 {
			min = v
		}
		if max < v {
			max = v
		}
	}
	if min == 0 && max == 0 {
		unit = "Byte"
	} else if min < 10000 {
		unit = "Byte"
	} else if min/1024 < 10000 {
		unit = "KiB"
	} else if min/1024/1024 < 10000 {
		unit = "MiB"
	} else if min/1024/1024/1024 < 10000 {
		unit = "GiB"
	} else {
		unit = "TiB"
	}
	for k, v := range valueMap {
		switch unit {
		case "Byte":
			valueMap[k], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
		case "KiB":
			valueMap[k], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v/1024), 64)
		case "MiB":
			valueMap[k], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v/1024/1024), 64)
		case "GiB":
			valueMap[k], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v/1024/1024/1024), 64)
		case "TiB":
			valueMap[k], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v/1024/1024/1024/1024), 64)
		}
		if v < 10000 {
			conversionMap[k] = fmt.Sprintf("%.2f", v) + "Byte"
		} else if v/1024 < 10000 {
			conversionMap[k] = fmt.Sprintf("%.2f", v/1024) + "KiB"
		} else if v/1024/1024 < 10000 {
			conversionMap[k] = fmt.Sprintf("%.2f", v/1024/1024) + "MiB"
		} else if v/1024/1024/1024 < 10000 {
			conversionMap[k] = fmt.Sprintf("%.2f", v/1024/1024/1024) + "GiB"
		} else {
			conversionMap[k] = fmt.Sprintf("%.2f", v/1024/1024/1024/1024) + "TiB"
		}
	}
	var trend []*form.LargeScreenStorageTrend
	for _, v := range timeList {
		trend = append(trend, &form.LargeScreenStorageTrend{
			Time:       v,
			Value:      valueMap[v],
			Unit:       unit,
			Conversion: conversionMap[v],
		})
	}
	return trend, nil
}

func (s *LargeScreenService) getResourcesByTag(tag string, cloudProductCode, resourceTypeCode string) ([]*form.LargeScreenResource, []string, error) {
	var response *form.LargeScreenResourceResponse
	param := map[string][]string{"tagKeyIdList": {tag}, "regionCodeList": {config.Cfg.Common.RegionName}}
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.ResourceTagApi + "/list-resource-by-tag-id")
	if err != nil || response == nil || response.Module == nil {
		return nil, nil, err
	}
	var resources []string
	if strutil.IsNotBlank(cloudProductCode) && strutil.IsNotBlank(resourceTypeCode) {
		for _, v := range response.Module.List {
			if v.CloudProductCode == cloudProductCode && v.ResourceTypeCode == resourceTypeCode {
				resources = append(resources, v.ResourceInstanceId)
			}
		}
	} else {
		for _, v := range response.Module.List {
			resources = append(resources, v.ResourceInstanceId)
		}
	}
	return response.Module.List, resources, nil
}

var intervalMap = map[string]int64{"1h": 3600, "6h": 21600, "12h": 43200, "24h": 86400}

// 获时间点位
func handleTimeInterval(start int64, end int64, step int64, firstTime int64) []int64 {
	var timeList []int64
	for firstTime-step >= start {
		firstTime -= step
	}
	for firstTime <= end {
		timeList = append(timeList, firstTime)
		firstTime += step
	}
	return timeList
}

// 数据保留两位小数
func changeDecimalFloat(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	v, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
	return v
}
