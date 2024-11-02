package slice_ptr

import (
	"testing"
)

func BenchmarkProps8PointersV1(b *testing.B) {
	Props8PointersV1(b.N)
}

func BenchmarkProps8PointersV2(b *testing.B) {
	Props8PointersV2(b.N)
}
