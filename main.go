package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Secret string `json:"secret"`
}

// type Response struct {
// 	Code string `json:"code"`
// }

func handleRequest(ctx context.Context, event Request) error {
	log.Println("Request: %v", event)

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
