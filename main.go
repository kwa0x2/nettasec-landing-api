package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"nettasec.co.uk/landing/backend/config"
	"nettasec.co.uk/landing/backend/routes"
)

func main() {

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	config.Connection()
	routes.WaitlistRoute(router)
	router.RunTLS(":9898", "myndfin.crt", "myndfin.key")
}
