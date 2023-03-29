package main

import (
	"fmt"
)

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func main() {
	var salary int = 2500000
	
	if salary < 150000{
		myError := MyError{msg: "El salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(myError.Error())
	} else{
		fmt.Println("Debe pagar impuestos")
	}

	


}
