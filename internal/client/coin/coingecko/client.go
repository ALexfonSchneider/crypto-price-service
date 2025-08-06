package coingecko

import (
	"context"
	"crypto-price-service/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	HttpTimeOut time.Duration
	Url         string
}

// Client — реализация MultiPriceFetcher
type Client struct {
	client *http.Client
	Url    string
}

func New(conf Config) *Client {
	return &Client{
		client: &http.Client{Timeout: conf.HttpTimeOut},
		Url:    conf.Url,
	}
}

// FetchPrices — получает цены для нескольких монет за один запрос
func (c *Client) FetchPrices(ctx context.Context, symbols []string) (dto.CurrentPrices, error) {
	// Формируем URL: ?symbols=btc,eth,sol&vs_currencies=usd
	url := fmt.Sprintf(
		"%s?symbols=%s&vs_currencies=usd",
		c.Url,
		strings.Join(symbols, ","),
	)

	fmt.Println(url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("coingecko api error: status %d", resp.StatusCode)
	}

	var result map[string]map[string]float64 // symbols->currency-|>price
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	for symbol, priceData := range result {
		if price, ok := priceData["usd"]; ok { // Для упрощения все цены будут в usd
			prices[symbol] = price
		}
	}

	return prices, nil
}
