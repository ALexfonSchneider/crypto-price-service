package watchlist

import (
	_ "crypto-price-service/internal/delivery/http/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	coins Coins
}

func New(coins Coins) *Handler {
	return &Handler{
		coins: coins,
	}
}

// watchlist godoc
// @Summary get coins watchlist
// @Tags Coins
// @Produce json
// @Success 200 {object} SuccessResponse "Watchable coins"
// @Failure 500 {object} dto.ErrorResponse "Internal error"
// @Router /currency/list [get]
func (h *Handler) Handle(c *gin.Context) {
	coins, err := h.coins.ListActive(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Coins: fromModel(coins),
	})
}
