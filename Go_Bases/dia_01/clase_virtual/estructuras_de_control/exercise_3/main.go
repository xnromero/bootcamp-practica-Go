package main 

import ("fmt"
datesI18N "github.com/davidgs/datesi18n"
)

/*Realizar una aplicación que reciba  una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.*/


func main (){

	//Opcion 1
	meses := make(map[int] string)
	
	meses[1] = "Enero"
	meses[2] = "Febrero"
	meses[3] = "Marzo"
	meses[4] = "Abril"
	meses[5] = "Mayo"
	meses[6] = "Junio"
	meses[7] = "Julio"
	meses[8] = "Agosto"
	meses[9] = "Septiembre"
	meses[10] = "Octubre"
	meses[11] = "Noviembre"
	meses[12] = "Diciembre"

	var numero int
	fmt.Println("Ingrese un número de mes: ")
	fmt.Scanln(&numero)
	//fmt.Println("Corresponde al mes: ", meses[numero])

	//Opcion 2
	sp := datesI18N.NewDatesI18N("sp")
	fmt.Println("Corresponde al mes: ", sp.MonthName(numero))
}