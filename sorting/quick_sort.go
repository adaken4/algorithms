package sorting

// QuickSort sorts a slice of integers using the quick sort algorithm.
func QuickSort(arr []int) []int {
  // Base case: arrays of length 0 or 1 are already sorted
  if len(arr) <= 1 {
		return arr
	}

  // Choose the first element as pivot
  // Note: For optimization, consider using median-of-three
	pivot := arr[0]

  // Create three partitions:
	lesser := []int{}
  equal := []int{}
	greater := []int{}

  // Partition the array into three parts handling duplicates efficiently
	for _, v := range arr {
		if v == pivot {
			equal = append(equal, v)
		} else if v < pivot {
			lesser = append(lesser, v)
		} else {
			greater = append(greater, v)
		}
	}

  // Recursively sort the sub-arrays
	sortedLesser := quickSort(lesser)
	sortedGreater := quickSort(greater)

  // Combine the sorted partitions with the pivot values
  result := append(sortedLesser, equal...)
  result = append(result, sortedGreater...)

	return result
}
