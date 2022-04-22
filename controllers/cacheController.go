package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/maki5/b4y_test/common"
	"github.com/maki5/b4y_test/models"
	"net/http"
)

type CacheRepo interface {
	GetCacheByKey(ctx context.Context, key string) (*models.Cache, error)
	CreateCache(ctx context.Context, key string) error
}

type CacheController struct {
	cacheRepo CacheRepo
}

func NewCacheController(repo CacheRepo) *CacheController {
	return &CacheController{cacheRepo: repo}
}

func (ctrl *CacheController) GetCache(c *gin.Context) {
	keyParam := c.Param("key")

	ctx := context.Background()
	cache, err := ctrl.cacheRepo.GetCacheByKey(ctx, keyParam)
	if err != nil && !common.NoDocuments(err) {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	output := make(map[string]interface{})

	if common.NoDocuments(err) {
		err := ctrl.cacheRepo.CreateCache(ctx, keyParam)
		if err != nil {
			output["error"] = err
			c.JSON(http.StatusInternalServerError, output)
			return
		}

		output["msg"] = "Cache miss"
		c.JSON(http.StatusOK, output)
		return
	}

	output["msg"] = "Cache hit"
	output["value"] = cache.Value

	c.JSON(http.StatusOK, output)
	return
}
