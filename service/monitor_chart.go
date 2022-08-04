package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/errors"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
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
	pql := strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+request.Instance+"',"+constant.FILTER)
	if strutil.IsNotBlank(request.Pid) {
		pql = strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+request.Instance+"',"+fmt.Sprintf(constant.PId, request.Pid))
	}
	prometheusResponse := s.prometheus.QueryRange(pql, strconv.Itoa(request.Start), strconv.Itoa(request.End), strconv.Itoa(request.Step))
	result := prometheusResponse.Data.Result
	labels := strings.Split(monitorItem.Labels, ",")
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
		ValueAxis: valueAxisFillEmptyData(result, timeList, labels, request.Instance),
	}
	return prometheusAxis, nil
}

func (s *MonitorChartService) GetTopData(request form.PrometheusRequest, instanceIdList []string, monitorItem form.MonitorItem) ([]form.PrometheusInstance, error) {
	if strutil.IsBlank(monitorItem.Expression) {
		return nil, errors.NewBusinessError("监控指标不存在")
	}
	var pql string
	if len(instanceIdList) > 0 {
		for i, v := range instanceIdList {
			instanceIdList[i] = fmt.Sprintf(constant.TopExpr, "1", strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+v+"'"))
		}
		pql = fmt.Sprintf(constant.TopExpr, strconv.Itoa(request.TopNum), strings.Join(instanceIdList, " or "))
	} else {
		pql = fmt.Sprintf(constant.TopExpr, strconv.Itoa(request.TopNum), strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, ""))
	}
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

func (s *MonitorChartService) GetProcessData(request form.PrometheusRequest) ([]form.ProcessData, error) {
	if strutil.IsBlank(request.Instance) {
		return nil, errors.NewBusinessError("instance为空")
	}
	if request.Start == 0 || request.End == 0 || request.Start > request.End {
		return nil, errors.NewBusinessError("时间参数错误")
	}
	cpu := dao.MonitorItem.GetMonitorItemCacheByMetricCode("ecs_processes_top5Cpus")
	mem := dao.MonitorItem.GetMonitorItemCacheByMetricCode("ecs_processes_top5Mems")
	fd := dao.MonitorItem.GetMonitorItemCacheByMetricCode("ecs_processes_top5Fds")
	cpuPql := strings.ReplaceAll(cpu.Expression, constant.MetricLabel, constant.INSTANCE+"='"+request.Instance+"',"+constant.FILTER)
	memPql := strings.ReplaceAll(mem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+request.Instance+"',"+constant.FILTER)
	fdPql := strings.ReplaceAll(fd.Expression, constant.MetricLabel, constant.INSTANCE+"='"+request.Instance+"',"+constant.FILTER)
	cpuResponse := s.prometheus.QueryRange(cpuPql, strconv.Itoa(request.Start), strconv.Itoa(request.End), strconv.Itoa(request.Step))
	memResponse := s.prometheus.QueryRange(memPql, strconv.Itoa(request.Start), strconv.Itoa(request.End), strconv.Itoa(request.Step))
	fdResponse := s.prometheus.QueryRange(fdPql, strconv.Itoa(request.Start), strconv.Itoa(request.End), strconv.Itoa(request.Step))
	memMap := make(map[string]*form.PrometheusResult)
	fdMap := make(map[string]*form.PrometheusResult)
	for _, v := range memResponse.Data.Result {
		memMap[v.Metric["pid"]] = v
	}
	for _, v := range fdResponse.Data.Result {
		fdMap[v.Metric["pid"]] = v
	}
	var processList []form.ProcessData
	for _, v := range cpuResponse.Data.Result {
		process := form.ProcessData{
			Pid:     v.Metric["pid"],
			CmdLine: v.Metric["cmd_line"],
			Name:    getProcessName(v.Metric["cmd_line"]),
		}
		if len(v.Values) != 0 {
			process.Time = util.TimestampToFullTimeFmtStr(int64(v.Values[len(v.Values)-1][0].(float64)))
			process.Cpu = changeDecimal(v.Values[len(v.Values)-1][1].(string))
			process.Memory = changeDecimal(memMap[process.Pid].Values[len(memMap[process.Pid].Values)-1][1].(string))
			process.Openfiles = fdMap[process.Pid].Values[len(fdMap[process.Pid].Values)-1][1].(string)
		}
		processList = append(processList, process)
	}
	return processList, nil
}

//获取区间数的值，为采集到的时间点位设为null
func valueAxisFillEmptyData(result []*form.PrometheusResult, timeList, labels []string, instanceId string) map[string][]string {
	resultMap := make(map[string][]string)
	for _, v := range result {
		timeMap := map[string]string{}
		for _, value := range v.Values {
			key := strconv.Itoa(int(value[0].(float64)))
			timeMap[key] = value[1].(string)
		}
		var arr []string
		for _, time := range timeList {
			arr = append(arr, changeDecimal(timeMap[time]))
		}
		key := instanceId
		for _, label := range labels {
			if label != "instance" && strutil.IsNotBlank(v.Metric[label]) {
				key = key + " - " + v.Metric[label]
			}
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

func getProcessName(cmdLine string) string {
	if strutil.IsBlank(cmdLine) {
		return "unknown"
	}
	list := strings.Split(strings.Split(cmdLine, " ")[0], "/")
	return list[len(list)-1]
}
