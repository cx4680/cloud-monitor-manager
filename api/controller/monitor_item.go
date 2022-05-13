package controller

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/validator/translate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MonitorItemCtl struct {
	service *service.MonitorItemService
}

func NewMonitorItemCtl() *MonitorItemCtl {
	return &MonitorItemCtl{service.NewMonitorItemService()}
}

var displayList = []string{"chart", "rule", "scaling"}

func (mic *MonitorItemCtl) GetMonitorItemByProductId(c *gin.Context) {
	var param form.MonitorItemParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	c.Set(global.ResourceName, param.ProductBizId)
	if strutil.IsNotBlank(param.Display) && !checkDisplay(param.Display) {
		c.JSON(http.StatusOK, global.NewError("查询失败，展示参数错误"))
		return
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", mic.service.GetMonitorItemByProductId(param)))
}

func checkDisplay(display string) bool {
	for _, v := range displayList {
		if display == v {
			return true
		}
	}
	return false
}
