package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Empleado struct {
	Id     int
	Nombre string
}

var empleados []Empleado

func main() {

	empleados = append(empleados,
		Empleado{
			Id:     1,
			Nombre: "Garma",
		},
		Empleado{
			Id:     2,
			Nombre: "Chacarron",
		})

	http.HandleFunc("/user", userFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar server", err)
	}
}

func userFunc(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getHandler(writer)
	case http.MethodPost:
		postHandler(request)
		break
	default:
		writer.WriteHeader(http.StatusBadRequest)
		break
	}
}

func postHandler(request *http.Request) {
	var empleadoRequest Empleado
	err := json.NewDecoder(request.Body).Decode(&empleadoRequest)
	if err != nil {
		log.Println("Error en parseo ", err)
	}

	empleados = append(empleados, empleadoRequest)
}

func getHandler(writer http.ResponseWriter) {

	err := json.NewEncoder(writer).Encode(empleados)
	if err != nil {
		log.Println("Error en parseo ", err)
	}

}
