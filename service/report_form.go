package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/errors"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ReportFormService struct {
	prometheus *PrometheusService
}

func NewReportFormService() *ReportFormService {
	return &ReportFormService{
		prometheus: NewPrometheusService(),
	}
}

func (s *ReportFormService) GetMonitorData(param form.ReportFormParam) ([]*form.ReportForm, error) {
	var instanceList []string
	instanceMap := make(map[string]*form.InstanceForm)
	for _, v := range param.InstanceList {
		instanceList = append(instanceList, v.InstanceId)
		instanceMap[v.InstanceId] = v
	}
	instances := strings.Join(instanceList, "|")
	item := dao.MonitorItem.GetMonitorItemCacheByMetricCode(param.ItemList[0])
	if strutil.IsNotBlank(item.Unit) {
		item.Name = fmt.Sprintf("%v（%v）", item.Name, item.Unit)
	}
	labels := strings.Split(item.Labels, ",")
	pql := strings.ReplaceAll(item.Expression, constant.MetricLabel, constant.INSTANCE+"=~'"+instances+"'")
	//获取单个指标的所有实例数据
	reportFormList := s.getOneItemData(param, item, instanceMap, pql, labels)
	return reportFormList, nil
}

func (s *ReportFormService) getOneItemData(param form.ReportFormParam, item form.MonitorItem, instanceMap map[string]*form.InstanceForm, pql string, labels []string) []*form.ReportForm {
	if len(param.Statistics) == 0 {
		return s.getOriginData(param, item, instanceMap, pql, labels)
	}
	return s.getAggregationData(param, item, instanceMap, pql, labels)
}

func (s *ReportFormService) getOriginData(param form.ReportFormParam, item form.MonitorItem, instanceMap map[string]*form.InstanceForm, pql string, labels []string) []*form.ReportForm {
	result := s.prometheus.QueryRange(pql, strconv.Itoa(param.Start), strconv.Itoa(param.End), strconv.Itoa(param.Step)).Data.Result
	if len(result) == 0 {
		return nil
	}
	var list []*form.ReportForm
	for _, prometheusResult := range result {
		for _, prometheusValue := range prometheusResult.Values {
			if f := s.buildOriginReportForm(param, instanceMap, item, labels, prometheusResult, prometheusValue); f != nil {
				list = append(list, f)
			}
		}
	}
	return list
}

func (s *ReportFormService) getAggregationData(param form.ReportFormParam, item form.MonitorItem, instanceMap map[string]*form.InstanceForm, pql string, labels []string) []*form.ReportForm {
	//计算开始时间当天的23时59分59秒
	start := param.Start + (86400 - (param.Start-57600)%86400)
	//计算结束时间当天的23时59分59秒
	end := param.End + (86400 - (param.End-57600)%86400)
	var result map[string]*form.PrometheusResult
	var ret = make(map[string]map[string]*form.PrometheusResult)
	//开启协程
	group := &sync.WaitGroup{}
	group.Add(len(param.Statistics))
	for _, statistics := range param.Statistics {
		m := make(map[string]*form.PrometheusResult)
		go s.getStatisticsMap(statistics, pql, start, end, param.Statistics, labels, group, m)
		ret[statistics] = m
		if result == nil {
			result = m
		}
	}
	group.Wait()
	var list []*form.ReportForm
	for k, v := range result {
		for i := range v.Values {
			dataMap := make(map[string][]interface{})
			for calcStyle, d := range ret {
				dataMap[calcStyle] = d[k].Values[i]
			}
			if f := s.buildAggregationReportForm(v.Metric["instance"], k, item, instanceMap, dataMap); f != nil {
				list = append(list, f)
			}
		}
	}
	return list
}

