package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/TalesPalma/web-app-golang/models"
	"github.com/TalesPalma/web-app-golang/services"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "Index", services.IndexService())
	if err != nil {
		fmt.Println(err)
	}
}

func FormPage(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "Form", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func SubmitPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("productName")
		desc := r.FormValue("productDescription")
		price := r.FormValue("productPrice")
		qtty := r.FormValue("productQuantity")

		fmt.Println(name, desc, price, qtty)
		convertPrice, _ := strconv.ParseFloat(price, 64)
		convertQtty, _ := strconv.Atoi(qtty)

		newProd := models.Product{
			Name:  name,
			Desc:  desc,
			Price: convertPrice,
			Qtty:  convertQtty,
		}
		services.FormService(newProd)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DetailsPage(w http.ResponseWriter, r *http.Request) {
	idFromUrl := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idFromUrl)
	temp.ExecuteTemplate(w, "Details", services.DetailsSerive(id))
}
