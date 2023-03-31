package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcularPromedio(t *testing.T) {
	//Arrange
	nota1 := 4.0
	nota2 := 10.0
	nota3 := 6.0
	nota4 := 8.0
	expectedAverage := 7.0
	var expectedErr error = nil

	//Act
	obtainedAverage, obtainedErr := CalcularPromedio(nota1, nota2, nota3, nota4)

	//Asserts
	assert.Equal(t, expectedAverage, obtainedAverage)
	assert.Equal(t, expectedErr, obtainedErr)
}
