package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularEstadisticas(t *testing.T) {
	t.Run("Mínimo de calificaciones", func(t *testing.T) {

		//Arrange
		notas := []float64{4.0, 5.0, 6.0, 5.0}
		expectedMin := 4.0
		expectedErr := ""

		//Act

		obtainedOperation, obtainedErr := CalcularEstadisticas(minimum, notas)
		obtainedMin := obtainedOperation(notas)

		//Asserts
		assert.Equal(t, expectedMin, obtainedMin)
		assert.Equal(t, expectedErr, obtainedErr)

	})

	t.Run("Máximo de calificaciones", func(t *testing.T) {

		//Arrange
		notas := []float64{4.0, 5.0, 6.0, 5.0}
		expectedMax := 6.0
		expectedErr := ""

		//Act

		obtainedOperation, obtainedErr := CalcularEstadisticas(maximum, notas)
		obtainedMax := obtainedOperation(notas)

		//Asserts
		assert.Equal(t, expectedMax, obtainedMax)
		assert.Equal(t, expectedErr, obtainedErr)

	})

	t.Run("Promedio de calificaciones", func(t *testing.T) {

		//Arrange
		notas := []float64{4.0, 5.0, 6.0, 5.0}
		expectedAverage := 5.0
		expectedErr := ""

		//Act

		obtainedOperation, obtainedErr := CalcularEstadisticas(average, notas)
		obtainedAverage:= obtainedOperation(notas)

		//Asserts
		assert.Equal(t, expectedAverage, obtainedAverage)
		assert.Equal(t, expectedErr, obtainedErr)

	})
}
