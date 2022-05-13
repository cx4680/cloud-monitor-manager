package web

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/openapi/v1.0"
	"github.com/gin-gonic/gin"
)

func loadOpenApiV1Routers() {
	group := Router.Group("/v1.0/")
	monitorProductOpenApiV1Routers(group)
	monitorItemOpenApiV1Routers(group)
	MonitorReportOpenApiV1Routers(group)
	ResourceOpenApiV1Routers(group)
}

func monitorProductOpenApiV1Routers(group *gin.RouterGroup) {
	monitorProductCtl := v1_0.NewMonitorProductCtl()
	group.GET("products", monitorProductCtl.GetMonitorProduct)
}

func monitorItemOpenApiV1Routers(group *gin.RouterGroup) {
	monitorItemCtl := v1_0.NewMonitorItemCtl()
	group.GET("products/:ProductAbbreviation/metrics", monitorItemCtl.GetMonitorItemsByProductAbbr)
}

func MonitorReportOpenApiV1Routers(group *gin.RouterGroup) {
	monitorReportFormCtl := v1_0.NewMonitorReportFormController()
	group.GET("resources/:ResourceId/metrics/:MetricCode/datas", monitorReportFormCtl.GetMonitorDatas)
	group.GET("resources/:ResourceId/metrics/:MetricCode/data", monitorReportFormCtl.GetMonitorData)
	group.GET("metrics/:MetricCode/:N/resources", monitorReportFormCtl.GetMonitorDataTop)
}

func ResourceOpenApiV1Routers(group *gin.RouterGroup) {
	resourceCtl := v1_0.NewResourceController()
	group.GET(":ProductAbbreviation/resources", resourceCtl.GetResourceList)
}
