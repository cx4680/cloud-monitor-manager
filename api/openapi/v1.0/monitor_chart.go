package v1_0

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/external"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/openapi"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type MonitorChartCtl struct {
	service *service.MonitorChartService
}

func NewMonitorChartController() *MonitorChartCtl {
	return &MonitorChartCtl{service.NewMonitorChartService()}
}

func (ctl *MonitorChartCtl) GetMonitorDatas(c *gin.Context) {
	resourceId := c.Param("ResourceId")
	metricCode := c.Param("MetricCode")
	var param = MonitorDataParam{Step: 60}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(openapi.GetErrorCode(err), c))
		return
	}
	c.Set(global.ResourceName, resourceId)
	if param.StartTime == 0 || param.EndTime == 0 || param.StartTime > param.EndTime {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(openapi.TimeParameterError, c))
		return
	}
	monitorItem := dao.MonitorItem.GetMonitorItemByMetricCode(metricCode)
	//校验参数
	errCode := checkParam(resourceId, monitorItem.Code)
	if errCode != nil {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(errCode, c))
		return
	}
	//查询Prometheus
	pql := strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+resourceId+"',"+constant.FILTER)
	prometheusResult := service.NewPrometheusService().QueryRange(pql, strconv.Itoa(param.StartTime), strconv.Itoa(param.EndTime), strconv.Itoa(param.Step)).Data.Result
	//构建数据
	var timeList []int
	if len(prometheusResult) == 0 {
		timeList = getTimeList(param.StartTime, param.EndTime, param.Step, param.StartTime)
	} else {
		timeList = getTimeList(param.StartTime, param.EndTime, param.Step, int(prometheusResult[0].Values[0][0].(float64)))
	}
	label := getLabel(monitorItem.Labels)
	result := MonitorRangeData{
		RequestId:           openapi.GetRequestId(c),
		MetricCode:          metricCode,
		ProductAbbreviation: monitorItem.ProductAbbreviation,
		Times:               timeList,
		StartTime:           util.TimestampToFullTimeFmtStr(int64(param.StartTime)),
		EndTime:             util.TimestampToFullTimeFmtStr(int64(param.EndTime)),
		Step:                param.Step,
		Dimension:           label,
		Points:              pointsFillEmptyRangeData(prometheusResult, timeList, label, resourceId),
	}
	c.JSON(http.StatusOK, result)
}

func (ctl *MonitorChartCtl) GetMonitorData(c *gin.Context) {
	resourceId := c.Param("ResourceId")
	metricCode := c.Param("MetricCode")
	c.Set(global.ResourceName, resourceId)
	monitorItem := dao.MonitorItem.GetMonitorItemByMetricCode(metricCode)
	//校验参数
	errCode := checkParam(resourceId, monitorItem.Code)
	if errCode != nil {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(errCode, c))
		return
	}
	//查询Prometheus
	pql := strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, constant.INSTANCE+"='"+resourceId+"',"+constant.FILTER)
	prometheusResult := service.NewPrometheusService().Query(pql, "").Data.Result
	label := getLabel(monitorItem.Labels)
	result := MonitorData{
		RequestId:           openapi.GetRequestId(c),
		MetricCode:          metricCode,
		ProductAbbreviation: monitorItem.ProductAbbreviation,
		Dimension:           label,
		CurrentTime:         util.GetNowStr(),
		Points:              pointsFillEmptyData(prometheusResult, label, resourceId),
	}
	c.JSON(http.StatusOK, result)
}

func (ctl *MonitorChartCtl) GetMonitorDataTop(c *gin.Context) {
	metricCode := c.Param("MetricCode")
	n := c.Param("N")
	i, err := strconv.Atoi(n)
	if err != nil || i <= 0 {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(openapi.InvalidParameter, c))
		return
	}
	monitorItem := dao.MonitorItem.GetMonitorItemByMetricCode(metricCode)
	if strutil.IsBlank(monitorItem.Code) {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(openapi.MetricCodeInvalid, c))
		return
	}
	//查询Prometheus
	pql := fmt.Sprintf(constant.TopExpr, n, strings.ReplaceAll(monitorItem.Expression, constant.MetricLabel, ""))
	prometheusResult := service.NewPrometheusService().Query(pql, "").Data.Result
	label := getLabel(monitorItem.Labels)
	result := MonitorTopData{
		RequestId:           openapi.GetRequestId(c),
		MetricCode:          metricCode,
		ProductAbbreviation: monitorItem.ProductAbbreviation,
		Dimension:           label,
		CurrentTime:         util.GetNowStr(),
		Tops:                []Top{},
	}
	for _, v := range prometheusResult {
		result.Tops = append(result.Tops, Top{ResourceId: v.Metric[constant.INSTANCE], Value: changeDecimal(v.Value[1].(string))})
	}
	c.JSON(http.StatusOK, result)
}

