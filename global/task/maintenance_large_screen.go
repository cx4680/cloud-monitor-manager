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
	"gorm.io/gorm"
	"net/url"
	"strconv"
	"strings"
)

func MaintenanceLargeScreen() BusinessTaskDTO {
	task := func() {
		region := config.Cfg.Common.RegionName
		//连接线上远程数据库
		db, err := getPushRemoteDb()
		if err != nil {
			logger.Logger().Error(err)
			return
		}
		//查询告警
		go getLargeScreenAlert(db, region)

		//设备
		go getDeviceData(db, region)

		//XGW
		go getXgwData(db, region)

		//云服务
		go getPodData(db, region)

		//ECS
		go getEcsDataMaintenance(db, region)

		//查询ECS top数据
		go getUsageDataMaintenance(db, region, MaintenanceCodeList)

		//SLB
		go getSlbData(db, region)

		//EIP
		go getEipData(db, region)

		//RDB
		go getRdbDataMaintenance(db, region)

		//存储磁盘
		go getStorageDiskData(db, region)

		//存储池
		go getStoragePoolData(db, region)
	}

	return BusinessTaskDTO{
		Cron: "0 */5 * * * ?",
		Name: "MaintenanceLargeScreen",
		Task: task,
	}
}

// 运维大屏告警数量表 maintenance_large_screen_resource_alert_total
func saveTableMaintenanceLargeScreenResourceAlertTotal(db *gorm.DB, resourceAlertTotal []*model.MaintenanceLargeScreenResourceAlertTotal) {
	if err := db.Save(resourceAlertTotal).Error; err != nil {
		logger.Logger().Error("saveTableMaintenanceLargeScreenResourceAlertTotal error: ", err)
	}
}

// 运维大屏告警列表 maintenance_large_screen_resource_alert_list
func saveTableMaintenanceLargeScreenResourceAlertList(db *gorm.DB, resourceAlertList []*model.MaintenanceLargeScreenResourceAlertList, deleteIdList []int64) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if len(deleteIdList) > 0 {
			if err := tx.Delete(&model.MaintenanceLargeScreenResourceAlertList{}, deleteIdList).Error; err != nil {
				return err
			}
		}
		if len(resourceAlertList) > 0 {
			if err := tx.Save(resourceAlertList).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		logger.Logger().Error("saveTableMaintenanceLargeScreenResourceAlertList error: ", err)
	}
}

// 运维大屏资源状态表 maintenance_large_screen_resource_running_status
func saveTableMaintenanceLargeScreenResourceRunningStatus(db *gorm.DB, newList []*model.MaintenanceLargeScreenResourceRunningStatus, deleteIdList []int64) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if len(deleteIdList) > 0 {
			if err := tx.Delete(&model.MaintenanceLargeScreenResourceRunningStatus{}, deleteIdList).Error; err != nil {
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
		logger.Logger().Error("saveTableMaintenanceLargeScreenResourceRunningStatus error: ", err)
	}
}

// 运维大屏资源分配表 maintenance_large_screen_resource_allocation_total
func saveTableMaintenanceLargeScreenResourceAllocationTotal(db *gorm.DB, newList []*model.MaintenanceLargeScreenResourceAllocationTotal, deleteIdList []int64) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if len(deleteIdList) > 0 {
			if err := tx.Delete(&model.MaintenanceLargeScreenResourceAllocationTotal{}, deleteIdList).Error; err != nil {
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
		logger.Logger().Error("saveTableMaintenanceLargeScreenResourceAllocationTotal error: ", err)
	}
}

// 运维大屏资源top表 maintenance_large_screen_resource_usage_top
func saveTableMaintenanceLargeScreenResourceUsageTop(db *gorm.DB, newList []*model.MaintenanceLargeScreenResourceUsageTop, deleteIdList []int64) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if len(deleteIdList) > 0 {
			if err := tx.Delete(&model.MaintenanceLargeScreenResourceUsageTop{}, deleteIdList).Error; err != nil {
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
		logger.Logger().Error("saveTableMaintenanceLargeScreenResourceUsageTop error: ", err)
	}
}

