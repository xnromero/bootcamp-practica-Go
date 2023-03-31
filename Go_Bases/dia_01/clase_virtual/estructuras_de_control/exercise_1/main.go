package main

import "fmt"

/*La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
Luego, imprimí cada una de las letras.*/

func main(){
	
	var word string

	fmt.Println("Ingrese una palabra: ")
	fmt.Scanln(&word)

	fmt.Println("Cantidad de letras: ", len(word))

	fmt.Println("Impresión de letras: ")
	for _, letter := range word {
		fmt.Println(string(letter))
	}
}