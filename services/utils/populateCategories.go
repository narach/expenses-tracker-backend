package utils

import (
	"fmt"
	"slices"
	"time"

	"github.com/shopspring/decimal"
	"narach1988.mne/expenses-tracker/services/structs"
)

func PopulateCategories(expenses []structs.Expense) ([]structs.Category) {
	categoriesMap := make(map[structs.CategoryKey]decimal.Decimal)
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
			categoriesMap[categoryKey].Add(decimal.NewFromFloat(expense.Amount))
		} else {
			categoriesMap[categoryKey] = decimal.NewFromFloat(expense.Amount)
		}
	}
	categories := []structs.Category{}
	for key, value := range categoriesMap {
		categoryItem := structs.Category {
			Amount: value,
			CategoryKey: key,
		}
		categories = append(categories, categoryItem)
	}

	// Sorting categories by amount
	slices.SortFunc(categories, func(i, j structs.Category) int {
		return i.Amount.Compare(j.Amount)
	})
	return categories
}
