package controller

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LargeScreenCtl struct {
	service *service.LargeScreenService
}

func NewLargeScreenCtl() *LargeScreenCtl {
	return &LargeScreenCtl{service: service.NewLargeScreenService()}
}

func (ctl *LargeScreenCtl) Tags(c *gin.Context) {
	data, err := ctl.service.Tags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceOverview(c *gin.Context) {
	tag := c.Query("tagKeyId")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	data, err := ctl.service.ResourceOverview(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceAlert(c *gin.Context) {
	tag := c.Query("tagKeyId")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	data, err := ctl.service.ResourceAlert(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceEcs(c *gin.Context) {
	tag := c.Query("tagKeyId")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	data, err := ctl.service.ResourceEcs(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceEcsTop(c *gin.Context) {
	tag := c.Query("tagKeyId")
	item := c.Query("item")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	data, err := ctl.service.ResourceEcsTop(tag, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceEip(c *gin.Context) {
	tag := c.Query("tagKeyId")
	interval := c.Query("interval")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	if strutil.IsBlank(interval) {
		interval = "1h"
	}
	data, err := ctl.service.ResourceEip(tag, interval)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceNat(c *gin.Context) {
	tag := c.Query("tagKeyId")
	interval := c.Query("interval")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	if strutil.IsBlank(interval) {
		interval = "1h"
	}
	data, err := ctl.service.ResourceNat(tag, interval)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceSlb(c *gin.Context) {
	tag := c.Query("tagKeyId")
	interval := c.Query("interval")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	if strutil.IsBlank(interval) {
		interval = "1h"
	}
	data, err := ctl.service.ResourceSlb(tag, interval)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceMonitorRdb(c *gin.Context) {
	tag := c.Query("tagKeyId")
	interval := c.Query("interval")
	productCode := c.Param("ProductCode")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	if strutil.IsBlank(interval) {
		interval = "1h"
	}
	data, err := ctl.service.ResourceRdb(tag, interval, productCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceOss(c *gin.Context) {
	tag := c.Query("tagKeyId")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	data, err := ctl.service.ResourceStorage(tag, constant.CloudProductCodeOss, constant.ResourceTypeCodeBucket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

func (ctl *LargeScreenCtl) ResourceEfs(c *gin.Context) {
	tag := c.Query("tagKeyId")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tagKeyId不能为空"))
		return
	}
	data, err := ctl.service.ResourceStorage(tag, constant.CloudProductCodeEfs, constant.ResourceTypeCodeShares)
	if err != nil {
		c.JSON(http.StatusInternalServerError, global.NewError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}
