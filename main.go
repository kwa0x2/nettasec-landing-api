package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"nettasec.co.uk/landing/backend/config"
	"nettasec.co.uk/landing/backend/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	allowedIPs := os.Getenv("ALLOWED_IPS")
	allowedIPList := strings.Split(allowedIPs, ",")
	fmt.Print("ips", allowedIPList)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedIPList,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	config.Connection()
	routes.WaitlistRoute(router)
	router.RunTLS(":9898", "myndfin.crt", "myndfin.key")
}
