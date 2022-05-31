package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"strconv"
)

type EbmsInstanceService struct {
	service.InstanceServiceImpl
}

type EbmsAdditional struct {
}

func (s *EbmsInstanceService) ConvertRealForm(f service.InstancePageForm) interface{} {
	param := InstanceRequest{
		CloudProductCode: "EBMS",
		//ResourceTypeCode: "instance",
		ResourceId: f.InstanceId,
		Name:       f.InstanceName,
		RegionCode: f.RegionCode,
		TenantId:   f.TenantId,
		StatusList: toStringList(f.StatusList),
		CurrPage:   strconv.Itoa(f.Current),
		PageSize:   strconv.Itoa(f.PageSize),
	}
	return param
}

func (s *EbmsInstanceService) DoRequest(url string, f interface{}) (interface{}, error) {
	respStr, err := httputil.HttpPostJson(url, f, nil)
	if err != nil {
		return nil, err
	}
	var resp InstanceResponse
	jsonutil.ToObject(respStr, &resp)
	return resp, nil
}

func (s *EbmsInstanceService) ConvertResp(realResp interface{}) (int, []service.InstanceCommonVO) {
	response := realResp.(InstanceResponse)
	var list []service.InstanceCommonVO
	if response.Data.Total > 0 {
		for _, d := range response.Data.List {
			var additional = &EbmsAdditional{}
			jsonutil.ToObject(d.Additional, additional)
			list = append(list, service.InstanceCommonVO{
				InstanceId:   d.ResourceId,
				InstanceName: d.ResourceName,
				TenantId:     d.TenantId,
				Labels: []service.InstanceLabel{{
					Name:  "status",
					Value: d.StatusDesc,
				}},
			})
		}
	}
	return response.Data.Total, list
}
