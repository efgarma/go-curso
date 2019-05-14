package main

import (
	"testing"
)

func TestAddTexto(t *testing.T) {

	peli := Peliculas{
		Id:     1,
		Nombre: "Avengers End Game",
	}

	peli.addTexto()

	if peli.Nombre != "Avengers End Game curso go" {
		t.Fail()
	}
}