func getLargeScreenAlert(db *gorm.DB, region string) {
	var response *form.LargeScreenAlertResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.LargeScreenAlertApi)
	if err != nil {
		logger.Logger().Error("getLargeScreenAlert error: ", err)
		return
	}
	logger.Logger().Infof("LargeScreenAlertResponse, P1: %v, P2: %v, P3: %v, P4: %v", response.P1, response.P2, response.P3, response.P4)
	var oldResourceAlertTotal []*model.MaintenanceLargeScreenResourceAlertTotal
	if err = db.Where("region = ?", region).Find(&oldResourceAlertTotal).Error; err != nil {
		logger.Logger().Error("getLargeScreenAlert error: ", err)
		return
	}
	var oldAlertLevelMap = make(map[string]int64)
	for _, v := range oldResourceAlertTotal {
		oldAlertLevelMap[v.AlertLevel] = v.Id
	}
	var resourceAlertTotal = []*model.MaintenanceLargeScreenResourceAlertTotal{
		{
			Id:          oldAlertLevelMap["紧急告警"],
			AlertLevel:  "紧急告警",
			AlertNumber: response.P1,
			Region:      region,
		},
		{
			Id:          oldAlertLevelMap["重要告警"],
			AlertLevel:  "重要告警",
			AlertNumber: response.P2,
			Region:      region,
		},
		{
			Id:          oldAlertLevelMap["次要告警"],
			AlertLevel:  "次要告警",
			AlertNumber: response.P3,
			Region:      region,
		},
		{
			Id:          oldAlertLevelMap["提示告警"],
			AlertLevel:  "提示告警",
			AlertNumber: response.P4,
			Region:      region,
		},
	}
	go saveTableMaintenanceLargeScreenResourceAlertTotal(db, resourceAlertTotal)

	var oldResourceAlertList []*model.MaintenanceLargeScreenResourceAlertList
	if err = db.Where("region = ?", region).Find(&oldResourceAlertList).Error; err != nil {
		logger.Logger().Error("getLargeScreenAlert error: ", err)
		return
	}
	var oldAlertMap = make(map[string]int64)
	for _, v := range oldResourceAlertList {
		oldAlertMap[fmt.Sprintf("%v-%v", v.RuleName, v.ResourceId)] = v.Id
	}
	var resourceAlertList []*model.MaintenanceLargeScreenResourceAlertList
	var newMap = make(map[string]interface{})
	for _, v := range response.List {
		newMap[v.Ip] = nil
		split := strings.Split(v.Target, ":")
		if len(split) > 1 {
			v.Target = split[1]
		}
		resourceAlertList = append(resourceAlertList, &model.MaintenanceLargeScreenResourceAlertList{
			Id:           oldAlertMap[fmt.Sprintf("%v-%v", v.Name, v.Target)],
			AlertLevel:   AlertLevelMap[v.Severity],
			RuleName:     v.Name,
			ResourceId:   v.Target,
			ResourceType: v.Type,
			AlertTime:    util.TimeToFullTimeFmtStr(v.StartsAt),
			Region:       region,
		})
	}
	var deleteIdList []int64
	for _, v := range oldResourceAlertList {
		if _, ok := newMap[fmt.Sprintf("%v-%v", v.RuleName, v.ResourceId)]; !ok {
			deleteIdList = append(deleteIdList, v.Id)
		}
	}
	go saveTableMaintenanceLargeScreenResourceAlertList(db, resourceAlertList, deleteIdList)
}

var AlertLevelMap = map[string]string{"1": "紧急", "2": "重要", "3": "次要", "4": "提示"}

var DeviceStatisticsParamList = []string{
	"{\"condition\":{\"device\":[{\"field\":\"category_name\",\"operator\":\"$regex\",\"value\":\"服务器\"}]},\"fields\":{},\"page\":{\"start\":0,\"limit\":5000,\"sort\":\"\"}}",
	"{\"condition\":{\"device\":[{\"field\":\"category_name\",\"operator\":\"$regex\",\"value\":\"交换机\"}]},\"fields\":{},\"page\":{\"start\":0,\"limit\":5000,\"sort\":\"\"}}",
}

