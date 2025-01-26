package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/s-uryansh/blockchain-gin/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/blockchain", handlers.GetBlockchain)
	r.POST("/blockchain", handlers.AddBlock)
	return r
}
