package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Test case 1: Regular unsorted slice
	input := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	result := BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestBubbleSort failed. Expected %v, got %v", expected, result)
	}

	// Test case 2: Already sorted slice
	input = []int{1, 2, 3, 4, 5, 6}
	expected = []int{1, 2, 3, 4, 5, 6}
	result = BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestBubbleSort failed. Expected %v, got %v", expected, result)
	}

	// Test case 3: Slice with all identical elements
	input = []int{5, 5, 5, 5}
	expected = []int{5, 5, 5, 5}
	result = BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestBubbleSort failed. Expected %v, got %v", expected, result)
	}

	// Test case 4: Slice with negative numbers
	input = []int{-1, -3, -2, 0, 2, 1}
	expected = []int{-3, -2, -1, 0, 1, 2}
	result = BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestBubbleSort failed. Expected %v, got %v", expected, result)
	}

	// Test case 5: Empty slice
	input = []int{}
	expected = []int{}
	result = BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestBubbleSort failed. Expected %v, got %v", expected, result)
	}

	// Test case 6: Slice with one element
	input = []int{42}
	expected = []int{42}
	result = BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestBubbleSort failed. Expected %v, got %v", expected, result)
	}
}
