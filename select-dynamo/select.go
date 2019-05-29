package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type tableDemo struct {
	Id               int
	Nombre           string
	NombreSecundario string `dynamodbav:"nombre_secundario"`
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("id", "secret", ""),
		Region:      aws.String("us-west-2"),
	})

	if err != nil {
		log.Fatal("ERROR AL CONECTARSE A AWS ", err)
	}

	sessDB := dynamodb.New(sess)

	request := &dynamodb.ScanInput{
		TableName: aws.String("table_demo"),
	}

	response, err := sessDB.Scan(request)
	if err != nil {
		log.Println("Error al obtener tablas ", err)

	}

	for _, v := range response.Items {
		log.Printf("%s", v)
		// Parseo
		item := tableDemo{}
		err = dynamodbattribute.UnmarshalMap(v, &item)

		if err != nil {
			log.Fatal("ERROR AL PARSEAR ", err)
		}
		log.Printf("Valores %v ", item)
	}

}
