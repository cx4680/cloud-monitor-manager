package web

import (
	"net/http"

	"code.cestc.cn/ccos-ops/oplog"
	"github.com/gin-gonic/gin"

	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/actuator"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/controller"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
)

const pathPrefix = "/api/cmm/"

const (
	ServiceName = "cloudMonitorManager"
	ApiVersion  = "1.0"

	RequestTypeRead  = "Read"
	RequestTypeWrite = "Write"
)

func NewV1OperatorInfo(api, name, requestType, level string) *oplog.OperatorInfo {
	return &oplog.OperatorInfo{
		EventRequestInfo: oplog.EventRequestInfo{
			ServiceName: ServiceName,
			EventApi:    api,
			EventName:   name,
			RequestType: requestType,
			ApiVersion:  ApiVersion,
			EventLevel:  level,
			EventRegion: config.Cfg.Common.RegionName,
			Utc:         false,
		},
		ResourceInfo: oplog.ResourceInfo{
			ResourceName: "",
		},
	}
}

func loadRouters() {
	inner()
	actuatorMapping()
	monitorProductRouters()
	monitorItemRouters()
	instance()
	monitorChart()
	reportForm()
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

func inner() {
	monitorChartCtl := controller.NewMonitorChartController()
	reportFormCtl := controller.NewReportFormController()
	instanceCtl := controller.NewInstanceCtl()
	group := Router.Group(pathPrefix + "inner")
	{
		group.GET("/monitorChart/getPrometheusData", monitorChartCtl.GetPrometheusData)
		group.POST("/monitorChart/postPrometheusData", monitorChartCtl.PostPrometheusData)
		group.GET("/monitorChart/getPrometheusRangeData", monitorChartCtl.GetPrometheusRangeData)
		group.POST("/reportForm/getMonitorData", reportFormCtl.GetMonitorData)
		group.GET("/instance/page", instanceCtl.GetPage)
	}
}

func monitorProductRouters() {
	monitorProductCtl := controller.NewMonitorProductCtl()
	group := Router.Group(pathPrefix + "monitorProduct/")
	{
		group.GET("/getMonitorProduct", oplog.GinTrail(NewV1OperatorInfo("GetMonitorProductsList", "获取监控云产品列表", RequestTypeRead, oplog.INFO)), monitorProductCtl.GetMonitorProduct)
	}
}

func monitorItemRouters() {
	monitorItemCtl := controller.NewMonitorItemCtl()
	group := Router.Group(pathPrefix + "monitorItem/")
	{
		group.GET("/getMonitorItemByProductBizId", oplog.GinTrail(NewV1OperatorInfo("GetMonitorItemsByIdList", "获取显示的监控项列表", RequestTypeRead, oplog.INFO)), monitorItemCtl.GetMonitorItemByProductBizId)
		group.GET("/getAllMonitorItemByProductBizId", oplog.GinTrail(NewV1OperatorInfo("GetAllMonitorItemsByIdList", "获取产品的所有监控项列表", RequestTypeRead, oplog.INFO)), monitorItemCtl.GetAllMonitorItemByProductBizId)
		group.POST("/openDisplay", oplog.GinTrail(NewV1OperatorInfo("OpenDisplay", "显示监控项", RequestTypeWrite, oplog.Warn)), monitorItemCtl.OpenDisplay)
	}
}

func monitorChart() {
	monitorChartCtl := controller.NewMonitorChartController()
	group := Router.Group(pathPrefix + "monitorChart/")
	{
		group.GET("/getData", oplog.GinTrail(NewV1OperatorInfo("GetMonitorChartData", "获取瞬时监控数据", RequestTypeRead, oplog.INFO)), monitorChartCtl.GetData)
		group.GET("/getRangeData", oplog.GinTrail(NewV1OperatorInfo("GetMonitorChartRangeData", "获取区间监控数据", RequestTypeRead, oplog.INFO)), monitorChartCtl.GetRangeData)
		group.GET("/getTopData", oplog.GinTrail(NewV1OperatorInfo("GetMonitorChartTop", "获取监控Top数据", RequestTypeRead, oplog.INFO)), monitorChartCtl.GetTopData)
		group.GET("/getProcessData", oplog.GinTrail(NewV1OperatorInfo("GetMonitorChartProcessData", "获取进程监控数据", RequestTypeRead, oplog.INFO)), monitorChartCtl.GetProcessData)
	}
}

func reportForm() {
	monitorChartCtl := controller.NewReportFormController()
	group := Router.Group(pathPrefix + "reportForm/")
	{
		group.POST("/getMonitorData", oplog.GinTrail(NewV1OperatorInfo("GetMonitorData", "获取报表导出数据", RequestTypeRead, oplog.INFO)), monitorChartCtl.GetMonitorData)
		group.POST("/export", oplog.GinTrail(NewV1OperatorInfo("Export", "报表导出", RequestTypeRead, oplog.INFO)), monitorChartCtl.Export)
	}
}

func instance() {
	instanceCtl := controller.NewInstanceCtl()
	group := Router.Group(pathPrefix + "instance/")
	{
		group.GET("/page", oplog.GinTrail(NewV1OperatorInfo("GetInstancePageList", "获取监控实例列表", RequestTypeRead, oplog.INFO)), instanceCtl.GetPage)
	}
}

func configItemRouters() {
	ctl := controller.NewConfigItemCtl()
	group := Router.Group(pathPrefix + "configItem/")
	{
		group.GET("/getMonitorRange", oplog.GinTrail(NewV1OperatorInfo("GetMonitorRangeList", "获取监控查询步长", RequestTypeRead, oplog.INFO)), ctl.GetMonitorRange)
	}
}
