package stockfetch

import (
	"fmt"
)

func (m *manager) GetData(symbol string) (string, error) {
	_, body, errs := m.requester.
		Get("https://www.alphavantage.co/query").
		Param("function", "TIME_SERIES_DAILY").
		Param("symbol", symbol).
		Param("apikey", m.apiKey).
		End()

	return body, fmt.Errorf("%v", errs)
}
