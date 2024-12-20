package main

import (
	_ "algos/sorting"
	"algos/useful"
	"fmt"
	"sort"
)

func main() {
	// names := []string{"ada", "okumu", "maina", "muchiri", "kuya", "olwal", "wambugu", "aketch"}
	// integers := []int{34, -1, 5, 90, -273, 89, 33, 1000}
	// sorting.InsertionSort(names)
	// sorting.BubbleSort(integers)
	// fmt.Println(integers)
	// fmt.Println(names)
	// fmt.Println(useful.CountOccurrences(integers))
	nums := []int{3,8,7,8,7,5}
	fmt.Println(findXSum(nums, 2, 2))
}

func findXSum(nums []int, k, x int) []int {
	n := len(nums)
	answer := make([]int, 0, n-k+1)
	for i := 0; i <= n-k; i++ {
		subArr := nums[i:i+k]
		occurences := useful.CountOccurrences(subArr)
		topFreqX := topMostFreq(occurences, x)
		xsum := 0
		for _, v := range topFreqX {
			xsum += v * occurences[v]
		}
		answer = append(answer, xsum)
	}
	return answer
}

func kLongSubArray(nums []int, i, k int) []int {
	return nums[i : i+k]
}

func topMostFreq(counts map[int]int, x int) []int {
	type kv struct {
		key int
		val int
	}

	var freqSlice []kv
	for key, val := range counts {
		freqSlice = append(freqSlice, kv{key, val})
	}
	sort.Slice(freqSlice, func(i, j int) bool {
		if freqSlice[i].val == freqSlice[j].val {
			return freqSlice[i].key > freqSlice[j].key
		}
		return freqSlice[i].val > freqSlice[j].val
	})
	var result []int
	for i := 0; i < x && i < len(freqSlice); i++ {
		result = append(result, freqSlice[i].key)
	}
	return result
}
