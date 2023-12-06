package handlers

import (
	"net/http"

	"crypt.com/v2/application/services"
	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	marketDataClient services.MarketDataClient
}

func NewCurrencyHandler(mdClient services.MarketDataClient) *CurrencyHandler {
	return &CurrencyHandler{
		marketDataClient: mdClient,
	}
}

// Get currency information by currency symbol.
func (h *CurrencyHandler) GetCurrency(c *gin.Context) {
	symbol := c.Param("symbol")
	resp, err := h.marketDataClient.GetCurrency(c, symbol)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Get all currency information.
func (h *CurrencyHandler) GetAllCurrency(c *gin.Context) {
	resp, err := h.marketDataClient.GetAllCurrency(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
