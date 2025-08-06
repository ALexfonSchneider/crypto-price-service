package AllForCoin

import (
	_ "crypto-price-service/internal/delivery/http/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	coins  Coins
	prices Prices
}

func New(coins Coins, prices Prices) *Handler {
	return &Handler{coins: coins, prices: prices}
}

// AllPricesForCoin godoc
// @Summary      Get all prices for a coin
// @Description  Returns all recorded prices for the given coin symbol
// @Tags         Prices
// @Param        symbol path string true "Coin symbol (e.g. BTC, ETH)"
// @Accept       json
// @Produce      json
// @Success      200 {object} SuccessResponse
// @Failure      404 {object} dto.ErrorResponse "Coin does not exist"
// @Failure      500 {object} dto.ErrorResponse "Internal server error"
// @Router       /coins/{symbol}/prices [get]
func (h *Handler) Handle(c *gin.Context) {
	symbol := c.Param("symbol")

	ctx := c.Request.Context()

	coin, err := h.coins.BySymbol(ctx, symbol)
	if err != nil {
		_ = c.Error(err)
		return
	}

	prices, err := h.prices.GetAllPricesForCoinByCoinID(ctx, coin.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, FromModel(prices))
}
