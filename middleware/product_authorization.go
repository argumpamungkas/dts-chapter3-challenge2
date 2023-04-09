package middleware

import (
	"DTS/Chapter-3/chapter3-challenge2/models"
	"DTS/Chapter-3/chapter3-challenge2/repo"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := repo.GetDB()

		productID, _ := strconv.Atoi(ctx.Param("productID"))

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		roleID := uint(userData["role"].(float64))

		if productID == 0 {
			var products []models.Product

			if roleID == 1 {
				err := db.Debug().Find(&products).Error
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"message": err.Error(),
					})
					return
				}
			} else if roleID == 2 {
				err := db.Debug().Where("user_id = ?", userID).Find(&products).Error
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"message": err.Error(),
					})
					return
				}
			}

			ctx.Next()
		}

		var product models.Product

		if productID != 0 {
			err := db.Select("user_id").First(&product, uint(productID)).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Data not found",
				})
				return
			}
			if roleID == 1 {
				ctx.Next()
			}

			if roleID == 2 && product.UserID != userID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "You are not allowed to access this data",
				})
				return
			}
			ctx.Next()
		}

	}
}
