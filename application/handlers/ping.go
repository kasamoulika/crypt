package handlers

import "github.com/gin-gonic/gin"

// PingResponse is the response structure of GetPing
type PingResponse struct {
	Message            string
	ServerHealthStatus int
}

// GetPing is the handler for GET /ping endpoint
func GetPing(ctx *gin.Context) {
	ctx.JSON(200, &PingResponse{
		ServerHealthStatus: 200,
		Message:            "Welcome to Crypt",
	})
}
