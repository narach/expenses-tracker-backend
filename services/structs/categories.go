package structs

type CategoryPick struct {
	Category string `dynamodbav:"Category"`
}

type CategoryItem struct {
	Category string  `dynamodbav:"Category"`
	Month    string  `dynamodbav:"Month"`
	Amount   float32 `dynamodbav:"Amount"`
}
