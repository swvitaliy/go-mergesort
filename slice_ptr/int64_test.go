package slice_ptr

import (
	"fmt"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 1_000},
	{input: 10_000},
	{input: 100_000},
	{input: 1_000_000},
	//{input: 10_000_000},
}

func BenchmarkInt64ValuesV1(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64ValuesV1(v.input)
			}
		})
	}
}

func BenchmarkInt64ValuesV2(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int64ValuesV2(v.input)
			}
		})
	}
}
