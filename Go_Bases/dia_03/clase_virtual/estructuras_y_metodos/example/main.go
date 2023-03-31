package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func main() {

	p:= Persona{
		Nombre:	"Gabriel",
		Edad: 	24,
	}

	p.Descripcion()
	fmt.Println(p.Nombre)

}

func (p *Persona) Descripcion() {
	fmt.Printf("%s tiene %d años de edad", p.Nombre, p.Edad)
	p.Nombre = "Maria"
}
