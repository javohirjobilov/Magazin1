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

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetCategory(w, r)
	case "POST":
		CreateCategory(w, r)
	case "PUT":
		UpdateCategory(w, r)
	case "DELETE":
		DeleteCategory(w, r)

	}
}
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.CategoryModel
	json.NewDecoder(r.Body).Decode(&newCategory)

	var CategoryData []models.CategoryModel
	categoryByte, _ := os.ReadFile("db/category.json")
	json.Unmarshal(categoryByte, &CategoryData)

	newCategory.ID = helper.MaxCategory(CategoryData)
	newCategory.CreatedAt = time.Now()
	newCategory.UpdatedAt = time.Now()

	if newCategory.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Type cannot be an empty string!")
		return
	}

	CategoryData = append(CategoryData, newCategory)

	res, _ := json.Marshal(CategoryData)
	os.WriteFile("db/category.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Category created!")

	json.NewEncoder(w).Encode(newCategory)
}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var updateCategory models.CategoryModel
	json.NewDecoder(r.Body).Decode(&updateCategory)

	var CategoryData []models.CategoryModel
	categoryByte, _ := os.ReadFile("db/category.json")
	json.Unmarshal(categoryByte, &CategoryData)

	var CategoryFound bool
	for i := 0; i < len(CategoryData); i++ {
		if CategoryData[i].ID == updateCategory.ID {
			if updateCategory.Type == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Type cannot be an empty string!")
				return
			}
			CategoryData[i].Type = updateCategory.Type
			CategoryData[i].UpdatedAt = updateCategory.UpdatedAt
			CategoryFound = true
			break
		}
	}
	if !CategoryFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Category's ID not found!")
		return
	}

	res, _ := json.Marshal(CategoryData)
	os.WriteFile("db/category.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Category updated!")

	json.NewEncoder(w).Encode(updateCategory)
}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var deleteteCategory models.CategoryModel
	json.NewDecoder(r.Body).Decode(&deleteteCategory)

	var CategoryData []models.CategoryModel
	categoryByte, _ := os.ReadFile("db/category.json")
	json.Unmarshal(categoryByte, &CategoryData)

	var CategoryFound bool
	for i := 0; i < len(CategoryData); i++ {
		if CategoryData[i].ID == deleteteCategory.ID {
			CategoryData = append(CategoryData[:i], CategoryData[i+1:]...)
			CategoryFound = true
			break
		}
	}
	if !CategoryFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Category's ID not found!")
		return
	}

	res, _ := json.Marshal(CategoryData)
	os.WriteFile("db/category.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Category deleted!")

	json.NewEncoder(w).Encode(deleteteCategory)
}
func GetCategory(w http.ResponseWriter, r *http.Request) {

	var getCategory models.CategoryModel
	json.NewDecoder(r.Body).Decode(&getCategory)

	var CategoryData []models.CategoryModel
	categoryByte, _ := os.ReadFile("db/category.json")
	json.Unmarshal(categoryByte, &CategoryData)

	if getCategory.ID > 0 {
		var CategoryFound bool
		for i := 0; i < len(CategoryData); i++ {
			if CategoryData[i].ID == getCategory.ID {
				fmt.Fprintln(w, "Category ID:", CategoryData[i].ID)
				fmt.Fprintln(w, "Category Type:", CategoryData[i].Type)
				fmt.Fprintln(w, "Category CreatedAt:", CategoryData[i].CreatedAt)
				fmt.Fprintln(w, "Category UpdatedAt:", CategoryData[i].UpdatedAt)
				fmt.Fprintln(w, "________________________________________________")
				for j := 0; j < len(CategoryData[i].Products); j++ {
					fmt.Fprintln(w, "  Product's ID:", CategoryData[i].Products[j].ID)
					fmt.Fprintln(w, "  Product's ProductType:", CategoryData[i].Products[j].ProductType)
					fmt.Fprintln(w, "  Product's Quantity:", CategoryData[i].Products[j].Quantity)
					fmt.Fprintln(w, "  Product's Price:", CategoryData[i].Products[j].Price)
					fmt.Fprintln(w, "  Product's CreatedAt:", CategoryData[i].Products[j].CreatedAt)
					fmt.Fprintln(w, "  Product's UpdatedAt:", CategoryData[i].Products[j].UpdatedAt)
					fmt.Fprintln(w, "  ________________________________________________")
				}
				CategoryFound = true
				break
			}
		}
		if !CategoryFound {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Category's ID not found!")
			return
		}
	} else if getCategory.ID == 0 {
		fmt.Fprintln(w, "________________________________________________")
		for i := 0; i < len(CategoryData); i++ {
			fmt.Fprintln(w, "Category ID:", CategoryData[i].ID)
			fmt.Fprintln(w, "Category Type:", CategoryData[i].Type)
			fmt.Fprintln(w, "Category CreatedAt:", CategoryData[i].CreatedAt)
			fmt.Fprintln(w, "Category UpdatedAt:", CategoryData[i].UpdatedAt)
			fmt.Fprintln(w, "________________________________________________")
			for j := 0; j < len(CategoryData[i].Products); j++ {
				fmt.Fprintln(w, "  Product's ID:", CategoryData[i].Products[j].ID)
				fmt.Fprintln(w, "  Product's ProductType:", CategoryData[i].Products[j].ProductType)
				fmt.Fprintln(w, "  Product's Quantity:", CategoryData[i].Products[j].Quantity)
				fmt.Fprintln(w, "  Product's Price:", CategoryData[i].Products[j].Price)
				fmt.Fprintln(w, "  Product's CreatedAt:", CategoryData[i].Products[j].CreatedAt)
				fmt.Fprintln(w, "  Product's UpdatedAt:", CategoryData[i].Products[j].UpdatedAt)
				fmt.Fprintln(w, "  ________________________________________________")
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	} else if getCategory.ID < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Category's ID not found!")
		return
	}

}
