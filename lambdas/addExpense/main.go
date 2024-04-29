package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"narach1988.mne/expenses-tracker/services/structs"
	"narach1988.mne/expenses-tracker/services/utils"
)

func addExpense(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newExpense structs.ExpensePayload

	if err := json.Unmarshal([]byte(req.Body), &newExpense); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	// Set default values for Category and Expense date if not specified
	// TODO Refactor to use default struct field values if possible
	if newExpense.Category == "" {
		newExpense.Category = "Other"
	}
	if newExpense.ExpenseDate == "" {
		currentTime := time.Now()
		newExpense.ExpenseDate = currentTime.Format("02-01-2006")
	}

	dynamoClient := utils.GetDynamoClient()

	result, err := dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item: map[string]types.AttributeValue{
			"ID":           &types.AttributeValueMemberS{Value: utils.GenerateUUID()},
			"Title":        &types.AttributeValueMemberS{Value: newExpense.Title},
			"Amount":       &types.AttributeValueMemberN{Value: fmt.Sprint(newExpense.Amount)},
			"Category":     &types.AttributeValueMemberS{Value: newExpense.Category},
			"expense_date": &types.AttributeValueMemberS{Value: newExpense.ExpenseDate},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	resultJson, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "OPTIONS,POST,GET",
			"Content-Type":                 "application/json",
		},
		Body: string(resultJson),
	}, nil
}

func main() {
	lambda.Start(addExpense)
}
