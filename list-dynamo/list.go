package main

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"net/http"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("DATOS REQUEST > ", request)

	list, err := getListado()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	body, err := json.Marshal(list)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, errors.New("ERROR AL PARSEAR RESPUESTA")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil

}

func getListado() ([]*string, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("id", "secret", ""),
		Region:      aws.String("us-west-2"),
	})

	if err != nil {
		log.Fatal("ERROR AL CONECTARSE A AWS ", err)
		return nil, err
	}

	sessDB := dynamodb.New(sess)

	request := &dynamodb.ListTablesInput{}

	response, err := sessDB.ListTables(request)
	if err != nil {
		log.Println("Error al obtener tablas ", err)
		return nil, err
	}

	if len(response.TableNames) <= 0 {
		log.Println("NO SE HAN ENCONTRADO TABLAS")
		return nil, errors.New("NO SE HAN ENCONTRADO TABLAS")
	}

	for _, v := range response.TableNames {
		log.Printf("%s", *v)
	}

	return response.TableNames, nil
}

func main() {
	lambda.Start(Handler)
}
