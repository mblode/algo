package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/algo/common"
	"github.com/algo/search"
	mySort "github.com/algo/sort"
)

func builtInSortStrings(words []string) {
	defer common.TimeTrack(time.Now())
	sort.Strings(words)
}

func builtInSortInts(nums []int) {
	defer common.TimeTrack(time.Now())
	sort.Ints(nums)
}

func main() {
	// Read files
	wordsLong := common.ReadFileAsStrings("words-long.txt")
	wordsShort := common.ReadFileAsStrings("words-short.txt")

	numsLong := common.GenerateSlice(10000)
	numsShort := common.GenerateSlice(20)

	// Built-in sort
	fmt.Println("\n--- Built-in sort ---")
	builtInSortStrings(wordsLong)
	builtInSortInts(numsLong)

	builtInSortStrings(wordsShort)
	builtInSortInts(numsShort)

	// Insertion sort
	wordsLong2 := wordsLong
	numsLong2 := numsLong
	wordsShort2 := wordsShort
	numsShort2 := numsShort

	fmt.Println("\n--- Insertion sort ---")
	mySort.InsertionSort(wordsLong2)
	mySort.InsertionSort(numsLong2)

	mySort.InsertionSort(wordsShort2)
	mySort.InsertionSort(numsShort2)

	// Quick sort
	wordsLong3 := wordsLong
	numsLong3 := numsLong
	wordsShort3 := wordsShort
	numsShort3 := numsShort

	fmt.Println("\n--- Quick sort ---")
	mySort.QuickSortStart(wordsLong3)
	mySort.QuickSortStart(numsLong3)

	mySort.QuickSortStart(wordsShort3)
	mySort.QuickSortStart(numsShort3)

	// String search
	fmt.Println("\n--- String search ---")
	search.LinearSearch("kqnplRdZsI", wordsLong)
	search.BinarySearch("kqnplRdZsI", wordsLong)

	// Number search
	fmt.Println("\n--- Number search ---")
	search.LinearSearch(6775845697, numsShort)
	search.BinarySearch(6775845697, numsShort)
}
