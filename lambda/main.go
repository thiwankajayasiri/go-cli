package main

import (
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func main() {
	lambda.Start(HandleRequest)
}


func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello World")

	response := events.APIGatewayProxyResponse{Body: "Hello World", StatusCode: 200}
	return response, nil
}