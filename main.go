package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maki5/b4y_test/controllers"
	"github.com/maki5/b4y_test/initialise"
	"github.com/maki5/b4y_test/router"
	"log"
)

func main() {
	log.Printf("app started")
	mongoClient, err := initialise.MongoClient()
	if err != nil {
		log.Fatal(err)
	}

	cacheRepo := initialise.CacheRepo(mongoClient)
	cacheController := controllers.NewCacheController(cacheRepo)

	r := gin.Default()
	api := r.Group("")
	router.RegisterCacheRoutes(api, *cacheController)

	r.Run()
}
