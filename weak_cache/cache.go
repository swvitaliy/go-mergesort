package weak_cache

import "sync"

// Cache represents a thread-safe cache .
type Cache[K comparable, V any] struct {
	mu    sync.Mutex
	items map[K]*V
}

// NewCache creates a new generic Cache instance.
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items: make(map[K]*V),
	}
}

// Get retrieves an item from the cache, if it's still alive.
func (c *Cache[K, V]) Get(key K) (*V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, exists := c.items[key]
	if !exists {
		return nil, false
	}

	return val, true
}

// Set adds an item to the cache.
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &value
}
