package controllers

import (
	"fmt"
	"net/http"
	"text/template"

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
