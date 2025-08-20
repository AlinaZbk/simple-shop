package handlers

import (
	"net/http"
	"shop/models"
	"shop/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductsHandler(c *gin.Context) {
	products := services.GetProducts
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	product := services.GetProductByID(id)
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func AddProductHandler(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	added := services.AddProduct(newProduct)
	c.JSON(http.StatusOK, added)
}

func UpdateProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var updated models.Product
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := services.UpdateProduct(id, updated)
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	if services.DeleteProduct(id) {
		c.JSON(http.StatusOK, gin.H{"error": "Товар удален"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
	}
}
