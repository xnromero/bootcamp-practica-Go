package main

/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.
*/

import ("fmt"
		"strings"
)

func main(){

	fmt.Println(CalcularSalario(120, "a"))
}

func CalcularSalario (minutos int, categoria string) float64 {
	var horasTrabajadas float64
	horasTrabajadas = float64(minutos / 60)
	var salario float64
	
	switch strings.ToLower(categoria){
	case "a":
		salario = (horasTrabajadas * 3000) * 1.5
	case "b":
		salario = (horasTrabajadas * 1500) * 1.2
	case "c":
		salario = (horasTrabajadas * 1000)
	default:
		salario = 0
	}
	return salario
}
