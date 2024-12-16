package weak_cache

import (
	"sync"
	"weak"
)

// WeakCache represents a thread-safe cache with weak pointers.
type WeakCache[K comparable, V any] struct {
	mu    sync.Mutex
	items map[K]weak.Pointer[V] // Weak pointers to cached objects
}

// NewWeakCache creates a new generic WeakCache instance.
func NewWeakCache[K comparable, V any]() *WeakCache[K, V] {
	return &WeakCache[K, V]{
		items: make(map[K]weak.Pointer[V]),
	}
}

// Get retrieves an item from the cache, if it's still alive.
func (c *WeakCache[K, V]) Get(key K) (*V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Retrieve the weak pointer for the given key
	ptr, exists := c.items[key]
	if !exists {
		return nil, false
	}

	// Attempt to dereference the weak pointer
	val := ptr.Value()
	if val == nil {
		// Object has been reclaimed by the garbage collector
		delete(c.items, key)
		return nil, false
	}

	return val, true
}

// Set adds an item to the cache.
func (c *WeakCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Create a weak pointer to the value
	c.items[key] = weak.Make(&value)
}
