package controller

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigItemCtl struct {
}

func NewConfigItemCtl() *ConfigItemCtl {
	return &ConfigItemCtl{}
}

func (ctl *ConfigItemCtl) GetMonitorRange(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", dao.ConfigItem.GetConfigItemList(dao.MonitorRange)))
}
