package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielsscti/aws-lambdas-practice-common"
	"github.com/gabrielsscti/aws-lambdas-practice-common/factory"
	customLambda "github.com/gabrielsscti/aws-lambdas-practice-common/lambda"
	"net/http"
)

type Insertion struct {
	customLambda.Base
}

func (i Insertion) handle(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := i.Repository.Insert([]byte(req.Body))
	if err != nil {
		if respErr, ok := err.(*common.ResponseError); ok {
			return events.APIGatewayProxyResponse{
				StatusCode: respErr.StatusCode,
				Body:       respErr.Body,
			}, err
		} else {
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	insertion := Insertion{
		Base: customLambda.Base{
			Repository: factory.CreateNewRepository(),
		},
	}

	lambda.Start(insertion.handle)
}
