package web

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/openapi/v1.0"
	"github.com/gin-gonic/gin"
)

func loadOpenApiV1Routers() {
	group := Router.Group("/v1.0/")
	monitorProductOpenApiV1Routers(group)
	monitorItemOpenApiV1Routers(group)
	monitorChartOpenApiV1Routers(group)
	resourceOpenApiV1Routers(group)
}

func monitorProductOpenApiV1Routers(group *gin.RouterGroup) {
	monitorProductCtl := v1_0.NewMonitorProductCtl()
	group.GET("products", monitorProductCtl.GetMonitorProduct)
}

func monitorItemOpenApiV1Routers(group *gin.RouterGroup) {
	monitorItemCtl := v1_0.NewMonitorItemCtl()
	group.GET("products/:ProductAbbreviation/metrics", monitorItemCtl.GetMonitorItemsByProductAbbr)
}

func monitorChartOpenApiV1Routers(group *gin.RouterGroup) {
	monitorChartCtl := v1_0.NewMonitorChartController()
	group.GET("resources/:ResourceId/metrics/:MetricCode/datas", monitorChartCtl.GetMonitorDatas)
	group.GET("resources/:ResourceId/metrics/:MetricCode/data", monitorChartCtl.GetMonitorData)
	group.GET("metrics/:MetricCode/:N/resources", monitorChartCtl.GetMonitorDataTop)
}

func resourceOpenApiV1Routers(group *gin.RouterGroup) {
	resourceCtl := v1_0.NewResourceController()
	group.GET(":ProductAbbreviation/resources", resourceCtl.GetResourceList)
}
