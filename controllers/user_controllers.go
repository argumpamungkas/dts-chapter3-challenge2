package controllers

import (
	"DTS/Chapter-3/chapter3-challenge2/helpers"
	"DTS/Chapter-3/chapter3-challenge2/models"
	"DTS/Chapter-3/chapter3-challenge2/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func RegisterUser(ctx *gin.Context) {
	db := repo.GetDB()
	var user models.User
	var role models.Role

	contentType := helpers.GetContentType(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	err := db.Debug().First(&role, user.RoleID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Role ID not found",
		})
		return
	}

	err = db.Debug().Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"full_name": user.FullName,
			"email":     user.Email,
			"role":      role.RoleName,
		},
		"message": "User succes created",
	})

}

func LoginUser(ctx *gin.Context) {
	db := repo.GetDB()
	var user models.User

	contentType := helpers.GetContentType(ctx)

	if contentType == appJson {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	password := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email",
		})
		return
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.RoleID, user.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
