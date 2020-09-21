package main

import (
	"github.com/huobirdcenter/huobi_golang/cmd/accountclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/commonclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/crossmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/etfclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/marketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/orderclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/walletclientexample"
	"github.com/huobirdcenter/huobi_golang/logging/perflogger"
	"os"
)

func main() {
	// 全局设置代理
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:1081")
	//os.Setenv("HTTPS_PROXY", "http://127.0.0.1:1081")
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	//accountclientexample.RunAllExamples()
	//orderclientexample.RunAllExamples()
	//algoorderclientexample.RunAllExamples()
	//marketclientexample.RunAllExamples()
	//isolatedmarginclientexample.RunAllExamples()
	//crossmarginclientexample.RunAllExamples()
	//walletclientexample.RunAllExamples()
	//subuserclientexample.RunAllExamples()
	//etfclientexample.RunAllExamples()
	//marketwebsocketclientexample.RunAllExamples()
	//accountwebsocketclientexample.RunAllExamples()
	//orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
