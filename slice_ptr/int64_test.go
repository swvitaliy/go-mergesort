package slice_ptr

import (
	"testing"
)

func BenchmarkInt64ValuesV1(b *testing.B) {
	Int64ValuesV1(b.N)
}

func BenchmarkInt64ValuesV2(b *testing.B) {
	Int64ValuesV2(b.N)
}
