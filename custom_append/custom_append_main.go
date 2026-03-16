package custom_append

import "fmt"

// creating custom append func (let's say we only work with int for now)
func CustomAppend(inputSlice []int, numberToAppend int) []int {
	capOfInputSlice := cap(inputSlice)
	lenOfInputSlice := len(inputSlice)

	// fmt.Printf("Current cap (before appending) %d, Current len %d\n", capOfInputSlice, lenOfInputSlice)

	// If we have to reallocated (more cap)
	if capOfInputSlice <= lenOfInputSlice {
		// fmt.Println("Inside if block - more cap needed")
		// create a new slice with more cap
		var newSliceCap int

		// handeling 0 cap slice
		if capOfInputSlice == 0 {
			newSliceCap = (capOfInputSlice + 1) * 2
		} else {
			newSliceCap = capOfInputSlice * 2
		}
		newSlice := make([]int, lenOfInputSlice, newSliceCap)

		// copy inputslice to new slice
		copy(newSlice, inputSlice)

		// now new slice has more cap so accesing one more index then allocate new number in that index
		newSlice = newSlice[:len(newSlice)+1]
		newSlice[lenOfInputSlice] = numberToAppend // this lenOfInputSlice = len(newSlice) - 1 (basically)
		return newSlice
	}

	// Append without re allocating
	// Creating a copy - don't need to copy when no reallocation is needed
	// var inputSliceCopy = make([]int, lenOfInputSlice, capOfInputSlice)
	// copy(inputSliceCopy, inputSlice)
	// fmt.Printf("Pointer of input slice copy %p\n", &inputSliceCopy[0])

	// fmt.Println(cap(inputSliceCopy), len(inputSliceCopy))

	// We already have extra cap so adjusting len by 1 index and allocating new number to that index
	// inputSliceCopy = inputSliceCopy[:lenOfInputSlice+1]
	// inputSliceCopy[lenOfInputSlice] = numberToAppend // this lenOfInputSlice = len(inputSliceCopy) - 1 (basically)

	fmt.Printf("Pointer of input slice %p\n", &inputSlice[0])
	inputSlice = inputSlice[:lenOfInputSlice+1]
	inputSlice[lenOfInputSlice] = numberToAppend
	return inputSlice
}
