package oplog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var userTypeMap = map[string]string{
	"0": "root-account",
	"1": "root-account",
	"2": "root-account",
	"3": "iam-user",
	"4": "assumed-role",
	"5": "system",
}

const (
	UserId       = "userId"
	UserName     = "userName"
	ResourceName = "resourceName"
)

type EventLevel = string

const (
	INFO  EventLevel = "Info"
	Warn  EventLevel = "Warn"
	Fatal EventLevel = "Fatal"
)

var (
	trailLogger      = zap.New(newTrail())
	timeFormatLayout = "2006/01/02 15:04:05"
)

type OperatorInfo struct {
	EventRequestInfo
	ResourceInfo
}
type EventRequestInfo struct {
	ServiceName string     //由产品统一定义的云服务名称，需与云资源模型中定义的产品名称一致
	EventApi    string     //自行或由产品统一定义的云服务事件名称，需与云资源模型中定义的action一致
	EventName   string     //API名称
	RequestType string     //根据当前接口的操作类型，定义当前接口是读接口还是写接口（ Write,Read）
	ApiVersion  string     //	接口版本当前定义为1.0，如接口升级需升级此版本号
	EventLevel  EventLevel //取值：“高危”:  Fatal ，如虚机的删除操作			 “危险”：Warn，如关机操作g			 “一般”：Info，如查看虚机列表
	EventRegion string     //当前服务发起的地域，如服务部署在北京则 cn-beijing
	Utc         bool       //是否UTC时间 默认false
}

type ResourceInfo struct {
	ResourceName string //操作资源的具体值，如操作的资源为用户(user)，那么记录用户id		 如果打印日志所属的API接入过IAM，最好与对应resourceId保持一致。
}

func GinTrail(operatorInfo *OperatorInfo) gin.HandlerFunc {
	return func(c *gin.Context) {
		replaceResponseWriter(c)
		requestParamJson, done := getRequestParams(c)
		if done {
			return
		}
		requestID := c.GetHeader("X-Request-ID")
		if len(requestID) == 0 {
			if newUUID, err := uuid.NewUUID(); err == nil {
				requestID = newUUID.String()
			}
		}
		ctx := context.WithValue(context.Background(), "X-Request-ID", requestID)
		c.Set("ctx", ctx)
		defer func() {
			loginId := c.GetString(UserId)
			userName := c.GetString(UserName)
			source := c.Request.Header["Origin"]
			eventSource := ""
			if len(source) > 0 {
				eventSource = source[0]
			}
			end := time.Now()
			if operatorInfo.Utc {
				end = end.UTC()
			}
			var response string
			var errs map[string]string
			if writer, ok := c.Writer.(*responseWriter); ok {
				response = string(writer.response)
				json.Unmarshal(writer.response, &errs)
			}
			result := getResult(c.Writer.Status())
			errMessage := errs["message"]
			var resError interface{}
			if err := recover(); err != nil {
				result = "Fail"
				errMessage = fmt.Sprint(err)
				resError = err
			}
			var resourceName string
			if resName, ok := c.Get(ResourceName); ok {
				resourceName = fmt.Sprintf("%v", resName)
			}
			if len(resourceName) == 0 {
				resourceName = operatorInfo.ResourceName
			}
			trailLogger.Info("[OP_ACTION_TRAIL_LOG]",
				zap.String("event_id", requestID),
				zap.String("event_version", "1"),
				zap.String("event_source", eventSource),
				zap.String("source_ip_address", c.ClientIP()),
				zap.String("user_agent", c.Request.UserAgent()),
				zap.String("service_name", operatorInfo.ServiceName),
				zap.String("event_api", operatorInfo.EventApi),
				zap.String("event_name", operatorInfo.EventName),
				zap.String("request_type", operatorInfo.RequestType),
				zap.String("api_version", operatorInfo.ApiVersion),
				zap.String("event_level", operatorInfo.EventLevel),
				zap.String("request_id", requestID),
				zap.String("event_time", end.Format(timeFormatLayout)),
				zap.String("event_region", operatorInfo.EventRegion),
				zap.String("resource_name", resourceName),
				zap.String("request_parameters", string(requestParamJson)),
				zap.String("result", result),
				zap.String("response_elements", response), //todo 脱敏
				zap.String("error_code", errs["code"]),
				zap.String("error_message", errMessage),
				zap.String("user_name", userName),
				zap.String("user_id", loginId),
			)
			if resError != nil {
				panic(resError)
			}
		}()
		c.Next()
	}

}

func getRequestParams(c *gin.Context) ([]byte, bool) {
	var data []byte
	if http.MethodGet == c.Request.Method {
		params := c.Request.URL.RawQuery
		requestParams, err := json.Marshal(params)
		if err != nil {
			c.Abort()
			log.Printf("parse params %v", err)
			return nil, true
		}
		data = requestParams
	} else {
		body, err := c.GetRawData()
		if err != nil {
			c.Abort()
			log.Printf("parse params %v", err)
			return nil, true
		}
		data = body
	}
	formatData := strings.Replace(strings.Replace(string(data), "\r\n", "", -1), " ", "", -1)
	requestParameters := make(map[string]string, 0)
	requestParameters["request"] = formatData
	requestParamJson, _ := json.Marshal(requestParameters)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return requestParamJson, false
}

func getResult(status int) string {
	if status != 200 {
		return "Fail"
	}
	return "Success"
}

func newTrail() zapcore.Core {
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	coreTrail := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(getEncoderConfig()), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), infoLevel),
	)
	return coreTrail
}

func getEncoderConfig() zapcore.EncoderConfig {
	EncodeLevel := zapcore.CapitalColorLevelEncoder
	return zapcore.EncoderConfig{
		FunctionKey:      "func",
		StacktraceKey:    "stack",
		NameKey:          "name",
		MessageKey:       "msg",
		LevelKey:         "level",
		ConsoleSeparator: " | ",
		EncodeLevel:      EncodeLevel,
		TimeKey:          "s",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(timeFormatLayout))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeName: func(n string, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(n)
		},
	}
}
