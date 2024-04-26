package structs

type ExpensePayload struct {
	Title       string  `dynamodbav:"title"`
	Amount      float64 `dynamodbav:"amount"`
	Category    string  `dynamodbav:"category"`
	ExpenseDate string  `dynamodbav:"expense_date"`
}

type Expense struct {
	ID string `dynamodbav:"id"`
	ExpensePayload
}
