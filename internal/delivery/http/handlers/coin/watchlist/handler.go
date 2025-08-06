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

// Watchlist godoc
// @Summary      Get list of watched coins
// @Description  Returns a list of all coins currently in the watchlist with their details
// @Tags         Coins
// @Produce      json
// @Success      200 {object} SuccessResponse "List of coins in the watchlist"
// @Failure      500 {object} dto.ErrorResponse "Internal server error"
// @Router       /coins [get]
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
