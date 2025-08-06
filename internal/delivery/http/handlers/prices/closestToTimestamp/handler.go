package closestToTimestamp

import (
	_ "crypto-price-service/internal/delivery/http/dto"
	apperrors "crypto-price-service/internal/errors"
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

// closestToTimestamp godoc
// @Summary get closest coin price
// @Param request body Request true "get closest price for coin"
// @Tags Prices
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse "ближайшая цена монеты"
// @Failure      400 {object} dto.ErrorResponse "Input error"
// @Failure      404 {object} dto.ErrorResponse "Coin does not exists"
// @Failure      404 {object} dto.ErrorResponse "Price does not exists yet"
// @Failure      500 {object} dto.ErrorResponse "Internal error"
// @Router /currency/price [post]
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

	t := time.Unix(req.Timestamp, 0)
	price, err := h.prices.ClosestByCoinID(ctx, coin.ID, t)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Symbol:    req.Symbol,
		Price:     price.Price,
		CreatedAt: price.CreatedAt,
	})

}
