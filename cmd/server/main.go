package main


import (
    "github.com/gin-gonic/gin"
    "github.com/MarianoLibre/go-web-capas/internal/products"
    "github.com/MarianoLibre/go-web-capas/cmd/server/handler"
)


func main() {
   repo := products.NewRepository()
   service := products.NewService(repo)
   p := handler.NewProduct(service)

   r := gin.Default()
   pr := r.Group("/products")
   pr.POST("/", p.Store())
   pr.GET("/", p.GetAll())
   r.Run()
}
