package web

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/controller"
)

const pathPrefix = "/hawkeye/"

func loadRouters() {
	monitorProductRouters()
	monitorItemRouters()
	instance()
	MonitorReportForm()
}

func monitorProductRouters() {
	monitorProductCtl := controller.NewMonitorProductCtl()
	group := Router.Group(pathPrefix + "monitorProduct/")
	{
		group.GET("/getAllMonitorProduct", monitorProductCtl.GetMonitorProduct)
	}
}

func monitorItemRouters() {
	monitorItemCtl := controller.NewMonitorItemCtl()
	group := Router.Group(pathPrefix + "monitorItem/")
	{
		group.GET("/getMonitorItemByProductBizId", monitorItemCtl.GetMonitorItemsById)
	}
}

func MonitorReportForm() {
	monitorReportFormCtl := controller.NewMonitorReportFormController()
	group := Router.Group(pathPrefix + "MonitorReportForm/")
	{
		group.GET("/getData", monitorReportFormCtl.GetData)
		group.GET("/getRangeData", monitorReportFormCtl.GetRangeData)
		group.GET("/getTop", monitorReportFormCtl.GetTop)
	}
}

func instance() {
	instanceCtl := controller.NewInstanceCtl()
	group := Router.Group(pathPrefix + "instance/")
	{
		group.GET("/page", instanceCtl.GetPage)
	}
}
