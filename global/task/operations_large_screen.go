package task

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/external"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
	"math"
	"net/url"
	"strconv"
	"strings"
)

func OperationsLargeScreen() BusinessTaskDTO {
	task := func() {
		region := config.Cfg.Common.RegionName
		//连接线上远程数据库
		db, err := getPushRemoteDb()
		if err != nil {
			logger.Logger().Error(err)
			return
		}
		var newResourceRunningStatusList []*model.OperationsLargeScreenResourceRunningStatus
		var deleteResourceRunningStatusIdList []string
		var newResourceUsageTopList []*model.OperationsLargeScreenResourceUsageTop
		var deleteResourceUsageTopIdList []string

		//查询物理服务器
		newPhysicalServerList, deletePhysicalServerIdList, err := getPhysicalServerData(db, region)
		if err != nil {
			logger.Logger().Errorf("getPhysicalServerData error: %v", err)
		}
		newResourceRunningStatusList = append(newResourceRunningStatusList, newPhysicalServerList...)
		deleteResourceRunningStatusIdList = append(deleteResourceRunningStatusIdList, deletePhysicalServerIdList...)

		//查询网络设备
		newNetworkDeviceList, deleteNetworkDeviceIdList, err := getNetworkDeviceData(db, region)
		if err != nil {
			logger.Logger().Errorf("getNetworkDeviceData error: %v", err)
		}
		newResourceRunningStatusList = append(newResourceRunningStatusList, newNetworkDeviceList...)
		deleteResourceRunningStatusIdList = append(deleteResourceRunningStatusIdList, deleteNetworkDeviceIdList...)

		//查询ECS
		newEcsList, deleteEcsIdList, err := getEcsData(db, region)
		if err != nil {
			logger.Logger().Errorf("getEcsData error: %v", err)
		}
		newResourceRunningStatusList = append(newResourceRunningStatusList, newEcsList...)
		deleteResourceRunningStatusIdList = append(deleteResourceRunningStatusIdList, deleteEcsIdList...)

		//查询数据库
		newRdbList, deleteRdbIdList, err := getRdbData(db, region)
		if err != nil {
			logger.Logger().Errorf("getDbData error: %v", err)
		}
		newResourceRunningStatusList = append(newResourceRunningStatusList, newRdbList...)
		deleteResourceRunningStatusIdList = append(deleteResourceRunningStatusIdList, deleteRdbIdList...)

		//裸金属
		newBmsList, deleteBmsIdList, err := getBmsData(db, region)
		if err != nil {
			logger.Logger().Errorf("getBmsData error: %v", err)
		}
		newResourceRunningStatusList = append(newResourceRunningStatusList, newBmsList...)
		deleteResourceRunningStatusIdList = append(deleteResourceRunningStatusIdList, deleteBmsIdList...)

		//查询top数据
		newResourceUsageTopList, deleteResourceUsageTopIdList, err = getUsageData(db, region)
		if err != nil {
			logger.Logger().Errorf("getUsageData error: %v", err)
		}

		//查询分配量和总量数据
		resourceAllocationTotalList, err := getResourceAllocationTotalDate(db, region)
		if err != nil {
			logger.Logger().Errorf("getResourceAllocationTotalDate error: %v", err)
		}

		//推送远程数据库
		if err = pushRemoteDb.Transaction(func(tx *gorm.DB) error {
			//operations_large_screen_resource_running_status
			if len(deleteResourceRunningStatusIdList) > 0 {
				if err = tx.Delete(&model.OperationsLargeScreenResourceRunningStatus{}, "resource_id IN ?", deleteResourceRunningStatusIdList).Error; err != nil {
					return err
				}
			}
			if len(newResourceRunningStatusList) > 0 {
				if err = tx.Save(newResourceRunningStatusList).Error; err != nil {
					return err
				}
			}
			//operations_large_screen_resource_usage_top
			if len(deleteResourceUsageTopIdList) > 0 {
				if err = tx.Delete(&model.OperationsLargeScreenResourceUsageTop{}, "resource_id IN (?)", deleteResourceUsageTopIdList).Error; err != nil {
					return err
				}
			}
			if len(newResourceUsageTopList) > 0 {
				if err = tx.Save(newResourceUsageTopList).Error; err != nil {
					return err
				}
			}
			//operations_large_screen_resource_allocation_total
			if len(resourceAllocationTotalList) > 0 {
				if err = tx.Save(resourceAllocationTotalList).Error; err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			logger.Logger().Error(err)
			return
		}
	}
	return BusinessTaskDTO{
		Cron: "0 */5 * * * ?",
		Name: "OperationsLargeScreen",
		Task: task,
	}
}

func getPhysicalServerData(db *gorm.DB, region string) ([]*model.OperationsLargeScreenResourceRunningStatus, []string, error) {
	param := "{\"condition\":{\"device\":[{\"field\":\"category_name\",\"operator\":\"$regex\",\"value\":\"服务器\"}]},\"fields\":{},\"page\":{\"start\":0,\"limit\":5000,\"sort\":\"\"}}"
	var response *form.CmdbResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.CmdbApi)
	if err != nil {
		return nil, nil, err
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "PhysicalServer").Find(&oldList).Error; err != nil {
		return nil, nil, err
	}
	var oldMap = make(map[string]*model.OperationsLargeScreenResourceRunningStatus)
	for _, old := range oldList {
		oldMap[old.ResourceId] = old
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.Info {
		newMap[v.Ip] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			ResourceId:  v.Ip,
			Status:      v.RunStatus,
			Type:        "PhysicalServer",
			Region:      region,
			FailureTime: util.TimeToStr(util.StrToTime(util.FullTimeFmt, v.CpuUpdateTime), util.MonthDayTimeFmt),
			CreateTime:  v.CreateTime,
		}
		if _, ok := oldMap[v.Ip]; ok {
			newModel.Id = oldMap[v.Ip].Id
		}
		if v.RunStatus != 1 && v.RunStatus != 0 {
			newModel.Status = 2
		}
		if strutil.IsBlank(v.CpuUpdateTime) {
			newModel.FailureTime = "-----"
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []string
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.ResourceId)
		}
	}
	return newList, deleteIdList, nil
}

