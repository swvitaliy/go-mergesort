package mergesort

import (
	"fmt"
	"math/rand"
)

// Sample Just to be sure that sorting algorithm is working
func Sample() {
	n := 42
	a := make([]int, n, n)
	fmt.Println("Unsorted:")
	for i := range a {
		a[i] = rand.Intn(100500)
	}
	for _, ai := range a {
		fmt.Print(ai, " ")
	}
	fmt.Println()
	fmt.Printf("isSorted? %v\n", IsSorted(a))

	ms := NewMergeSort(n)
	ms.Run(2, a, 0, n-1, false)

	fmt.Println("Sorted:")
	for _, ai := range a {
		fmt.Print(ai, " ")
	}
	fmt.Println()
	fmt.Printf("isSorted? %v\n", IsSorted(a))
}
