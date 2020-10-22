package main

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

const size = 12
const sym = "btcusdt"
const (
	BTCUSDT = "btcusdt"
	UNIUSDT = "uniusdt"
	ETHHUSDT = "ethusdt"
)

const period = market.DAY1

var symbols = []string{BTCUSDT, UNIUSDT}

func main() {
	for {
		os.Setenv("HTTP_PROXY", "http://127.0.0.1:1081")
		//symbols := getSymbols()

		//symbols = []string{sym}
		m := make(map[string][]market.Candlestick, 0)

		for _, symbol := range symbols {
			//println(symbol)
			b, res := getCandlestick(symbol)
			if b {
				m[symbol] = res
			}
		}

		// test1 -
		_ = func() {
			for sym, _ := range m {
				price := 1.0
				x := 0
				for i := 0; i < size; i++ {
					o, _ := m[sym][i].Open.Float64()
					c, _ := m[sym][i].Close.Float64()

					if x >= 3 {
						price *= 1 + ((c - o) / o)
						x = 0
						continue
					}

					if c-o >= 0 {
						x++
					} else {
						x = 0
					}
				}

				println(fmt.Sprintf("symbol: %s, price: %f", sym, price))
			}
		}

		//test2, -
		_ = func() {
			maxPrice := 1.0
			minPrice := 1.0
			avgPrice := 1.0
			for i := 0; i < size-1; i++ {
				max := -1000.0
				min := 10000.0
				maxSym := ""
				minSym := ""

				for symbol, list := range m {
					o, _ := list[i].Open.Float64()
					c, _ := list[i].Close.Float64()
					rate := (c - o) / o
					if rate > max {
						max = rate
						maxSym = symbol
					}
					if rate < min {
						min = rate
						minSym = symbol
					}
				}

				avg := 0.0
				avgSym := ""

				for symbol, list := range m {
					o, _ := list[i].Open.Float64()
					c, _ := list[i].Close.Float64()

					rate := math.Abs((max+min)/2 - (c-o)/o)

					if avgSym == "" || avg > rate {
						avgSym = symbol
						avg = rate
					}

				}

				o, _ := m[maxSym][i+1].Open.Float64()
				c, _ := m[maxSym][i+1].Close.Float64()

				o2, _ := m[minSym][i+1].Open.Float64()
				c2, _ := m[minSym][i+1].Close.Float64()

				o3, _ := m[avgSym][i+1].Open.Float64()
				c3, _ := m[avgSym][i+1].Close.Float64()

				//println(fmt.Sprintf("symbol: %s,  cha: %f, cha: %f, symbol: %s,  cha: %f, cha: %f", maxSym, max, (c-o)/o, minSym, min, (c2-o2)/o2))
				maxPrice *= 1 + (c-o)/o
				minPrice *= 1 + (c2-o2)/o2
				avgPrice *= 1 + (c3-o3)/o3
				println(fmt.Sprintf("%f, %f, %f", maxPrice, minPrice, avgPrice), maxSym, minSym, avgSym)
			}
		}

		// test3, +
		_ = func() {

			a, b := 0, 0

			all := 0.0
			for sym, _ := range m {
				price := 1.0
				for i := 0; i < size-1; i++ {

					o, _ := m[sym][i].Open.Float64()
					c, _ := m[sym][i].Close.Float64()
					h, _ := m[sym][i].High.Float64()
					l, _ := m[sym][i].Low.Float64()

					if c-o > 0 && math.Abs(h-c) < math.Abs(l-o) {
						o1, _ := m[sym][i+1].Open.Float64()
						c1, _ := m[sym][i+1].Close.Float64()
						price *= 1 + ((c1 - o1) / o1)
					}
					//println(fmt.Sprintf("symbol: %s, price: %f", sym, price))
				}
				println(fmt.Sprintf("%f", price-1))
				if price-1 < 0 {
					a++
				} else if price-1 > 0 {
					b++
				}
				all += price - 1
			}
			println(fmt.Sprintf("a: %d, b:%d, all :%f", a, b, all))
		}

		// test4 还行 不太靠谱
		_ = func() {
			price := 1.0
			for i := 0; i < size; i++ {
				maxSym := ""
				maxRate := 0.0
				for sym, _ := range m {
					o, _ := m[sym][i].Open.Float64()
					c, _ := m[sym][i].Close.Float64()
					h, _ := m[sym][i].High.Float64()
					l, _ := m[sym][i].Low.Float64()

					rate := (math.Abs(l-o) - math.Abs(h-c)) / o

					if c > o && math.Abs(h-c) < math.Abs(l-o) {
						if maxSym == "" || maxRate < rate {
							maxSym = sym
							maxRate = rate
						}
					}
				}
				if maxSym != "" {
					o, _ := m[maxSym][i+1].Open.Float64()
					c, _ := m[maxSym][i+1].Close.Float64()
					price *= 1 + (c-o)/o
					println(fmt.Sprintf("maxSym: %s, cha: %f", maxSym, (c-o)/o))
				}
			}

			println(fmt.Sprintf("price: %f", price))
		}

		// test5
		_ = func() {

			mRes := make([][]*Record, 0)

			for i := 0; i < size; i++ {
				rates := make([]*Record, 0)
				for sym, _ := range m {
					o, _ := m[sym][i].Open.Float64()
					c, _ := m[sym][i].Close.Float64()
					h, _ := m[sym][i].High.Float64()
					l, _ := m[sym][i].Low.Float64()
					if c > o && h-c < o-l {

						rate := (math.Abs(o-l) - math.Abs(h-c)) / o
						rates = append(rates, &Record{
							Symbol: sym,
							Rate:   rate,
						})
					}
				}
				sort.Sort(Records(rates))
				mRes = append(mRes, rates)
			}
			println("===========")
		}

		// test6
		_ = func() {
			for sym, list := range m {
				o1, _ := list[0].Open.Float64()
				c1, _ := list[6].Close.Float64()

				o2, _ := list[7].Open.Float64()
				c2, _ := list[13].Close.Float64()

				o3, _ := list[14].Open.Float64()
				c3, _ := list[20].Close.Float64()

				println(fmt.Sprintf("sym: %s, ch1: %.2f, ch1:%.2f, ch1:%.2f", sym, (c1-o1)/o1, (c2-o2)/o2, (c3-o3)/o3))

			}
		}

		// test7
		_ = func() {
			one := m[sym][size-1]
			//for _, one := range list {
			c, _ := one.Close.Float64()
			o, _ := one.Open.Float64()
			println(fmt.Sprintf("%0.2f", (c-o)/o*100))
			//}
		}

		// test8

		test := func() {
			// MA(n) = 最近的N天收盘价之和 / n

			for _, sym := range symbols {
				ma := 0.0
				for _, one := range m[sym] {
					c, _ := one.Close.Float64()
					ma += c
				}
				ma = ma / float64(len(m[sym]))
				c, _ := m[sym][size-1].Close.Float64()
				o, _ := m[sym][size-1].Open.Float64()
				//h, _ := m[sym][size-1].Close.Float64()
				bias := (c - ma) / ma * 100
				println(fmt.Sprintf("sym: %s, bias: %0.4f, c: %0.4f", sym, bias, (c-o)/o))
			}
		}

		test()
		time.Sleep(time.Second * 1)
	}

}

