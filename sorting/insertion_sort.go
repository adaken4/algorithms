// sorting/insertion_sort.go
package sorting

// InsertionSort sorts an array of strings in ascending order using the insertion sort algorithm.
func InsertionSort(slice []string) {
	// Declare key to hold the current string being compared
	var key string

	// Declare variables for the current and previous indices
	var current, previous int

	// Outer loop: iterate through the array starting from the second element
	for current = 1; current < len(slice); current++ {
		key = slice[current]   // Pick the current element (key) to be inserted
		previous = current - 1 // Initialize previous to point to the last element of the sorted portion

		// Inner loop: shift elements of the sorted portion bigger than the key to the right to create the correct position for key
		for previous >= 0 && slice[previous] > key {
			slice[previous+1] = slice[previous] // Shift the element to the right
			previous--                          // Move to the next element on the left
		}
		slice[previous+1] = key // Insert the key into its correct position
	}
}
