package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func main() {
	var salary int = 250

	result, err := evaluateSalary(salary)

	if errors.Is(err, MyError{msg: "ERROR: el salario es menor a 10000"}) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func evaluateSalary(salary int) (string, error) {
	if salary < 10000 {
		return "", MyError{msg: "ERROR: el salario es menor a 10000"}
	} else {
		return "Debe pagar impuestos", nil
	}

}