func (s *ReportFormService) buildOriginReportForm(param form.ReportFormParam, instanceMap map[string]*form.InstanceForm, item form.MonitorItem, labels []string, prometheusResult *form.PrometheusResult, prometheusValue []interface{}) (f *form.ReportForm) {
	defer func() {
		if e := recover(); e != nil {
			logger.Logger().Error(e)
		}
	}()
	f = &form.ReportForm{
		Region:       param.RegionCode,
		InstanceName: instanceMap[prometheusResult.Metric["instance"]].InstanceName,
		InstanceId:   prometheusResult.Metric["instance"],
		Status:       instanceMap[prometheusResult.Metric["instance"]].Status,
		ItemName:     item.Name,
		Time:         util.TimestampToFullTimeFmtStr(int64(prometheusValue[0].(float64))),
		Timestamp:    int64(prometheusValue[0].(float64)),
		Value:        changeDecimal(prometheusValue[1].(string)),
	}
	for _, label := range labels {
		if label != "instance" && strutil.IsNotBlank(prometheusResult.Metric[label]) {
			f.InstanceId = f.InstanceId + " - " + prometheusResult.Metric[label]
		}
	}
	return
}

func (s *ReportFormService) buildAggregationReportForm(instanceId, key string, item form.MonitorItem, instanceMap map[string]*form.InstanceForm, dataMap map[string][]interface{}) (f *form.ReportForm) {
	defer func() {
		if e := recover(); e != nil {
			logger.Logger().Error(e)
		}
	}()
	time := ""
	timestamp := int64(0)
	for _, v := range dataMap {
		time = util.TimestampToDayTimeFmtStr(int64(v[0].(float64)) - 1)
		timestamp = int64(v[0].(float64) - 1)
		break
	}
	f = &form.ReportForm{
		Region:       config.Cfg.Common.RegionName,
		InstanceName: instanceMap[instanceId].InstanceName,
		InstanceId:   key,
		Status:       instanceMap[instanceId].Status,
		ItemName:     item.Name,
		Time:         time,
		Timestamp:    timestamp,
	}

	for calcStyle, d := range dataMap {
		rf := reflect.ValueOf(f)
		ff := rf.Elem().FieldByName(firstUpper(calcStyle) + "Value")
		ff.SetString(changeDecimal(d[1].(string)))
	}
	return
}

func (s *ReportFormService) getStatisticsMap(aggregation, pql string, start, end int, statistics, labels []string, group *sync.WaitGroup, resultMap map[string]*form.PrometheusResult) {
	defer group.Done()
	for _, v := range statistics {
		if aggregation == v {
			expr := fmt.Sprintf("%s_over_time((%s)[1d:1h])", aggregation, pql)
			result := s.prometheus.QueryRange(expr, strconv.Itoa(start), strconv.Itoa(end), "86400").Data.Result
			for _, prometheusResult := range result {
				key := prometheusResult.Metric["instance"]
				for _, label := range labels {
					if label != "instance" && strutil.IsNotBlank(prometheusResult.Metric[label]) {
						key = key + " - " + prometheusResult.Metric[label]
					}
				}
				resultMap[key] = prometheusResult
			}
		}
	}
}

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]

}

func (s *ReportFormService) Export(param form.ReportFormParam, userInfo string) error {
	var url = config.Cfg.Common.AsyncExportApi
	var num = len(param.InstanceList)
	var header = map[string]string{"user-info": userInfo}
	var count = 1
	var countString string
	var sheetParamList []string
	var newParam form.ReportFormParam
	for i, instance := range param.InstanceList {
		for _, item := range param.ItemList {
			newParam = param
			newParam.ItemList = []string{item}
			newParam.InstanceList = []*form.InstanceForm{instance}
			sheetParamList = append(sheetParamList, jsonutil.ToString(newParam))
		}
		if (i+1)%200 == 0 || i+1 == num {
			asyncParams := []form.AsyncExportParam{
				{
					SheetSeq:       0,
					SheetName:      "云资源监控",
					SheetParamList: sheetParamList,
				},
			}
			if num > 200 {
				countString = "-" + strconv.Itoa(count)
			}
			asyncRequest := form.AsyncExportRequest{
				TemplateId: "cloud_monitor_manager",
				Params:     asyncParams,
				FileName:   "云上资源监控" + countString,
			}
			result, err := httputil.HttpPostJson(url, asyncRequest, header)
			logger.Logger().Infof("AsyncExport：%v", result)
			if err != nil {
				logger.Logger().Infof("AsyncExportError：%v", err)
				return errors.NewBusinessError("异步导出API调用失败")
			}
			sheetParamList = []string{}
			count++
		}
	}
	return nil
}

