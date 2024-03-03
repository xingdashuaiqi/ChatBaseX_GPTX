package router

import (
	"ChatBaseX-GPTX0113/dao"
	"ChatBaseX-GPTX0113/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		root, _ := dao.GetCommunityEarnings(8315778129200896)
		c.JSON(200, gin.H{
			"message": root,
		})
	})

	return r
}
