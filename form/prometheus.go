package form

type PrometheusRequest struct {
	InstanceType string `form:"instanceType"`
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
	XAxis []string            `json:"xaxis"`
	YAxis map[string][]string `json:"yaxis"`
}

type PrometheusInstance struct {
	Instance string `json:"instance"`
	Value    string `json:"value"`
}
