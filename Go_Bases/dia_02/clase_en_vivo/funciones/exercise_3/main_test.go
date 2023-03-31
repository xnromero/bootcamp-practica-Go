package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularioSalario(t *testing.T) {
	t.Run("Categoría A", func(t *testing.T) {

		//Arrange
		minutes := 120
		category := "a"
		expectedSalary := 9000.0

		//Act
		obtainedSalary := CalcularSalario(minutes, category)

		//Assert
		assert.Equal(t, expectedSalary, obtainedSalary)
	})

	t.Run("Categoría B", func(t *testing.T) {

		//Arrange
		minutes := 120
		category := "b"
		expectedSalary := 3600.0

		//Act
		obtainedSalary := CalcularSalario(minutes, category)

		//Assert
		assert.Equal(t, expectedSalary, obtainedSalary)
	})

	t.Run("Categoría C", func(t *testing.T) {

		//Arrange
		minutes := 120
		category := "c"
		expectedSalary := 2000.0

		//Act
		obtainedSalary := CalcularSalario(minutes, category)

		//Assert
		assert.Equal(t, expectedSalary, obtainedSalary)
	})

	


}
