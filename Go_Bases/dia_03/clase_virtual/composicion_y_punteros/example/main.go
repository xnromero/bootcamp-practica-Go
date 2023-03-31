package main


import (
	"fmt"
)

func main(){
	var v int = 19
	//var p *int
	p := &v
	fmt.Println("El puntero p referencia a la dirección de memoria: ", p)
	fmt.Println("Al desreferenciar el puntero se obtiene: ", *p)
}