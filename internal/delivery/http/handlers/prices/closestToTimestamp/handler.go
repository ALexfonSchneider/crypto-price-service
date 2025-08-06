package closestToTimestamp

import (
	_ "crypto-price-service/internal/delivery/http/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	coins  Coins
	prices Prices
}

func New(coins Coins, prices Prices) *Handler {
	return &Handler{
		coins:  coins,
		prices: prices,
	}
}

// ClosestToTimestamp godoc
// @Summary      Get the closest price for a coin at a given timestamp
// @Description  Returns the price record closest to the specified timestamp for a given coin symbol
// @Tags         Prices
// @Param        symbol      path    string  true   "Coin symbol (e.g. BTC, ETH)"
// @Param        timestamp   query    int64   true   "Timestamp in milliseconds"
// @Produce      json
// @Success      200 {object} SuccessResponse
// @Failure      400 {object} dto.ErrorResponse "Invalid timestamp or missing parameters"
// @Failure      404 {object} dto.ErrorResponse "Coin not found or no price data exists"
// @Failure      500 {object} dto.ErrorResponse "Internal server error"
// @Router       /coins/{symbol}/price/closest [get]
func (h *Handler) Handle(c *gin.Context) {
	symbol := c.Param("symbol")
	timestamp := c.GetInt64("timestamp")

	ctx := c.Request.Context()

	coin, err := h.coins.BySymbol(ctx, symbol)
	if err != nil {
		_ = c.Error(err)
		return
	}

	t := time.Unix(timestamp, 0)
	price, err := h.prices.ClosestByCoinID(ctx, coin.ID, t)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Symbol:    coin.Symbol,
		Price:     price.Price,
		CreatedAt: price.CreatedAt,
	})

}
