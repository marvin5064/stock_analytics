package main

import (
	"github.com/marvin5064/stock-analytics/lib/logger"
	"github.com/marvin5064/stock-analytics/lib/stockfetch"
	"github.com/spf13/viper"
)

type Server struct {
	stockFetchManager stockfetch.Manager
}

func main() {
	srv := &Server{}
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

	url := viper.GetString("url")
	if url != "" {
		logger.Info("successfully loaded url from config")
	}

	srv.stockFetchManager = stockfetch.New(url, apikey)
	data, err := srv.stockFetchManager.GetData("0941.hk")
	if err != nil {
		logger.Error(err)
	}
	logger.Info(data)
}
