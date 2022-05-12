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
	var product model.MonitorProduct
	db.Where(model.MonitorProduct{Abbreviation: abbreviation}).First(&product)
	return &product

}

func (d *MonitorProductDao) GetMonitorProductByBizId(BizId string) model.MonitorProduct {
	var product = model.MonitorProduct{}
	global.DB.Where("biz_id = ?", BizId).First(&product)
	return product
}

func (d *MonitorProductDao) GetMonitorProduct() *[]model.MonitorProduct {
	var product = &[]model.MonitorProduct{}
	global.DB.Where("status = ?", "1").Find(product)
	return product
}

func (d *MonitorProductDao) GetAllMonitorProduct() *[]model.MonitorProduct {
	var product = &[]model.MonitorProduct{}
	global.DB.Find(product)
	return product
}

func (d *MonitorProductDao) GetMonitorProductDTO() *[]model.MonitorProductDTO {
	var product = &[]model.MonitorProductDTO{}
	global.DB.Model(&model.MonitorProduct{}).Where("status = ?", "1").Order("sort ASC").Find(product)
	return product
}

func (d *MonitorProductDao) ChangeStatus(db *gorm.DB, bizId []string, status uint8) {
	global.DB.Model(&model.MonitorProduct{}).Where("biz_id IN (?)", bizId).Update("status", status)
}

func (d *MonitorProductDao) GetByName(db *gorm.DB, name string) *model.MonitorProduct {
	if strutil.IsBlank(name) {
		return nil
	}
	var product model.MonitorProduct
	db.Where(model.MonitorProduct{Name: name}).First(&product)
	return &product

}
