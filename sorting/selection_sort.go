package sorting

// SelectionSort sorts an array of strings in ascending order using the selection sort algorithm.
func SelectionSort(arr []string) {
	var i, j, minIdx int
	// Traverse through array elements to the second last element
	for i = 0; i < len(arr)-1; i++ {
		// Assume the current index is the smallest
		minIdx = i
		// Find the smallest element in the remaining unsorted array
		for j = i + 1; j < len(arr); j++ {
			// Keep updating minIdx as a smaller element is found
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Swap the found minimum element with the first element of the unsorted array
		if minIdx != i {
			arr[minIdx], arr[i] = arr[i], arr[minIdx]
		}
	}
}
