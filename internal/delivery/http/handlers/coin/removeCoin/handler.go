package removeCoin

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

// RemoveCoin godoc
// @Summary      Remove coin from watchlist
// @Description  Removes a coin from the watchlist by its symbol
// @Tags         Coins
// @Param        symbol  path     string  true  "Coin symbol (e.g. BTC, ETH)"
// @Produce      json
// @Success      204 "Coin removed successfully (no content)"
// @Failure      500  {object} dto.ErrorResponse "Internal server error"
// @Router       /coins/{symbol} [delete]
func (h *Handler) Handle(c *gin.Context) {
	symbol := c.Param("symbol")

	if err := h.coins.Deactivate(c.Request.Context(), symbol); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
