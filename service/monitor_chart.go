package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/errors"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"fmt"
	"strconv"
	"strings"
)

type MonitorChartService struct {
	prometheus *PrometheusService
}

func NewMonitorChartService() *MonitorChartService {
	return &MonitorChartService{
		prometheus: NewPrometheusService(),
	}
}

func (s *MonitorChartService) GetData(request form.PrometheusRequest) (*form.PrometheusValue, error) {
	if strutil.IsBlank(request.Instance) {
		return nil, errors.NewBusinessError("instance为空")
	}
	monitorItem := dao.MonitorItem.GetMonitorItemCacheByMetricCode(request.Name)
	if strutil.IsBlank(monitorItem.Code) {
		return nil, errors.NewBusinessError("指标不存在")
	}
	if monitorItem.ProductAbbreviation == "ecs" {
		request.Instance = changeEcsInstanceId(monitorItem.Host, request.TenantId, request.Instance)
	}
	pql := strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+request.Instance+"',"+constant.FILTER)
	prometheusResponse := s.prometheus.Query(pql, request.Time)
	if len(prometheusResponse.Data.Result) == 0 {
		return nil, nil
	}
	result := prometheusResponse.Data.Result[0].Value
	prometheusValue := &form.PrometheusValue{
		Time:  strconv.Itoa(int(result[0].(float64))),
		Value: changeDecimal(result[1].(string)),
	}
	return prometheusValue, nil
}

func (s *MonitorChartService) GetRangeData(request form.PrometheusRequest) (*form.PrometheusAxis, error) {
	if strutil.IsBlank(request.Instance) {
		return nil, errors.NewBusinessError("instance为空")
	}
	monitorItem := dao.MonitorItem.GetMonitorItemCacheByMetricCode(request.Name)
	if monitorItem.ProductAbbreviation == "ecs" {
		request.Instance = changeEcsInstanceId(monitorItem.Host, request.TenantId, request.Instance)
	}
	pql := strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+request.Instance+"',"+constant.FILTER)
	prometheusResponse := s.prometheus.QueryRange(pql, strconv.Itoa(request.Start), strconv.Itoa(request.End), strconv.Itoa(request.Step))
	result := prometheusResponse.Data.Result
	labels := strings.Split(monitorItem.Labels, ",")
	var label string
	for i := range labels {
		if labels[i] != "instance" {
			label = labels[i]
		}
	}
	start := request.Start
	end := request.End
	step := request.Step
	var timeList []string
	if len(result) == 0 {
		timeList = getTimeList(start, end, step, start)
	} else {
		timeList = getTimeList(start, end, step, int(result[0].Values[0][0].(float64)))
	}
	prometheusAxis := &form.PrometheusAxis{
		TimeAxis:  timeList,
		ValueAxis: valueAxisFillEmptyData(result, timeList, label, request.Instance),
	}
	return prometheusAxis, nil
}

func (s *MonitorChartService) GetTopData(request form.PrometheusRequest) ([]form.PrometheusInstance, error) {
	if strutil.IsBlank(request.Name) {
		return nil, errors.NewBusinessError("监控指标不能为空")
	}
	pql := fmt.Sprintf(constant.TopExpr, request.TopNum, strings.ReplaceAll(dao.MonitorItem.GetMonitorItemCacheByMetricCode(request.Name).Expression, constant.MetricLabel, ""))
	result := s.prometheus.Query(pql, request.Time).Data.Result
	var instanceList []form.PrometheusInstance
	for i := range result {
		instanceDTO := form.PrometheusInstance{
			Instance: result[i].Metric[constant.INSTANCE],
			Value:    changeDecimal(result[i].Value[1].(string)),
		}
		instanceList = append(instanceList, instanceDTO)
	}
	return instanceList, nil
}

//获取区间数的值，为采集到的时间点位设为null
func valueAxisFillEmptyData(result []form.PrometheusResult, timeList []string, label string, instanceId string) map[string][]string {
	resultMap := make(map[string][]string)
	for i := range result {
		timeMap := map[string]string{}
		for j := range result[i].Values {
			key := strconv.Itoa(int(result[i].Values[j][0].(float64)))
			timeMap[key] = result[i].Values[j][1].(string)
		}
		var key string
		var arr []string
		for k := range timeList {
			arr = append(arr, changeDecimal(timeMap[timeList[k]]))
		}
		if strutil.IsBlank(result[i].Metric[label]) {
			key = instanceId
		} else {
			key = result[i].Metric[label]
		}
		resultMap[key] = arr
	}
	return resultMap
}

//获取区间数据的时间点位
func getTimeList(start int, end int, step int, firstTime int) []string {
	var timeList []string
	if start > end {
		return timeList
	}
	for firstTime-step >= start {
		firstTime -= step
	}
	for firstTime <= end {
		timeList = append(timeList, strconv.Itoa(firstTime))
		firstTime += step
	}
	return timeList
}

//数据保留两位小数
func changeDecimal(value string) string {
	if strutil.IsBlank(value) {
		return ""
	}
	v, _ := strconv.ParseFloat(value, 64)
	return fmt.Sprintf("%.2f", v)
}

//ecs实例ID转换
func changeEcsInstanceId(host, tenantId, instanceId string) string {
	response, _ := httputil.HttpGet(host + "/compute/ecs/ops/v1/" + tenantId + "/servers/" + instanceId)
	var result EcsVo
	jsonutil.ToObject(response, &result)
	return result.Data.SerialNumber
}

type EcsVo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SerialNumber string `json:"serialNumber"`
	} `json:"data"`
}
