package web

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/api/openapi/v1.0"
	"code.cestc.cn/ccos-ops/cloud-monitor/pkg/business-common/global/iam"
	"code.cestc.cn/ccos-ops/cloud-monitor/pkg/business-common/global/logs"
	"code.cestc.cn/yyptb-group_tech/iam-sdk-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func loadOpenApiV1Routers() {
	group := Router.Group("/v1.0/")
	monitorProductOpenApiV1Routers(group)
	monitorItemOpenApiV1Routers(group)
	MonitorReportOpenApiV1Routers(group)
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

func instanceOpenApiRouters(group *gin.RouterGroup) {
	ctl := v1_0.NewInstanceCtl()
	group.GET("resources/:ResourceId/rules", logs.GinTrailzap(false, Read, logs.INFO, logs.Resource), iam.AuthIdentify(&models.Identity{Product: iam.ProductMonitor, Action: "GetInstanceRulePageList", ResourceType: "*", ResourceId: "*"}), ctl.Page)
	group.DELETE("resources/:ResourceId/rules", logs.GinTrailzap(false, Write, logs.Warn, logs.Resource), iam.AuthIdentify(&models.Identity{Product: iam.ProductMonitor, Action: "UnbindInstanceRule", ResourceType: "*", ResourceId: "*"}), ctl.Unbind)
	group.PUT("resources/:ResourceId/rules", logs.GinTrailzap(false, Write, logs.Warn, logs.Resource), iam.AuthIdentify(&models.Identity{Product: iam.ProductMonitor, Action: "BindInstanceRule", ResourceType: "*", ResourceId: "*"}), ctl.Bind)
}
