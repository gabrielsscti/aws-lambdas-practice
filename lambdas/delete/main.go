package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/common"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/common/factory"
	customLambda "github.com/gabrielsscti/gabriel/aws-lambdas-practice/common/lambda"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/common/movie"
	"log"
	"net/http"
)

type Delete struct {
	customLambda.Base
}

func (d Delete) handle(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movieObj movie.Movie
	err := json.Unmarshal([]byte(req.Body), &movieObj)
	if err != nil {
		log.Println("Error while unmarshaling request body: " + err.Error())
		log.Println(req.Body)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid payload",
		}, err
	}

	err = d.Repository.Delete(movieObj.ID)
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
	deleteStruct := Delete{
		Base: customLambda.Base{
			Repository: factory.CreateNewRepository(),
		},
	}

	lambda.Start(deleteStruct.handle)
}
