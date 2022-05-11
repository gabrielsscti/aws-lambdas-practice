package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielsscti/gabriel/aws-lambdas-practice/Common/Movie"
	"net/http"
)

var movieRepository Movie.Repository

func handler() (events.APIGatewayProxyResponse, error) {
	response, err := movieRepository.GetAll()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
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
	movieRepository = Movie.CreateNewMockMovies()

	lambda.Start(handler)
}
