package main

import (
	"fmt"
)

type MyError struct {
	salary int
}

func (e MyError) Error() string {
	message := fmt.Sprintf("Error: el mínimo imponible es de 150000 y el salario ingresado es de %d.", e.salary)
	return message
}

func main() {
	var salary int = 12
	result, err := evaluateSalary(salary)
	
	if err != nil {
		error := fmt.Errorf(err.Error())
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

}

func evaluateSalary(salary int) (string, error) {
	if salary < 10000 {
		return "", MyError{salary: salary}
	} else {
		return "Debe pagar impuestos", nil
	}

}
