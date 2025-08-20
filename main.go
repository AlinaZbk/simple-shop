package main

import (
	"fmt"
	"shop/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // объект роутера Gin

	router.Static("/static", "./static") // подключаем фронтенд

	// Главная страница
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Товары
	router.GET("/products", handlers.GetProductsHandler)
	router.GET("/products/:id", handlers.GetProductByIDHandler)
	router.POST("/products", handlers.AddProductHandler)
	router.PUT("/products/:id", handlers.UpdateProductHandler)
	router.DELETE("/products/:id", handlers.DeleteProductHandler)

	//Корзина
	router.GET("/cart/total", handlers.GetCartTotalHandler)
	router.GET("/cart", handlers.GetCartHandler)
	router.POST("/cart", handlers.AddToCartHandler)
	router.DELETE("/cart/:id", handlers.DeleteFromCartHandler)

	fmt.Println("Server running at http://localhost:8080")

	router.Run(":8080") // запуск сервера
}
