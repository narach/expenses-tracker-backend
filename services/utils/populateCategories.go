package utils

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"time"

	"narach1988.mne/expenses-tracker/services/structs"
)

func PopulateCategories(expenses []structs.Expense) ([]structs.Category) {
	categoriesMap := make(map[structs.CategoryKey]float64)
	for _, expense := range expenses {
		category := expense.Category
		unixTime := time.Unix(expense.ExpenseDateLong/1000, 0)
		year := unixTime.Year()
		month := unixTime.Month()
		categoryMonth := fmt.Sprintf("%.2d-%.4d", month, year)
		categoryKey := structs.CategoryKey {
			Name: category,
			Month: categoryMonth,
		}
		_, alreadyAdded := categoriesMap[categoryKey]
		if (alreadyAdded) {
			categoriesMap[categoryKey] += expense.Amount
		} else {
			categoriesMap[categoryKey] = expense.Amount
		}
	}
	categories := []structs.Category{}
	for key, value := range categoriesMap {
		categoryItem := structs.Category {
			Amount: toFixed(value, 2),
			CategoryKey: key,
		}
		categories = append(categories, categoryItem)
	}

	// Sorting categories by amount
	slices.SortFunc(categories, func(i, j structs.Category) int {
		return cmp.Compare(j.Amount, i.Amount)
	})
	return categories
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}
