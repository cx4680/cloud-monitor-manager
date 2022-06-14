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

type SlbAdditional struct {
	Address   string `json:"address"`
	Spec      string `json:"spec"`
	Vpc       string `json:"vpc"`
	VpcName   string `json:"vpcName"`
	Container []struct {
	} `json:"container"`
	Listeners []struct {
		ListenerId   string `json:"listenerId"`
		ListenerName string `json:"listenerName"`
		Port         int    `json:"port"`
		Protocol     string `json:"protocol"`
	} `json:"listeners"`
	Eip struct {
		Name string `json:"name"`
		Ip   string `json:"ip"`
	} `json:"eip"`
}

func (s *SlbInstanceService) ConvertRealForm(f service.InstancePageForm) interface{} {
	param := InstanceRequest{
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

func (s *SlbInstanceService) DoRequest(url string, f interface{}) (interface{}, error) {
	respStr, err := httputil.HttpPostJson(url, f, nil)
	if err != nil {
		return nil, err
	}
	var resp InstanceResponse
	jsonutil.ToObject(respStr, &resp)
	return resp, nil
}

func (s *SlbInstanceService) ConvertResp(realResp interface{}) (int, []service.InstanceCommonVO) {
	response := realResp.(InstanceResponse)
	var list []service.InstanceCommonVO
	if response.Data.Total > 0 {
		for _, d := range response.Data.List {
			var additional = &SlbAdditional{}
			jsonutil.ToObject(d.Additional, additional)
			list = append(list, service.InstanceCommonVO{
				InstanceId:   d.ResourceId,
				InstanceName: d.ResourceName,
				TenantId:     d.TenantId,
				Labels: []service.InstanceLabel{{
					Name:  "status",
					Value: d.StatusDesc,
				}, {
					Name:  "privateIp",
					Value: additional.Address,
				}, {
					Name:  "vpcName",
					Value: additional.VpcName,
				}, {
					Name:  "vpcId",
					Value: additional.Vpc,
				}, {
					Name:  "spec",
					Value: additional.Spec,
				}, {
					Name:  "listener",
					Value: getListenerList(additional),
				}, {
					Name:  "eipIp",
					Value: additional.Eip.Ip,
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