func getDeviceData(db *gorm.DB, region string) {
	var responseDataInfo []*form.CmdbInfo
	for _, param := range DeviceStatisticsParamList {
		var response *form.CmdbResponse
		_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.CmdbApi)
		if err != nil {
			logger.Logger().Error("getDeviceData: ", err)
			return
		}
		if response != nil && response.Data != nil {
			responseDataInfo = append(responseDataInfo, response.Data.Info...)
		}
	}
	//查询数据库的旧数据
	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err := db.Where("region = ? AND type = ?", region, "Device").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getDeviceData: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range responseDataInfo {
		newMap[v.Ip] = nil
		var newModel = &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.Ip],
			ResourceId:   v.Ip,
			ResourceName: v.Ip,
			Status:       v.RunStatus,
			Type:         "Device",
			Region:       region,
			FailureTime:  util.TimeToStr(util.StrToTime(util.FullTimeFmt, v.CpuUpdateTime), util.MonthDayTimeFmt),
			CreateTime:   v.CreateTime,
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
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getXgwData(db *gorm.DB, region string) {
	param := "{\"condition\":{\"device\":[{\"field\":\"business_type\",\"operator\":\"$regex\",\"value\":\"XGW\"},{\"field\":\"category_name\",\"operator\":\"$regex\",\"value\":\"服务器\"}]},\"fields\":{},\"page\":{\"start\":0,\"limit\":5000,\"sort\":\"\"}}"
	var response *form.CmdbResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.CmdbApi)
	if err != nil {
		logger.Logger().Error("getXgwData error: ", err)
		return
	}
	//查询数据库的旧数据
	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "XGW").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getXgwData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.Info {
		newMap[v.Ip] = nil
		var newModel = &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.Ip],
			ResourceId:   v.Ip,
			ResourceName: v.Ip,
			Status:       v.RunStatus,
			Type:         "XGW",
			Region:       region,
			FailureTime:  util.TimeToStr(util.StrToTime(util.FullTimeFmt, v.CpuUpdateTime), util.MonthDayTimeFmt),
			CreateTime:   v.CreateTime,
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
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getPodData(db *gorm.DB, region string) {
	cocClusterBigScreenApi := config.Cfg.Common.CocClusterBigScreenApi
	var deploymentResponse *form.CocClusterDeployment
	_, err := httputil.GetHttpClient().R().SetResult(&deploymentResponse).Get(cocClusterBigScreenApi + "/deployments?page=1&size=50000")
	if err != nil {
		logger.Logger().Error("getPodData error: ", err)
		return
	}
	var statefulSetResponse *form.CocClusterStatefulSet
	_, err = httputil.GetHttpClient().R().SetResult(&statefulSetResponse).Get(cocClusterBigScreenApi + "/statefulsets?page=1&size=50000")
	if err != nil {
		logger.Logger().Error("getPodData error: ", err)
		return
	}

	var cocClusterList []*form.CocCluster
	cocClusterList = append(cocClusterList, deploymentResponse.List...)
	cocClusterList = append(cocClusterList, statefulSetResponse.List...)

	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "Server").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getPodData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}

	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range cocClusterList {
		if strings.Contains(v.Namespace, "tenant") {
			continue
		}
		newMap[v.Name] = nil
		resourceRunningStatus := &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.Name],
			ResourceId:   v.Name,
			ResourceName: v.Name,
			Type:         "Server",
			Region:       region,
			CreateTime:   util.TimeToFullTimeFmtStr(v.CreateTime),
		}
		if v.InstanceTotal == v.InstanceReady {
			resourceRunningStatus.Status = 1
		} else {
			//查询故障时间
			var labelList []string
			for key, value := range v.Labels {
				labelList = append(labelList, key+"="+value)
			}
			param := "label=" + strings.Join(labelList, ",")
			var pods *form.CocClusterPod
			_, err = httputil.GetHttpClient().R().SetResult(&pods).Get(cocClusterBigScreenApi + "/pods?page=1&size=10&namespace=" + v.Namespace + "&" + param)
			if err != nil {
				logger.Logger().Error(err)
			}
			if pods != nil && pods.Data != nil {
				resourceRunningStatus.Status = 2
				if strutil.IsNotBlank(pods.Data.Time) {
					resourceRunningStatus.FailureTime = util.TimeToStr(util.StrToTime(util.FullTimeFmt, pods.Data.Time), util.MonthDayTimeFmt)
				} else {
					resourceRunningStatus.FailureTime = "-----"
				}
			} else {
				resourceRunningStatus.Status = 1
			}
		}
		newList = append(newList, resourceRunningStatus)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getEcsDataMaintenance(db *gorm.DB, region string) {
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
		logger.Logger().Error("getPodData error: ", err)
		return
	}
	if response == nil || response.Msg != "success" {
		return
	}
	//查询数据库的旧数据
	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "ECS").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getPodData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.List {
		newMap[v.ResourceId] = nil
		var newModel = &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.ResourceId],
			ResourceId:   v.ResourceId,
			ResourceName: v.ResourceName,
			Type:         "ECS",
			Region:       v.RegionCode,
			FailureTime:  util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:   util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
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
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getSlbData(db *gorm.DB, region string) {
	param := external.InstanceRequest{
		CloudProductCode: "SLB",
		ResourceTypeCode: "instance",
		CurrPage:         "1",
		PageSize:         "9999",
		RegionCode:       region,
	}
	var response *external.InstanceResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
	if err != nil {
		logger.Logger().Error("getSlbData error: ", err)
		return
	}
	if response == nil || response.Msg != "success" {
		logger.Logger().Error("getSlbData error: ", err)
		return
	}
	//查询数据库的旧数据
	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "SLB").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getSlbData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.List {
		if v.RegionCode != config.Cfg.Common.RegionName {
			continue
		}
		newMap[v.ResourceId] = nil
		var newModel = &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.ResourceId],
			ResourceId:   v.ResourceId,
			ResourceName: v.ResourceName,
			Type:         "SLB",
			Region:       v.RegionCode,
			FailureTime:  util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:   util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if v.StatusDesc == "failed" {
			newModel.Status = 2
		} else {
			newModel.Status = 1
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getEipData(db *gorm.DB, region string) {
	param := external.InstanceRequest{
		CloudProductCode: "EIP",
		ResourceTypeCode: "instance",
		CurrPage:         "1",
		PageSize:         "9999",
		RegionCode:       region,
	}
	var response *external.InstanceResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.Rc)
	if err != nil {
		logger.Logger().Error("getEipData error: ", err)
		return
	}
	if response == nil || response.Msg != "success" {
		logger.Logger().Error("getEipData error: ", err)
		return
	}
	//查询数据库的旧数据
	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "EIP").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getEipData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.List {
		newMap[v.ResourceId] = nil
		var newModel = &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.ResourceId],
			ResourceId:   v.ResourceId,
			ResourceName: v.ResourceName,
			Type:         "EIP",
			Region:       v.RegionCode,
			FailureTime:  util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:   util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
		}
		if v.StatusDesc == "错误" {
			newModel.Status = 2
		} else {
			newModel.Status = 1
		}
		newList = append(newList, newModel)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getRdbDataMaintenance(db *gorm.DB, region string) {
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
			logger.Logger().Error("getRdbDataMaintenance error: ", err)
			continue
		}
		if response == nil || response.Msg != "success" {
			continue
		}
		responseDataList = append(responseDataList, response.Data.List...)
	}
	//查询数据库的旧数据
	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err := db.Where("region = ? AND type = ?", region, "云数据库").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getRdbDataMaintenance error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}
	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range responseDataList {
		newMap[v.ResourceId] = nil
		var newModel = &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.ResourceId],
			ResourceId:   v.ResourceId,
			ResourceName: v.ResourceName,
			Type:         "云数据库",
			Region:       v.RegionCode,
			FailureTime:  util.TimestampToFmtStr(int64(v.UpdateTime)/1000, "01-02 15:04:05"),
			CreateTime:   util.TimestampToFullTimeFmtStr(int64(v.CreateTime) / 1000),
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
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

