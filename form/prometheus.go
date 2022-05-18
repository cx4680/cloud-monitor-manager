package form

type PrometheusRequest struct {
	TenantId     string `form:"tenantId"`
	InstanceType string `form:"instanceType"`
	Abbreviation string `form:"abbreviation"`
	Name         string `form:"name"`
	Instance     string `form:"instance"`
	TopNum       string `form:"topNum"`
	/**
	 * 瞬时数据查询参数 时间戳
	 */
	Time string `form:"time"`
	/**
	 * 区间数据查询参数 时间戳
	 */
	Start int `form:"start"`
	End   int `form:"end"`
	Step  int `form:"step"`
}

type PrometheusResponse struct {
	Status string         `json:"status"`
	Data   PrometheusData `json:"data"`
}

type PrometheusData struct {
	ResultType string             `json:"resultType"`
	Result     []PrometheusResult `json:"result"`
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
