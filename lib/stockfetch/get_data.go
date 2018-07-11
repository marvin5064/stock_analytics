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

	if len(errs) != 0 {
		return "", fmt.Errorf("%v", errs)
	}
	return body, nil
}
