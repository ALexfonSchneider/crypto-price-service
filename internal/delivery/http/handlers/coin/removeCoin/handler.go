package removeCoin

import (
	_ "crypto-price-service/internal/delivery/http/dto"
	apperrors "crypto-price-service/internal/errors"
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

// removeCoin godoc
// @Summary remove coin from watchlist
// @Tags Coins
// @Param request body Request true "remove coin request"
// @Accept json
// @Produce json
// @Success 200 "No content (only status)"
// @Failure 500 {object} dto.ErrorResponse "Internal error"
// @Router /currency/remove [delete]
func (h *Handler) Handle(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperrors.NewInvalidRequest().Wrap(err))
		return
	}

	if err := h.coins.Deactivate(c.Request.Context(), req.Symbol); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
