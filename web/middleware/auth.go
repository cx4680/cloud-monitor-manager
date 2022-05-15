package middleware

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/openapi"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"path"
)

//忽略认证的路径列表
var ignoreList = []string{"/actuator/**"}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI

		for _, i := range ignoreList {
			match, err := path.Match(i, uri)
			if err != nil {
				if openapi.OpenApiRouter(c) {
					c.JSON(http.StatusUnauthorized, openapi.NewRespError(openapi.AuthorizedFailed, c))
				} else {
					c.JSON(http.StatusUnauthorized, global.NewError("权限认证失败"))
				}
				c.Abort()
			}
			if match {
				req := c.Request
				tenantId := req.Header.Get("tenantId")
				c.Set(global.TenantId, tenantId)
				c.Set(global.UserId, tenantId)
				c.Next()
				return
			}
		}
		if err := ParsingAndSetUserInfo(c); err != nil {
			if openapi.OpenApiRouter(c) {
				c.JSON(http.StatusUnauthorized, openapi.NewRespError(openapi.AuthorizedFailed, c))
			} else {
				c.JSON(http.StatusUnauthorized, global.NewError("权限认证失败"))
			}
			c.Abort()
		}
		c.Next()
	}

}

func ParsingAndSetUserInfo(c *gin.Context) error {
	if config.Cfg.Common.Env == "local" {
		c.Set(global.UserType, "1")
		c.Set(global.TenantId, "1")
		c.Set(global.UserId, "1")
		c.Set(global.UserName, "jim")
		return nil
	}

	req := c.Request
	userInfoEncode := req.Header.Get("user-info")
	var userMap map[string]string
	if len(userInfoEncode) != 0 {
		userInfoDecode, err := base64.StdEncoding.DecodeString(userInfoEncode)
		if err != nil {
			return err
		}
		if strutil.IsNotBlank(string(userInfoDecode)) {
			logger.Logger().Infof(`userInfo: %s`, string(userInfoDecode))
			jsonutil.ToObject(string(userInfoDecode), &userMap)
			if strutil.IsBlank(userMap["staffId"]) {
				userMap["staffId"] = "0"
			}
			if strutil.IsBlank(userMap["staffName"]) {
				userMap["staffName"] = "inner service"
			}
			c.Set(global.UserID, userMap["staffId"])
			c.Set(global.UserName, userMap["staffName"])
			return nil
		}
	}
	userIdentity := req.Header.Get("userIdentity")
	if len(userIdentity) != 0 {
		jsonutil.ToObject(userIdentity, &userMap)
		c.Set(global.UserType, userMap["typeCode"])
		c.Set(global.TenantId, userMap["accountId"])
		c.Set(global.UserId, userMap["principalId"])
		c.Set(global.UserName, userMap["userName"])
		return nil
	}
	return errors.New("无用户信息")

}
