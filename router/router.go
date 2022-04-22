package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maki5/b4y_test/controllers"
)

func RegisterCacheRoutes(router *gin.RouterGroup, ctrl controllers.CacheController) {
	router.GET("/:key", ctrl.GetCache)
}
