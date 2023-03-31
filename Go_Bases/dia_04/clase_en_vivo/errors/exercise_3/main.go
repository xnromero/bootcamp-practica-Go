package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary int = 120000

	result, err := evaluateSalary(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func evaluateSalary(salary int) (string, error) {
	if salary < 10000 {
		return "", errors.New("ERROR: el salario es menor a 10000")
	} else {
		return "Debe pagar impuestos", nil
	}

}
