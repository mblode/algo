package common

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileAsStrings(fname string) []string {
	b, err := os.ReadFile(fname)
	check(err)

	splitLines := strings.Split(string(b), "\n")

	return splitLines
}

func ReadFileAsNumbers(fname string) []int {
	b, err := os.ReadFile(fname)
	check(err)

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		check(err)
		nums = append(nums, n)
	}

	return nums
}
