package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maki5/b4y_test/router"
	"log"
)

func main() {
	log.Printf("app started")

	r := gin.Default()
	api := r.Group("")
	router.RegisterRoutes(api)

	r.Run()
}
