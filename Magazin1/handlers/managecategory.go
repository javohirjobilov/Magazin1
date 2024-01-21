package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"magazin/models"
	"time"
)

func ManageCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	case "POST":
		AddProduct(w, r)
	case "PUT":
	case "DELETE":
		RemoweProduct(w, r)
	}
}
func AddProduct(w http.ResponseWriter, r *http.Request) {
	var addProduct models.ManageProduct
	json.NewDecoder(r.Body).Decode(&addProduct)

	var CategoryData []models.CategoryModel
	categoryByte, _ := os.ReadFile("db/category.json")
	json.Unmarshal(categoryByte, &CategoryData)

	var ProductData []models.ProductModel
	productByte, _ := os.ReadFile("db/product.json")
	json.Unmarshal(productByte, &ProductData)

	var FoundCategory bool
	var FoundProduct bool

	for j := 0; j < len(CategoryData); j++ {
		if CategoryData[j].ID == addProduct.CategoryID {

			for i := 0; i < len(ProductData); i++ {
				if ProductData[i].ID == addProduct.ProductID {
					ProductData[i].UpdatedAt = time.Now()
					CategoryData[j].Products = append(CategoryData[j].Products, ProductData[i])
					FoundProduct = true
					break
				}
			}
			if !FoundProduct {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Product's ID not found!")
				return
			}
			fmt.Fprintln(w, "")
			FoundCategory = true
			break
		}
	}
	if !FoundCategory {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Category's ID not found!")
		return
	}

	res, _ := json.Marshal(CategoryData)
	os.WriteFile("db/category.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintln(w, "Product has been added to the category")

	json.NewEncoder(w).Encode(addProduct)
}
func RemoweProduct(w http.ResponseWriter, r *http.Request) {
	var remoweProduct models.ManageProduct
	json.NewDecoder(r.Body).Decode(&remoweProduct)

	var CategoryData []models.CategoryModel
	categoryByte, _ := os.ReadFile("db/category.json")
	json.Unmarshal(categoryByte, &CategoryData)

	var ProductData []models.ProductModel
	productByte, _ := os.ReadFile("db/product.json")
	json.Unmarshal(productByte, &ProductData)

	var FoundCategory bool
	var FoundProduct bool

	for j := 0; j < len(CategoryData); j++ {
		if CategoryData[j].ID == remoweProduct.CategoryID {

			for i := 0; i < len(CategoryData[i].Products); i++ {
				if CategoryData[i].Products[j].ID == remoweProduct.ProductID {
					CategoryData[j].Products = append(CategoryData[j].Products[:i], CategoryData[j].Products[i+1:]...)
					FoundProduct = true
					break
				}
			}
			if !FoundProduct {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Product's ID not found!")
				return
			}
			FoundCategory = true
			break
		}
	}
	if !FoundCategory {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Category's ID not found!")
		return
	}

	res, _ := json.Marshal(CategoryData)
	os.WriteFile("db/category.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintln(w, "Product removed from category")

	json.NewEncoder(w).Encode(remoweProduct)
}
