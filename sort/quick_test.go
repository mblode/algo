package sort

import (
	"reflect"
	"testing"
)

func TestQuickSortInts(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 2, 4, 1}, expected: []int{1, 2, 3, 4}},
		{input: []int{40, 20, 30}, expected: []int{20, 30, 40}},
		{input: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		output := QuickSort(test.input)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("got %v want %v", output, test.expected)
		}
	}
}

func TestQuickSortStrings(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{input: []string{"a", "c", "b"}, expected: []string{"a", "b", "c"}},
		{input: []string{"ac", "ab", "aa"}, expected: []string{"aa", "ab", "ac"}},
		{input: []string{"pancakes", "burgers", "cheese"}, expected: []string{"burgers", "cheese", "pancakes"}},
	}

	for _, test := range tests {
		output := QuickSort(test.input)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("got %v want %v", output, test.expected)
		}
	}
}
