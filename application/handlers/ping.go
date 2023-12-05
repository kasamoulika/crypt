package handlers

import "github.com/gin-gonic/gin"

// PingResponse is the response structure of PingHandler
type PingResponse struct {
	Message            string
	ServerHealthStatus int
}

// PingHandler is the handler for GET /ping endpoint
func PingHandler(ctx *gin.Context) {
	ctx.JSON(200, &PingResponse{
		ServerHealthStatus: 200,
		Message:            "Welcome to Crypt",
	})
}
