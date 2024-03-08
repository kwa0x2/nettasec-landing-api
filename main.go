package main

import (
	"github.com/gin-gonic/gin"
	"nettasec.co.uk/landing/backend/config"
	"nettasec.co.uk/landing/backend/routes"
)

func main() {
	router := gin.New()
	config.Connection()
	routes.WaitlistRoute(router)
	router.Run(":9898")
}

