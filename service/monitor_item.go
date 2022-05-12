package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
)

type MonitorItemService struct {
	dao *dao.MonitorItemDao
}

func NewMonitorItemService() *MonitorItemService {
	return &MonitorItemService{
		dao: dao.MonitorItem,
	}
}

func (s *MonitorItemService) GetMonitorItem(param form.MonitorItemParam) []model.MonitorItem {
	return s.dao.GetMonitorItem(param.ProductBizId, param.OsType, param.Display)
}

func (s *MonitorItemService) GetMonitorItemPage(pageSize int, pageNum int, productAbbr string) *form.PageVO {
	var monitorItemList []model.MonitorItem
	sql := "select item.* from t_monitor_item item  ,t_monitor_product product  where item.product_biz_id=product.biz_id  and product.abbreviation=?"
	paginate := util.Paginate(pageSize, pageNum, sql, []interface{}{productAbbr}, &monitorItemList)
	return paginate
}
