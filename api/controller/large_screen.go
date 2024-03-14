package controller

import (
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
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
}

func (ctl *LargeScreenCtl) ResourceOverview(c *gin.Context) {
	tag := c.Query("tag")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tag不能为空"))
	}
	data, err := ctl.service.ResourceOverview(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
}

func (ctl *LargeScreenCtl) ResourceAlert(c *gin.Context) {
	tag := c.Query("tag")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tag不能为空"))
	}
	data, err := ctl.service.ResourceAlert(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
}

func (ctl *LargeScreenCtl) ResourceEcs(c *gin.Context) {
	tag := c.Query("tag")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tag不能为空"))
	}
	data, err := ctl.service.ResourceEcs(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
}

func (ctl *LargeScreenCtl) ResourceEip(c *gin.Context) {
	tag := c.Query("tag")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tag不能为空"))
	}
	data, err := ctl.service.ResourceEip(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
}

func (ctl *LargeScreenCtl) ResourceRdb(c *gin.Context) {
	tag := c.Query("tag")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tag不能为空"))
	}
	data, err := ctl.service.ResourceRdb(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
}

func (ctl *LargeScreenCtl) ResourceStorage(c *gin.Context) {
	tag := c.Query("tag")
	if strutil.IsBlank(tag) {
		c.JSON(http.StatusBadRequest, global.NewError("tag不能为空"))
	}
	data, err := ctl.service.ResourceStorage(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, global.NewSuccess("查询成功", data))
}
