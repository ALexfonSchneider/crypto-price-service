package AllForCoin

import (
	_ "crypto-price-service/internal/delivery/http/dto"
	apperrors "crypto-price-service/internal/errors"
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
// @Summary get all prices for coin
// @Tags Prices
// @Param request body Request true "get all prices for coin"
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} dto.ErrorResponse "Coin doew not exists"
// @Failure 500 {object} dto.ErrorResponse "Internal error"
// @Router /currency/prices [post]
func (h *Handler) Handle(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperrors.NewInvalidRequest().Wrap(err))
		return
	}

	ctx := c.Request.Context()

	coin, err := h.coins.BySymbol(ctx, req.Symbol)
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
