package web

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/actuator"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

const pathPrefix = "/hawkeye/"

func loadRouters() {
	actuatorMapping()
	monitorProductRouters()
	monitorItemRouters()
	instance()
	MonitorReportForm()
}

func actuatorMapping() {
	group := Router.Group("/actuator")
	{
		group.GET("/env", func(c *gin.Context) {
			c.JSON(http.StatusOK, actuator.Env())
		})
		group.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, actuator.Info())
		})
		group.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, actuator.Health())
		})
		group.GET("/metrics", func(c *gin.Context) {
			c.JSON(http.StatusOK, actuator.Metrics())
		})
	}
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
