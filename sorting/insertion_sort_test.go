package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	names := []string{"ada", "okumu", "maina", "muchiri", "kuya", "olwal", "wambugu", "aketch"}
	expected := []string{"ada", "aketch", "kuya", "maina", "muchiri", "okumu", "olwal", "wambugu"}

	InsertionSort(names)

	if !reflect.DeepEqual(names, expected) {
		t.Errorf("SelectionSort() = %v, want %v", names, expected)
	}
}
