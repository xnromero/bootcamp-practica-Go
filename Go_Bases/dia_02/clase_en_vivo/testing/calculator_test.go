package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	//Arrange
	expectedResult := 4.0
	args := []float64{2, 2}

	//Act
	obtainedResult := Add(args...)
	
	//Assert
	assert.Equal(t, expectedResult, obtainedResult)

}

func TestDivide(t *testing.T) {
	t.Run("Divide two positive numbers", func(t *testing.T){
		//Arrange
		num1 := 10.0
		num2 := 2.0
		expectedResult := 5.0
		expectedOk := true

		//Act
		obtainedResult, obtainedOk := Divide(num1, num2)

		//Assert
		assert.Equal(t, expectedResult, obtainedResult)
		assert.Equal(t, expectedOk, obtainedOk)
	})

	t.Run("Cannot divide by zero", func(t *testing.T){
		num1 := 10.0
		num2 := 0.0
		expectedResult := 0.0
		expectedOk := false

		obtainedResult, obtainedOk := Divide(num1, num2)

		//Assert
		assert.Equal(t, expectedResult, obtainedResult)
		assert.Equal(t, expectedOk, obtainedOk)
	})


}