//检验参数
func checkParam(resourceId, metricCode string) *openapi.ErrorCode {
	if strutil.IsBlank(resourceId) {
		return openapi.MissingResource
	}
	if strutil.IsBlank(metricCode) {
		return openapi.MetricCodeInvalid
	}
	return nil
}

//获取实例ID列表
func getInstances(product string) string {
	f := service.InstancePageForm{
		Product:  product,
		Current:  1,
		PageSize: 10000,
	}
	instanceService := external.ProductInstanceServiceMap[f.Product]
	stage, _ := instanceService.(service.InstanceStage)
	page, err := instanceService.GetPage(f, stage)
	if err != nil {
		return ""
	}
	var instanceList []string
	for _, v := range page.Records.([]service.InstanceCommonVO) {
		instanceList = append(instanceList, v.InstanceId)
	}
	return strings.Join(instanceList, "|")
}

func getLabel(labels string) string {
	var label string
	for _, v := range strings.Split(labels, ",") {
		if strutil.IsBlank(label) || v != "instance" {
			label = v
		}
	}
	return label
}

//区间数据时间点
func getTimeList(start, end, step, firstTime int) []int {
	var timeList []int
	if start > end {
		return timeList
	}
	for firstTime-step >= start {
		firstTime -= step
	}
	for firstTime <= end {
		timeList = append(timeList, firstTime)
		firstTime += step
	}
	return timeList
}

//构建区间监控数据，未采集到数据则补null
func pointsFillEmptyRangeData(result []form.PrometheusResult, timeList []int, label, resourceId string) []RangePoint {
	var points []RangePoint
	if len(result) == 0 {
		return []RangePoint{}
	}
	for _, v := range result {
		var point RangePoint
		timeMap := map[int]string{}
		for _, b := range v.Values {
			timeMap[int(b[0].(float64))] = b[1].(string)
		}
		for _, n := range timeList {
			point.Values = append(point.Values, changeDecimal(timeMap[n]))
		}
		if strutil.IsBlank(v.Metric[label]) {
			point.Name = resourceId
		} else {
			point.Name = v.Metric[label]
		}
		points = append(points, point)
	}
	return points
}

//构建瞬时监控数据
func pointsFillEmptyData(result []form.PrometheusResult, label, resourceId string) []Point {
	var points []Point
	if len(result) == 0 {
		return []Point{}
	}
	for _, v := range result {
		var point Point
		point.Value = changeDecimal(v.Value[1].(string))
		if strutil.IsBlank(v.Metric[label]) {
			point.Name = resourceId
		} else {
			point.Name = v.Metric[label]
		}
		points = append(points, point)
	}
	return points
}

//数据保留2位小数
func changeDecimal(value string) string {
	if strutil.IsBlank(value) {
		return ""
	}
	v, _ := strconv.ParseFloat(value, 64)
	return fmt.Sprintf("%.2f", v)
}

type MonitorDataParam struct {
	StartTime int `json:"StartTime"`
	EndTime   int `json:"EndTime"`
	Step      int `json:"Step"`
}

type MonitorRangeData struct {
	RequestId           string       `json:"RequestId"`
	MetricCode          string       `json:"MetricCode"`
	ProductAbbreviation string       `json:"ProductAbbreviation"`
	Times               []int        `json:"Times"`
	StartTime           string       `json:"StartTime"`
	EndTime             string       `json:"EndTime"`
	Step                int          `json:"Step"`
	Dimension           string       `json:"Dimension"`
	Points              []RangePoint `json:"Points"`
}

type RangePoint struct {
	Name   string   `json:"Name"`
	Values []string `json:"Values"`
}

type MonitorData struct {
	RequestId           string  `json:"RequestId"`
	MetricCode          string  `json:"MetricCode"`
	ProductAbbreviation string  `json:"ProductAbbreviation"`
	Dimension           string  `json:"Dimension"`
	CurrentTime         string  `json:"CurrentTime"`
	Points              []Point `json:"Points"`
}

type Point struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type MonitorTopData struct {
	RequestId           string `json:"RequestId"`
	MetricCode          string `json:"MetricCode"`
	ProductAbbreviation string `json:"ProductAbbreviation"`
	Dimension           string `json:"Dimension"`
	CurrentTime         string `json:"CurrentTime"`
	Tops                []Top  `json:"Tops"`
}

type Top struct {
	ResourceId string `json:"ResourceId"`
	Value      string `json:"Value"`
}
