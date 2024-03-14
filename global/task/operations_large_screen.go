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
		//查询分配量和总量数据
		go getResourceAllocationTotalDate(db, region)

		//查询top数据
		go getUsageData(db, region, OperationsCodeList)

		//查询物理服务器
		go getPhysicalServerData(db, region)

		//查询网络设备
		go getNetworkDeviceData(db, region)

		//查询ECS
		go getEcsData(db, region)

		//查询数据库
		go getRdbData(db, region)

		//裸金属
		go getBmsData(db, region)
	}
	return BusinessTaskDTO{
		Cron: "0 */5 * * * ?",
		Name: "OperationsLargeScreen",
		Task: task,
	}
}

// 运营大屏资源状态表 operations_large_screen_resource_running_status
func saveTableOperationsLargeScreenResourceRunningStatus(db *gorm.DB, newList []*model.OperationsLargeScreenResourceRunningStatus, deleteIdList []int64) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if len(deleteIdList) > 0 {
			if err := tx.Delete(&model.OperationsLargeScreenResourceRunningStatus{}, deleteIdList).Error; err != nil {
				return err
			}
		}
		if len(newList) > 0 {
			if err := tx.Save(newList).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		logger.Logger().Error("saveTableOperationsLargeScreenResourceRunningStatus error: ", err)
		return
	}
}

// 运营大屏资源top表 operations_large_screen_resource_usage_top
func saveTableOperationsLargeScreenResourceUsageTop(db *gorm.DB, newList []*model.OperationsLargeScreenResourceUsageTop, deleteIdList []int64) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if len(deleteIdList) > 0 {
			if err := tx.Delete(&model.OperationsLargeScreenResourceUsageTop{}, deleteIdList).Error; err != nil {
				return err
			}
		}
		if len(newList) > 0 {
			if err := tx.Save(newList).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		logger.Logger().Error("saveTableOperationsLargeScreenResourceUsageTop error: ", err)
		return
	}
}

// 运营大屏资源分配表 operations_large_screen_resource_allocation_total
func saveTableOperationsLargeScreenResourceAllocationTotal(db *gorm.DB, newList []*model.OperationsLargeScreenResourceAllocationTotal, deleteIdList []int64) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if len(deleteIdList) > 0 {
			if err := tx.Delete(&model.OperationsLargeScreenResourceAllocationTotal{}, deleteIdList).Error; err != nil {
				return err
			}
		}
		if len(newList) > 0 {
			if err := tx.Save(newList).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		logger.Logger().Error("saveTableOperationsLargeScreenResourceAllocationTotal error: ", err)
		return
	}
}

