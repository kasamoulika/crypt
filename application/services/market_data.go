package services

import (
	"crypt.com/v2/application/models"
	"github.com/gin-gonic/gin"
)

// MarketDataClient is the interface providing all the required methods to serve market data
type MarketDataClient interface {

	// GetCurrency retrieves information about a specific currency by its symbol.
	GetCurrency(ctx *gin.Context, symbol string) (*models.Currency, error)
}
