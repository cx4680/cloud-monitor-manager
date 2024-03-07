package form

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
