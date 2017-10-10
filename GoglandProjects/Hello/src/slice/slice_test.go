package slice_test

import (
	"testing"
	"slice"
	"github.com/stretchr/testify/assert"
)

func TestCreateASliceOfIntegers(t *testing.T) {

	//act
	var actualIntegers []int = slice.CreateASliceOfIntegers()

	//assert
	assert.Equal(t, 2, len(actualIntegers), "because slice contains two integers")
}

func TestReturnASliceOfASliceOfIntegers(t *testing.T) {

	//act
	upperBound := 4
	var actualPartialListOfIntegers []int = slice.ReturnASliceOfASliceOfIntegers(upperBound)

	//assert
	assert.Equal(t, 4, len(actualPartialListOfIntegers))
	assert.Equal(t, 1, actualPartialListOfIntegers[0])
	assert.Equal(t, 5, actualPartialListOfIntegers[3])
}

func TestAppendAnIntegerToASliceOfIntegers(t *testing.T) {

	//arrange
	newInteger := 8

	//act
	var actualListOfIntegers = slice.AppendAnIntegerToASliceOfIntegers(newInteger)

	//assert
	assert.Equal(t, 5, len(actualListOfIntegers), "because we appended one integer to it")
	assert.Equal(t, 8, actualListOfIntegers[len(actualListOfIntegers)-1])
}