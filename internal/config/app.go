package config

import "time"

type CoinPriceCollector struct {
	Interval time.Duration `koanf:"interval" validate:"required"`
}

type CoinPriceFetcher struct {
	Timeout   time.Duration   `koanf:"timeout" validate:"required"`
	Coingecko CoingeckoSource `koanf:"coingecko"`
}

type CoingeckoSource struct {
	Url string `koanf:"url" validate:"required"`
}

type App struct {
	CoinPriceCollector CoinPriceCollector `koanf:"coin_price_collector"`
	CoinPriceFetcher   CoinPriceFetcher   `koanf:"coin_price_fetcher"`
}
