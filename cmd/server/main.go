package main

import (
	"github.com/MarianoLibre/go-web-capas/cmd/server/handler"
	"github.com/MarianoLibre/go-web-capas/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
    pr.PUT("/:id", p.Update())
	r.Run()
}
