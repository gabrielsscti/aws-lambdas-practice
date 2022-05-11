package main

import "github.com/aws/aws-lambda-go/lambda"

func handler() (string, error) {
	return "Hello Serveless!", nil
}

func main() {
	lambda.Start(handler)
}
