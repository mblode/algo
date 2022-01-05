package search

import (
	"testing"
)

func TestLinearSearchIntInList(t *testing.T) {
	var itemsNumbers = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	ans := LinearSearch(1, itemsNumbers)

	if ans != true {
		t.Errorf("LinearSearch(1, itemsNumbers) = %t; want true", ans)
	}
}

func TestLinearSearchIntNotInList(t *testing.T) {
	var itemsNumbers = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	ans := LinearSearch(0, itemsNumbers)

	if ans != false {
		t.Errorf("LinearSearch(0, itemsNumbers) = %t; want false", ans)
	}
}

func TestLinearSearchStringInList(t *testing.T) {
	var itemsWords = []string{"hello", "my", "name", "is", "bob"}
	ans := LinearSearch("hello", itemsWords)

	if ans != true {
		t.Errorf("LinearSearch(\"hello\", itemsWords) = %t; want true", ans)
	}
}

func TestLinearSearchStringNotInList(t *testing.T) {
	var itemsWords = []string{"hello", "my", "name", "is", "bob"}
	ans := LinearSearch("foo", itemsWords)

	if ans != false {
		t.Errorf("LinearSearch(\"foo\", itemsWords) = %t; want false", ans)
	}
}
