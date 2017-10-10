package slice

func CreateASliceOfIntegers() []int {
	var someIntegers []int = make([]int, 5, 10)
	//NOTE:  You can create a slice with the capacity the same as the size
	//var someIntegers []int = make([]int, 5)
	//shorthand
	//someIntegers := []int { 1, 2, 3 }

	someIntegers[0] = 1
	someIntegers[1] = 2

	return someIntegers
}

func ReturnASliceOfASliceOfIntegers(upperBound int) []int {
	someIntegers := []int {1,2,3,5,8,13,21}

	//NOTE:  you can leave off either boundary to either start at the beg or run to the end
	//someIntegers[:4]
	//someIntegers[1:]
	//you can also use something like someIntegers[5:len(someIntegers)]
	var partialListOfIntegers = someIntegers[0:4]

	return partialListOfIntegers
}

func AppendAnIntegerToASliceOfIntegers(newValue int) []int {
	someIntegers := []int{1, 2, 3, 5}

	moreIntegers := append(someIntegers, newValue)

	return moreIntegers
}
