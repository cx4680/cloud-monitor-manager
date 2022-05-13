package controller

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/validator/translate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MonitorChartCtl struct {
	service *service.MonitorChartService
}

func NewMonitorChartController() *MonitorChartCtl {
	return &MonitorChartCtl{
		service: service.NewMonitorChartService(),
	}
}

func (mpc *MonitorChartCtl) GetData(c *gin.Context) {
	var param form.PrometheusRequest
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	data, err := mpc.service.GetData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (mpc *MonitorChartCtl) GetRangeData(c *gin.Context) {
	var param = form.PrometheusRequest{Step: 60}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	data, err := mpc.service.GetRangeData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (mpc *MonitorChartCtl) GetTopData(c *gin.Context) {
	var param = form.PrometheusRequest{TopNum: "5"}
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	data, err := mpc.service.GetTopData(param)
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}
