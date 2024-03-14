package task

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/constant"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/sys_component/sys_redis"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"github.com/google/uuid"
	"github.com/robfig/cron"
	"time"
)

type BusinessTask interface {
	Add(string, func()) error
	Start()
}

type BusinessTaskDTO struct {
	Cron string
	Name string
	Task func()
}

type BusinessTaskImpl struct {
	c *cron.Cron
}

var BusinessTaskList = []BusinessTaskDTO{OperationsLargeScreen(), MaintenanceLargeScreen()}

var podValue string

func NewBusinessTaskImpl() *BusinessTaskImpl {
	return &BusinessTaskImpl{c: cron.New()}
}

func (t *BusinessTaskImpl) Add(bt BusinessTaskDTO) error {
	var err error
	if strutil.IsEmpty(bt.Name) {
		err = t.c.AddFunc(bt.Cron, bt.Task)
	} else {
		err = t.c.AddFunc(bt.Cron, func() {
			logger.Logger().Info(bt.Name + " start running")
			bt.Task()
			logger.Logger().Info(bt.Name + " running over")
		})
	}
	if err != nil {
		return err
	}
	return nil
}

func (t *BusinessTaskImpl) Start() {
	if canItRun() {
		t.c.Start()
	} else {
		logger.Logger().Info("其它pod正在运行定时任务，该pod已跳过")
	}
}

func (t *BusinessTaskImpl) Stop() {
	t.c.Stop()
}

func canItRun() bool {
	if strutil.IsBlank(podValue) {
		podValue = uuid.New().String()
	}
	value, err := sys_redis.Get(constant.LargeScreenKey)
	if err != nil {
		logger.Logger().Error("获取定时任务redis缓存错误: ", err)
	}
	if strutil.IsBlank(value) {
		if err = sys_redis.SetByTimeOut(constant.LargeScreenKey, podValue, time.Minute*10); err != nil {
			logger.Logger().Error("设置定时任务redis缓存错误: ", err)
		}
	}
	value, err = sys_redis.Get(constant.LargeScreenKey)
	if err != nil {
		logger.Logger().Error("获取定时任务redis缓存错误: ", err)
	}
	if value == podValue {
		return true
	} else {
		return false
	}
}
