package services

import (
	"github.com/TalesPalma/web-app-golang/db"
	"github.com/TalesPalma/web-app-golang/models"
)

func IndexService() []models.Product {

	db := db.ConnectDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM product")

	if err != nil {
		panic(err)
	}

	var products []models.Product
	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Desc, &p.Price, &p.Qtty)

		if err != nil {
			panic(err)
		}

		products = append(products, p)
	}
	//Get all products
	return products
}
