package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/Common/Movie"
	"net/http"
	"strconv"
)

var repository Movie.Repository

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(req.PathParameters["id"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "ID must be a number",
		}, nil
	}

	response, err := repository.GetByID(id)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
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
	repository = Movie.CreateNewMockMovies()

	lambda.Start(handler)
}
