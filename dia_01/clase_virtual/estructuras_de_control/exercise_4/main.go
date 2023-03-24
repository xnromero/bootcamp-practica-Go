package main

import "fmt"

/*Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin. 
Por otro lado también es necesario: 
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.*/

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int

	fmt.Println("Edad de Benjamín: ", employees["Benjamin"])
	
	for _, employee := range employees {
		if employee > 21 {
			count ++
		}
	}
	fmt.Println("Empleados mayores a 21 años: ", count)

	fmt.Println("Agregando a Federico: ")
	employees["Federico"] = 25
	fmt.Println(employees)

	fmt.Println("Eliminando a Pedro: ")
	delete(employees, "Pedro")
	fmt.Println(employees)
}	