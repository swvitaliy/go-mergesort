package slice_ptr

type Props8 struct {
	Len   int64
	Index int64
	Ptr   *int
	Next  *Props8
	Str   string
	Sl    []Props8
	Dbl   float64
	Bool  bool
}

func Props8PointersV1(n int) []*Props8 {
	values := make([]*Props8, n)
	for i := 0; i < n; i++ {
		values[i] = &Props8{
			Len:   int64(i),
			Index: int64(i),
			Ptr:   &i,
			Str:   "blablabla",
			Sl:    nil,
			Dbl:   float64(i),
			Bool:  true,
		}
	}
	return values
}

func Props8PointersV2(n int) []*Props8 {
	values := make([]*Props8, n)
	props := make([]Props8, n)
	for i := 0; i < n; i++ {
		props[i].Len = int64(i)
		props[i].Index = int64(i)
		props[i].Ptr = &i
		props[i].Str = "blablabla"
		props[i].Sl = nil
		props[i].Dbl = float64(i)
		props[i].Bool = true
		values[i] = &props[i]
	}
	return values
}
