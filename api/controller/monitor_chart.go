package controller

import (
	"net/http"
	"strings"

	"code.cestc.cn/ccos-ops/oplog"
	"github.com/gin-gonic/gin"

	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/errors"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/external"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/validator/translate"
)

type MonitorChartCtl struct {
	service *service.MonitorChartService
}

func NewMonitorChartController() *MonitorChartCtl {
	return &MonitorChartCtl{
		service: service.NewMonitorChartService(),
	}
}

func (ctl *MonitorChartCtl) GetData(c *gin.Context) {
	var param form.PrometheusRequest
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	c.Set(oplog.ResourceName, param.Instance)
	data, err := ctl.service.GetData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (ctl *MonitorChartCtl) GetRangeData(c *gin.Context) {
	var param = form.PrometheusRequest{Step: 60}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	c.Set(oplog.ResourceName, param.Instance)
	data, err := ctl.service.GetRangeData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (ctl *MonitorChartCtl) GetTopData(c *gin.Context) {
	var param = form.PrometheusRequest{TopNum: 5}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	if param.TopNum <= 0 {
		c.JSON(http.StatusBadRequest, global.NewError("topNum参数错误"))
		return
	}
	if strutil.IsBlank(param.Name) {
		c.JSON(http.StatusBadRequest, global.NewError("监控指标不能为空"))
		return
	}
	monitorItem := dao.MonitorItem.GetMonitorItemCacheByMetricCode(param.Name)
	var instanceIdList []string
	if len(strings.Split(monitorItem.Labels, ",")) > 1 {
		instanceIdList, err = getInstanceList(monitorItem.ProductAbbreviation)
		if err != nil {
			c.JSON(http.StatusBadRequest, global.NewError(err.Error()))
			return
		}
	}
	data, err := ctl.service.GetTopData(param, instanceIdList, monitorItem)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (ctl *MonitorChartCtl) GetProcessData(c *gin.Context) {
	var param = form.PrometheusRequest{Step: 60}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	data, err := ctl.service.GetProcessData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (ctl *MonitorChartCtl) GetPrometheusData(c *gin.Context) {
	var param = form.PrometheusRequest{Step: 60}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	data, err := ctl.service.GetPrometheusData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (ctl *MonitorChartCtl) PostPrometheusData(c *gin.Context) {
	var param = form.PrometheusRequest{Step: 60}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	data, err := ctl.service.PostPrometheusData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (ctl *MonitorChartCtl) GetPrometheusRangeData(c *gin.Context) {
	var param = form.PrometheusRequest{Step: 60}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	data, err := ctl.service.GetPrometheusRangeData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

// 获取实例ID列表
func getInstanceList(product string) ([]string, error) {
	instanceService := external.ProductInstanceServiceMap[product]
	page, err := instanceService.GetPage(service.InstancePageForm{
		Current:  1,
		PageSize: 9999,
	}, instanceService.(service.InstanceStage))
	if err != nil {
		return nil, errors.NewBusinessError("获取监控产品服务失败")
	}
	var instanceList []string
	for _, v := range page.Records.([]service.InstanceCommonVO) {
		instanceList = append(instanceList, v.InstanceId)
	}
	return instanceList, nil
}
