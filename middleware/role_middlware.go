package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RoleMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		roleID := uint(userData["role"].(float64))

		if roleID != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not admin",
			})
			return
		}

		ctx.Next()
	}
}
