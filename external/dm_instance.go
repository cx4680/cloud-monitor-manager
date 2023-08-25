package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"strconv"
)

type DmInstanceService struct {
	service.InstanceServiceImpl
}

type DmAdditional struct {
	InstanceMode string `json:"instanceMode"`
	RdsStateTime string `json:"RdsStateTime"`
}

func (s *DmInstanceService) ConvertRealForm(f service.InstancePageForm) interface{} {
	param := InstanceRequest{
		CloudProductCode: "RDB",
		ResourceTypeCode: "dm",
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

func (s *DmInstanceService) DoRequest(url string, f interface{}) (interface{}, error) {
	respStr, err := httputil.HttpPostJson(url, f, nil)
	if err != nil {
		return nil, err
	}
	var resp InstanceResponse
	jsonutil.ToObject(respStr, &resp)
	return resp, nil
}

func (s *DmInstanceService) ConvertResp(realResp interface{}) (int, []service.InstanceCommonVO) {
	response := realResp.(InstanceResponse)
	var list []service.InstanceCommonVO
	if response.Data.Total > 0 {
		for _, d := range response.Data.List {
			var additional = &DmAdditional{}
			jsonutil.ToObject(d.Additional, additional)
			list = append(list, service.InstanceCommonVO{
				InstanceId:   d.ResourceId,
				InstanceName: d.ResourceName,
				TenantId:     d.TenantId,
				Labels: []service.InstanceLabel{{
					Name:  "status",
					Value: d.StatusDesc,
				}, {
					Name:  "instanceMode",
					Value: additional.InstanceMode,
				}, {
					Name:  "rdsStateTime",
					Value: additional.RdsStateTime,
				}},
			})
		}
	}
	return response.Data.Total, list
}
