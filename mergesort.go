package mergesort

import (
	"sync"
)

type MergeSort struct {
	b []int
}

func NewMergeSort(n int) *MergeSort {
	return &MergeSort{
		b: make([]int, n, n),
	}
}

type sortFn = func(a []int, i, j int)

func (ms *MergeSort) Run(k int, a []int, i, j int, optimized bool) {
	var fn sortFn
	if optimized {
		fn = ms.mergeSortOptimized
	} else {
		fn = ms.mergeSort
	}

	ms.run(k, a, i, j, fn)
}

func merge(a, b []int, l, m, r int) {
	// b := make([]int, n-i+1, n-i+1)
	i := l
	k := l
	j := m + 1
	for i <= m && j <= r {
		if a[i] < a[j] {
			b[k] = a[i]
			i++
			k++
		} else {
			b[k] = a[j]
			j++
			k++
		}
	}
	if i <= m {
		copy(b[k:], a[i:m+1])
	}
	if j <= r {
		copy(b[k:], a[j:r+1])
	}
	copy(a[l:], b[l:r+1]) // copy(a[l:], b[:r-l+1]) when using a local b slice
}

// number_of_goroutines = 2^k
func (ms *MergeSort) run(k int, a []int, i, j int, sort sortFn) {
	if k == 0 {
		sort(a, i, j)
	} else {
		m := (i + j) / 2
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			ms.run(k-1, a, i, m, sort)
		}()
		go func() {
			defer wg.Done()
			ms.run(k-1, a, m+1, j, sort)
		}()
		wg.Wait()
		merge(a, ms.b, i, m, j)
	}
}

func insertionSort(a []int, i, j int) {
	for k := i + 1; k <= j; k++ {
		for l := k; l > i && a[l-1] > a[l]; l-- {
			a[l], a[l-1] = a[l-1], a[l]
		}
	}
}

// TimSort (switch to insertion sort for sub arrays of size 64 or less)
func (ms *MergeSort) mergeSortOptimized(a []int, i, j int) {
	if j-i < 64 {
		insertionSort(a, i, j)
		return
	}

	m := (i + j) / 2
	ms.mergeSortOptimized(a, i, m)
	ms.mergeSortOptimized(a, m+1, j)
	merge(a, ms.b, i, m, j)
}

func (ms *MergeSort) mergeSort(a []int, i, j int) {
	if i < j {
		m := (i + j) / 2
		ms.mergeSort(a, i, m)
		ms.mergeSort(a, m+1, j)
		merge(a, ms.b, i, m, j)
	}
}

func IsSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
