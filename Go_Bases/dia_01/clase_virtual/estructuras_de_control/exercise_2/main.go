package main

import "fmt"

/*Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000. 
Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.*/


func main() {
	var edad int
	var isEmpleado bool
	var antiguedad int
	var sueldo float64

	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&edad)

	if edad < 22 {
		fmt.Println("No se puede otorgar el préstamo porque no cumple con el requisito de edad")
	} else {
		fmt.Println("Es empleado?: ")
		fmt.Scanln(&isEmpleado)
		if !isEmpleado {
			fmt.Println("No se puede otorgar el préstamo porque no se encuentra empleado")
		} else { 
			fmt.Println("Ingrese antigüedad en años en la empresa: ")
			fmt.Scanln(&antiguedad)
			if antiguedad < 1 {
				fmt.Println("No se puede otorgar el préstamo porque no cumple con el requisito de antigüedad")
			} else {
				fmt.Println("Ingrese su salario: ")
				fmt.Scanln(&sueldo)
				if sueldo >= 100000 {
					fmt.Println("Es elegible para obtener el préstamo sin intereses")
				} else {
					fmt.Println("Es elegible para obtener el préstamo con intereses")
				}
			}
		}
	}
}