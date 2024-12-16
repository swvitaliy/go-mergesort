package weak_cache

import (
	"fmt"
	"testing"
)

func BenchmarkSyncWeakCacheRW1(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			c := NewSyncWeakCache[int, int]()
			for i := 0; i < b.N; i++ {
				readWrite1(c, v.input, v.input)
			}
		})
	}
}

func BenchmarkSyncWeakCacheRW2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			c := NewSyncWeakCache[int, int]()
			for i := 0; i < b.N; i++ {
				readWrite2(c, v.input)
			}
		})
	}
}
