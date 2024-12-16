package weak_cache

type ICache[K comparable, V any] interface {
	Get(key K) (*V, bool)
	Set(key K, value V)
}

func readWrite1(c ICache[int, int], n, m int) {
	for i := 0; i < n; i++ {
		c.Set(i, i)
	}
	for i := 0; i < m; i++ {
		c.Get(i)
	}
}

func readWrite2(c ICache[int, int], n int) {
	for i := 0; i < n; i++ {
		c.Set(i, i)
		c.Get(i)
	}
}
