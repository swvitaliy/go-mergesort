package slice_ptr

type Int64Value struct {
	value int64
}

func Int64ValuesV1(n int) []*Int64Value {
	values := make([]*Int64Value, n)
	for i := 0; i < n; i++ {
		values[i] = &Int64Value{value: int64(i)}
	}
	return values
}

func Int64ValuesV2(n int) []*Int64Value {
	values := make([]*Int64Value, n)
	ints := make([]Int64Value, n)
	for i := 0; i < n; i++ {
		ints[i].value = int64(i)
		values[i] = &ints[i]
	}
	return values
}
