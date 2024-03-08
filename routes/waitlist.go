package routes

import (
	"github.com/gin-gonic/gin"
	"nettasec.co.uk/landing/backend/controller"
)

func WaitlistRoute(router *gin.Engine) {
	waitlist := router.Group("/api")
	waitlist.POST("/waitlist", controller.CreateWaitlist)
}
