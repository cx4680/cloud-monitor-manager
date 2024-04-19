package form

import "time"

type CmdbResponse struct {
	Result  bool      `json:"result"`
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    *CmdbData `json:"data"`
}

type CmdbData struct {
	Count int         `json:"count"`
	Info  []*CmdbInfo `json:"info"`
}

type CmdbInfo struct {
	BkInstId      int64  `json:"bk_inst_id"`
	BkInstName    string `json:"bk_inst_name"`
	CategoryId    int64  `json:"category_id"`
	CategoryName  string `json:"category_name"`
	CreateTime    string `json:"create_time"`
	CpuUpdateTime string `json:"cpu_update_time"`
	DeviceModel   string `json:"device_model"`
	DeviceStatus  string `json:"device_status"`
	Ip            string `json:"ip"`
	RunStatus     int    `json:"run_status"`
}

type EcsCusInventory struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		Cores        int   `json:"cores"`
		UsedCores    int   `json:"usedCores"`
		SurplusCores int   `json:"surplusCores"`
		Ram          int64 `json:"ram"`
		UsedRam      int64 `json:"usedRam"`
		SurplusRam   int64 `json:"surplusRam"`
	} `json:"data"`
}

type EbsCusInventory struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		TotalCapacityGb     int64 `json:"totalCapacityGb"`     // 底层实际容量，单位gb
		UsedCapacityGb      int64 `json:"usedCapacityGb"`      // 客户数据实际占用容量，单位gb
		AllocatedCapacityGb int64 `json:"allocatedCapacityGb"` // 客户开盘容量，单位gb
		RatioCapacityGb     int64 `json:"ratioCapacityGb"`     // 超分之后的容量，单位gb
		AllocatedVolumeNum  int64 `json:"allocatedVolumeNum"`  // 客户开盘数量
	} `json:"data"`
}

type RdbManageResource struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		DatabaseTotal int64 `json:"databaseTotal"`
		DatabaseError int64 `json:"databaseError"`
		CpuTotal      int64 `json:"cpuTotal"`
		CpuUsed       int64 `json:"cpuUsed"`
	} `json:"data"`
}

type CmqCpuResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		RegionCode string `json:"regionCode"`
		RegionName string `json:"regionName"`
		CpuCap     int64  `json:"cpuCap"`
		CpuUsed    int64  `json:"cpuUsed"`
		MemCap     int64  `json:"memCap"`
		MemUsed    int64  `json:"memUsed"`
	} `json:"data"`
}

type CmqCountResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		RegionCode           string `json:"regionCode"`
		RegionName           string `json:"regionName"`
		AllInstanceNum       int64  `json:"allInstanceNum"`
		ExceptionInstanceNum int64  `json:"exceptionInstanceNum"`
	} `json:"data"`
}

type LargeScreenAlertResponse struct {
	List []*struct {
		Name                string    `json:"name"`
		Expr                string    `json:"expr"`
		Target              string    `json:"target"`
		Severity            string    `json:"severity"`
		Type                string    `json:"type"`
		MonitorType         string    `json:"monitor_type"`
		Status              string    `json:"status"`
		Description         string    `json:"description"`
		StartsAt            time.Time `json:"starts_at"`
		Start               int       `json:"start"`
		AlertId             string    `json:"alert_id"`
		Ip                  string    `json:"ip"`
		AlertAttr           string    `json:"alert_attr"`
		InstId              string    `json:"inst_id"`
		OperationType       string    `json:"operation_type"`
		ResponsiblePerson   string    `json:"responsible_person"`
		ResponsiblePersonId string    `json:"responsible_person_id"`
	} `json:"list"`
	P1 int `json:"p1"`
	P2 int `json:"p2"`
	P3 int `json:"p3"`
	P4 int `json:"p4"`
}

type CocClusterDeployment struct {
	List []*CocCluster `json:"list"`
}

