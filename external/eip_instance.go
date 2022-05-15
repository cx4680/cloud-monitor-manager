package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"strconv"
)

type EipInstanceService struct {
	service.InstanceServiceImpl
}

type EipRequest struct {
	CloudProductCode string   `json:"cloudProductCode"`
	ResourceTypeCode string   `json:"resourceTypeCode"`
	ResourceId       string   `json:"resourceId"`
	StatusList       []string `json:"statusList"`
	RegionCode       string   `json:"regionCode"`
	Name             string   `json:"name"`
	TenantId         string   `json:"tenantId"`
	PageSize         string   `json:"pageSize"`
	CurrPage         string   `json:"currPage"`
}

type EipResponse struct {
	Code    string  `json:"code"`
	Msg     string  `json:"msg"`
	TraceId string  `json:"traceId"`
	Data    EipPage `json:"data"`
}

type EipPage struct {
	Total int        `json:"total"`
	List  []*EipList `json:"list"`
}

type EipList struct {
	Id               int    `json:"id"`
	UuidStr          string `json:"uuidStr"`
	RegionCode       string `json:"regionCode"`
	RegionName       string `json:"regionName"`
	ResourceTypeCode string `json:"resourceTypeCode"`
	CloudProductCode string `json:"cloudProductCode"`
	TenantId         string `json:"tenantId"`
	TenantName       string `json:"tenantName"`
	ResourceId       string `json:"resourceId"`
	ResourceName     string `json:"resourceName"`
	OrderId          string `json:"orderId"`
	ResourceUrl      string `json:"resourceUrl"`
	AvailabilityZone string `json:"availabilityZone"`
	Status           int    `json:"status"`
	StatusDesc       string `json:"statusDesc"`
	Deleted          int    `json:"deleted"`
	CreateTime       string `json:"createTime"`
	UpdateTime       string `json:"updateTime"`
	Additional       string `json:"additional"`
	ResCreateTime    string `json:"resCreateTime"`
	ResUpdateTime    string `json:"resUpdateTime"`
	Creator          string `json:"creator"`
	Modifier         string `json:"modifier"`
}

type EipAdditional struct {
	BandWidth struct {
		BandwidthId   string `json:"bandwidthId"`
		BandWidthSize string `json:"BandWidthSize"`
	} `json:"bandWidth"`
	BindInstanceId string `json:"bindInstanceId"`
	EipIpAddress   string `json:"eipIpAddress"`
}

func (ecs *EipInstanceService) ConvertRealForm(f service.InstancePageForm) interface{} {
	param := EipRequest{
		CloudProductCode: f.Product,
		ResourceTypeCode: "instance",
		ResourceId:       f.InstanceId,
		Name:             f.InstanceName,
		RegionCode:       f.RegionCode,
		TenantId:         f.TenantId,
		StatusList:       toStringList(f.StatusList),
		CurrPage:         strconv.Itoa(f.Current),
		PageSize:         strconv.Itoa(f.PageSize),
	}
	return param
}

func (ecs *EipInstanceService) DoRequest(url string, f interface{}) (interface{}, error) {
	respStr, err := httputil.HttpPostJson(url, f, nil)
	if err != nil {
		return nil, err
	}
	var resp EipResponse
	jsonutil.ToObject(respStr, &resp)
	return resp, nil
}

func (ecs *EipInstanceService) ConvertResp(realResp interface{}) (int, []service.InstanceCommonVO) {
	response := realResp.(EipResponse)
	var list []service.InstanceCommonVO
	if response.Data.Total > 0 {
		for _, d := range response.Data.List {
			var eipAdditional = &EipAdditional{}
			jsonutil.ToObject(d.Additional, eipAdditional)
			list = append(list, service.InstanceCommonVO{
				InstanceId:   d.ResourceId,
				InstanceName: d.ResourceName,
				Labels: []service.InstanceLabel{{
					Name:  "status",
					Value: d.StatusDesc,
				}, {
					Name:  "eipAddress",
					Value: eipAdditional.EipIpAddress,
				}, {
					Name:  "bandWidth",
					Value: eipAdditional.BandWidth.BandWidthSize,
				}, {
					Name:  "bindInstanceId",
					Value: eipAdditional.BindInstanceId,
				}},
			})
		}
	}
	return response.Data.Total, list
}
