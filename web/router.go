package web

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/actuator"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

const pathPrefix = "/api/cmm/"

func loadRouters() {
	actuatorMapping()
	monitorProductRouters()
	monitorItemRouters()
	instance()
	monitorChart()
	configItemRouters()
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
		group.GET("/getMonitorProduct", monitorProductCtl.GetMonitorProduct)
	}
}

func monitorItemRouters() {
	monitorItemCtl := controller.NewMonitorItemCtl()
	group := Router.Group(pathPrefix + "monitorItem/")
	{
		group.GET("/getMonitorItemByProductBizId", monitorItemCtl.GetMonitorItemByProductBizId)
		group.GET("/getAllMonitorItemByProductBizId", monitorItemCtl.GetAllMonitorItemByProductBizId)
		group.POST("/openDisplay", monitorItemCtl.OpenDisplay)
	}
}

func monitorChart() {
	monitorChartCtl := controller.NewMonitorChartController()
	group := Router.Group(pathPrefix + "monitorChart/")
	{
		group.GET("/getData", monitorChartCtl.GetData)
		group.GET("/getRangeData", monitorChartCtl.GetRangeData)
		group.GET("/getTopData", monitorChartCtl.GetTopData)
	}
}

func instance() {
	instanceCtl := controller.NewInstanceCtl()
	group := Router.Group(pathPrefix + "instance/")
	{
		group.GET("/page", instanceCtl.GetPage)
	}
}

func configItemRouters() {
	ctl := controller.NewConfigItemCtl()
	group := Router.Group(pathPrefix + "configItem/")
	{
		group.GET("/getMonitorRange", ctl.GetMonitorRange)
	}
}
