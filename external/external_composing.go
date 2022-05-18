package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
)

//已接入的产品简称
const (
	ECS = "ecs"
	EIP = "eip"
	SLB = "slb"
)

var ProductInstanceServiceMap = map[string]service.InstanceService{
	ECS: &EcsInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	EIP: &EipInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	SLB: &SlbInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
}