func getNetworkDeviceData(db *gorm.DB, region string) ([]*model.OperationsLargeScreenResourceRunningStatus, []string, error) {
	param := "{\"condition\":{\"device\":[{\"field\":\"business_type\",\"operator\":\"$regex\",\"value\":\"XGW\"},{\"field\":\"category_name\",\"operator\":\"$regex\",\"value\":\"服务器\"}]},\"fields\":{},\"page\":{\"start\":0,\"limit\":5000,\"sort\":\"\"}}"
	var response *form.CmdbResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.CmdbApi)
	if err != nil {
		return nil, nil, err
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "NetworkDevice").Find(&oldList).Error; err != nil {
		return nil, nil, err
	}
	var oldMap = make(map[string]*model.OperationsLargeScreenResourceRunningStatus)
	for _, old := range oldList {
		oldMap[old.ResourceId] = old
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.Info {
		newMap[v.Ip] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			ResourceId:  v.Ip,
			Status:      v.RunStatus,
			Type:        "NetworkDevice",
			Region:      region,
			FailureTime: util.TimeToStr(util.StrToTime(util.FullTimeFmt, v.CpuUpdateTime), util.MonthDayTimeFmt),
			CreateTime:  v.CreateTime,
		}
		if _, ok := oldMap[v.Ip]; ok {
			newModel.Id = oldMap[v.Ip].Id
		}
		if v.RunStatus != 1 && v.RunStatus != 0 {
			newModel.Status = 2
		}
		if strutil.IsBlank(v.CpuUpdateTime) {
			newModel.FailureTime = "-----"
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []string
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.ResourceId)
		}
	}
	return newList, deleteIdList, nil
}

