package v1_0

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/external"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/openapi"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResourceCtl struct {
}

func NewResourceController() *ResourceCtl {
	return &ResourceCtl{}
}
func (ctl *ResourceCtl) GetResourceList(c *gin.Context) {
	param := openapi.NewPageQuery()
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(openapi.GetErrorCode(err), c))
		return
	}
	productAbbreviation := c.Param("ProductAbbreviation")
	tenantId := c.Param("tenantId")
	f := service.InstancePageForm{Product: productAbbreviation, TenantId: tenantId, PageSize: param.PageSize, Current: param.PageNumber}
	instanceService := external.ProductInstanceServiceMap[f.Product]
	if instanceService == nil {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(openapi.ProductAbbreviationInvalid, c))
		return
	}
	page, err := instanceService.GetPage(f, instanceService.(service.InstanceStage))
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.NewRespError(openapi.SystemError, c))
		return
	}
	var resources []ResourceInfo
	if records, ok := page.Records.([]service.InstanceCommonVO); ok {
		for _, record := range records {
			resources = append(resources, ResourceInfo{
				ResourceId:   record.InstanceId,
				ResourceName: record.InstanceName,
			})
		}
	}

	resourcePage := ResourcePage{
		ResCommonPage: &openapi.ResCommonPage{
			ResCommon: openapi.ResCommon{
				RequestId: openapi.GetRequestId(c),
			},
			TotalCount: page.Total,
			PageSize:   page.Size,
			PageNumber: page.Current,
		},
		Resources: resources,
	}
	c.JSON(http.StatusOK, resourcePage)
}

type ResourcePage struct {
	*openapi.ResCommonPage
	Resources []ResourceInfo
}

type ResourceInfo struct {
	ResourceId   string
	ResourceName string
}
