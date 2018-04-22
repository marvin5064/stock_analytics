package lib

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func main() {
	resp, body, errs := gorequest.New().
		Get("https://www.alphavantage.co/query").
		Param("function", "TIME_SERIES_DAILY").
		Param("symbol", "0941.hk").
		Param("apikey", "demo").
		End()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(resp, body)
}
