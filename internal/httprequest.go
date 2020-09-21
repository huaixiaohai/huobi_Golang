package internal

import (
	"github.com/huobirdcenter/huobi_golang/logging/perflogger"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("GET", url)

	return string(result), err
}

func HttpPost(url string, body string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	logger.StopAndLog("POST", url)

	return string(result), err
}


//
//var proxyConf = "127.0.0.1:1081"
//
//func buildHttpClient(isProxy bool) *http.Client {
//	var proxy func(*http.Request) (*url.URL, error) = nil
//	if isProxy {
//		proxy = func(_ *http.Request) (*url.URL, error) {
//			return url.Parse("http://" + proxyConf)
//		}
//	}
//	transport := &http.Transport{Proxy: proxy}
//	client := &http.Client{Transport: transport}
//	return client
//}
//
//func HttpGet(urlStr string) (string, error) {
//	logger := perflogger.GetInstance()
//	logger.Start()
//
//	client := buildHttpClient(true)
//
//
//	resp, err := client.Get(urlStr)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//	result, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	logger.StopAndLog("GET", urlStr)
//
//	return string(result), err
//}
//
//func HttpPost(url string, body string) (string, error) {
//	logger := perflogger.GetInstance()
//	logger.Start()
//
//	resp, err := http.Post(url, "application/json", strings.NewReader(body))
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//	result, err := ioutil.ReadAll(resp.Body)
//
//	logger.StopAndLog("POST", url)
//
//	return string(result), err
//}

