package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/dao"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/errors"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
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

func (s *MonitorItemService) ChangeMonitorItemDisplay(param form.MonitorItemParam, userId string) error {
	if len(param.MonitorItemList) == 0 {
		return errors.NewBusinessError("监控项不能为空")
	}
	OldMonitorItemMap := make(map[string]int)
	for _, v := range dao.MonitorItem.GetCloseMonitorItem(userId) {
		OldMonitorItemMap[v.ItemBizId] = 1
	}
	var newMonitorItemMap []model.MonitorItemClose
	if param.Active == "close" {
		for _, v := range param.MonitorItemList {
			if strutil.IsNotBlank(v) && OldMonitorItemMap[v] == 0 {
				newMonitorItemMap = append(newMonitorItemMap, model.MonitorItemClose{UserId: userId, ItemBizId: v})
			}
		}
		if len(newMonitorItemMap) > 0 {
			s.dao.CloseMonitorItem(newMonitorItemMap)
		}
	} else if param.Active == "open" {
		s.dao.OpenMonitorItem(param.MonitorItemList)
	}
	return nil
}

func (s *MonitorItemService) GetMonitorItemPage(pageSize int, pageNum int, productAbbr string) *form.PageVO {
	var monitorItemList []model.MonitorItem
	sql := "select item.* from t_monitor_item item  ,t_monitor_product product  where item.product_biz_id=product.biz_id  and product.abbreviation=?"
	paginate := util.Paginate(pageSize, pageNum, sql, []interface{}{productAbbr}, &monitorItemList)
	return paginate
}
