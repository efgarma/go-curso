package main

import (
	"log"
	"net/http"
)

type Usuarios struct {
}

func main() {

	http.HandleFunc("/usuarios", usersFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func usersFunc(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodGet:
		handleGet(writer, request)
		return
	case http.MethodPost:
		handlePost(writer, request)
		return
	}
}

func handleGet(writer http.ResponseWriter, request *http.Request) {

}

func handlePost(writer http.ResponseWriter, request *http.Request) {

}
