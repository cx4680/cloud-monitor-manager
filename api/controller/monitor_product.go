package controller

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MonitorProductCtl struct {
	service *service.MonitorProductService
}

func NewMonitorProductCtl() *MonitorProductCtl {
	return &MonitorProductCtl{service.NewMonitorProductService()}
}

func (mpc *MonitorProductCtl) GetMonitorProduct(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", mpc.service.GetMonitorProduct()))
}

func (mpc *MonitorProductCtl) GetAllMonitorProduct(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", mpc.service.GetMonitorProduct()))
}
