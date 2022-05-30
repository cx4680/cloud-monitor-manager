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

func (s *MonitorItemService) GetMonitorItemByProductBizId(param form.MonitorItemParam, userId string) []model.MonitorItem {
	return s.dao.GetMonitorItem(param.ProductBizId, param.OsType, param.Display, userId)
}

func (s *MonitorItemService) GetAllMonitorItemByProductBizId(param form.MonitorItemParam) []model.MonitorItem {
	return s.dao.GetAllMonitorItem(param.ProductBizId, param.OsType, param.Display)
}

func (s *MonitorItemService) OpenDisplay(param form.MonitorItemParam, userId string) error {
	newMonitorItemMap := make(map[string]int)
	for _, v := range param.MonitorItemList {
		newMonitorItemMap[v] = 1
	}
	monitorItemList := dao.MonitorItem.GetAllMonitorItem(param.ProductBizId, "", "")
	var closeList []model.MonitorItemClose
	var allBizId []string
	for _, v := range monitorItemList {
		allBizId = append(allBizId, v.BizId)
		if newMonitorItemMap[v.BizId] != 1 {
			closeList = append(closeList, model.MonitorItemClose{UserId: userId, ItemBizId: v.BizId})
		}
	}
	s.dao.OpenMonitorItem(allBizId)
	if len(closeList) > 0 {
		s.dao.CloseMonitorItem(closeList)
	}
	return nil
}

func (s *MonitorItemService) GetMonitorItemPage(pageSize int, pageNum int, productAbbr string) *form.PageVO {
	var monitorItemList []model.MonitorItem
	sql := "select item.* from t_monitor_item item  ,t_monitor_product product  where item.product_biz_id=product.biz_id  and product.abbreviation=?"
	paginate := util.Paginate(pageSize, pageNum, sql, []interface{}{productAbbr}, &monitorItemList)
	return paginate
}
