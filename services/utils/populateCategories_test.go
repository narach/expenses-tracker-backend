package utils

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
	"narach1988.mne/expenses-tracker/services/structs"
)

func TestPopulateCategories(t * testing.T) {
	expenseData := []structs.Expense {
		{
			ID: "1",
			ExpensePayload: structs.ExpensePayload {
				Title: "Products", Amount: 50.0, Category: "Food", ExpenseDate: "05-01-2024",ExpenseDateLong: 1714521600000,
			},
		},
		{
			ID: "1",
			ExpensePayload: structs.ExpensePayload {
				Title: "Petrol", Amount: 75.0, Category: "Vehicles", ExpenseDate: "05-01-2024",ExpenseDateLong: 1714521600000,
			},
		},
	}

	categoriesWanted := []structs.Category {
		{
			CategoryKey: structs.CategoryKey{
				Name: "Food",
				Month: "05-2024",
			},
			Amount: decimal.NewFromFloat(50.0),
		},
		{
			CategoryKey: structs.CategoryKey{
				Name: "Vehicles",
				Month: "05-2024",
			},
			Amount: decimal.NewFromFloat(75.0),
		},
	}

	categoriesActual := PopulateCategories(expenseData)
	if (!reflect.DeepEqual(categoriesWanted, categoriesActual)) {
		t.Errorf("Expected %v, but got %v", categoriesWanted, categoriesActual)
	}
}