func getPhysicalServerData(db *gorm.DB, region string) {
	param := "{\"condition\":{\"device\":[{\"field\":\"category_name\",\"operator\":\"$regex\",\"value\":\"服务器\"}]},\"fields\":{},\"page\":{\"start\":0,\"limit\":5000,\"sort\":\"\"}}"
	var response *form.CmdbResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.CmdbApi)
	if err != nil {
		logger.Logger().Error("getPhysicalServerData error: ", err)
		return
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "PhysicalServer").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getPhysicalServerData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.Info {
		newMap[v.Ip] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			Id:          oldMap[v.Ip],
			ResourceId:  v.Ip,
			Status:      v.RunStatus,
			Type:        "PhysicalServer",
			Region:      region,
			FailureTime: util.TimeToStr(util.StrToTime(util.FullTimeFmt, v.CpuUpdateTime), util.MonthDayTimeFmt),
			CreateTime:  v.CreateTime,
		}
		if v.RunStatus != 1 && v.RunStatus != 0 {
			newModel.Status = 2
		}
		if strutil.IsBlank(v.CpuUpdateTime) {
			newModel.FailureTime = "-----"
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableOperationsLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getNetworkDeviceData(db *gorm.DB, region string) {
	param := "{\"condition\":{\"device\":[{\"field\":\"business_type\",\"operator\":\"$regex\",\"value\":\"XGW\"},{\"field\":\"category_name\",\"operator\":\"$regex\",\"value\":\"服务器\"}]},\"fields\":{},\"page\":{\"start\":0,\"limit\":5000,\"sort\":\"\"}}"
	var response *form.CmdbResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.CmdbApi)
	if err != nil {
		logger.Logger().Error("getNetworkDeviceData error: ", err)
		return
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "NetworkDevice").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getNetworkDeviceData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.Info {
		newMap[v.Ip] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			Id:          oldMap[v.Ip],
			ResourceId:  v.Ip,
			Status:      v.RunStatus,
			Type:        "NetworkDevice",
			Region:      region,
			FailureTime: util.TimeToStr(util.StrToTime(util.FullTimeFmt, v.CpuUpdateTime), util.MonthDayTimeFmt),
			CreateTime:  v.CreateTime,
		}
		if v.RunStatus != 1 && v.RunStatus != 0 {
			newModel.Status = 2
		}
		if strutil.IsBlank(v.CpuUpdateTime) {
			newModel.FailureTime = "-----"
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableOperationsLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getEcsData(db *gorm.DB, region string) {
	param := external.InstanceRequest{
		CloudProductCode: "ECS",
		ResourceTypeCode: "instance",
		CurrPage:         "1",
		PageSize:         "9999",
		RegionCode:       region,
	}
	var response *external.InstanceResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
	if err != nil {
		logger.Logger().Error("getEcsData error: ", err)
		return
	}
	if response == nil || response.Msg != "success" {
		return
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "ECS").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getEcsData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.List {
		newMap[v.ResourceId] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			Id:          oldMap[v.ResourceId],
			ResourceId:  v.ResourceId,
			Type:        "ECS",
			Region:      v.RegionCode,
			FailureTime: util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:  util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if v.StatusDesc == "active" {
			newModel.Status = 1
		} else {
			newModel.Status = 2
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableOperationsLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getRdbData(db *gorm.DB, region string) {
	var paramList = []*external.InstanceRequest{
		{
			CloudProductCode: "RDB",
			ResourceTypeCode: "mysql",
			CurrPage:         "1",
			PageSize:         "9999",
			RegionCode:       region,
		},
		{
			CloudProductCode: "RDB",
			ResourceTypeCode: "dm",
			CurrPage:         "1",
			PageSize:         "9999",
			RegionCode:       region,
		},
		{
			CloudProductCode: "RDB",
			ResourceTypeCode: "pg",
			CurrPage:         "1",
			PageSize:         "9999",
			RegionCode:       region,
		},
	}
	var responseDataList []*external.InstanceList
	for _, param := range paramList {
		var response *external.InstanceResponse
		_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
		if err != nil {
			logger.Logger().Error("getRdbData error: ", err)
			continue
		}
		if response == nil || response.Msg != "success" {
			continue
		}
		responseDataList = append(responseDataList, response.Data.List...)
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err := db.Where("region = ? AND type = ?", region, "数据库").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getRdbData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range responseDataList {
		newMap[v.ResourceId] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			Id:          oldMap[v.ResourceId],
			ResourceId:  v.ResourceId,
			Type:        "数据库",
			Region:      v.RegionCode,
			FailureTime: util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:  util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if v.StatusDesc == "运行中" || v.StatusDesc == "创建中" || v.StatusDesc == "扩缩容中" || v.StatusDesc == "主库运行中" ||
			v.StatusDesc == "停止" || v.StatusDesc == "维护中" || v.StatusDesc == "重启中" || v.StatusDesc == "删除中" ||
			v.StatusDesc == "备份中" || v.StatusDesc == "备份恢复中" || v.StatusDesc == "实例被冻结" || v.StatusDesc == "从库重建中" ||
			v.StatusDesc == "初始化中" || v.StatusDesc == "启动中" || v.StatusDesc == "停止中" || v.StatusDesc == "正在删除状态" {
			newModel.Status = 1
		} else {
			newModel.Status = 2
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableOperationsLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getBmsData(db *gorm.DB, region string) {
	var paramList = []*external.InstanceRequest{
		{
			CloudProductCode: "BMS",
			ResourceTypeCode: "TBMS",
			CurrPage:         "1",
			PageSize:         "9999",
			RegionCode:       region,
		},
		{
			CloudProductCode: "BMS",
			ResourceTypeCode: "EBMS",
			CurrPage:         "1",
			PageSize:         "9999",
			RegionCode:       region,
		},
	}
	var responseDataList []*external.InstanceList
	for _, param := range paramList {
		var response *external.InstanceResponse
		_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
		if err != nil {
			logger.Logger().Error("getBmsData error: ", err)
			continue
		}
		if response == nil || response.Msg != "success" {
			continue
		}
		responseDataList = append(responseDataList, response.Data.List...)
	}
	//查询数据库的旧数据
	var oldList []*model.OperationsLargeScreenResourceRunningStatus
	if err := db.Where("region = ? AND type = ?", region, "裸金属").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getBmsData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.OperationsLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range responseDataList {
		newMap[v.ResourceId] = nil
		var newModel = &model.OperationsLargeScreenResourceRunningStatus{
			Id:          oldMap[v.ResourceId],
			ResourceId:  v.ResourceId,
			Type:        "裸金属",
			Region:      v.RegionCode,
			FailureTime: util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:  util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if v.StatusDesc == "running" {
			newModel.Status = 1
		} else {
			newModel.Status = 2
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableOperationsLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

var OperationsCodeList = []string{
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

func getUsageData(db *gorm.DB, region string, codeList []string) {
	var newList []*model.OperationsLargeScreenResourceUsageTop
	var deleteIdList []int64
	for _, code := range codeList {
		var resourceType, attribute, unit string
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
			unit = "个"
		case "eip_upstream_bandwidth_usage":
			resourceType = "EIP"
			attribute = "带宽"
		}
		//查询指标
		var monitorItem *model.MonitorItem
		if err := global.DB.Where("metric_name = ?", code).Find(&monitorItem).Error; err != nil {
			logger.Logger().Error("getUsageData error: ", err)
			continue
		}
		pql := fmt.Sprintf(constant.TopExpr, "10", strings.ReplaceAll(monitorItem.MetricsLinux, constant.MetricLabel, ""))
		prometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(pql)
		response := &form.PrometheusResponse{}
		_, err := httputil.GetHttpClient().R().SetResult(&response).Get(prometheusUrl)
		if err != nil {
			logger.Logger().Error("getUsageData error: ", err)
			continue
		}
		if response == nil || response.Data == nil || response.Data.Result == nil || len(response.Data.Result) == 0 {
			logger.Logger().Infof("query prometheus empty, pql: %s, origin response:%s", pql, response)
			continue
		}
		//查询数据库的旧数据
		var oldList []*model.OperationsLargeScreenResourceUsageTop
		if err = db.Where("region = ? AND type = ? AND attribute = ?", region, resourceType, attribute).Find(&oldList).Error; err != nil {
			logger.Logger().Error("getUsageData error: ", err)
			continue
		}
		var oldMap = make(map[string]int64)
		for _, v := range oldList {
			oldMap[v.ResourceId] = v.Id
		}
		if strutil.IsBlank(unit) {
			unit = monitorItem.Unit
		}
		var newMap = make(map[string]interface{})
		for _, v := range response.Data.Result {
			resourceId := v.Metric[constant.INSTANCE]
			number := changeDecimal(v.Value[1].(string))
			newMap[resourceId] = nil
			newModel := &model.OperationsLargeScreenResourceUsageTop{
				Id:         oldMap[resourceId],
				ResourceId: resourceId,
				Type:       resourceType,
				Attribute:  attribute,
				Number:     number,
				Unit:       unit,
				Region:     config.Cfg.Common.RegionName,
			}
			newList = append(newList, newModel)
		}
		for _, old := range oldList {
			if _, ok := newMap[old.ResourceId]; !ok {
				deleteIdList = append(deleteIdList, old.Id)
			}
		}
	}
	saveTableOperationsLargeScreenResourceUsageTop(db, newList, deleteIdList)
}

// 数据保留两位小数
func changeDecimal(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return math.Trunc(v*1e2+0.5) * 1e-2
}

func getResourceAllocationTotalDate(db *gorm.DB, region string) {
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
		logger.Logger().Error("getResourceAllocationTotalDate error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[fmt.Sprintf("%s-%s", v.Type, v.Attribute)] = v.Id
	}
	var newMap = make(map[string]int64)
	for i, v := range newList {
		newMap[fmt.Sprintf("%s-%s", v.Type, v.Attribute)] = v.Id
		newList[i].Id = oldMap[fmt.Sprintf("%s-%s", v.Type, v.Attribute)]
	}
	var deleteIdList []int64
	for _, v := range oldList {
		if _, ok := newMap[fmt.Sprintf("%s-%s", v.Type, v.Attribute)]; !ok {
			deleteIdList = append(deleteIdList, v.Id)
		}
	}
	saveTableOperationsLargeScreenResourceAllocationTotal(db, newList, deleteIdList)
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
			Allocation: float64(response.Data.UsedRam) / 1000 / 1000 / 1000,
			Total:      float64(response.Data.Ram) / 1000 / 1000 / 1000,
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
			Allocation: float64(response.Data.AllocatedCapacityGb) / 1024,
			Total:      float64(response.Data.TotalCapacityGb) / 1024,
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
	//计算当天的23时59分59秒
	now := util.GetNow().Unix()
	t := now + (86400 - (now-57600)%86400)
	timeString := strconv.FormatInt(t, 10)
	//总带宽
	totalPql := "max_over_time(sum(avg(eip_config_upstream_bandwidth{eipType='external_eip'})by(instance))[1d:1m])"
	totalPrometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(totalPql) + "&time=" + timeString
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
	upstreamPql := "max_over_time(sum(eip_upstream_bits_rate{eipType='external_eip'})[1d:1m])"
	upstreamPrometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(upstreamPql) + "&time=" + timeString
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
	downstreamPql := "max_over_time(sum(eip_downstream_bits_rate{eipType='external_eip'})[1d:1m])"
	downstreamPrometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(downstreamPql) + "&time=" + timeString
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
		var remoteDatabase *model.LargeScreenRemoteDatabase
		if err := global.DB.Where("region = ?", config.Cfg.Common.RegionName).First(&remoteDatabase).Error; err != nil {
			logger.Logger().Error("getPushRemoteDb error: ", err)
			return nil, err
		}
		dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", remoteDatabase.User, remoteDatabase.Pass, remoteDatabase.Host, remoteDatabase.Port, remoteDatabase.Db)
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
