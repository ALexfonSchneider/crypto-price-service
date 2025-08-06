// @title Crypto Price API
// @version 1.0
// @description API for watching crypto prices
// @BasePath /api/v1

package main

import (
	"context"
	_ "crypto-price-service/docs"
	coingeckoclient "crypto-price-service/internal/client/coin/coingecko"
	"crypto-price-service/internal/config"
	db "crypto-price-service/internal/db/gen"
	"crypto-price-service/internal/delivery/http/handlers/coin/addCoin"
	"crypto-price-service/internal/delivery/http/handlers/coin/removeCoin"
	"crypto-price-service/internal/delivery/http/handlers/coin/watchlist"
	"crypto-price-service/internal/delivery/http/handlers/prices/AllForCoin"
	"crypto-price-service/internal/delivery/http/handlers/prices/closestToTimestamp"
	middleware "crypto-price-service/internal/delivery/http/middlewares"
	"crypto-price-service/internal/repository/postgres/coin"
	"crypto-price-service/internal/repository/postgres/price"
	coinservice "crypto-price-service/internal/services/coin"
	priceservice "crypto-price-service/internal/services/price"
	watchers "crypto-price-service/internal/watcher"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	conf := config.MustConfig()

	pool, err := pgxpool.New(ctx, conf.Postgres.DSN())
	if err != nil {
		panic(err)
	}

	querier := db.New(pool)

	coinRepo := coin.New(querier)
	priceRepo := price.New(querier)

	coinService := coinservice.NewService(coinRepo)
	priceService := priceservice.NewService(priceRepo)

	r := gin.Default()

	addCoinHandler := addCoin.New(coinService)
	listCoinsHandler := watchlist.New(coinService)
	removeCoinHandler := removeCoin.New(coinService)

	closestCoinPriceHandler := closestToTimestamp.New(coinService, priceService)
	allPricesForCoin := AllForCoin.New(coinService, priceService)

	r.Use(middleware.Logger(logger), middleware.ErrorMiddleware(logger))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/currency/add", addCoinHandler.Handle)
		v1.GET("/currency/list", listCoinsHandler.Handle)
		v1.DELETE("/currency/remove", removeCoinHandler.Handle)

		v1.POST("/currency/price", closestCoinPriceHandler.Handle)
		v1.POST("/currency/prices", allPricesForCoin.Handle)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().UTC()})
	})

	// swagger доступен по /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fetcher := coingeckoclient.New(coingeckoclient.Config{
		HttpTimeOut: conf.App.CoinPriceFetcher.Timeout,
		Url:         conf.App.CoinPriceFetcher.Coingecko.Url,
	})
	watcher := watchers.New(priceService, coinService, fetcher, logger, time.Second*5)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := watcher.Watch(ctx); err != nil {
			logger.Error(err.Error())
		}
	}()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.HttpServer.Host, conf.HttpServer.Port),
		Handler: r,
	}

	go func() {
		slog.Info("Server started", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server error", "error", err)
		}
	}()

	<-ctx.Done()
	slog.Info("Got stop signal, stopping...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("Got graceful shutdown error", "error", err)
		srv.Close()
	}

	wg.Wait()

	slog.Info("Server stopped correctly")
}
