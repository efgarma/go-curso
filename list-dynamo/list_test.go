package main

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var request events.APIGatewayProxyRequest

	response, err := Handler(request)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal("Error en test")
	}

}
