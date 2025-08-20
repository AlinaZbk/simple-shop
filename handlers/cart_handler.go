package handlers

import (
	"net/http"
	"shop/models"
	"shop/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCartHandler(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetCart())
}

func AddToCartHandler(c *gin.Context) {
	var item models.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if services.AddToCart(item) {
		c.JSON(http.StatusOK, gin.H{"message": "Товар добавлен в корзину"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недостаточно товара на складе"})
	}
}

func DeleteFromCartHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
	}

	if services.DeleteFromCart(id) {
		c.JSON(http.StatusOK, gin.H{"message": "Товар удален из корзины"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден в корзине"})
	}
}