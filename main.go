package main

import (
	"fmt"
	"github.com/TalesPalma/web-app-golang/controllers"
	"net/http"
)

func main() {
	fmt.Println("Server running on port 8080...")
	fmt.Println("http://localhost:8080")

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/products", controllers.FormPage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
