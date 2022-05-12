package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
)

//已接入的产品简称
const (
	ECS        = "ecs"
	CBR        = "cbr"
	EIP        = "eip"
	NAT        = "nat"
	SLB        = "slb"
	BMS        = "bms"
	EBMS       = "ebms"
	MYSQL      = "mysql"
	DM         = "dm"
	POSTGRESQL = "postgresql"
	KAFKA      = "kafka"
	REDIS      = "redis"
	MONGO      = "mongo"
	CGW        = "cgw"
)

var ProductInstanceServiceMap = map[string]service.InstanceService{
	CBR: &CbrInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	ECS: &EcsInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	EIP: &EipInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	NAT: &NatInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	SLB: &SlbInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	BMS: &BmsInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	EBMS: &EbmsInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	MYSQL: &MySqlInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	DM: &MySqlInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	POSTGRESQL: &MySqlInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	KAFKA: &KafkaInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	REDIS: &RedisInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	MONGO: &MongoInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	CGW: &CgwInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
}
