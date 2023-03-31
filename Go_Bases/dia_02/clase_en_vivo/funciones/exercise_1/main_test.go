package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuestoSalario(t *testing.T) {

	t.Run("Salario por debajo de 50000", func(t *testing.T){

		//Arrange
		salary := 40000.0
		expectedTax := 0.0

		//Act
		obtainedTax := CalcularImpuestoSalario(salary)

		//Assert
		assert.Equal(t, expectedTax, obtainedTax)
	})

	t.Run("Salario por encima de 50000", func(t *testing.T){

		//Arrange
		salary := 60000.0
		expectedTax := 10200.0

		//Act
		obtainedTax := CalcularImpuestoSalario(salary)

		//Assert
		assert.Equal(t, expectedTax, obtainedTax)
	})

	t.Run("Salario por encima de 150000", func(t *testing.T){

		//Arrange
		salary := 160000.0
		expectedTax := 43200.0

		//Act
		obtainedTax := CalcularImpuestoSalario(salary)

		//Assert
		assert.Equal(t, expectedTax, obtainedTax)
	})

}


