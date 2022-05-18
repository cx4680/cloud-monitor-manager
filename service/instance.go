package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
)

type InstanceLabel struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InstanceCommonVO struct {
	InstanceId   string          `json:"instanceId"`
	InstanceName string          `json:"instanceName"`
	TenantId     string          `json:"tenantId"`
	Labels       []InstanceLabel `json:"labels"`
}

type InstancePageForm struct {
	InstanceId   string `form:"instanceId"`
	InstanceName string `form:"instanceName"`
	RegionCode   string `form:"regionCode"`
	TenantId     string `form:"tenantId"`
	StatusList   string `form:"statusList"`
	Current      int    `form:"current,default=1"`
	PageSize     int    `form:"pageSize,default=10"`
	Product      string `form:"product"`
}

type InstanceStage interface {
	ConvertRealForm(InstancePageForm) interface{}
	DoRequest(string, interface{}) (interface{}, error)
	ConvertResp(realResp interface{}) (int, []InstanceCommonVO)
}

type InstanceService interface {
	GetPage(InstancePageForm, InstanceStage) (*form.PageVO, error)
}

type InstanceServiceImpl struct {
}

func (is *InstanceServiceImpl) GetPage(page InstancePageForm, stage InstanceStage) (*form.PageVO, error) {
	var err error
	f := stage.ConvertRealForm(page)

	url := config.Cfg.Common.Rc
	logger.Logger().Infof(" request  %+v ,%s", page, url)
	resp, err := stage.DoRequest(url, f)
	if err != nil {
		return nil, err
	}
	logger.Logger().Infof(" resp:%+v", resp)
	total, list := stage.ConvertResp(resp)
	return &form.PageVO{
		Records: list,
		Total:   total,
		Size:    page.PageSize,
		Current: page.Current,
		Pages:   (total / page.PageSize) + 1,
	}, nil
}
