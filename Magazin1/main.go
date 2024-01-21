package main

import (
	"fmt"
	"net/http"
	"magazin/handlers"
)

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/category", handlers.CategoryHandler)
	http.HandleFunc("/product", handlers.ProductHandler)
	http.HandleFunc("/managecategory", handlers.ManageCategory)

	http.ListenAndServe(":8080", nil)
}
