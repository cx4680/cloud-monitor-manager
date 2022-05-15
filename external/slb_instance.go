package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"strconv"
	"strings"
)

type SlbInstanceService struct {
	service.InstanceServiceImpl
}

type SlbRequest struct {
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

type SlbResponse struct {
	Code    string  `json:"code"`
	Msg     string  `json:"msg"`
	TraceId string  `json:"traceId"`
	Data    SlbPage `json:"data"`
}

type SlbPage struct {
	Total int        `json:"total"`
	List  []*SlbList `json:"list"`
}

type SlbList struct {
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
	OsType           string `json:"osType"`
}

type SlbAdditional struct {
	Address   string `json:"address"`
	Spec      string `json:"spec"`
	Vpc       string `json:"vpc"`
	Container []struct {
	} `json:"container"`
	Listeners []struct {
		ListenerId   string `json:"listenerId"`
		ListenerName string `json:"listenerName"`
		Port         int    `json:"port"`
		Protocol     string `json:"protocol"`
	} `json:"listeners"`
}

func (slb *SlbInstanceService) ConvertRealForm(f service.InstancePageForm) interface{} {
	param := SlbRequest{
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

func (slb *SlbInstanceService) DoRequest(url string, f interface{}) (interface{}, error) {
	respStr, err := httputil.HttpPostJson(url, f, nil)
	if err != nil {
		return nil, err
	}
	var resp SlbResponse
	jsonutil.ToObject(respStr, &resp)
	return resp, nil
}

func (slb *SlbInstanceService) ConvertResp(realResp interface{}) (int, []service.InstanceCommonVO) {
	response := realResp.(SlbResponse)
	var list []service.InstanceCommonVO
	if response.Data.Total > 0 {
		for _, d := range response.Data.List {
			var slbAdditional = &SlbAdditional{}
			jsonutil.ToObject(d.Additional, slbAdditional)
			list = append(list, service.InstanceCommonVO{
				InstanceId:   d.ResourceId,
				InstanceName: d.ResourceName,
				Labels: []service.InstanceLabel{{
					Name:  "status",
					Value: d.StatusDesc,
				}, {
					Name:  "address",
					Value: slbAdditional.Address,
				}, {
					Name:  "vpcName",
					Value: slbAdditional.Vpc,
				}, {
					Name:  "vpcId",
					Value: slbAdditional.Vpc,
				}, {
					Name:  "spec",
					Value: slbAdditional.Spec,
				}, {
					Name:  "listener",
					Value: getListenerList(slbAdditional),
				}},
			})
		}
	}
	return response.Data.Total, list
}

func getListenerList(slb *SlbAdditional) string {
	var listener []string
	for _, v := range slb.Listeners {
		listener = append(listener, v.ListenerName)
	}
	return strings.Join(listener, ",")
}
