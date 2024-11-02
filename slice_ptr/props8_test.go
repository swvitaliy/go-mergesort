package slice_ptr

import (
	"fmt"
	"testing"
)

func BenchmarkProps8PointersV1(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Props8PointersV1(v.input)
			}
		})
	}
}

func BenchmarkProps8PointersV2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Props8PointersV2(v.input)
			}
		})
	}
}
