package sorting

// BubbleSort sorts a slice of integers using the bubble sort algorithm.
func BubbleSort(slice []int) []int {
	var i, j int     // Loop counters
	var swapped bool // Flag to check if any elements were swapped

	// Outer loop goes through each element in the slice from first to last comparing pairs
	for i = 0; i < len(slice)-1; i++ {
		swapped = false // Reset the swapped flag for this pass

		// Inner loop compares adjacent elements up to len(slice)-i-1
		// Since the last i elements are already sorted
		for j = 0; j < len(slice)-i-1; j++ {
			// If current element is greater than the next, swap them
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true // Set swapped flag to true since we made a swap
			}
		}

		// If no elments were swapped in this pass, the slice is sorted
		// We can exit early for optimization
		if !swapped {
			break
		}
	}
	return slice // Return the sorted slice
}
