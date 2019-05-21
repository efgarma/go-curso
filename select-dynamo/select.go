package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("AKIAJXULSU247S3IHMDQ", "jYzvb/AJDmNAcjxNW+w3EcttJjL/hbu2ShbjBNzx", ""),
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

	log.Println(*(response.Items[0]["id"].N))
	/*for _, v := range response.Items {
		log.Printf("%s", v)
	}*/

}
