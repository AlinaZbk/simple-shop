package services

import "shop/models"

func GetProducts() []models.Product { return products }

func GetProductByID(id int) *models.Product {
	for i := range products {
		if products[i].ID == id {
			return &products[i]
		}
	}
	return nil // товар не найден
}

func AddProduct(product models.Product) models.Product {
	product.ID = getNextID()
	products = append(products, product)
	return product
}

func UpdateProduct(id int, updated models.Product) *models.Product {
	for i := range products {
		if products[i].ID == id {
			updated.ID = id
			products[i] = updated
			return &products[i]
		}
	}
	return nil // товар не найден
}

func DeleteProduct(id int) bool {
	for i := range products {
		if products[i].ID == id {
			products = append(products[:i], products[i+1:]...)
			return true // успешно удалили
		}
	}
	return false // товар не найден
}


func getNextID() int {
	maxID := 0
	for _, product := range products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}