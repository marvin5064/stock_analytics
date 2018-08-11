package stockfetch

import (
	"encoding/json"
	"fmt"

	stock "github.com/marvin5064/stock-analytics/protobuf/stock"
)

type AlphavantageQueryReturn struct {
	Metadata         AlphavantageMetadata              `json:"Meta Data"`
	TimeSeriesPrices map[string]AlphavantageDailyPrice `json:"Time Series (Daily)"`
}
type AlphavantageMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}
type AlphavantageDailyPrice struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func (m *manager) GetData(request *stock.StockPriceRequest) (*stock.StockPriceResponse, error) {
	_, body, errs := m.requester.
		Get("https://www.alphavantage.co/query").
		Param("function", "TIME_SERIES_DAILY").
		Param("symbol", request.GetSymbol()).
		Param("apikey", m.apiKey).
		End()

	if len(errs) != 0 {
		return nil, fmt.Errorf("%v", errs)
	}

	return parseDataReturn(body)
}

func parseDataReturn(body string) (*stock.StockPriceResponse, error) {
	dataReturned := &AlphavantageQueryReturn{}
	err := json.Unmarshal([]byte(body), dataReturned)
	if err != nil {
		return nil, err
	}
	fmt.Println("dataReturned", dataReturned)
	return nil, nil
}
