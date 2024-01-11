package router

import (
	"ChatBaseX_GPPTX/global"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/database", func(c *gin.Context) {
		c.JSON(200, global.DbConfig)
	})
	// router.POST("/api/ai-pool/rewardAIX", UserrewardAIX)
	// return router
	router.POST("/api/ai-pool/datereward", DepositHandler)
	return router
}
