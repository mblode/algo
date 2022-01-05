package sort

import (
	"time"
	"math/rand"
	"github.com/algo/common"
)


func QuickSort[K common.Generic](a []K) []K {
	if len(a) < 2 {
		return a
	}
	
	left, right := 0, len(a)-1
	
	pivot := rand.Int() % len(a)
	
	a[pivot], a[right] = a[right], a[pivot]
	
	for i, _ := range a {
			if a[i] < a[right] {
					a[left], a[i] = a[i], a[left]
					left++
			}
	}
	
	a[left], a[right] = a[right], a[left]
	
	QuickSort(a[:left])
	QuickSort(a[left+1:])
	
	return a
}

func QuickSortStart[K common.Generic](a []K) []K {
	defer common.TimeTrack(time.Now())
	return QuickSort(a)
}