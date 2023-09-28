package form

type PrometheusRequest struct {
	TenantId     string `form:"tenantId" json:"tenantId"`
	InstanceType string `form:"instanceType" json:"instanceType"`
	Abbreviation string `form:"abbreviation" json:"abbreviation"`
	Name         string `form:"name" json:"name"`
	Instance     string `form:"instance" json:"instance"`
	TopNum       int    `form:"topNum" json:"topNum"`
	Time         string `form:"time" json:"time"`             // 瞬时数据查询参数 时间戳
	Start        int    `form:"start" json:"start"`           // 开始时间 时间戳
	End          int    `form:"end" json:"end"`               // 结束时间 时间戳
	Step         int    `form:"step" json:"step"`             // 步长 时间戳
	Statistics   string `form:"statistics" json:"statistics"` // 聚合函数 sum(求和)  min(最小值)  max (最大值)  avg (平均值)  stddev (标准差)  stdvar (标准差异)  count (计数)
	Scope        string `form:"scope" json:"scope"`           // 聚合范围 1s(1秒)  1m(1分钟)  1h(1小时)  1d(1天)  1w(1周)  1y(1年)
	Pid          string `form:"pid" json:"pid"`
	Promql       string `form:"promql" json:"promql"`
}

type PrometheusResponse struct {
	Status string          `json:"status"`
	Data   *PrometheusData `json:"data"`
}

type PrometheusData struct {
	ResultType string              `json:"resultType"`
	Result     []*PrometheusResult `json:"result"`
}

type PrometheusResult struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
	Values [][]interface{}   `json:"values"`
}

type PrometheusValue struct {
	Time  string `json:"time"`
	Value string `json:"value"`
}

type PrometheusAxis struct {
	TimeAxis  []string            `json:"timeAxis"`
	ValueAxis map[string][]string `json:"valueAxis"`
}

type PrometheusInstance struct {
	Instance string `json:"instance"`
	Value    string `json:"value"`
}

type ProcessData struct {
	Pid       string `json:"pid"`
	Name      string `json:"name"`
	CmdLine   string `json:"cmdLine"`
	Cpu       string `json:"cpu"`
	Memory    string `json:"memory"`
	Openfiles string `json:"openfiles"`
	Time      string `json:"time"`
}
