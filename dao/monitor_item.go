package dao

import (
	"bytes"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/sys_component/sys_redis"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"strconv"
	"text/template"
	"time"
)

type MonitorItemDao struct {
}

var MonitorItem = new(MonitorItemDao)

func (d *MonitorItemDao) GetMonitorItem(productBizId, osType, display string) []model.MonitorItem {
	var monitorItemList []model.MonitorItem
	if strutil.IsNotBlank(display) {
		global.DB.Where("status = ? AND is_display = ? AND product_biz_id = ? AND display LIKE ?", "1", "1", productBizId, "%"+display+"%").Find(&monitorItemList)
	} else {
		global.DB.Where("status = ? AND is_display = ? AND product_biz_id = ?", "1", "1", productBizId).Find(&monitorItemList)
	}
	if strutil.IsBlank(osType) {
		return monitorItemList
	}
	var newMonitorItemList []model.MonitorItem
	for _, v := range monitorItemList {
		if strutil.IsNotBlank(v.ShowExpression) && !isShow(v.ShowExpression, osType) {
			continue
		}
		newMonitorItemList = append(newMonitorItemList, v)
	}
	return newMonitorItemList
}

func (d *MonitorItemDao) GetMonitorItemCacheByMetricCode(name string) model.MonitorItem {
	value, err := sys_redis.Get(name)
	if err != nil {
		logger.Logger().Info("key=" + name + ", error:" + err.Error())
	}
	var monitorItemModel = model.MonitorItem{}
	if strutil.IsNotBlank(value) {
		jsonutil.ToObject(value, &monitorItemModel)
		return monitorItemModel
	}
	monitorItemModel = d.GetMonitorItemByMetricCode(name)
	if monitorItemModel == (model.MonitorItem{}) {
		logger.Logger().Info("获取监控项为空")
		return monitorItemModel
	}
	if e := sys_redis.SetByTimeOut(name, jsonutil.ToString(monitorItemModel), time.Hour); e != nil {
		logger.Logger().Error("设置监控项缓存错误, key=" + name)
	}
	return monitorItemModel
}

func (d *MonitorItemDao) GetMonitorItemByMetricCode(name string) model.MonitorItem {
	var monitorItemModel = model.MonitorItem{}
	global.DB.Where("metric_name = ?", name).First(&monitorItemModel)
	return monitorItemModel
}

func isShow(exp string, os string) bool {
	m := map[string]string{"OSTYPE": os}
	var buf bytes.Buffer
	temp, _ := template.New("exp").Parse(exp)
	if err := temp.Execute(&buf, m); err != nil {
		logger.Logger().Errorf("展示表达式解析失败：%v", err)
		return true
	}
	isShowBool, err := strconv.ParseBool(buf.String())
	if err != nil {
		logger.Logger().Errorf("展示表达式解析失败：%v", err)
		return true
	}
	return isShowBool
}
