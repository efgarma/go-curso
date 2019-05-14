package main

import (
	"fmt"
)

type Peliculas struct {
	Id     int
	Nombre string
}

func (p *Peliculas) addTexto() {
	p.Nombre += " curso go"
}

func main() {
	pelicula := Peliculas{
		Id:     1,
		Nombre: "Gol",
	}

	fmt.Println(pelicula)
	pelicula.addTexto()
	fmt.Println(pelicula)

}