var isStorageLogin bool

func getStorageDiskData(db *gorm.DB, region string) {
	//新增存储登录信息接口
	if !isStorageLogin {
		var storageLoginList []*model.LargeScreenStorageLogin
		if err := global.DB.Where("region = ?", region).Find(&storageLoginList).Error; err != nil {
			logger.Logger().Error("getStorageDiskData error: ", err)
			return
		}
		for _, v := range storageLoginList {
			param := map[string]string{"vendor": v.Vendor, "type": v.Type, "username": v.Username, "password": v.Password, "manageUrl": v.Password}
			response, err := httputil.GetHttpClient().R().SetBody(param).Post(config.Cfg.Common.EbsApi + "/storage-disk-status/loginInfo")
			logger.Logger().Info("LargeScreenStorageLogin response: ", response)
			if err != nil {
				logger.Logger().Error("LargeScreenStorageLogin error: ", err)
				continue
			}
			isStorageLogin = true
		}
	}

	var response *form.DiskResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.EbsApi + "/storage-disk-status")
	if err != nil {
		logger.Logger().Error("getStorageDiskData error: ", err)
		return
	}
	if response == nil || response.Data == nil {
		return
	}

	var oldList []*model.MaintenanceLargeScreenResourceRunningStatus
	if err = db.Where("region = ? AND type = ?", region, "StorageDisk").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getStorageDiskData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[v.ResourceId] = v.Id
	}

	var newList []*model.MaintenanceLargeScreenResourceRunningStatus
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.Stats {
		newMap[v.DiskId] = nil
		resourceRunningStatus := &model.MaintenanceLargeScreenResourceRunningStatus{
			Id:           oldMap[v.DiskId],
			ResourceId:   v.DiskId,
			ResourceName: v.DiskId,
			Type:         "StorageDisk",
			Region:       region,
		}
		if v.DiskStatus == "health" {
			resourceRunningStatus.Status = 1
		} else if v.DiskStatus == "unhealth" {
			resourceRunningStatus.Status = 2
		} else {
			resourceRunningStatus.Status = 0
		}
		if strutil.IsNotBlank(v.FaultTime) {
			resourceRunningStatus.FailureTime = util.TimeToStr(util.StrToTime(util.FullTimeFmt, v.FaultTime), util.MonthDayTimeFmt)
		} else {
			resourceRunningStatus.FailureTime = "-----"
		}
		newList = append(newList, resourceRunningStatus)
	}
	var deleteIdList []int64
	for _, old := range oldList {
		if _, ok := newMap[old.ResourceId]; !ok {
			deleteIdList = append(deleteIdList, old.Id)
		}
	}
	saveTableMaintenanceLargeScreenResourceRunningStatus(db, newList, deleteIdList)
}

