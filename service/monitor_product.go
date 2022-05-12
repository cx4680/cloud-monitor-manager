package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
)

type MonitorProductService struct {
	dao *dao.MonitorProductDao
}

func NewMonitorProductService() *MonitorProductService {
	return &MonitorProductService{
		dao: dao.MonitorProduct,
	}
}

func (s *MonitorProductService) GetMonitorProduct() *[]model.MonitorProductDTO {
	return s.dao.GetMonitorProductDTO()
}

func (s *MonitorProductService) GetAllMonitorProduct() *[]model.MonitorProduct {
	return s.dao.GetAllMonitorProduct()
}

func (s *MonitorProductService) GetMonitorProductPage(pageSize int, pageNum int) *form.PageVO {
	var productList []model.MonitorProduct
	var total int64
	global.DB.Model(productList).Where("status = ?", "1").Count(&total)
	if total != 0 {
		offset := (pageNum - 1) * pageSize
		global.DB.Where("status = ?", "1").Order("sort ASC").Offset(offset).Limit(pageSize).Find(&productList)
	}
	return &form.PageVO{
		Records: productList,
		Current: pageNum,
		Size:    pageSize,
		Total:   int(total),
	}
}
