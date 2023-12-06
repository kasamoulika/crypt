package handlers

import (
	"net/http"

	"crypt.com/v2/application/config"
	"crypt.com/v2/application/constants"
	"crypt.com/v2/application/errors"
	"crypt.com/v2/application/models"
	"crypt.com/v2/application/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
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

	if !isValidSymbol(symbol) {
		c.JSON(http.StatusBadRequest, &models.ErrorResponse{
			Code:        http.StatusBadRequest,
			Description: "Invalid symbol",
		})
		return
	}

	resp, err := h.marketDataClient.GetCurrency(c, symbol)

	if err != nil {
		if e, ok := err.(*errors.SError); ok {
			switch e.Description {
			case constants.Error404:
				c.JSON(e.Code, models.GetErrorResponse(e))
				return
			}
		}
		log.WithContext(c).WithError(err).Error("error in fetching currency data")
		sErr := errors.NewCurrencyError(http.StatusInternalServerError, constants.Error500, err)
		c.JSON(http.StatusInternalServerError, models.GetErrorResponse(sErr))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func isValidSymbol(symbol string) bool {
	supportedSymbols := config.GetConfig().Symbol.SupportedSymbols
	return slices.Contains(supportedSymbols, symbol)
}

// Get all currency information.
func (h *CurrencyHandler) GetAllCurrency(c *gin.Context) {
	resp, err := h.marketDataClient.GetAllCurrency(c)

	if err != nil {
		if e, ok := err.(*errors.SError); ok {
			switch e.Description {
			case constants.Error404:
				c.JSON(e.Code, models.GetErrorResponse(e))
				return
			}
		}
		log.WithContext(c).WithError(err).Error("error in fetching currency data")
		sErr := errors.NewCurrencyError(http.StatusInternalServerError, constants.Error500, err)
		c.JSON(http.StatusInternalServerError, models.GetErrorResponse(sErr))
		return
	}

	c.JSON(http.StatusOK, resp)
}
