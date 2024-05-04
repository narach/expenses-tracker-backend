package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
				Name: "Vehicles",
				Month: "05-2024",
			},
			Amount: decimal.NewFromFloat(75.0),
		},
		{
			CategoryKey: structs.CategoryKey{
				Name: "Food",
				Month: "05-2024",
			},
			Amount: decimal.NewFromFloat(50.0),
		},
	}

	categoriesActual := PopulateCategories(expenseData)
	if (!reflect.DeepEqual(categoriesWanted, categoriesActual)) {
		t.Errorf("Expected %v, but got %v", categoriesWanted, categoriesActual)
	}
}

func TestCategoriesByMonth(t * testing.T) {
	// Read Expenses data from file
	expensesBinary := readFile("./testData/inputExpenses.json")
	
	var expenses []structs.Expense
	json.Unmarshal(expensesBinary, &expenses)

	expCategoriesBinary := readFile("./testData/outputCategories.json")
	var expCategoreis []structs.Category
	json.Unmarshal(expCategoriesBinary, &expCategoreis)

	actualCategories := PopulateCategories(expenses)
	for index, expectedCategory := range expCategoreis {
		actualCategory := actualCategories[index]
		isEquals := expectedCategory.CategoryKey == actualCategory.CategoryKey && expectedCategory.Amount.Equal(actualCategory.Amount)
		if !isEquals {
			t.Errorf("Expected %v, but got %v", expectedCategory, actualCategory)
		}
	}
}

func readFile(filename string) ([]byte) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + filename)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	return byteValue
}