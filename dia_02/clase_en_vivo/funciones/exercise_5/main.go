package main

/*Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

import ("fmt"
)

const(
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	tarantula = "tarantula"
)

func main(){

	animal, err := Animal(hamster)
	if err != ""{
		fmt.Println(err)
	} else {
		fmt.Println(animal(2))
	}

}

func Dog(cantidad float64) float64{
	return cantidad * 10
}

func Cat(cantidad float64) float64{
	return cantidad * 5
}

func Hamster(cantidad float64) float64{
	return cantidad * 0.25
}

func Tarantula(cantidad float64) float64{
	return cantidad * 0.15
}

func Animal(animal string) (func(float64) float64, string){
	switch animal{
	case dog:
		return Dog, ""
	case cat:
		return Cat, ""
	case hamster:
		return Hamster, ""
	case tarantula:
		return Tarantula, ""
	default:
		return nil, "Animal no válido"
	}
}

