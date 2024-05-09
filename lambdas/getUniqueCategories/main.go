package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"narach1988.mne/expenses-tracker/services/structs"
	"narach1988.mne/expenses-tracker/services/utils"
)

func getUniqueCategories() (events.APIGatewayProxyResponse, error) {
	dynamoDbClient := utils.GetDynamoClient()

	projExp := expression.NamesList(expression.Name("Category"))
	expr, err := expression.NewBuilder().WithProjection(projExp).Build()
	if err != nil {
		log.Printf("Couldn't build expressions for scan. Here's why: %v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprint("Error while retrieving data from DynamoDB", err),
		}, nil
	}

	response, err := dynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:            aws.String(os.Getenv("CATEGORY_TABLE_NAME")),
		ProjectionExpression: expr.Projection(),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprint("Error while retrieving data from DynamoDB", err),
		}, nil
	}

	var categories []structs.CategoryPick

	if err = attributevalue.UnmarshalListOfMaps(response.Items, &categories); err != nil {
		log.Fatalf("Error occured while umashalling, %v", err)
	}

	categories = slices.Compact(categories)

	jsonResponse, err := json.Marshal(categories)

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

func main() {
	lambda.Start(getUniqueCategories)
}
