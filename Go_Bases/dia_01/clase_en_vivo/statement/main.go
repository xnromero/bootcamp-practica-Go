package main

import "fmt"

func main () {
	//IfStatement()
	SwitchStatement()
}

func IfStatement(){
	var edad int
	edad = 20

	if edad <= 14 {
		fmt.Println("No puede hacer nada")
		return
	} 
	if edad >= 18 {
		fmt.Println("Puede votar")
	} 
	if edad >= 16{
		fmt.Println("Puede conducir")
	}
	if edad >=14 {
		fmt.Println("puede ver peliculas clasificadas PG-13")
	}
}

func SwitchStatement() {

	edad :=15

	if edad >= 0 {
		switch{
		case edad >= 18:
			fmt.Println("Puede votar")
			fallthrough
		case edad >= 16:
			fmt.Println("Puede conducir")
			fallthrough
		case edad >=14 :
			fmt.Println("puede ver peliculas clasificadas PG-13")
		default:
			fmt.Print("No puede hacer nada")
		}
	} else {
		fmt.Println("Edad negativa")
	}

}