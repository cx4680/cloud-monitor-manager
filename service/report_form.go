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
	"strconv"
	"strings"
)

type ReportFormService struct {
	prometheus *PrometheusService
}

func NewReportFormService() *ReportFormService {
	return &ReportFormService{
		prometheus: NewPrometheusService(),
	}
}

func (s *ReportFormService) GetMonitorData(param form.ReportFormParam) ([]form.ReportForm, error) {
	var reportForm []form.ReportForm
	var instanceList []string
	instanceMap := make(map[string]form.InstanceForm)
	for _, v := range param.InstanceList {
		instanceList = append(instanceList, v.InstanceId)
		instanceMap[v.InstanceId] = v
	}
	instances := strings.Join(instanceList, "|")
	for _, v := range param.ItemList {
		item := dao.MonitorItem.GetMonitorItemCacheByMetricCode(v)
		labels := strings.Split(item.Labels, ",")
		pql := strings.ReplaceAll(item.Expression, constant.MetricLabel, constant.INSTANCE+"=~'"+instances+"'")
		result := s.prometheus.QueryRange(pql, strconv.Itoa(param.Start), strconv.Itoa(param.End), strconv.Itoa(param.Step)).Data.Result
		if len(result) == 0 {
			continue
		}
		if len(param.Statistics) == 0 {
			for _, v1 := range result {
				for _, v2 := range v1.Values {
					f := form.ReportForm{
						Region:       param.RegionCode,
						InstanceName: instanceMap[v1.Metric["instance"]].InstanceName,
						InstanceId:   v1.Metric["instance"],
						Status:       instanceMap[v1.Metric["instance"]].Status,
						ItemName:     item.Name,
						Time:         util.TimestampToFullTimeFmtStr(int64(v2[0].(float64))),
						Timestamp:    int64(v2[0].(float64)),
						Value:        changeDecimal(v2[1].(string)),
					}
					for _, v3 := range labels {
						if v3 != "instance" && strutil.IsNotBlank(v1.Metric[v3]) {
							f.InstanceId = f.InstanceId + " - " + v1.Metric[v3]
						}
					}
					reportForm = append(reportForm, f)
				}
			}
		} else {
			maxResult := make(map[string]form.PrometheusResult)
			minResult := make(map[string]form.PrometheusResult)
			avgResult := make(map[string]form.PrometheusResult)
			start := param.Start + (86400 - (param.Start-57600)%86400)
			end := param.End + (86400 - (param.End-57600)%86400)
			for _, v1 := range param.Statistics {
				expr := fmt.Sprintf("%s_over_time((%s)[1d:1m])", v1, pql)
				result = s.prometheus.QueryRange(expr, strconv.Itoa(start), strconv.Itoa(end), "86400").Data.Result
				if v1 == "max" {
					for _, v2 := range result {
						maxResult[v2.Metric["instance"]] = v2
					}
				}
				if v1 == "min" {
					for _, v2 := range result {
						minResult[v2.Metric["instance"]] = v2
					}
				}
				if v1 == "avg" {
					for _, v2 := range result {
						avgResult[v2.Metric["instance"]] = v2
					}
				}
			}
			for _, v1 := range result {
				for i, v2 := range v1.Values {
					f := form.ReportForm{
						Region:       param.RegionCode,
						InstanceName: instanceMap[v1.Metric["instance"]].InstanceName,
						InstanceId:   v1.Metric["instance"],
						Status:       instanceMap[v1.Metric["instance"]].Status,
						ItemName:     item.Name,
						Time:         util.TimestampToDayTimeFmtStr(int64(v2[0].(float64)) - 1),
						Timestamp:    int64(v2[0].(float64) - 1),
					}
					if len(maxResult[v1.Metric["instance"]].Values) != 0 {
						f.MaxValue = changeDecimal(maxResult[v1.Metric["instance"]].Values[i][1].(string))
					}
					if len(minResult[v1.Metric["instance"]].Values) != 0 {
						f.MinValue = changeDecimal(minResult[v1.Metric["instance"]].Values[i][1].(string))
					}
					if len(avgResult[v1.Metric["instance"]].Values) != 0 {
						f.AvgValue = changeDecimal(avgResult[v1.Metric["instance"]].Values[i][1].(string))
					}
					for _, v3 := range labels {
						if v3 != "instance" && strutil.IsNotBlank(v1.Metric[v3]) {
							f.InstanceId = f.InstanceId + " - " + v1.Metric[v3]
						}
					}
					reportForm = append(reportForm, f)
				}
			}
		}
	}
	return reportForm, nil
}

func (s *ReportFormService) Export(param form.ReportFormParam, userInfo string) error {
	url := config.Cfg.Common.AsyncExportApi
	asyncParams := []form.AsyncExportParam{
		{
			SheetSeq:   0,
			SheetName:  "云资源监控",
			SheetParam: jsonutil.ToString(param),
		},
	}
	asyncRequest := form.AsyncExportRequest{
		TemplateId: "cloud_monitor_manager",
		Params:     asyncParams,
	}
	header := map[string]string{"user-info": userInfo}
	result, err := httputil.HttpPostJson(url, asyncRequest, header)
	logger.Logger().Infof("AsyncExport：%v", result)
	if err != nil {
		logger.Logger().Infof("AsyncExportError：%v", err)
		return errors.NewBusinessError("异步导出API调用失败")
	}
	return nil
}
