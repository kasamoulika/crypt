package hitbtc

import (
	"crypt.com/v2/application/models"
	"github.com/gin-gonic/gin"
)

type TestClient struct {
	GetCurrencyData    *models.Currency
	GetCurrencyDataErr error
}

func (c TestClient) GetCurrency(ctx *gin.Context, symbol string) (*models.Currency, error) {
	return c.GetCurrencyData, c.GetCurrencyDataErr
}
