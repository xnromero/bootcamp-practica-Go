package main

/*Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. 
No se pueden introducir notas negativas.*/

import ("fmt"
		"errors"
)

func main(){

	resultado, error := CalcularPromedio(5,6,5,7)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(resultado)
	}

}

func CalcularPromedio(calificaciones ...float64) (float64, error){
	var promedio float64

	for _, calificacion := range calificaciones {
		if calificacion > 0 {
			promedio += calificacion
		} else {
			return 0, errors.New("No se permiten notas negativas")
		}
		
	}
	return promedio / float64(len(calificaciones)), nil
}
