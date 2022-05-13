package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/common"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/common/factory"
	customLambda "github.com/gabrielsscti/gabriel/aws-lambdas-practice/common/lambda"
	"net/http"
)

type FindAll struct {
	customLambda.Base
}

func (f FindAll) handle() (events.APIGatewayProxyResponse, error) {
	response, err := f.Repository.GetAll()
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
		Body: string(response),
	}, nil
}

func main() {
	findAll := FindAll{
		Base: customLambda.Base{
			Repository: factory.CreateNewRepository(),
		},
	}

	lambda.Start(findAll.handle)
}
