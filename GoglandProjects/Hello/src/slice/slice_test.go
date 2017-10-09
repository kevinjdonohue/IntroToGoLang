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