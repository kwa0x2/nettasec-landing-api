package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nettasec.co.uk/landing/backend/config"
	"nettasec.co.uk/landing/backend/models"
)


func CreateWaitlist(ctx *gin.Context) {

	var waitlist models.Waitlists
	var count int64


	if err := ctx.BindJSON(&waitlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input format"})
		return
	}

	if waitlist.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{ "error": "Missing email field" })
		return
	}

	config.DB.Model(&models.Waitlists{}).Where("email = ?", waitlist.Email).Count(&count)

	if count  > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email must be unique"})
		return
	}

	result := config.DB.Create(&waitlist)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{"rowAffected": result.RowsAffected})
	

}
