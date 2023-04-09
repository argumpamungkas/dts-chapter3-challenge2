package controllers

import (
	"DTS/Chapter-3/chapter3-challenge2/models"
	"DTS/Chapter-3/chapter3-challenge2/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRole(ctx *gin.Context) {
	db := repo.GetDB()
	var role models.Role

	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Debug().Create(&role).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid create data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

func GetRole(ctx *gin.Context) {
	db := repo.GetDB()
	var roleDatas []models.Role

	err := db.Debug().Find(&roleDatas).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": roleDatas,
	})

}
