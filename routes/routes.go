package routes

import (
	"productfc/cmd/product/handler"
	"productfc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, productHandler handler.ProductHandler) {
	router.Use(middleware.RequestLogger())
	router.POST("/v1/product_category", productHandler.ProductCategoryManagement)
}
