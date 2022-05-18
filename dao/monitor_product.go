package dao

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"gorm.io/gorm"
)

type MonitorProductDao struct {
}

var MonitorProduct = new(MonitorProductDao)

func (d *MonitorProductDao) GetByAbbreviation(db *gorm.DB, abbreviation string) *model.MonitorProduct {
	if strutil.IsBlank(abbreviation) {
		return nil
	}
	var product = &model.MonitorProduct{}
	db.Where(model.MonitorProduct{Abbreviation: abbreviation}).First(product)
	return product

}

func (d *MonitorProductDao) GetMonitorProduct() *[]model.MonitorProduct {
	var product = &[]model.MonitorProduct{}
	global.DB.Where(model.MonitorProduct{Status: 1}).Order("sort ASC").Find(product)
	return product
}
