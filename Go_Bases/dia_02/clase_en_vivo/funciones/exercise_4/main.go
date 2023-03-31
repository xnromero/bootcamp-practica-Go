package main

/*Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar 
una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
*/

import ("fmt"
)

const(
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main(){

	notas := []float64 {4.0, 5.0, 6.0, 5.0}

	operacionDefinida, err := CalcularEstadisticas(average, notas)
	if err != ""{
		fmt.Println(err)
	} else {
		fmt.Println(operacionDefinida(notas))
	}

}

func Minimo(notas []float64) float64 {
    min := 999.99

	for _, nota := range notas {
		if nota < min {
			min = nota
		} 
	}
    return min   
}

func Maximo(notas []float64) float64 {
    max := 0.0

	for _, nota := range notas {
		if nota > max {
			max = nota
		} 
	}  
    return max 
}

func Promedio (notas []float64) float64 { 
	var promedio float64
	for _, nota := range notas {
		promedio += nota
	}
	return promedio / float64(len(notas))
}

func CalcularEstadisticas(operador string, notas []float64) (func([]float64) float64, string)  {

    switch operador{
    case minimum:
        return Minimo, ""
    case maximum:
        return Maximo, ""
    case average:
      return Promedio, ""
	default:
		return nil, "Operación no definida"
    }

}
