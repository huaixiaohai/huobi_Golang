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
	"github.com/huobirdcenter/huobi_golang/view"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
	"time"
)

func main() {
	go func() {
		// 全局设置代理
		os.Setenv("HTTP_PROXY", "http://127.0.0.1:1081")
		//os.Setenv("HTTPS_PROXY", "http://127.0.0.1:1081")
		//for {
			runAll()
			time.Sleep(time.Second)
		//}
	}()

	view.UIManagerIns().Init()
}

// plot test
func plotTest() {
	var a, b float64 = 0.7, 3
	points := plotter.XYs{}
	for i := 0; i <= 10; i++ {
		points = append(points, plotter.XY{
			X: float64(i),
			Y: a*float64(i) + b,
		})
	}

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}
	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 10, 10

	if err := plotutil.AddLines(plt,
		"line1", points,
	); err != nil {
		panic(err)
	}

	if err := plt.Save(5*vg.Inch, 5*vg.Inch, "01-draw-line.png"); err != nil {
		panic(err)
	}
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	//accountclientexample.RunAllExamples()
	//orderclientexample.RunAllExamples()
	//algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
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
