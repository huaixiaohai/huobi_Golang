package market

import "github.com/shopspring/decimal"

type GetAllSymbolsLast24hCandlesticksAskBidResponse struct {
	Status string              `json:"status"`
	Ts     int64               `json:"ts"`
	Data   []SymbolCandlestick `json:"data"`
}
type SymbolCandlestick struct {
	Amount  decimal.Decimal `json:"amount"`
	Open    decimal.Decimal `json:"open"`
	Close   decimal.Decimal `json:"close"`
	High    decimal.Decimal `json:"high"`
	Symbol  string          `json:"symbol"`
	Count   int64           `json:"count"`
	Low     decimal.Decimal `json:"low"`
	Vol     decimal.Decimal `json:"vol"`
	Bid     decimal.Decimal `json:"bid"`
	BidSize decimal.Decimal `json:"bidSize"`
	Ask     decimal.Decimal `json:"ask"`
	AskSize decimal.Decimal `json:"askSize"`
}

type SymbolCandlesticks []SymbolCandlestick

//Len()
func (s SymbolCandlesticks) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s SymbolCandlesticks) Less(i, j int) bool {
	oi, _ := s[i].Open.Float64()
	ci, _ := s[i].Close.Float64()

	oj, _ := s[j].Open.Float64()
	cj, _ := s[j].Close.Float64()

	return  (ci-oi)/oi  <  (cj-oj)/oj
}

//Swap()
func (s SymbolCandlesticks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
