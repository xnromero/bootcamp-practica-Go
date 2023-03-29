package main

import (
	"errors"
	"fmt"
)

type Persona struct {
	Legajo    int
	Nombre    string
	DNI       int
	Telefono  int
	Domicilio string
}

var Clientes = []Persona{}

func main() {

	persona1 := Persona{
		Legajo:    1234,
		Nombre:    "Juan",
		DNI:       48374857,
		Telefono:  473847330,
		Domicilio: "Av siempre viva 123",
	}

	persona2 := Persona{
		Legajo:    12345,
		Nombre:    "José",
		DNI:       46374859,
		Telefono:  938474632,
		Domicilio: "Calle falsa 123",
	}

	RegistrarCliente(persona1)
	RegistrarCliente(persona2)

	fmt.Println("Ejecución finalizada")
	fmt.Println(Clientes)

}
func RegistrarCliente(person Persona) {

	defer func() {

		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}

	}()

	for i := range Clientes {
		if Clientes[i].Legajo == person.Legajo {
			panic("El cliente ya existe")
		}
	}

	_, err := ValidarDatos(person)
	if err != nil {
		panic(err)
	} else {
		Clientes = append(Clientes, person)
	}

}

func ValidarDatos(person Persona) (string, error) {

	if person.Legajo == 0 || person.DNI == 0 || person.Telefono == 0 || person.Nombre == "" || person.Domicilio == "" {
		return "", errors.New("No se admiten valores en cero")
	}

	return "Valores válidos", nil

}