type CocClusterStatefulSet struct {
	List []*CocCluster `json:"list"`
}

type CocCluster struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	InstanceTotal int32             `json:"instance_total"`
	InstanceReady int32             `json:"instance_ready"`
	Containers    []string          `json:"containers"`
	Images        []string          `json:"images"`
	Labels        map[string]string `json:"labels"`
	CreateTime    time.Time         `json:"create_time"`
}

type CocClusterPod struct {
	Data *Pod `json:"data"`
}

type Pod struct {
	Name   string `json:"name"`
	Time   string `json:"time"`
	Status string `json:"status"`
}

type DiskResponse struct {
	Data *struct {
		TotalCount int `json:"total_count"`
		Stats      []*struct {
			DiskId     string `json:"disk_id"`
			NodeName   string `json:"node_name"`
			NodeIp     string `json:"node_ip"`
			DiskStatus string `json:"disk_status"`
			FaultTime  string `json:"fault_time"`
		} `json:"stats"`
	} `json:"data"`
}

type EbsResponse struct {
	Data *struct {
		TotalCount int `json:"total_count"`
		PoolStatus []*struct {
			Name              string  `json:"name"`
			TotalCapacity     float64 `json:"totalCapacity"`
			UsedCapacity      float64 `json:"usedCapacity"`
			AvailableCapacity float64 `json:"availableCapacity"`
			LicenseLeftTime   string  `json:"license_left_day"`
		} `json:"pool_status"`
	} `json:"data"`
}

type LargeScreenResourceTagResponse struct {
	Module *struct {
		List  []*LargeScreenResourceTag `json:"list"`
		Count string                    `json:"count"`
	} `json:"module"`
}

type LargeScreenResourceTag struct {
	TagKey   string `json:"tagKey"`
	YagKeyId string `json:"tagKeyId"`
}

type LargeScreenResourceResponse struct {
	Module *struct {
		List  []*LargeScreenResource `json:"list"`
		Count string                 `json:"count"`
	} `json:"module"`
}

type LargeScreenResource struct {
	ResourceId         string `json:"resourceId"`
	ResourceInstanceId string `json:"resourceInstanceId"`
	ResourceName       string `json:"resourceName"`
	CloudProductCode   string `json:"cloudProductCode"`
	ResourceTypeCode   string `json:"resourceTypeCode"`
	StatusDesc         string `json:"statusDesc"`
	Additional         string `json:"additional"`
}

type LargeScreenResourceStorageAdditional struct {
	Storage   int64 `json:"storage"`
	ObjectNum int64 `json:"objectNum"`
}

type LargeScreenResourceOverview struct {
	Ecs struct {
		Total  int `json:"total"`
		Normal int `json:"normal"`
	} `json:"ecs"`
	Eip struct {
		Total     int `json:"total"`
		Bandwidth int `json:"bandwidth"`
	} `json:"eip"`
	Rdb struct {
		Mysql int `json:"mysql"`
		Dm    int `json:"dm"`
		Pg    int `json:"pg"`
	} `json:"rdb"`
	Slb struct {
		Total  int `json:"total"`
		Normal int `json:"normal"`
	} `json:"slb"`
	Nat struct {
		Total  int `json:"total"`
		Normal int `json:"normal"`
	} `json:"nat"`
}

type EipAdditional struct {
	BandWidth struct {
		BandwidthId   string `json:"bandwidthId"`
		BandWidthSize int    `json:"BandWidthSize"`
	} `json:"bandWidth"`
	BindInstanceId string `json:"bindInstanceId"`
	EipIpAddress   string `json:"eipIpAddress"`
}

type YyLargeScreen struct {
	P1    int                   `json:"p1"`
	P2    int                   `json:"p2"`
	P3    int                   `json:"p3"`
	P4    int                   `json:"p4"`
	Trend []*YyLargeScreenTrend `json:"trend"`
	List  []*YyLargeScreenAlert `json:"list"`
}

