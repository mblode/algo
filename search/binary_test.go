package search

import (
	"testing"
)

func TestBinarySearchIntInList(t *testing.T) {
	var itemsNumbers = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	ans := BinarySearch(1, itemsNumbers)

	if ans != true {
		t.Errorf("BinarySearch(1, itemsNumbers) = %t; want true", ans)
	}
}

func TestBinarySearchIntNotInList(t *testing.T) {
	var itemsNumbers = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	ans := BinarySearch(0, itemsNumbers)

	if ans != false {
		t.Errorf("BinarySearch(0, itemsNumbers) = %t; want false", ans)
	}
}

func TestBinarySearchStringInList(t *testing.T) {
	var itemsWords = []string{"bob", "hello", "is", "my", "name"}
	ans := BinarySearch("is", itemsWords)

	if ans != true {
		t.Errorf("BinarySearch(\"hello\", itemsWords) = %t; want true", ans)
	}
}

func TestBinarySearchStringNotInList(t *testing.T) {
	var itemsWords = []string{"bob", "hello", "is", "my", "name"}
	ans := BinarySearch("foo", itemsWords)

	if ans != false {
		t.Errorf("BinarySearch(\"foo\", itemsWords) = %t; want false", ans)
	}
}