type Record struct {
	Symbol string
	Rate   float64
}

type Records []*Record

func (s Records) Len() int { return len(s) }

func (s Records) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Records) Less(i, j int) bool { return s[i].Rate < s[j].Rate }

func getSymbols() []string {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSymbols()
	symbol := make([]string, 0)
	if err != nil {
		applogger.Error("Get symbols error: %s", err)
	} else {
		//applogger.Info("Get symbols, count=%d", len(resp))
		for _, result := range resp {
			//applogger.Info("symbol=%s, BaseCurrency=%s, QuoteCurrency=%s", result.Symbol, result.BaseCurrency, result.QuoteCurrency)
			has := strings.Contains(result.Symbol, "usdt")
			if has {
				symbol = append(symbol, result.Symbol)
			}
		}
	}
	return symbol
}

func getCandlestick(symbol string) (bool, []market.Candlestick) {

	client := new(client.MarketClient).Init(config.Host)

	optionalRequest := market.GetCandlestickOptionalRequest{Period: period, Size: size}
	resp, err := client.GetCandlestick(symbol, optionalRequest)

	b := false
	res := make([]market.Candlestick, 0)

	if len(resp) == size {
		if err != nil {
			applogger.Error(err.Error())
		} else {
			for _, kline := range resp {
				//c, _ := kline.Close.Float64()
				//o, _ := kline.Open.Float64()
				//if cg != c {
				//	cg = c
				//}
				//applogger.Info("High=%v, Low=%v, open=%v, close=%v, cha=%f", kline.High, kline.Low, kline.Open, kline.Close, (c-o)/c*100)
				res = append(res, kline)

			}
			for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
				res[i], res[j] = res[j], res[i]
			}
		}
		b = true
	}

	return b, res
}


// bias 乖离率
// kdj 随机指标
// bool 布林线
// obv 能量潮 OBV指标又称为能量潮，也叫成交量净额指标，OBV能量潮又称为平衡交易量，是由美国投资分析家葛兰碧在1981年创立的,它的理论基础是“能量是因,股价是果”。
// ccr
//