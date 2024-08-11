package main

import (
	"fmt"
	"net/http"

	"github.com/TalesPalma/web-app-golang/controllers"
	"github.com/TalesPalma/web-app-golang/db"
)

func main() {

	db := db.ConnectDatabase()
	db.Ping()

	fmt.Println("Server running on port 8080...")
	fmt.Println("http://localhost:8080")

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/products", controllers.FormPage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
