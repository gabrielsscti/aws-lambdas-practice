package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gabrielsscti/aws-lambdas-practice-common/movie"
	"io/ioutil"
	"log"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(opt *config.LoadOptions) error {
		opt.Region = "sa-east-1"
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	movies, err := readMovies("C:\\repos\\aws-lambdas-practice\\dynamo-saver\\movies.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		fmt.Println("Inserting " + movie.Name + " " + movie.ID)
		err = insertMovie(cfg, movie)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func insertMovie(cfg aws.Config, movie movie.Movie) error {
	svc := dynamodb.NewFromConfig(cfg)

	val, err := attributevalue.MarshalMap(movie)
	if err != nil {
		return err
	}

	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("movies"),
		Item:      val,
	})
	if err != nil {
		return err
	}

	return nil
}

func readMovies(fileName string) ([]movie.Movie, error) {
	movies := make([]movie.Movie, 0)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
