package main

import "shop/models"

var products = []models.Product{
	{ID: 1, Name: "Ноутбук", Description: "Macbook air m4 - подходит для работы и повседневных задач", Price: 78500, Quantity: 25},
	{ID: 2, Name: "Чистый код", Description: "Книга для настоящих разработчиков", Price: 3250, Quantity: 8},
	{ID: 3, Name: "Ridmi Buds 5", Description: "Беспроводные наушники с активным шумоподавлением", Price: 3800, Quantity: 20},
}

var cart = []models.CartItem{}

func getNextID() int {
	maxID := 0
	for _, product := range products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}
