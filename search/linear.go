package search

import (
	"time"

	"github.com/algo/common"
)

func LinearSearch[K common.Generic](needle K, haystack []K) bool {
	defer common.TimeTrack(time.Now())

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}
