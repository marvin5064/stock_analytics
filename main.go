package main

import (
	"github.com/marvin5064/stock-analytics/lib/logger"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

func main() {
	defer logger.Sync()

	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("fail to load config")
	}

	apikey := viper.GetString("apikey")
	if apikey != "" {
		logger.Info("successfully loaded api key from config")
	}
	resp, body, errs := gorequest.New().
		Get("https://www.alphavantage.co/query").
		Param("function", "TIME_SERIES_DAILY").
		Param("symbol", "0941.hk").
		Param("apikey", apikey).
		End()
	if errs != nil {
		logger.Error(errs)
	}
	logger.Info(resp, body)
}
