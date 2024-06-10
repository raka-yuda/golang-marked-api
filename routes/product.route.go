package routes

import (
	"golang-marked-api/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	productController controllers.ProductController
}

func NewRouteProduct(productController controllers.ProductController) ProductRoutes {
	return ProductRoutes{productController}
}

func (pr *ProductRoutes) ProductRoute(rg *gin.RouterGroup) {

	router := rg.Group("products")
	// router.POST("/", pr.productController.CreateProduct)
	router.GET("/", pr.productController.GetAllProducts)
	// router.PATCH("/:productId", pr.productController.UpdateProduct)
	// router.GET("/:productId", pr.productController.GetProductById)
	// router.DELETE("/:productId", pr.productController.DeleteProductById)
}
