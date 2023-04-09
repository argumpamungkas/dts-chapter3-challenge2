package controllers

import (
	"DTS/Chapter-3/chapter3-challenge2/helpers"
	"DTS/Chapter-3/chapter3-challenge2/models"
	"DTS/Chapter-3/chapter3-challenge2/repo"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	db := repo.GetDB()
	var product models.Product

	contentType := helpers.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		ctx.ShouldBindJSON(&product)
	} else {
		ctx.ShouldBind(&product)
	}

	product.UserID = userID

	err := db.Debug().Create(&product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)

}

func GetProductById(ctx *gin.Context) {
	db := repo.GetDB()
	var product models.Product

	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	err = db.Debug().First(&product, "id = ?", productID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func GetAllProduct(ctx *gin.Context) {
	db := repo.GetDB()
	var product []models.Product

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	roleID := uint(userData["role"].(float64))

	if roleID == 1 {
		err := db.Debug().Find(&product).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	if roleID == 2 {
		err := db.Debug().Where("user_id", userID).Find(&product).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data_product": product,
	})

}

func UpdateProduct(ctx *gin.Context) {
	db := repo.GetDB()
	var product, findProduct models.Product

	contentType := helpers.GetContentType(ctx)

	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	roleID := uint(userData["role"].(float64))

	if contentType == appJson {
		ctx.ShouldBindJSON(&product)
	} else {
		ctx.ShouldBind(&product)
	}

	product = models.Product{
		Title:       product.Title,
		Description: product.Description,
	}

	err = db.Where("id = ?", productID).First(&findProduct).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product.ID = uint(productID)

	if roleID == 1 {

		product.UserID = findProduct.UserID

		err := db.Model(&product).Where("id = ?", productID).Updates(product).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	} else if roleID == 2 {
		product.UserID = userID
		err := db.Model(&product).Where("id = ?", productID).Updates(product).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusCreated, product)
}

func DeleteProduct(ctx *gin.Context) {
	db := repo.GetDB()
	var product models.Product

	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	err = db.Debug().Where("id = ?", productID).First(&product).Delete(&product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product %s success deleted", product.Title),
	})
}
