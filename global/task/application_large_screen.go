package task

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/model"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"time"
)

func ApplicationLargeScreen() BusinessTaskDTO {
	task := func() {
		ossResources, efsResources, err := getResourcesByTag()
		if err != nil {
			logger.Logger().Error(err)
			return
		}
		now := util.GetNow()
		deleteTime := util.TimestampToFullTimeFmtStr(now.Unix() + (86400 - (now.Unix()-57600)%86400) - 8*86400)
		go getOssData(now, ossResources)
		go getEfsData(now, efsResources)
		if err = global.DB.Where("create_time < ?", deleteTime).Delete(&model.LargeScreenResourceStorage{}).Error; err != nil {
			logger.Logger().Error(err)
		}
	}

	return BusinessTaskDTO{
		Cron: "0 */5 * * * ?",
		Name: "ApplicationLargeScreen",
		Task: task,
	}
}

func getOssData(now time.Time, resources []*form.LargeScreenResource) {
	var old []*model.LargeScreenResourceStorage
	if err := global.DB.Where("type = ? AND time = ?", constant.CloudProductCodeOss, util.TimeToStr(now, util.DayTimeFmt)).Find(&old).Error; err != nil {
		logger.Logger().Error(err)
	}
	var oldMap = make(map[string]*model.LargeScreenResourceStorage)
	for _, v := range old {
		oldMap[v.ResourceId] = v
	}
	var newList []*model.LargeScreenResourceStorage
	var newMap = make(map[string]interface{})
	for _, v := range resources {
		var additional = &form.LargeScreenResourceStorageAdditional{}
		jsonutil.ToObject(v.Additional, additional)
		newMap[v.ResourceInstanceId] = nil
		resourceStorage := oldMap[v.ResourceInstanceId]
		if resourceStorage != nil {
			if resourceStorage.Value < additional.Storage {
				resourceStorage.Value = additional.Storage
			}
		} else {
			resourceStorage = &model.LargeScreenResourceStorage{
				ResourceId: v.ResourceInstanceId,
				Type:       constant.CloudProductCodeOss,
				Time:       util.TimeToStr(now, util.DayTimeFmt),
				CreateTime: now,
				Value:      additional.Storage,
			}
		}
		newList = append(newList, resourceStorage)
	}
	if err := global.DB.Save(newList).Error; err != nil {
		logger.Logger().Error(err)
	}
}

func getEfsData(now time.Time, resources []string) {
	var response *form.LargeScreenStorageResponse
	param := map[string][]string{"instanceIds": resources}
	_, err := httputil.GetHttpClient().R().SetResult(&response).SetBody(param).Post(config.Cfg.Common.EfsApi)
	if err != nil || response == nil || response.Data == nil {
		logger.Logger().Error(err)
		return
	}
	var old []*model.LargeScreenResourceStorage
	if err = global.DB.Where("type = ? AND time = ?", constant.CloudProductCodeEfs, util.TimeToStr(time.Now(), util.DayTimeFmt)).Find(&old).Error; err != nil {
		logger.Logger().Error(err)
	}
	var oldMap = make(map[string]*model.LargeScreenResourceStorage)
	for _, v := range old {
		oldMap[v.ResourceId] = v
	}
	var newList []*model.LargeScreenResourceStorage
	var newMap = make(map[string]interface{})
	for _, v := range response.Data {
		newMap[v.InstanceId] = nil
		resourceStorage := oldMap[v.InstanceId]
		if resourceStorage != nil {
			if resourceStorage.Value < v.Size {
				resourceStorage.Value = v.Size
			}
		} else {
			resourceStorage = &model.LargeScreenResourceStorage{
				ResourceId: v.InstanceId,
				Type:       constant.CloudProductCodeEfs,
				Time:       util.TimeToStr(now, util.DayTimeFmt),
				CreateTime: now,
				Value:      v.Size,
			}
		}
		newList = append(newList, resourceStorage)
	}
	if len(newList) > 0 {
		if err = global.DB.Save(newList).Error; err != nil {
			logger.Logger().Error(err)
		}
	}
}

func getResourcesByTag() ([]*form.LargeScreenResource, []string, error) {
	var resourceTagResponse *form.LargeScreenResourceTagResponse
	_, err := httputil.GetHttpClient().R().SetResult(&resourceTagResponse).Get(config.Cfg.Common.ResourceTagApi + "/list-all")
	if err != nil || resourceTagResponse == nil || resourceTagResponse.Module == nil {
		return nil, nil, err
	}
	var tagKeyIdList []string
	for _, v := range resourceTagResponse.Module.List {
		tagKeyIdList = append(tagKeyIdList, v.YagKeyId)
	}
	var resourceResponse *form.LargeScreenResourceResponse
	param := map[string][]string{"tagKeyIdList": tagKeyIdList}
	_, err = httputil.GetHttpClient().R().SetResult(&resourceResponse).SetBody(param).Post(config.Cfg.Common.ResourceTagApi + "/list-resource-by-tag-id")
	if err != nil || resourceResponse == nil || resourceResponse.Module == nil {
		return nil, nil, err
	}
	var ossResources []*form.LargeScreenResource
	var efsResources []string
	for _, v := range resourceResponse.Module.List {
		if v.CloudProductCode == constant.CloudProductCodeOss && v.ResourceTypeCode == constant.ResourceTypeCodeBucket {
			ossResources = append(ossResources, v)
		}
		if v.CloudProductCode == constant.CloudProductCodeEfs && v.ResourceTypeCode == constant.ResourceTypeCodeShares {
			efsResources = append(efsResources, v.ResourceInstanceId)
		}
	}
	return ossResources, efsResources, nil
}
