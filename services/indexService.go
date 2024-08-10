package services

import "github.com/TalesPalma/web-app-golang/models"

func IndexService() []models.Product {
	var products = []models.Product{
		{ID: 1, Name: "Product 1", Desc: "Description 1", Price: 100.0, Qtty: 10},
		{ID: 2, Name: "Product 2", Desc: "Description 2", Price: 200.0, Qtty: 20},
		{ID: 3, Name: "Product 3", Desc: "Description 3", Price: 300.0, Qtty: 30},
		{ID: 3, Name: "Product 3", Desc: "Description 3", Price: 1.99, Qtty: 30},
	}

	return products
}
