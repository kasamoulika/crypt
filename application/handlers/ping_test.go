package handlers

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	testCases := []struct {
		description string
		expectCode  int
	}{
		{
			description: "200 Success response",
			expectCode:  200,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			r := httptest.NewRequest("GET", "/ping", nil)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			PingHandler(c)
			assert.Equal(t, tc.expectCode, w.Code)
			assert.NotNil(t, w.Body.String())
		})
	}
}
