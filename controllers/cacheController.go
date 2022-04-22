package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CacheController struct {

}

func NewCacheController() *CacheController {
	return &CacheController{}
}

func (c *CacheController) GetCache(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "sample str")
	return
}
