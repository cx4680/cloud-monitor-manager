package httputil

import (
	"bytes"
	"code.cestc.cn/ccos-ops/cloud-monitor-manager/util/jsonutil"
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(path string) (string, error) {
	resp, err := http.Get(path)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpHeaderGet(path string, header map[string]string) (string, error) {
	req, _ := http.NewRequest("GET", path, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// HttpPost params 格式 a=123&b=234
func HttpPost(path string, params string) (string, error) {
	resp, err := http.Post(path,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPostForm(path string, params map[string][]string) (string, error) {
	resp, err := http.PostForm(path, params)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

func HttpPostJson(path string, params interface{}, headers map[string]string) (string, error) {
	req, err := http.NewRequest("POST", path, bytes.NewBuffer([]byte(jsonutil.ToString(params))))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

var httpClient *resty.Client

func GetHttpClient() *resty.Client {
	if httpClient == nil {
		client := resty.New()
		client.SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
		httpClient = client
	}
	return httpClient
}