type YyLargeScreenTrend struct {
	Time     string `json:"time"`
	Alert    int    `json:"alert"`
	Recovery int    `json:"recovery"`
}

type YyLargeScreenAlert struct {
	AlertLevel   string `json:"alertLevel" gorm:"column:severity"`
	ResourceId   string `json:"resourceId" gorm:"column:name"`
	ResourceType string `json:"resourceType" gorm:"column:target_type"`
	CreateTime   string `json:"createTime" gorm:"column:create_time"`
}

type LargeScreenEcs struct {
	Cpu struct {
		Value string `json:"value"`
		Unit  string `json:"unit"`
	} `json:"cpu"`
	Memory struct {
		Value string `json:"value"`
		Unit  string `json:"unit"`
	} `json:"memory"`
	Disk struct {
		Value string `json:"value"`
		Unit  string `json:"unit"`
	} `json:"disk"`
}

type LargeScreenEcsTop struct {
	Name string                    `json:"name"`
	Data []*LargeScreenEcsTopValue `json:"data"`
}

type LargeScreenEcsTopValue struct {
	ResourceId string `json:"resourceId"`
	Value      string `json:"value"`
}

type LargeScreenMonitorEip struct {
	Time                string  `json:"time"`
	UpstreamBandwidth   float64 `json:"upstreamBandwidth"`
	DownstreamBandwidth float64 `json:"downstreamBandwidth"`
	Upstream            float64 `json:"upstream"`
	Downstream          float64 `json:"downstream"`
}

type LargeScreenMonitorNat struct {
	Time                string  `json:"time"`
	UpstreamBandwidth   float64 `json:"upstreamBandwidth"`
	DownstreamBandwidth float64 `json:"downstreamBandwidth"`
	Upstream            float64 `json:"upstream"`
	Downstream          float64 `json:"downstream"`
}

type LargeScreenMonitorSlb struct {
	Time                      string  `json:"time"`
	AllConnectionCount        float64 `json:"allConnectionCount"`
	AllEstConnection          float64 `json:"allEstConnection"`
	AllNoneEstConnectionCount float64 `json:"allNoneEstConnectionCount"`
	NewConnectionRate         float64 `json:"newConnectionRate"`
	DropConnectionRate        float64 `json:"dropConnectionRate"`
	RequestRate               float64 `json:"requestRate"`
	Http2xxRate               float64 `json:"http2xxRate"`
	Http3xxRate               float64 `json:"http3xxRate"`
	Http4xxRate               float64 `json:"http4xxRate"`
	Http5xxRate               float64 `json:"http5xxRate"`
}

type LargeScreenMonitor struct {
	Name  string                     `json:"name"`
	Code  string                     `json:"code"`
	Unit  string                     `json:"unit"`
	Chart []*LargeScreenMonitorChart `json:"chart"`
}

type LargeScreenMonitorChart struct {
	ResourceId string                    `json:"resourceId"`
	Data       []*LargeScreenMonitorData `json:"data"`
}

type LargeScreenMonitorData struct {
	Time  string `json:"time"`
	Value string `json:"value"`
}

type LargeScreenStorageTrend struct {
	Time       string  `json:"time"`
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	Conversion string  `json:"conversion"`
}

type LargeScreenEfsResponse struct {
	Data []*struct {
		RegionCode       string `json:"regionCode"`
		CloudProductCode string `json:"cloudProductCode"`
		ResourceTypeCode string `json:"resourceTypeCode"`
		InstanceId       string `json:"instanceId"`
		ResourceName     string `json:"resourceName"`
		Status           string `json:"status"`
		UsedCapacity     int64  `json:"usedCapacity"`
		TotalCapacity    int64  `json:"totalCapacity"`
		CreateTime       string `json:"createTime"`
		UpdateTime       string `json:"updateTime"`
	} `json:"data"`
}