var ecsCpuBaseUsageDownSampling = "100 * avg by(instance,instanceType)(rate(ecs_base_vcpu_seconds{$INSTANCE}[3h]))"

func (s *ReportFormService) GetReportFormData(param form.ReportFormParam) ([]*form.ReportForm, error) {
	if len(param.InstanceList) == 0 {
		return nil, errors.NewBusinessError("实例不能为空")
	}
	if strutil.IsBlank(param.Item) {
		return nil, errors.NewBusinessError("指标不能为空")
	}
	var instanceList []string
	instanceMap := make(map[string]*form.InstanceForm)
	for _, v := range param.InstanceList {
		instanceList = append(instanceList, v.InstanceId)
		instanceMap[v.InstanceId] = v
	}
	instances := strings.Join(instanceList, "|")
	item := dao.MonitorItem.GetMonitorItemCacheByMetricCode(param.Item)
	labels := strings.Split(item.Labels, ",")
	var downSampling = false
	if int(time.Now().Unix())-param.Start >= 3024000 {
		downSampling = true
		if param.Item == "ecs_cpu_base_usage" {
			item.Expression = ecsCpuBaseUsageDownSampling
		}
	}

	pql := strings.ReplaceAll(item.Expression, constant.MetricLabel, constant.INSTANCE+"=~'"+instances+"'")
	//获取单个指标的所有实例数据
	//计算开始时间当天的23时59分59秒
	start := param.Start + (86400 - (param.Start-57600)%86400)
	//计算结束时间当天的23时59分59秒
	end := param.End + (86400 - (param.End-57600)%86400)
	var result map[string]*form.PrometheusResult
	var ret = make(map[string]map[string]*form.PrometheusResult)
	//开启协程
	group := &sync.WaitGroup{}
	group.Add(len(param.Statistics))
	for _, statistics := range param.Statistics {
		m := make(map[string]*form.PrometheusResult)
		go func(aggregation, pql string, start, end int, statistics, labels []string, group *sync.WaitGroup, resultMap map[string]*form.PrometheusResult) {
			defer group.Done()
			for _, v := range statistics {
				if aggregation == v {
					expr := fmt.Sprintf("%s_over_time((%s)[1d:1h])", aggregation, pql)
					var response *form.PrometheusResponse
					if downSampling {
						response = s.prometheus.QueryFrontendRangeDownSampling(expr, strconv.Itoa(start), strconv.Itoa(end), "86400")
					} else {
						response = s.prometheus.QueryFrontendRange(expr, strconv.Itoa(start), strconv.Itoa(end), "86400")
					}
					if response.Data == nil || response.Data.Result == nil {
						continue
					}
					for _, prometheusResult := range response.Data.Result {
						key := prometheusResult.Metric["instance"]
						for _, label := range labels {
							if label != "instance" && strutil.IsNotBlank(prometheusResult.Metric[label]) {
								key = key + " - " + prometheusResult.Metric[label]
							}
						}
						resultMap[key] = prometheusResult
					}
				}
			}
		}(statistics, pql, start, end, param.Statistics, labels, group, m)

		ret[statistics] = m
		if result == nil {
			result = m
		}
	}
	group.Wait()
	var list []*form.ReportForm
	for k, v := range result {
		for i := range v.Values {
			dataMap := make(map[string][]interface{})
			for calcStyle, d := range ret {
				dataMap[calcStyle] = d[k].Values[i]
			}
			if f := s.buildAggregationReportForm(v.Metric["instance"], k, item, instanceMap, dataMap); f != nil {
				list = append(list, f)
			}
		}
	}
	return list, nil
}
