package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDog(t *testing.T) {
	//Arrange
	cantidad := 10.0
	expectedKg := 100.0

	//Act
	obtainedKg := Dog(cantidad)

	//Assert
	assert.Equal(t, expectedKg, obtainedKg)
}

func TestCat(t *testing.T) {
	//Arrange
	cantidad := 10.0
	expectedKg := 50.0

	//Act
	obtainedKg := Cat(cantidad)

	//Assert
	assert.Equal(t, expectedKg, obtainedKg)
}

func TestHamster(t *testing.T) {
	//Arrange
	cantidad := 10.0
	expectedKg := 2.5

	//Act
	obtainedKg := Hamster(cantidad)

	//Assert
	assert.Equal(t, expectedKg, obtainedKg)
}

func TestTarantula(t *testing.T) {
	//Arrange
	cantidad := 10.0
	expectedKg := 1.5

	//Act
	obtainedKg := Tarantula(cantidad)

	//Assert
	assert.Equal(t, expectedKg, obtainedKg)
}