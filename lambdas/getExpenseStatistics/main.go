package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"narach1988.mne/expenses-tracker/services/structs"
	"narach1988.mne/expenses-tracker/services/utils"
)

func getExpenseStatistics() (events.APIGatewayProxyResponse, error) {
	dynamoDbClient := utils.GetDynamoClient()

	response, err := dynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprint("Error while retrieving data from DynamoDB", err),
		}, nil
	}

	var expenses []structs.Expense

	if err = attributevalue.UnmarshalListOfMaps(response.Items, &expenses); err != nil {
		log.Fatalf("Error occured while umashalling, %v", err)
	}

	expenseStats := groupExpenses(expenses)

	jsonResponse, err := json.Marshal(expenseStats)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(jsonResponse),
	}, nil
}

// Group expenses by categories
func groupExpenses(expenses []structs.Expense) map[string]float64 {
	expensesStats := make(map[string]float64)
	for _, expense := range expenses {
		_, exists := expensesStats[expense.Category]
		if exists {
			expensesStats[expense.Category] += expense.Amount
		} else {
			expensesStats[expense.Category] = expense.Amount
		}
	}
	return expensesStats
}

func main() {
	lambda.Start(getExpenseStatistics)
}
