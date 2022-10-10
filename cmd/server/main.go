package main

import (
	"log"

	"github.com/MarianoLibre/go-web-capas/cmd/server/handler"
	"github.com/MarianoLibre/go-web-capas/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("error: failed to load .env file")
	}
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNameAndPrice())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
