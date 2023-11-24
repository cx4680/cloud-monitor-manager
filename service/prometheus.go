package service

import (
	"net/url"
	"strings"

	"code.cestc.cn/ccos-ops/cloud-monitor-manager/config"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/form"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/logger"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/httputil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/strutil"
)

type PrometheusService struct {
}

func NewPrometheusService() *PrometheusService {
	return &PrometheusService{}
}

func (s *PrometheusService) Query(pql string, time string) *form.PrometheusResponse {
	var cfg = config.Cfg.Prometheus
	requestUrl := cfg.Url + cfg.Query
	logger.Logger().Info(requestUrl + pql)
	pql = url.QueryEscape(pql)
	if strutil.IsNotBlank(time) {
		pql += "&time=" + time
	}
	return sendRequest(requestUrl, pql)
}

func (s *PrometheusService) PostQuery(pql string, time string) *form.PrometheusResponse {
	var cfg = config.Cfg.Prometheus
	query := strings.Split(cfg.Query, "?")
	requestUrl := cfg.Url + query[0]
	logger.Logger().Info("requestUrl: " + requestUrl + " pql: " + pql)
	// pql = url.QueryEscape(pql)
	queryParam := make(map[string][]string)
	if strutil.IsNotBlank(time) {
		queryParam["time"] = []string{time}
	}
	queryParam["query"] = []string{pql}
	return sendPostRequest(requestUrl, queryParam)
}

func (s *PrometheusService) QueryRange(pql string, start string, end string, step string) *form.PrometheusResponse {
	var cfg = config.Cfg.Prometheus
	requestUrl := cfg.Url + cfg.QueryRange
	logger.Logger().Info(requestUrl + pql)
	pql = url.QueryEscape(pql) + "&start=" + start + "&end=" + end + "&step=" + step
	return sendRequest(requestUrl, pql)
}

func (s *PrometheusService) QueryFrontendRange(pql string, start string, end string, step string) *form.PrometheusResponse {
	var cfg = config.Cfg.Prometheus
	requestUrl := cfg.ThanosQueryFrontend + cfg.QueryRange
	pql = url.QueryEscape(pql) + "&start=" + start + "&end=" + end + "&step=" + step + "&partial_response=false"
	logger.Logger().Info(requestUrl + pql)
	return sendRequest(requestUrl, pql)
}

func (s *PrometheusService) QueryFrontendRangeDownSampling(pql string, start string, end string, step string) *form.PrometheusResponse {
	var cfg = config.Cfg.Prometheus
	requestUrl := cfg.ThanosQueryFrontend + cfg.QueryRange
	pql = url.QueryEscape(pql) + "&start=" + start + "&end=" + end + "&step=" + step + "&partial_response=false&max_source_resolution=1h"
	logger.Logger().Info(requestUrl + pql)
	return sendRequest(requestUrl, pql)
}

func sendRequest(requestUrl string, pql string) *form.PrometheusResponse {
	prometheusResponse := &form.PrometheusResponse{}
	response, err := httputil.HttpGet(requestUrl + pql)
	if err != nil {
		logger.Logger().Errorf("query prometheus error:%v, pql: %s, origin response:%s\n", err, pql, response)
		return prometheusResponse
	}
	jsonutil.ToObject(response, &prometheusResponse)
	if prometheusResponse == nil || prometheusResponse.Data == nil || prometheusResponse.Data.Result == nil || len(prometheusResponse.Data.Result) == 0 {
		logger.Logger().Infof("query prometheus empty, pql: %s, origin response:%s\n", pql, response)
	}
	return prometheusResponse
}

func sendPostRequest(requestUrl string, params map[string][]string) *form.PrometheusResponse {
	prometheusResponse := &form.PrometheusResponse{}
	response, err := httputil.HttpPostForm(requestUrl, params)
	if err != nil {
		logger.Logger().Errorf("error:%v\n", err)
		return prometheusResponse
	}
	jsonutil.ToObject(response, &prometheusResponse)
	return prometheusResponse
}
