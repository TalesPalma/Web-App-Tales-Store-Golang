package services

import (
	"log"

	"github.com/TalesPalma/web-app-golang/db"
	"github.com/TalesPalma/web-app-golang/models"
)

func IndexService() []models.Product {

	db := db.ConnectDatabase()

	rows, err := db.Query("SELECT * FROM product")

	if err != nil {
		panic(err)
	}

	//Get all products
	var products []models.Product
	for rows.Next() {

		p := models.Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Desc, &p.Price, &p.Qtty)

		if err != nil {
			panic(err)
		}

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func FormService(newProd models.Product) {
	db := db.ConnectDatabase()

	name := newProd.Name
	desc := newProd.Desc
	price := newProd.Price
	qtty := newProd.Qtty

	result, err := db.Exec(
		`INSERT INTO product (nameproduct, description, price, qtty)
		VALUES ($1,$2,$3,$4)`,
		name, desc, price, qtty,
	)

	if err != nil {
		log.Fatalln("Error on insert: ", err)
	}

	log.Println(result)

	defer db.Close()

}

func DetailsService(id int) models.Product {
	db := db.ConnectDatabase()
	var p models.Product

	row := db.QueryRow("Select * from product where id = $1", id)
	if err := row.Scan(&p.ID, &p.Name, &p.Desc, &p.Price, &p.Qtty); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return p
}

func DeleteService(id int) {
	db := db.ConnectDatabase()
	db.Exec("Delete from product where id=$1", id)
	defer db.Close()
}
