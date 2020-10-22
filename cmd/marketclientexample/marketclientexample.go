package marketclientexample

import (
	"fmt"
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"sort"
	"strings"
)

func RunAllExamples() {
	getCandlestick()
	//getLast24hCandlestickAskBid()
	//getLast24hCandlesticks()
	//getDepth()
	//getLatestTrade()
	//getHistoricalTrade()
	//getLast24hCandlestick()
}

func GetCandlestick(symbol string) (res float64) {
	client := new(client.MarketClient).Init(config.Host)

	optionalRequest := market.GetCandlestickOptionalRequest{Period: market.DAY1, Size: 1}
	resp, err := client.GetCandlestick(symbol, optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, kline := range resp {
			o, _ := kline.Open.Float64()
			c, _ := kline.Close.Float64()
			//applogger.Info("High=%v, Low=%v, open=%v, close=%v, cha=%f", kline.High, kline.Low, kline.Open, kline.Close, (c-o)/o*100)
			res = (c - o) / o * 100
		}
	}
	return
}

//  Get the candlestick/kline for the btcusdt. The specified data number is 10 .
var cg float64
var closeG = 0.256984

func getCandlestick() {
	client := new(client.MarketClient).Init(config.Host)

	optionalRequest := market.GetCandlestickOptionalRequest{Period: market.DAY1, Size: 30}
	resp, err := client.GetCandlestick("titanusdt", optionalRequest)
	if len(resp) == 30 {
		if err != nil {
			applogger.Error(err.Error())
		} else {
			for _, kline := range resp {
				c, _ := kline.Close.Float64()
				if cg != c {
					cg = c
					applogger.Info("High=%v, Low=%v, open=%v, close=%v, cha=%f", kline.High, kline.Low, kline.Open, kline.Close, (c-closeG)/closeG*100)
				}
			}
		}
	}
}

func GetLast24hCandlestickAskBid() (res float64) {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestickAskBid("forusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		//applogger.Info("Bid=%+v, Ask=%+v", resp.Bid, resp.Ask)
		c, _ := resp.Close.Float64()
		o, _ := resp.Open.Float64()
		res = (c - o) / o
		//applogger.Info("High: %v, Low: %v, Close: %v, Open: %v, cha[%s]",
		//	resp.High, resp.Low, resp.Close, resp.Open, (c-o)/o)
	}
	return
}

//  Get the latest ticker with some important 24h aggregated market data for btcusdt.
func getLast24hCandlestickAskBid() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestickAskBid("cruusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		//applogger.Info("Bid=%+v, Ask=%+v", resp.Bid, resp.Ask)
		c, _ := resp.Close.Float64()
		o, _ := resp.Open.Float64()
		applogger.Info("High: %v, Low: %v, Close: %v, Open: %v, cha[%s]",
			resp.High, resp.Low, resp.Close, resp.Open, (c-o)/o)
	}
}

//  Get the latest tickers for all supported pairs
func getLast24hCandlesticks() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetAllSymbolsLast24hCandlesticksAskBid()

	rmIndex := make([]int, 0)

	if err != nil {
		applogger.Error(err.Error())
	} else {
		for i, tick := range resp {
			if tick.Open.IsZero() || tick.Close.IsZero() {
				rmIndex = append(rmIndex, i)
			}
		}
	}

	for _, index := range rmIndex {
		resp = append(resp[:index], resp[index+1:]...)
	}

	sort.Sort(market.SymbolCandlesticks(resp))
	str := ""
	i := 0
	for _, tick := range resp {
		if strings.Contains(tick.Symbol, "usdt") {
			c, _ := tick.Close.Float64()
			o, _ := tick.Open.Float64()
			//applogger.Info("Symbol: %s, High: %v, Low: %v, Close: %v, Open: %v, Ask[%v, %v], Bid[%v, %v], cha[%s]",
			//	tick.Symbol, tick.High, tick.Low, tick.Close, tick.Open, tick.Ask, tick.AskSize, tick.Bid, tick.BidSize, (c-o)/o)
			str += fmt.Sprintf("<symbol: %s, cha: %f>", tick.Symbol, (c-o)/o*100)
			i++
		}

		if i > 4 {
			break
		}
	}

	applogger.Info(str)

}

//  Get the current order book of the btcusdt.
func getDepth() {
	optionalRequest := market.GetDepthOptionalRequest{10}
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetDepth("btcusdt", market.STEP0, optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, ask := range resp.Asks {
			applogger.Info("ask: %+v", ask)
		}
		for _, bid := range resp.Bids {
			applogger.Info(": %+v", bid)
		}

	}
}

//  Get the latest trade with btucsdt price, volume, and direction.
func getLatestTrade() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLatestTrade("btcusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, trade := range resp.Data {
			applogger.Info("Id=%v, Price=%v", trade.Id, trade.Price)
		}
	}
}

//  Get the most recent trades with btcusdt price, volume, and direction.
func getHistoricalTrade() {
	client := new(client.MarketClient).Init(config.Host)
	optionalRequest := market.GetHistoricalTradeOptionalRequest{5}
	resp, err := client.GetHistoricalTrade("btcusdt", optionalRequest)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		for _, tradeData := range resp {
			for _, trade := range tradeData.Data {
				applogger.Info("price: %v", trade.Price)
			}
		}
	}
}

//  Get the summary of trading in the market for the last 24 hours.
func getLast24hCandlestick() {
	client := new(client.MarketClient).Init(config.Host)

	resp, err := client.GetLast24hCandlestick("btcusdt")
	if err != nil {
		applogger.Error(err.Error())
	} else {
		applogger.Info("Close=%v, Open=%v", resp.Close, resp.Open)
	}
}
