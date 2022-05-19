package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"strconv"
	"strings"
)

type EcsInstanceService struct {
	service.InstanceServiceImpl
}

type EcsRequest struct {
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

type EcsResponse struct {
	Code    string  `json:"code"`
	Msg     string  `json:"msg"`
	TraceId string  `json:"traceId"`
	Data    EcsPage `json:"data"`
}

type EcsPage struct {
	Total int        `json:"total"`
	List  []*EcsList `json:"list"`
}

type EcsList struct {
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

type EcsAdditional struct {
	SystemType string `json:"systemType"`
}

func (ecs *EcsInstanceService) ConvertRealForm(f service.InstancePageForm) interface{} {
	param := EcsRequest{
		CloudProductCode: "ECS",
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

func (ecs *EcsInstanceService) DoRequest(url string, f interface{}) (interface{}, error) {
	logger.Logger().Infof("form:%s", f.(EcsRequest))
	respStr, err := httputil.HttpPostJson(url, f, nil)
	if err != nil {
		return nil, err
	}
	var resp EcsResponse
	jsonutil.ToObject(respStr, &resp)
	return resp, nil
}

func (ecs *EcsInstanceService) ConvertResp(realResp interface{}) (int, []service.InstanceCommonVO) {
	response := realResp.(EcsResponse)
	var list []service.InstanceCommonVO
	if response.Data.Total > 0 {
		for _, d := range response.Data.List {
			var ecsAdditional = &EcsAdditional{}
			jsonutil.ToObject(d.Additional, ecsAdditional)
			list = append(list, service.InstanceCommonVO{
				InstanceId:   d.ResourceId,
				InstanceName: d.ResourceName,
				TenantId:     d.TenantId,
				Labels: []service.InstanceLabel{{
					Name:  "status",
					Value: d.StatusDesc,
				}, {
					Name:  "osType",
					Value: ecsAdditional.SystemType,
				}},
			})
		}
	}
	return response.Data.Total, list
}

func toStringList(s string) []string {
	statusList := strings.Split(s, ",")
	var list []string
	for _, v := range statusList {
		list = append(list, v)
	}
	return list
}
