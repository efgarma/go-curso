package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Person{})
	if err != nil {
		fmt.Println(err)
	}
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		fmt.Println(err)
	}
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	err := json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		fmt.Println(err)
	}
	people = append(people, person)
	err = json.NewEncoder(w).Encode(people)
	if err != nil {
		fmt.Println(err)
	}
}

// Metodo para eliminar
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	router := mux.NewRouter()
	people = append(people,
		Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}},
		Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"},
		Person{ID: "3", Firstname: "Jose", Lastname: "Perez"})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
