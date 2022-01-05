package sort

import (
	"time"

	"github.com/algo/common"
)

func InsertionSort[K common.Generic](items []K) []K {
	defer common.TimeTrack(time.Now())
	n := len(items)

	for i := 1; i < n; i++ {
		j := i

		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}

			j = j - 1
		}
	}

	return items
}
