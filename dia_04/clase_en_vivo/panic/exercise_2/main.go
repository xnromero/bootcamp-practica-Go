package main

import (
	"io"
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

	file, err := os.Open(name)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	text, err:= io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(text))

	defer file.Close()
}
