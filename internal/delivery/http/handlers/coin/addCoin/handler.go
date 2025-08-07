package addCoin

import (
	_ "crypto-price-service/internal/delivery/http/dto"
	apperrors "crypto-price-service/internal/errors"
	"crypto-price-service/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	coins Coins
}

func New(coins Coins) *Handler {
	return &Handler{coins: coins}
}

// AddCoin godoc
// @Summary      Add a coin to the watchlist
// @Description  Adds a new coin to the watchlist and returns the created coin
// @Tags         Coins
// @Accept       json
// @Produce      json
// @Param        request  body     Request  true  "Add coin request"
// @Success      201 {object} SuccessResponse "Created coin object"
// @Failure      400 {object} dto.ErrorResponse "Invalid request"
// @Failure      500 {object} dto.ErrorResponse "Internal server error"
// @Router       /coins [post]
func (h *Handler) Handle(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperrors.NewInvalidRequest().Wrap(err))
		return
	}

	coin := models.NewCoin(req.Name, req.Symbol, true, time.Now())
	if err := h.coins.Activate(c.Request.Context(), coin); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, responseFromModel(*coin))
}