func getEcsData(db *gorm.DB, region string) ([]*model.OperationsLargeScreenResourceRunningStatus, []string, error) {
	param := external.InstanceRequest{
		CloudProductCode: "ECS",
		ResourceTypeCode: "instance",
		CurrPage:         "1",
		PageSize:         "9999",
	}
	var response *external.InstanceResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
	if err != nil {
		return nil, nil, err
	}
	if response == nil || response.Msg != "success" {
		return nil, nil, nil
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "ECS").Find(&oldList).Error; err != nil {
		return nil, nil, err
	}
	var oldMap = make(map[string]*model.OperationsLargeScreenResourceRunningStatus)
	for _, old := range oldList {
		oldMap[old.ResourceId] = old
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.List {
		newMap[v.ResourceId] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			ResourceId:  v.ResourceId,
			Type:        "ECS",
			Region:      config.Cfg.Common.RegionName,
			FailureTime: util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:  util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if _, ok := oldMap[v.ResourceId]; ok {
			newModel.Id = oldMap[v.ResourceId].Id
		}
		if v.StatusDesc == "active" {
			newModel.Status = 1
		} else {
			newModel.Status = 0
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []string
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.ResourceId)
		}
	}
	return newList, deleteIdList, nil
}

func getRdbData(db *gorm.DB, region string) ([]*model.OperationsLargeScreenResourceRunningStatus, []string, error) {
	var paramList = []*external.InstanceRequest{
		{
			CloudProductCode: "RDB",
			ResourceTypeCode: "mysql",
			CurrPage:         "1",
			PageSize:         "9999",
		},
		{
			CloudProductCode: "RDB",
			ResourceTypeCode: "dm",
			CurrPage:         "1",
			PageSize:         "9999",
		},
		{
			CloudProductCode: "RDB",
			ResourceTypeCode: "pg",
			CurrPage:         "1",
			PageSize:         "9999",
		},
	}
	var responseDataList []*external.InstanceList
	for _, param := range paramList {
		var response *external.InstanceResponse
		_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
		if err != nil {
			return nil, nil, err
		}
		if response == nil || response.Msg != "success" {
			continue
		}
		responseDataList = append(responseDataList, response.Data.List...)
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err := db.Where("region = ? AND type = ?", region, "RDB").Find(&oldList).Error; err != nil {
		return nil, nil, err
	}
	var oldMap = make(map[string]*model.OperationsLargeScreenResourceRunningStatus)
	for _, old := range oldList {
		oldMap[old.ResourceId] = old
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var deleteIdList []string
	var newMap = make(map[string]interface{})
	for _, v := range responseDataList {
		newMap[v.ResourceId] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			ResourceId:  v.ResourceId,
			Type:        "数据库",
			Region:      config.Cfg.Common.RegionName,
			FailureTime: util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:  util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if _, ok := oldMap[v.ResourceId]; ok {
			newModel.Id = oldMap[v.ResourceId].Id
		}
		if v.StatusDesc == "运行中" {
			newModel.Status = 1
		} else {
			newModel.Status = 0
		}
		newList = append(newList, newModel)
	}
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.ResourceId)
		}
	}
	return newList, deleteIdList, nil
}