func getStoragePoolData(db *gorm.DB, region string) {
	var response *form.EbsResponse
	_, err := httputil.GetHttpClient().R().SetResult(&response).Get(config.Cfg.Common.EbsApi + "/storage-pool-status")
	if err != nil {
		logger.Logger().Error("getStoragePoolData error: ", err)
		return
	}
	if response == nil || response.Data == nil {
		return
	}

	var oldList []*model.MaintenanceLargeScreenResourceAllocationTotal
	if err = db.Where("region = ? AND type = ?", region, "StoragePool").Find(&oldList).Error; err != nil {
		logger.Logger().Error("getStoragePoolData error: ", err)
		return
	}
	var oldMap = make(map[string]int64)
	for _, v := range oldList {
		oldMap[fmt.Sprintf("%s-%s", v.Type, v.Name)] = v.Id
	}

	var newList []*model.MaintenanceLargeScreenResourceAllocationTotal
	var newMap = make(map[string]interface{})
	for _, v := range response.Data.PoolStatus {
		newMap[fmt.Sprintf("%s-%s", "StoragePool", v.Name)] = nil
		resourceRunningStatus := &model.MaintenanceLargeScreenResourceAllocationTotal{
			Id:         oldMap[fmt.Sprintf("%s-%s", "StoragePool", v.Name)],
			Type:       "StoragePool",
			Name:       v.Name,
			Allocation: v.UsedCapacity,
			Total:      v.TotalCapacity,
			Unit:       region,
			Time:       v.LicenseLeftTime,
			Region:     region,
		}
		_, err = strconv.Atoi(v.LicenseLeftTime)
		if err == nil {
			resourceRunningStatus.Time += "天"
		}
		newList = append(newList, resourceRunningStatus)
	}
	var deleteIdList []int64
	for _, v := range oldList {
		if _, ok := newMap[fmt.Sprintf("%s-%s", v.Type, v.Name)]; !ok {
			deleteIdList = append(deleteIdList, v.Id)
		}
	}
	saveTableMaintenanceLargeScreenResourceAllocationTotal(db, newList, deleteIdList)
}

