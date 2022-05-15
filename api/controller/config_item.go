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

func (ctl *ConfigItemCtl) GetStatisticalPeriodList(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", dao.ConfigItem.GetConfigItemList(dao.StatisticalPeriodPid)))
}

func (ctl *ConfigItemCtl) GetContinuousCycleList(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", dao.ConfigItem.GetConfigItemList(dao.ContinuousCyclePid)))
}
func (ctl *ConfigItemCtl) GetStatisticalMethodsList(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", dao.ConfigItem.GetConfigItemList(dao.StatisticalMethodsPid)))
}
func (ctl *ConfigItemCtl) GetComparisonMethodList(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", dao.ConfigItem.GetConfigItemList(dao.ComparisonMethodPid)))
}
func (ctl *ConfigItemCtl) GetOverviewItemList(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", dao.ConfigItem.GetConfigItemList(dao.OverviewItemPid)))
}
func (ctl *ConfigItemCtl) GetMonitorRange(c *gin.Context) {
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", dao.ConfigItem.GetConfigItemList(dao.MonitorRange)))
}
