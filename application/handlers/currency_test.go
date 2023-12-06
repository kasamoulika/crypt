package handlers

import (
	"errors"
	"fmt"
	"net/http/httptest"

	"testing"

	"crypt.com/v2/application/config"
	"crypt.com/v2/application/models"
	"crypt.com/v2/application/services/external/hitbtc"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadConfig("../..", "dev")
}

func TestGetCurrency(t *testing.T) {
	testCases := []struct {
		description      string
		symbol           string
		hitbtctestclient *hitbtc.TestClient
		expectCode       int
	}{
		{
			description: "200 Success response",
			symbol:      "AAPL",
			expectCode:  200,
			hitbtctestclient: &hitbtc.TestClient{
				GetCurrencyData: &models.Currency{
					Id:        "ETHBTC",
					LastPrice: "111.5",
				},
			},
		},
		{
			description: "Error response from hitbtc",
			symbol:      "ETHBTC",
			expectCode:  500,
			hitbtctestclient: &hitbtc.TestClient{
				GetCurrencyDataErr: errors.New("HTTP ERROR"),
			},
		},
		{
			description: "400 BadRequest - Invalid symbol",
			symbol:      "AAPL",
			expectCode:  400,
			hitbtctestclient: &hitbtc.TestClient{
				GetCurrencyData: &models.Currency{
					Id:        "ETHBTC",
					LastPrice: "111.5",
				},
			},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			r := httptest.NewRequest("GET", fmt.Sprintf("/v1/currency/%s", tc.symbol), nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			currencyHandler := NewCurrencyHandler(tc.hitbtctestclient)
			currencyHandler.GetCurrency(c)
			fmt.Println(w.Body.String())
			assert.Equal(t, tc.expectCode, w.Code)
			assert.NotNil(t, w.Body.String())
		})
	}
}

func TestGetAllCurrency(t *testing.T) {
	testCases := []struct {
		description      string
		hitbtctestclient *hitbtc.TestClient
		expectCode       int
	}{
		{
			description: "200 Success response",
			expectCode:  200,
			hitbtctestclient: &hitbtc.TestClient{
				GetAllCurrencyData: []models.Currency{
					{
						Id:        "ETHBTC",
						LastPrice: "111.5",
					},
				},
			},
		},
		{
			description: "Error response from hitbtc",
			expectCode:  500,
			hitbtctestclient: &hitbtc.TestClient{
				GetCurrencyDataErr: errors.New("HTTP ERROR"),
			},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			r := httptest.NewRequest("GET", fmt.Sprintf("/v1/currency/all"), nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			currencyHandler := NewCurrencyHandler(tc.hitbtctestclient)
			currencyHandler.GetAllCurrency(c)
			fmt.Println(w.Body.String())
			assert.Equal(t, tc.expectCode, w.Code)
			assert.NotNil(t, w.Body.String())
		})
	}
}
