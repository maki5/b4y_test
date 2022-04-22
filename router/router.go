package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maki5/b4y_test/controllers"
)

func RegisterRoutes(router *gin.RouterGroup) {
	cacheController := controllers.NewCacheController()

	router.GET("/", cacheController.GetCache)
}