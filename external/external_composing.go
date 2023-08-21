package external

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/service"
	"strings"
)

// 已接入的产品简称
const (
	ECS        = "ecs"
	EIP        = "eip"
	SLB        = "slb"
	CBR        = "cbr"
	NAT        = "nat"
	MYSQL      = "mysql"
	DM         = "dm"
	POSTGRESQL = "postgresql"
	KAFKA      = "kafka"
	BMS        = "bms"
	EBMS       = "ebms"
	REDIS      = "redis"
	MONGO      = "mongo"
	CGW        = "cgw"
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
	CBR: &CbrInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	NAT: &NatInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	MYSQL: &MySqlInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	DM: &DmInstanceService{
		InstanceServiceImpl: service.InstanceServiceImpl{},
	},
	POSTGRESQL: &PgInstanceService{
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

func toStringList(s string) []string {
	var list []string
	if len(s) == 0 {
		return list
	}
	statusList := strings.Split(s, ",")
	for _, v := range statusList {
		list = append(list, v)
	}
	return list
}

type InstanceRequest struct {
	CloudProductCode string   `json:"cloudProductCode"`
	ResourceTypeCode string   `json:"resourceTypeCode"`
	ResourceId       string   `json:"resourceId"`
	StatusList       []string `json:"statusList"`
	RegionCode       string   `json:"regionCode"`
	Name             string   `json:"name"`
	TenantId         string   `json:"tenantId"`
	PageSize         string   `json:"pageSize"`
	CurrPage         string   `json:"currPage"`
}

type InstanceResponse struct {
	Code    string       `json:"code"`
	Msg     string       `json:"msg"`
	TraceId string       `json:"traceId"`
	Data    InstancePage `json:"data"`
}

type InstancePage struct {
	Total int             `json:"total"`
	List  []*InstanceList `json:"list"`
}

type InstanceList struct {
	Id               int    `json:"id"`
	UuidStr          string `json:"uuidStr"`
	RegionCode       string `json:"regionCode"`
	RegionName       string `json:"regionName"`
	ResourceTypeCode string `json:"resourceTypeCode"`
	CloudProductCode string `json:"cloudProductCode"`
	TenantId         string `json:"tenantId"`
	TenantName       string `json:"tenantName"`
	ResourceId       string `json:"resourceId"`
	ResourceName     string `json:"resourceName"`
	OrderId          string `json:"orderId"`
	ResourceUrl      string `json:"resourceUrl"`
	AvailabilityZone string `json:"availabilityZone"`
	Status           int    `json:"status"`
	StatusDesc       string `json:"statusDesc"`
	Deleted          int    `json:"deleted"`
	CreateTime       int    `json:"createTime"`
	UpdateTime       int    `json:"updateTime"`
	Additional       string `json:"additional"`
	ResCreateTime    int    `json:"resCreateTime"`
	ResUpdateTime    int    `json:"resUpdateTime"`
	Creator          string `json:"creator"`
	Modifier         string `json:"modifier"`
}