var MaintenanceCodeList = []string{
	"ecs_cpu_base_usage",    //ECS cpu使用率
	"ecs_memory_base_usage", //ECS 内存使用率
}

func getUsageDataMaintenance(db *gorm.DB, region string, codeList []string) {
	var newList []*model.MaintenanceLargeScreenResourceUsageTop
	var deleteIdList []int64
	var unit string
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
			unit = "个"
		case "eip_upstream_bandwidth_usage":
			resourceType = "EIP"
			attribute = "带宽"
		}
		//查询指标
		var monitorItem *model.MonitorItem
		if err := global.DB.Where("metric_name = ?", code).Find(&monitorItem).Error; err != nil {
			logger.Logger().Error("getUsageDataMaintenance error: ", err)
			continue
		}
		pql := fmt.Sprintf(constant.TopExpr, "5", strings.ReplaceAll(monitorItem.MetricsLinux, constant.MetricLabel, ""))
		prometheusUrl := config.Cfg.Prometheus.Url + config.Cfg.Prometheus.Query + url.QueryEscape(pql)
		response := &form.PrometheusResponse{}
		_, err := httputil.GetHttpClient().R().SetResult(&response).Get(prometheusUrl)
		if err != nil {
			logger.Logger().Error("getUsageDataMaintenance error: ", err)
			continue
		}
		if response == nil || response.Data == nil || response.Data.Result == nil || len(response.Data.Result) == 0 {
			logger.Logger().Infof("query prometheus empty, pql: %s, origin response:%s", pql, response)
			continue
		}
		//查询数据库的旧数据
		var oldList []*model.MaintenanceLargeScreenResourceUsageTop
		if err = db.Where("region = ? AND type = ? AND attribute = ?", region, resourceType, attribute).Find(&oldList).Error; err != nil {
			logger.Logger().Error("getUsageDataMaintenance error: ", err)
			continue
		}
		var oldMap = make(map[string]int64)
		for _, v := range oldList {
			oldMap[v.ResourceId] = v.Id
		}
		var newMap = make(map[string]interface{})
		for _, v := range response.Data.Result {
			resourceId := v.Metric[constant.INSTANCE]
			number := changeDecimal(v.Value[1].(string))
			newMap[resourceId] = nil
			newModel := &model.MaintenanceLargeScreenResourceUsageTop{
				Id:         oldMap[resourceId],
				ResourceId: resourceId,
				Type:       resourceType,
				Attribute:  attribute,
				Number:     number,
				Unit:       monitorItem.Unit,
				Region:     config.Cfg.Common.RegionName,
			}
			if strutil.IsBlank(newModel.Unit) {
				newModel.Unit = unit
			}
			newList = append(newList, newModel)
		}
		for _, old := range oldList {
			if _, ok := newMap[old.ResourceId]; !ok {
				deleteIdList = append(deleteIdList, old.Id)
			}
		}
	}
	saveTableMaintenanceLargeScreenResourceUsageTop(db, newList, deleteIdList)
}
