package main

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/pipeline"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/sys_component/sys_db"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/global/sys_component/sys_redis"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/validator/translate"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/web"
	"context"
	"fmt"
	"os"
)

func main() {
	loader := pipeline.NewMainLoader()

	loader.AddStage(func(*context.Context) error {
		if err := sys_db.InitDb(config.Cfg.Db); err != nil {
			logger.Logger().Errorf("init database error: %v\n", err)
			return err
		}
		return nil
	})

	loader.AddStage(func(*context.Context) error {
		if err := sys_redis.InitClient(config.Cfg.Redis); err != nil {
			logger.Logger().Errorf("init redis error: %v\n", err)
			return err
		}
		return nil
	})

	loader.AddStage(func(*context.Context) error {
		return sys_db.InitData(config.Cfg.Db, "hawkeye", "file://./migrations")
	})

	loader.AddStage(func(*context.Context) error {
		return translate.InitTrans("zh")
	})

	loader.AddStage(func(*context.Context) error {
		return web.Start(config.Cfg.Serve)
	})

	_, err := loader.Start()
	if err != nil {
		fmt.Printf("exit error: %v", err)
		os.Exit(1)
	}

}
