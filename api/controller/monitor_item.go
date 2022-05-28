package controller

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
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

func (mic *MonitorItemCtl) GetMonitorItemByProductBizId(c *gin.Context) {
	var param form.MonitorItemParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	userId, err := util.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
		return
	}
	c.Set(global.ResourceName, param.ProductBizId)
	if strutil.IsNotBlank(param.Display) && !checkDisplay(param.Display) {
		c.JSON(http.StatusOK, global.NewError("查询失败，展示参数错误"))
		return
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", mic.service.GetMonitorItemByProductBizId(param, userId)))
}

func (mic *MonitorItemCtl) GetAllMonitorItemByProductBizId(c *gin.Context) {
	var param form.MonitorItemParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	c.Set(global.ResourceName, param.ProductBizId)
	if strutil.IsNotBlank(param.Display) && !checkDisplay(param.Display) {
		c.JSON(http.StatusOK, global.NewError("查询失败，展示参数错误"))
		return
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", mic.service.GetAllMonitorItemByProductBizId(param)))
}

func (mic *MonitorItemCtl) ChangeMonitorItemDisplay(c *gin.Context) {
	var param form.MonitorItemParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	if strutil.IsBlank(param.Active) {
		c.JSON(http.StatusBadRequest, global.NewError("active不能为空"))
	}
	userId, err := util.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
		return
	}
	if err := mic.service.ChangeMonitorItemDisplay(param, userId); err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(err.Error()))
	}
	c.JSON(http.StatusOK, global.NewSuccess("切换成功", true))
}

func checkDisplay(display string) bool {
	for _, v := range displayList {
		if display == v {
			return true
		}
	}
	return false
}
