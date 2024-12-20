package useful

func CountOccurrences(arr []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}
	return counts
}