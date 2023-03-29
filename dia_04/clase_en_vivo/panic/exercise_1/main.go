package main

import (
	"fmt"
	"os"
)

func main() {

	OpenArchive("customers.txt")
	fmt.Println("Ejecución finalizada")

}
func OpenArchive(name string) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.Open(name)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

}
