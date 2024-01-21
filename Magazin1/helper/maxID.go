package helper

import "magazin/models"

func MaxCategory(Array []models.CategoryModel) int {
	var maxID = 0
	for i := 0; i < len(Array); i++ {
		if maxID < Array[i].ID {
			maxID = Array[i].ID
		}
	}
	return maxID + 1
}
func MaxProduct(Array []models.ProductModel) int {
	var maxID = 0
	for i := 0; i < len(Array); i++ {
		if maxID < Array[i].ID {
			maxID = Array[i].ID
		}
	}
	return maxID + 1
}
