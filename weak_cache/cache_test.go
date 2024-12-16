package weak_cache

import (
	"fmt"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 10_000},
	{input: 50_000},
	{input: 250_000},
}

func BenchmarkCacheRW1(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			c := NewCache[int, int]()
			for i := 0; i < b.N; i++ {
				readWrite1(c, v.input, v.input)
			}
		})
	}
}

func BenchmarkCacheRW2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			c := NewCache[int, int]()
			for i := 0; i < b.N; i++ {
				readWrite2(c, v.input)
			}
		})
	}
}
