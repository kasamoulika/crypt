package hitbtc

import (
	"crypt.com/v2/application/models"
	"github.com/gin-gonic/gin"
)

type TestClient struct {
	GetCurrencyData    *models.Currency
	GetCurrencyDataErr error
	GetAllCurrencyData []models.Currency
}

func (c TestClient) GetCurrency(ctx *gin.Context, symbol string) (*models.Currency, error) {
	return c.GetCurrencyData, c.GetCurrencyDataErr
}

func (c TestClient) GetAllCurrency(ctx *gin.Context) ([]models.Currency, error) {
	return c.GetAllCurrencyData, c.GetCurrencyDataErr
}
