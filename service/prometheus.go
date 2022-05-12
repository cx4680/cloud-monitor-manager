package service

import (
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
	"net/url"
)

type PrometheusService struct {
}

func NewPrometheusService() *PrometheusService {
	return &PrometheusService{}
}

func (s *PrometheusService) Query(pql string, time string) form.PrometheusResponse {
	var cfg = config.Cfg.Prometheus
	requestUrl := cfg.Url + cfg.Query
	logger.Logger().Info(requestUrl + pql)
	pql = url.QueryEscape(pql)
	if strutil.IsNotBlank(time) {
		pql += "&time=" + time
	}
	return sendRequest(requestUrl, pql)
}

func (s *PrometheusService) QueryRange(pql string, start string, end string, step string) form.PrometheusResponse {
	var cfg = config.Cfg.Prometheus
	requestUrl := cfg.Url + cfg.QueryRange
	logger.Logger().Info(requestUrl + pql)
	pql = url.QueryEscape(pql) + "&start=" + start + "&end=" + end + "&step=" + step
	return sendRequest(requestUrl, pql)
}

func sendRequest(requestUrl string, pql string) form.PrometheusResponse {
	var prometheusResponse form.PrometheusResponse
	response, err := httputil.HttpGet(requestUrl + pql)
	if err != nil {
		logger.Logger().Errorf("error:%v\n", err)
		return prometheusResponse
	}
	jsonutil.ToObject(response, &prometheusResponse)
	return prometheusResponse
}
