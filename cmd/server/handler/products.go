package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MarianoLibre/go-web-capas/internal/products"
	"github.com/MarianoLibre/go-web-capas/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name      string  `json:"name"`
	Colour    string  `json:"colour"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	CreatedAt string  `json:"created-at"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid token"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid token"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		//fmt.Println("HANDLER>>> ", req)
		p, err := c.service.Store(req.Name, req.Colour, req.Code, req.CreatedAt, req.Stock, req.Price, req.Published)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid token"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid Id"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product 'name' is required"))
			return
		}
		if req.Colour == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product 'colour' is required"))
			return
		}
		if req.CreatedAt == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "'creation date' is required"))
			return
		}
		if req.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product 'code' is required"))
			return
		}
		if req.Stock == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Product 'stock' is required"))
			return
		}
		if req.Price == 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Product 'price' is required"))
			return
		}
		p, err := c.service.Update(int(id), req.Name, req.Colour, req.Code, req.CreatedAt, req.Stock, req.Price, req.Published)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid token"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid Id"))
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "'Name' is required"))
			return
		}
		if req.Price == 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, "'Price' is required"))
			return
		}
		p, err := c.service.UpdateNameAndPrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Invalid token"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid id"))
			return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("Product '%d' has been deleted.", id), ""))
	}
}