func getBmsData(db *gorm.DB, region string) ([]*model.OperationsLargeScreenResourceRunningStatus, []string, error) {
	var paramList = []*external.InstanceRequest{
		{
			CloudProductCode: "BMS",
			ResourceTypeCode: "TBMS",
			CurrPage:         "1",
			PageSize:         "9999",
		},
		{
			CloudProductCode: "BMS",
			ResourceTypeCode: "EBMS",
			CurrPage:         "1",
			PageSize:         "9999",
		},
	}
	var responseDataList []*external.InstanceList
	for _, param := range paramList {
		var response *external.InstanceResponse
		_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
		if err != nil {
			return nil, nil, err
		}
		if response == nil || response.Msg != "success" {
			continue
		}
		responseDataList = append(responseDataList, response.Data.List...)
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err := db.Where("region = ? AND type = ?", region, "BMS").Find(&oldList).Error; err != nil {
		return nil, nil, err
	}
	var oldMap = make(map[string]*model.OperationsLargeScreenResourceRunningStatus)
	for _, old := range oldList {
		oldMap[old.ResourceId] = old
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range responseDataList {
		newMap[v.ResourceId] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			ResourceId:  v.ResourceId,
			Type:        "裸金属",
			Region:      config.Cfg.Common.RegionName,
			FailureTime: util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:  util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if _, ok := oldMap[v.ResourceId]; ok {
			newModel.Id = oldMap[v.ResourceId].Id
		}
		if v.StatusDesc == "running" {
			newModel.Status = 1
		} else {
			newModel.Status = 0
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []string
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.ResourceId)
		}
	}
	return newList, deleteIdList, nil
}

var codeList = []string{
	"ecs_cpu_base_usage",            //ECS cpu使用率
	"ecs_memory_base_usage",         //ECS 内存使用率
	"mysql_cpu_usage",               //RDB mysql cpu使用率
	"dm_global_status_cpu_use_rate", //RDB dm cpu使用率
	"pg_cpu_usage",                  //RDB pg cpu使用率
	"mysql_mem_usage",               //RDB mysql 内存使用率
	"dm_global_status_mem_use_rate", //RDB dm 内存使用率
	"pg_mem_usage",                  //RDB pg 内存使用率
	"mysql_current_cons_num",        //RDB mysql 连接数
	"dm_global_status_sessions",     //RDB dm 连接数
	"pg_open_ct_num",                //RDB pg 连接数
	"eip_upstream_bandwidth_usage",  //EIP 出网带宽使用率
}

func getUsageData(db *gorm.DB, region string) ([]*model.OperationsLargeScreenResourceUsageTop, []string, error) {
	var newList []*model.OperationsLargeScreenResourceUsageTop
	var deleteIdList []string
	for _, code := range codeList {
		var resourceType, attribute string
		switch code {
		case "ecs_cpu_base_usage":
			resourceType = "ECS"
			attribute = "CPU"
		case "ecs_memory_base_usage":
			resourceType = "ECS"
			attribute = "内存"
		case "mysql_cpu_usage", "dm_global_status_cpu_use_rate", "pg_cpu_usage":
			resourceType = "RDB"
			attribute = "CPU"
		case "mysql_mem_usage", "dm_global_status_mem_use_rate", "pg_mem_usage":
			resourceType = "RDB"
			attribute = "内存"
		case "mysql_current_cons_num", "dm_global_status_sessions", "pg_open_ct_num":
			resourceType = "RDB"
			attribute = "连接数"
		case "eip_upstream_bandwidth_usage":
			resourceType = "EIP"
			attribute = "带宽"
		}
		//查询指标
		var monitorItem *model.MonitorItem
		if err := global.DB.Where("metric_name = ?", code).Find(&monitorItem).Error; err != nil {
			return nil, nil, err
		}
		pql := fmt.Sprintf(constant.TopExpr, "10", strings.ReplaceAll(monitorItem.MetricsLinux, constant.MetricLabel, ""))
		prometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(pql)
		response := &form.PrometheusResponse{}
		_, err := httputil.GetHttpClient().R().SetResult(&response).Get(prometheusUrl)
		if err != nil {
			logger.Logger().Error(err)
			continue
		}
		if response == nil || response.Data == nil || response.Data.Result == nil || len(response.Data.Result) == 0 {
			logger.Logger().Infof("query prometheus empty, pql: %s, origin response:%s", pql, response)
			continue
		}
		//查询数据库的旧数据
		var oldList []*model.OperationsLargeScreenResourceUsageTop
		if err = db.Where("region = ? AND type = ? AND attribute = ?", region, resourceType, attribute).Find(&oldList).Error; err != nil {
			return nil, nil, err
		}
		var oldMap = make(map[string]*model.OperationsLargeScreenResourceUsageTop)
		for _, old := range oldList {
			oldMap[old.ResourceId] = old
		}
		var newMap = make(map[string]interface{})
		for _, v := range response.Data.Result {
			resourceId := v.Metric[constant.INSTANCE]
			number := changeDecimal(v.Value[1].(string))
			newMap[resourceId] = nil
			newModel := &model.OperationsLargeScreenResourceUsageTop{
				ResourceId: resourceId,
				Type:       resourceType,
				Attribute:  attribute,
				Number:     number,
				Unit:       monitorItem.Unit,
				Region:     config.Cfg.Common.RegionName,
			}
			if _, ok := oldMap[resourceId]; ok {
				newModel.Id = oldMap[resourceId].Id
			}
			newList = append(newList, newModel)
		}
		for _, old := range oldList {
			if _, ok := newMap[old.ResourceId]; !ok {
				deleteIdList = append(deleteIdList, old.ResourceId)
			}
		}
	}
	return newList, deleteIdList, nil
}

// 数据保留两位小数
func changeDecimal(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return math.Trunc(v*1e2+0.5) * 1e-2
}

func getResourceAllocationTotalDate(db *gorm.DB, region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	var newList []*model.OperationsLargeScreenResourceAllocationTotal
	//ECS
	ecsResourceAllocationTotalDate, err := getEcsResourceAllocationTotalDate(region)
	if err != nil {
		logger.Logger().Errorf("getEcsResourceAllocationTotalDate error: %v", err)
	}
	newList = append(newList, ecsResourceAllocationTotalDate...)
	//EBS
	ebsResourceAllocationTotalDate, err := getEbsResourceAllocationTotalDate(region)
	if err != nil {
		logger.Logger().Errorf("getEbsResourceAllocationTotalDate error: %v", err)
	}
	newList = append(newList, ebsResourceAllocationTotalDate...)
	//RDB cpu
	rdbCpuResourceAllocationTotalDate, err := getRdbCpuResourceAllocationTotalDate(region)
	if err != nil {
		logger.Logger().Errorf("getRdbCpuResourceAllocationTotalDate error: %v", err)
	}
	newList = append(newList, rdbCpuResourceAllocationTotalDate...)
	//RDB 实例个数
	rdbCountResourceAllocationTotalDate, err := getRdbCountResourceAllocationTotalDate(region)
	if err != nil {
		logger.Logger().Errorf("getRdbCountResourceAllocationTotalDate error: %v", err)
	}
	newList = append(newList, rdbCountResourceAllocationTotalDate...)
	//CMQ cpu
	cmqCpuResourceAllocationTotalDate, err := getCmqCpuResourceAllocationTotalDate(region)
	if err != nil {
		logger.Logger().Errorf("getCmqCpuResourceAllocationTotalDate error: %v", err)
	}
	newList = append(newList, cmqCpuResourceAllocationTotalDate...)
	//CMQ 实例个数
	cmqCountResourceAllocationTotalDate, err := getCmqCountResourceAllocationTotalDate(region)
	if err != nil {
		logger.Logger().Errorf("getCmqCountResourceAllocationTotalDate error: %v", err)
	}
	newList = append(newList, cmqCountResourceAllocationTotalDate...)
	//EIP 带宽
	eipResourceAllocationTotalDate, err := getEipResourceAllocationTotalDate(region)
	if err != nil {
		logger.Logger().Errorf("getEipResourceAllocationTotalDate error: %v", err)
	}
	newList = append(newList, eipResourceAllocationTotalDate...)
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceAllocationTotal
	if err = db.Where("region = ?", region).Find(&oldList).Error; err != nil {
		return nil, err
	}
	var oldMap = make(map[string]*model.OperationsLargeScreenResourceAllocationTotal)
	for _, old := range oldList {
		oldMap[fmt.Sprintf("%s-%s", old.Type, old.Attribute)] = old
	}
	for i, v := range newList {
		key := fmt.Sprintf("%s-%s", v.Type, v.Attribute)
		if _, ok := oldMap[key]; ok {
			newList[i].Id = oldMap[key].Id
		}
	}
	return newList, nil
}

func getEcsResourceAllocationTotalDate(region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	var response *form.EcsCusInventory
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.CusInventoryApi + "/large/cpu/mem")
	if err != nil {
		return nil, err
	}
	if response == nil || response.Data == nil {
		return nil, nil
	}
	var list = []*model.OperationsLargeScreenResourceAllocationTotal{
		{
			Type:       "ECS",
			Attribute:  "cpu",
			Allocation: math.Ceil(float64(response.Data.UsedCores) / 1000),
			Total:      math.Ceil(float64(response.Data.Cores) / 1000),
			Unit:       "核",
			Region:     region,
		},
		{
			Type:       "ECS",
			Attribute:  "memory",
			Allocation: math.Ceil(float64(response.Data.UsedRam) / 1000 / 1000 / 1000),
			Total:      math.Ceil(float64(response.Data.Ram) / 1000 / 1000 / 1000),
			Unit:       "GB",
			Region:     region,
		},
	}
	return list, nil
}

