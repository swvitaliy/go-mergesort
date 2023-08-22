package mergesort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

var a100k []int
var a1M []int

func setup() {
	a100k = make([]int, 100000, 100000)
	for i := 0; i < len(a100k); i++ {
		a100k[i] = rand.Intn(100500)
	}
	a1M = make([]int, 1000000, 1000000)
	for i := 0; i < len(a1M); i++ {
		a1M[i] = rand.Intn(100500)
	}
}

func invalidateCPUCache(a []int) {
	for i := 0; i < 1000; i++ {
		a[rand.Intn(len(a))] = rand.Intn(100500)
	}
}

func tearDown() {

}

// ------------------------------------------------------------------------------
// Benchmark with K=8 (number of goroutines = 2^k = 256)

func BenchmarkRunA1M_K8_Optimized(e *testing.B) {
	a := make([]int, len(a1M), len(a1M))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a1M)
		invalidateCPUCache(a)
		ms.Run(8, a, 0, n-1, true)
	}
}

func BenchmarkRunA1M_K8(e *testing.B) {
	a := make([]int, len(a1M), len(a1M))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a1M)
		invalidateCPUCache(a)
		ms.Run(8, a, 0, n-1, false)
	}
}

func BenchmarkRunA100k_K8_Optimized(e *testing.B) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a100k)
		invalidateCPUCache(a)
		ms.Run(8, a, 0, n-1, true)
	}
}

func TestRunA100k_K8(t *testing.T) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	copy(a, a100k)
	invalidateCPUCache(a)
	ms.Run(8, a, 0, n-1, false)
	assert.True(t, IsSorted(a), "Sorted")
}

func BenchmarkRunA100k_K8(e *testing.B) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a100k)
		invalidateCPUCache(a)
		ms.Run(8, a, 0, n-1, false)
	}
}

// ------------------------------------------------------------------------------
// Benchmark with K=3 (number of goroutines = 2^k = 8)

func BenchmarkRunA1M_K3_Optimized(e *testing.B) {
	a := make([]int, len(a1M), len(a1M))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a1M)
		invalidateCPUCache(a)
		ms.Run(3, a, 0, n-1, true)
	}
}

func BenchmarkRunA1M_K3(e *testing.B) {
	a := make([]int, len(a1M), len(a1M))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a1M)
		invalidateCPUCache(a)
		ms.Run(3, a, 0, n-1, false)
	}
}

func TestRunA100k_K3_Optimized(t *testing.T) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	copy(a, a100k)
	invalidateCPUCache(a)
	ms.Run(3, a, 0, n-1, true)
	assert.True(t, IsSorted(a), "Sorted")
}

func BenchmarkRunA100k_K3_Optimized(e *testing.B) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a100k)
		invalidateCPUCache(a)
		ms.Run(3, a, 0, n-1, true)
	}
}

func BenchmarkRunA100k_K3(e *testing.B) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a100k)
		invalidateCPUCache(a)
		ms.Run(3, a, 0, n-1, false)
	}
}

// ------------------------------------------------------------------------------
// Benchmark with K=0

func TestRunA100k_K0_Optimized(t *testing.T) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	copy(a, a100k)
	invalidateCPUCache(a)
	ms.Run(0, a, 0, n-1, true)
	assert.True(t, IsSorted(a), "Sorted")
}

func BenchmarkRunA100k_K0_Optimized(e *testing.B) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a100k)
		invalidateCPUCache(a)
		ms.Run(0, a, 0, n-1, true)
	}
}

func BenchmarkRunA100k_K0(e *testing.B) {
	a := make([]int, len(a100k), len(a100k))
	n := len(a)
	ms := NewMergeSort(n)
	for i := 0; i < e.N; i++ {
		copy(a, a100k)
		invalidateCPUCache(a)
		ms.Run(0, a, 0, n-1, false)
	}
}
