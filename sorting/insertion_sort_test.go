package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	// Test case 1: Regular unsorted slice
	names := []string{"ada", "okumu", "maina", "muchiri", "kuya", "olwal", "wambugu", "aketch"}
	expected := []string{"ada", "aketch", "kuya", "maina", "muchiri", "okumu", "olwal", "wambugu"}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}

	// Test case 2: Already sorted slice
	names = []string{"apple", "banana", "cherry", "date"}
	expected = []string{"apple", "banana", "cherry", "date"}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}

	// Test case 3: Slice with all identical elements
	names = []string{"same", "same", "same", "same"}
	expected = []string{"same", "same", "same", "same"}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}

	// Test case 4: Slice with one element
	names = []string{"single"}
	expected = []string{"single"}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}

	// Test case 5: Empty slice
	names = []string{}
	expected = []string{}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}

	// Test case 6: Slice with mixed case strings
	names = []string{"Zebra", "apple", "Orange", "banana"}
	expected = []string{"Orange", "Zebra", "apple", "banana"}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}

	// Test case 7: Slice with numbers as strings
	names = []string{"10", "2", "33", "21"}
	expected = []string{"10", "2", "21", "33"}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}

	// Test case 8: Slice with special characters
	names = []string{"!", "@", "#", "$", "%", "^", "&", "*", "("}
	expected = []string{"!", "#", "$", "%", "&", "(", "*", "@", "^"}
	InsertionSort(names)
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("TestInsertionSort failed. Expected %v, got %v", expected, names)
	}
}
