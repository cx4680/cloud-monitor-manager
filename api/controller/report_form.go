package controller

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/validator/translate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReportFormCtl struct {
	service *service.ReportFormService
}

func NewReportFormController() *ReportFormCtl {
	return &ReportFormCtl{
		service: service.NewReportFormService(),
	}
}

func (mrc *ReportFormCtl) GetMonitorData(c *gin.Context) {
	var callback = form.CallbackReportForm{}
	err := c.ShouldBindJSON(&callback)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	var param = form.ReportFormParam{}
	jsonutil.ToObject(callback.Param, &param)
	param.RegionCode = config.Cfg.Common.RegionName
	if len(param.InstanceList) == 0 {
		c.JSON(http.StatusBadRequest, global.NewError("实例不能为空"))
		return
	}
	if len(param.ItemList) == 0 {
		c.JSON(http.StatusBadRequest, global.NewError("监控指标不能为空"))
		return
	}
	if param.Start == 0 || param.End == 0 || param.Start > param.End {
		c.JSON(http.StatusBadRequest, global.NewError("时间参数有误"))
		return
	}
	//用监控项数组的下标充当分页
	param.Current = callback.Current - 1
	result, err := mrc.service.GetMonitorData(param)
	if err == nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":      http.StatusOK,
			"message":   "success",
			"pageCount": len(param.ItemList),
			"result":    result,
		})
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}

func (mrc *ReportFormCtl) Export(c *gin.Context) {
	var param = form.ReportFormParam{Step: 60}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.NewError(translate.GetErrorMsg(err)))
		return
	}
	if len(param.InstanceList) == 0 {
		c.JSON(http.StatusBadRequest, global.NewError("实例ID不能为空"))
		return
	}
	if len(param.ItemList) == 0 {
		c.JSON(http.StatusBadRequest, global.NewError("监控指标不能为空"))
		return
	}
	if param.Start == 0 || param.End == 0 || param.Start > param.End {
		c.JSON(http.StatusBadRequest, global.NewError("时间参数有误"))
		return
	}
	param.RegionCode = config.Cfg.Common.RegionName
	err = mrc.service.Export(param, c.Request.Header.Get("user-info"))
	if err == nil {
		c.JSON(http.StatusOK, global.NewSuccess("导入任务已下发", true))
	} else {
		c.JSON(http.StatusOK, global.NewError(err.Error()))
	}
}
