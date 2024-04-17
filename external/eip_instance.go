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

type EipAdditional struct {
	BandWidth struct {
		BandwidthId   string `json:"bandwidthId"`
		BandWidthSize int    `json:"BandWidthSize"`
	} `json:"bandWidth"`
	BindInstanceId string `json:"bindInstanceId"`
	EipIpAddress   string `json:"eipIpAddress"`
}

func (s *EipInstanceService) ConvertRealForm(f service.InstancePageForm) interface{} {
	param := InstanceRequest{
		CloudProductCode: "EIP",
		ResourceTypeCode: "instance",
		ResourceId:       f.InstanceId,
		Name:             f.InstanceName,
		RegionCode:       f.RegionCode,
		TenantId:         f.TenantId,
		StatusList:       toStringList(f.StatusList),
		CurrPage:         strconv.Itoa(f.Current),
		PageSize:         strconv.Itoa(f.PageSize),
		TagKeyId:         f.TagKeyId,
	}
	return param
}

func (s *EipInstanceService) DoRequest(url string, f interface{}) (interface{}, error) {
	respStr, err := httputil.HttpPostJson(url, f, nil)
	if err != nil {
		return nil, err
	}
	var resp InstanceResponse
	jsonutil.ToObject(respStr, &resp)
	return resp, nil
}

func (s *EipInstanceService) ConvertResp(realResp interface{}) (int, []service.InstanceCommonVO) {
	response := realResp.(InstanceResponse)
	var list []service.InstanceCommonVO
	if response.Data.Total > 0 {
		for _, d := range response.Data.List {
			var additional = &EipAdditional{}
			jsonutil.ToObject(d.Additional, additional)
			list = append(list, service.InstanceCommonVO{
				InstanceId:   d.ResourceId,
				InstanceName: d.ResourceName,
				TenantId:     d.TenantId,
				Labels: []service.InstanceLabel{{
					Name:  "status",
					Value: d.StatusDesc,
				}, {
					Name:  "eipAddress",
					Value: additional.EipIpAddress,
				}, {
					Name:  "bandWidth",
					Value: strconv.Itoa(additional.BandWidth.BandWidthSize),
				}, {
					Name:  "bindInstanceId",
					Value: additional.BindInstanceId,
				}},
			})
		}
	}
	return response.Data.Total, list
}
