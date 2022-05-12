package v1_0

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/openapi"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MonitorProductCtl struct {
	service *service.MonitorProductService
}

func NewMonitorProductCtl() *MonitorProductCtl {
	return &MonitorProductCtl{service.NewMonitorProductService()}
}

func (mpc *MonitorProductCtl) GetMonitorProduct(c *gin.Context) {
	pageQuery := openapi.NewPageQuery()
	if err := c.ShouldBindQuery(pageQuery); err != nil {
		c.JSON(http.StatusBadRequest, openapi.NewRespError(openapi.GetErrorCode(err), c))
		return
	}
	productPageVo := mpc.service.GetMonitorProductPage(pageQuery.PageSize, pageQuery.PageNumber)
	var productMetaList []ProductMeta
	productListVo := productPageVo.Records.([]model.MonitorProduct)
	for _, productVo := range productListVo {
		productMeta := ProductMeta{
			Name:         productVo.Name,
			Abbreviation: productVo.Abbreviation,
			Description:  productVo.Description,
			MonitorType:  productVo.MonitorType,
		}
		productMetaList = append(productMetaList, productMeta)
	}
	page := ProductPage{
		ResCommonPage: *openapi.NewResCommonPage(c, productPageVo),
		Resources:     productMetaList,
	}
	c.JSON(http.StatusOK, page)
}

type ProductMeta struct {
	Name         string // 监控产品名称
	Description  string // 描述
	Abbreviation string // 简称
	MonitorType  string // 监控类型
}

type ProductPage struct {
	openapi.ResCommonPage
	Resources []ProductMeta
}
