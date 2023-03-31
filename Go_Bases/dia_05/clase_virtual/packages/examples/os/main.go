package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//readFile()
	//createFile()
	fileDetails()
}

func readFile() {
	data, err := os.ReadFile("prueba.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}

func createFile(){
	f, err := os.Create("nuevo_archivo.txt")
	if err != nil{
		log.Fatal(err)
	}

	_, err2 := f.WriteString("Texto generado desde código")
	if err2 != nil{
		log.Fatal(err)
	}

	fmt.Println("Done!")
}

func fileDetails(){
	file := "prueba.txt"
	f, err := os.Stat(file)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Es un directorio: ", f.IsDir())
	fmt.Println("Nombre del archivo/directorio: ", f.Name())
	fmt.Println("Tamaño del archivo en bytes: ", f.Size())
	fmt.Println("Fecha y hora del archivo: ", f.ModTime())
	fmt.Println("Permisos del archivo: ", f.Mode())
}
