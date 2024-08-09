package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/TalesPalma/web-app-golang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	var products = []models.Product{
		{ID: 1, Name: "Product 1", Desc: "Description 1", Price: 100.0, Qtty: 10},
		{ID: 2, Name: "Product 2", Desc: "Description 2", Price: 200.0, Qtty: 20},
		{ID: 3, Name: "Product 3", Desc: "Description 3", Price: 300.0, Qtty: 30},
		{ID: 3, Name: "Product 3", Desc: "Description 3", Price: 1.99, Qtty: 30},
	}
	err := temp.ExecuteTemplate(w, "Index", products)

	if err != nil {
		fmt.Println(err)
	}

}