func getEbsResourceAllocationTotalDate(region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	var response *form.EbsCusInventory
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.CusInventoryApi + "/super/sum")
	if err != nil {
		return nil, err
	}
	if response == nil || response.Data == nil {
		return nil, nil
	}
	var list = []*model.OperationsLargeScreenResourceAllocationTotal{
		{
			Type:       "EBS",
			Attribute:  "capacity",
			Allocation: math.Ceil(float64(response.Data.AllocatedCapacityGb) / 1000),
			Total:      math.Ceil(float64(response.Data.TotalCapacityGb) / 1000),
			Unit:       "TB",
			Region:     region,
		},
	}
	return list, nil
}

func getRdbCpuResourceAllocationTotalDate(region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	var response *form.RdbManageResource
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.RdbApi + "/cpu")
	if err != nil {
		return nil, err
	}
	if response == nil || response.Data == nil {
		return nil, nil
	}
	var list = []*model.OperationsLargeScreenResourceAllocationTotal{
		{
			Type:       "RDB",
			Attribute:  "cpu",
			Allocation: float64(response.Data.CpuUsed),
			Total:      float64(response.Data.CpuTotal),
			Unit:       "核",
			Region:     region,
		},
	}
	return list, nil
}

func getRdbCountResourceAllocationTotalDate(region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	var response *form.RdbManageResource
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.RdbApi + "/count")
	if err != nil {
		return nil, err
	}
	if response == nil || response.Data == nil {
		return nil, nil
	}
	var list = []*model.OperationsLargeScreenResourceAllocationTotal{
		{
			Type:       "RDB",
			Attribute:  "status",
			Allocation: float64(response.Data.DatabaseError),
			Total:      float64(response.Data.DatabaseTotal),
			Unit:       "个",
			Region:     region,
		},
	}
	return list, nil
}

