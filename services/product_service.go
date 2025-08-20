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

func GetCart() []models.CartItem { return cart }

func AddToCart(item models.CartItem) bool {
	for i, product := range products {
		if product.ID == item.ProductID {
			if product.Quantity >= item.Quantity {
				products[i].Quantity -= item.Quantity
				cart = append(cart, models.CartItem{
					ProductID: product.ID,
					Name: product.Name,
					Description: product.Description,
					Price: product.Price,
					Quantity: item.Quantity,
				})
				return true
			}
			return false
		}
	}
	return false
}

func DeleteFromCart(productID int) bool {
	for i, item := range cart {
		if item.ProductID == productID {
			for j, product := range products {
				if product.ID == productID {
					products[j].Quantity += item.Quantity
					break
				}
			}
			cart = append(cart[:i], cart[i+1:]...)
			return true
		}
	}
	return false // не нашли товар в корзине
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