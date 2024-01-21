package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"magazin/helper"
	"magazin/models"
	"time"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetProduct(w, r)
	case "POST":
		CreateProduct(w, r)
	case "PUT":
		UpdateProduct(w, r)
	case "DELETE":
		DeleteProduct(w, r)
	}
}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.ProductModel
	json.NewDecoder(r.Body).Decode(&newProduct)

	var ProductData []models.ProductModel
	productByte, _ := os.ReadFile("db/product.json")
	json.Unmarshal(productByte, &ProductData)

	newProduct.ID = helper.MaxProduct(ProductData)
	newProduct.CreatedAt = time.Now()
	newProduct.UpdatedAt = time.Now()

	if newProduct.ProductType == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "ProductType cannot be an empty string!")
		return
	}
	if newProduct.Quantity == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "ProductType cannot be an empty string!")
		return
	}
	if newProduct.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "ProductType cannot be an empty string!")
		return
	}

	ProductData = append(ProductData, newProduct)

	res, _ := json.Marshal(ProductData)
	os.WriteFile("db/product.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Product created!")

	json.NewEncoder(w).Encode(newProduct)

}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct models.ProductModel
	json.NewDecoder(r.Body).Decode(&updateProduct)

	var ProductData []models.ProductModel
	productByte, _ := os.ReadFile("db/product.json")
	json.Unmarshal(productByte, &ProductData)

	var ProductFound bool
	for i := 0; i < len(ProductData); i++ {
		if ProductData[i].ID == updateProduct.ID {
			if updateProduct.ProductType == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "ProductType cannot be an empty string!")
				return
			}
			if updateProduct.Quantity == 0 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Product Quantity  cannot be an empty string!")
				return
			}
			if updateProduct.Price == 0 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Product Price cannot be an empty string!")
				return
			}
			ProductData[i].ProductType = updateProduct.ProductType
			ProductData[i].Quantity = updateProduct.Quantity
			ProductData[i].Price = updateProduct.Price
			ProductFound = true
			break
		}
	}
	if !ProductFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Product ID not found!")
		return
	}

	res, _ := json.Marshal(ProductData)
	os.WriteFile("db/product.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Product updated!")

	json.NewEncoder(w).Encode(updateProduct)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var deleteProduct models.ProductModel
	json.NewDecoder(r.Body).Decode(&deleteProduct)

	var ProductData []models.ProductModel
	productByte, _ := os.ReadFile("db/product.json")
	json.Unmarshal(productByte, &ProductData)

	var ProductFound bool
	for i := 0; i < len(ProductData); i++ {
		if ProductData[i].ID == deleteProduct.ID {
			ProductData = append(ProductData[:i], ProductData[i+1:]...)
			ProductFound = true
			break
		}
	}
	if !ProductFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Product ID not found!")
		return
	}

	res, _ := json.Marshal(ProductData)
	os.WriteFile("db/product.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Product deleted!")

	json.NewEncoder(w).Encode(deleteProduct)
}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	var getProduct models.ProductModel
	json.NewDecoder(r.Body).Decode(&getProduct)

	var ProductData []models.ProductModel
	productByte, _ := os.ReadFile("db/product.json")
	json.Unmarshal(productByte, &ProductData)

	if getProduct.ID > 0 {
		var ProductFound bool
		for i := 0; i < len(ProductData); i++ {
			if ProductData[i].ID == getProduct.ID {
				fmt.Fprintln(w, "________________________________________________")
				fmt.Fprintln(w, "Product ID:", ProductData[i].ID)
				fmt.Fprintln(w, "Product ProductType:", ProductData[i].ProductType)
				fmt.Fprintln(w, "Product Quantity:", ProductData[i].Quantity)
				fmt.Fprintln(w, "Product Price:", ProductData[i].Price)
				fmt.Fprintln(w, "Product CreatedAt:", ProductData[i].CreatedAt)
				fmt.Fprintln(w, "Product UpdatedAt:", ProductData[i].UpdatedAt)
				fmt.Fprintln(w, "________________________________________________")
				ProductFound = true
				break
			}
		}
		if !ProductFound {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Product ID not found!")
			return
		}
	} else if getProduct.ID == 0 {
		for i := 0; i < len(ProductData); i++ {
			fmt.Fprintln(w, "________________________________________________")
			fmt.Fprintln(w, "Product ID:", ProductData[i].ID)
			fmt.Fprintln(w, "Product ProductType:", ProductData[i].ProductType)
			fmt.Fprintln(w, "Product Quantity:", ProductData[i].Quantity)
			fmt.Fprintln(w, "Product Price:", ProductData[i].Price)
			fmt.Fprintln(w, "Product CreatedAt:", ProductData[i].CreatedAt)
			fmt.Fprintln(w, "Product UpdatedAt:", ProductData[i].UpdatedAt)
			fmt.Fprintln(w, "________________________________________________")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	} else if getProduct.ID < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Product ID not found!")
		return
	}

}
