package router

import (
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	{
		buildProductsRoutes(rg)
	}

}

func buildProductsRoutes(r *gin.RouterGroup) {
	db := []products.Product{
		{ID: "mock", SellerID: "1", Description: "generic product", Price: 123.55},
		{ID: "mock2", SellerID: "2", Description: "generic product2", Price: 1000.1},
	}
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	handler := products.NewHandler(service)

	prodRoute := r.Group("/products")
	{
		prodRoute.GET("", handler.GetProducts)
	}

}
