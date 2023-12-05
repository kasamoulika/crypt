package hitbtc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sync"

	"crypt.com/v2/application/config"
	"crypt.com/v2/application/constants"
	"crypt.com/v2/application/errors"
	"crypt.com/v2/application/models"
	"crypt.com/v2/application/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var client services.MarketDataClient
var once sync.Once

// HitBtcClient represents the client structure of hitbtc service
type HitBtcClient struct {
	baseUrl string
}

// currencySymbolResult represents result object of currency symbol API
type currencySymbolResult struct {
	Ask  string `json:"ask"`
	Bid  string `json:"bid"`
	Last string `json:"last"`
	Low  string `json:"low"`
	High string `json:"high"`
	Open string `json:"open"`
}

// NewHitBtcClient initializes new hitbtc market data client
func NewHitBtcClient() services.MarketDataClient {
	once.Do(func() {
		hitbtcConfig := config.GetConfig().HitBtc
		if hitbtcConfig.Mock == true {
			client = &TestClient{
				GetCurrencyData: &models.Currency{Id: "AAPL", LastPrice: "173.73"},
			}
			return
		}

		client = &HitBtcClient{
			hitbtcConfig.BaseUrl,
		}
	})
	return client
}

// GetCurrency returns currency stock price of the given currency symbol
func (c *HitBtcClient) GetCurrency(ctx *gin.Context, symbol string) (*models.Currency, error) {
	url := fmt.Sprintf("%s/api/3/public/ticker/%s", c.baseUrl, symbol)
	// TODO: change below log level to Debug
	log.Info(ctx, url)
	response, err := http.Get(url)
	if err != nil {
		log.WithContext(ctx).Error(ctx, "HitBtc GetCurrency request failed", err, nil)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.WithContext(ctx).Error(ctx, "error: Non-OK status code received ", response.Body)
		return nil, errors.NewCurrencyError(response.StatusCode, constants.HitBtcError, nil)
	}

	var result currencySymbolResult
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		log.WithContext(ctx).Error(ctx, "error decoding currency response ", response.Body)
		return nil, errors.NewCurrencyError(response.StatusCode, constants.HitBtcError, nil)
	}

	finalResponse := &models.Currency{
		Id:           symbol,
		AskPrice:     result.Ask,
		BidPrice:     result.Bid,
		HighestPrice: result.High,
		LowestPrice:  result.Low,
		OpenPrice:    result.Open,
		LastPrice:    result.Last,
	}

	return finalResponse, nil
}
