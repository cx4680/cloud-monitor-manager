package dao

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
)

type ConfigItemDao struct {
}

var ConfigItem = new(ConfigItemDao)

func (dao *ConfigItemDao) GetConfigItem(code interface{}, pBizId int64, data string) *model.ConfigItem {
	item := model.ConfigItem{}
	db := global.DB
	if code != nil {
		db = db.Where("code", code)
	}
	if pBizId > 0 {
		db = db.Where("p_biz_id", pBizId)
	}
	if len(data) > 0 {
		db = db.Where("data", data)
	}
	db.Find(&item)
	return &item
}

func (dao *ConfigItemDao) GetConfigItemList(pBizId int64) []*model.ConfigItem {
	var list []*model.ConfigItem
	db := global.DB
	db = db.Where("p_biz_id", pBizId).Order("sort_id ASC")
	db.Find(&list)
	return list
}

const (
	MonitorRange = 1 //监控周期
)
