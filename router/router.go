package router

import (
	"DTS/Chapter-3/chapter3-challenge2/controllers"
	"DTS/Chapter-3/chapter3-challenge2/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	roleRouter := r.Group("/role")
	{
		roleRouter.POST("/", controllers.CreateRole)
		roleRouter.GET("/", controllers.GetRole)
	}

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authentication())

		productRouter.POST("/create", controllers.CreateProduct)

		productRouter.GET("/all", middleware.ProductAuthorization(), controllers.GetAllProduct)

		productRouter.GET("/:productID", middleware.ProductAuthorization(), controllers.GetProductById)

		productRouter.PUT("/:productID", middleware.RoleMiddleWare(), middleware.ProductAuthorization(), controllers.UpdateProduct)

		productRouter.DELETE("/:productID", middleware.RoleMiddleWare(), middleware.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
