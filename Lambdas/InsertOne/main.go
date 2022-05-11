package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/Common/Movie"
	"net/http"
)

var repository Movie.Repository

func handle(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := repository.InsertMovie([]byte(req.Body))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid payload",
		}, nil
	}

	movies, err := repository.GetAll()
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
		Body: string(movies),
	}, nil
}

func main() {
	repository = Movie.CreateNewMockMovies()

	lambda.Start(handle)
}