func getCmqCpuResourceAllocationTotalDate(region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	var response *form.CmqCpuResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.CmqNodeApi)
	if err != nil {
		return nil, err
	}
	if response == nil || response.Data == nil {
		return nil, nil
	}
	var list = []*model.OperationsLargeScreenResourceAllocationTotal{
		{
			Type:       "CMQ",
			Attribute:  "cpu",
			Allocation: float64(response.Data.CpuUsed),
			Total:      float64(response.Data.CpuCap),
			Unit:       "核",
			Region:     region,
		},
	}
	return list, nil
}

func getCmqCountResourceAllocationTotalDate(region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	var response *form.CmqCountResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.CmqInsApi)
	if err != nil {
		return nil, err
	}
	if response == nil || response.Data == nil {
		return nil, nil
	}
	var list = []*model.OperationsLargeScreenResourceAllocationTotal{
		{
			Type:       "CMQ",
			Attribute:  "status",
			Allocation: float64(response.Data.ExceptionInstanceNum),
			Total:      float64(response.Data.AllInstanceNum),
			Unit:       "个",
			Region:     region,
		},
	}
	return list, nil
}

func getEipResourceAllocationTotalDate(region string) ([]*model.OperationsLargeScreenResourceAllocationTotal, error) {
	//总带宽
	totalPql := "sum(avg(eip_config_upstream_bandwidth{eipType='external_eip'})by(instance))"
	totalPrometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(totalPql)
	totalResponse := &form.PrometheusResponse{}
	_, err := httputil.GetHttpClient().R().SetResult(&totalResponse).Get(totalPrometheusUrl)
	if err != nil {
		logger.Logger().Error(err)
		return nil, err
	}
	if totalResponse == nil || totalResponse.Data == nil || totalResponse.Data.Result == nil || len(totalResponse.Data.Result) == 0 {
		logger.Logger().Infof("query prometheus empty, pql: %s, origin response:%s", totalPql, totalResponse)
		return nil, nil
	}
	total := changeDecimal(totalResponse.Data.Result[0].Value[1].(string)) / 1000 / 1000
	//上行流量
	upstreamPql := "sum(eip_upstream_bits_rate{eipType='external_eip'})"
	upstreamPrometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(upstreamPql)
	upstreamResponse := &form.PrometheusResponse{}
	_, err = httputil.GetHttpClient().R().SetResult(&upstreamResponse).Get(upstreamPrometheusUrl)
	if err != nil {
		logger.Logger().Error(err)
		return nil, err
	}
	if upstreamResponse == nil || upstreamResponse.Data == nil || upstreamResponse.Data.Result == nil || len(upstreamResponse.Data.Result) == 0 {
		logger.Logger().Infof("query prometheus empty, pql: %s, origin response:%s", upstreamPql, upstreamResponse)
		return nil, nil
	}
	upstream := changeDecimal(upstreamResponse.Data.Result[0].Value[1].(string)) / 1000 / 1000
	//下行流量
	downstreamPql := "sum(eip_upstream_bits_rate{eipType='external_eip'})"
	downstreamPrometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(downstreamPql)
	downstreamResponse := &form.PrometheusResponse{}
	_, err = httputil.GetHttpClient().R().SetResult(&downstreamResponse).Get(downstreamPrometheusUrl)
	if err != nil {
		logger.Logger().Error(err)
		return nil, err
	}
	if downstreamResponse == nil || downstreamResponse.Data == nil || downstreamResponse.Data.Result == nil || len(downstreamResponse.Data.Result) == 0 {
		logger.Logger().Infof("query prometheus empty, pql: %s, origin response:%s", downstreamPql, downstreamResponse)
		return nil, nil
	}
	downstream := changeDecimal(downstreamResponse.Data.Result[0].Value[1].(string)) / 1000 / 1000
	var list = []*model.OperationsLargeScreenResourceAllocationTotal{
		{
			Type:       "EIP",
			Attribute:  "upstream",
			Allocation: upstream,
			Total:      total,
			Unit:       "个",
			Region:     region,
		},
		{
			Type:       "EIP",
			Attribute:  "downstream",
			Allocation: downstream,
			Total:      total,
			Unit:       "个",
			Region:     region,
		},
	}
	return list, nil
}

var pushRemoteDb *gorm.DB

func getPushRemoteDb() (*gorm.DB, error) {
	if pushRemoteDb == nil {
		dsn := config.Cfg.Db.LargeScreenDbDsn
		logger.Logger().Infof("LargeScreenDbDsn: %v", dsn)
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{Logger: dbLogger.Default.LogMode(dbLogger.Info)})
		if err != nil {
			return nil, err
		}
		pushRemoteDb = db
	}
	return pushRemoteDb, nil
}
