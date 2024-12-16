package weak_cache

import (
	"sync"
	"weak"
)

// SyncWeakCache represents a thread-safe cache with weak pointers.
type SyncWeakCache[K comparable, V any] struct {
	items sync.Map
}

// NewSyncWeakCache creates a new generic SyncWeakCache instance.
func NewSyncWeakCache[K comparable, V any]() *SyncWeakCache[K, V] {
	return &SyncWeakCache[K, V]{
		items: sync.Map{},
	}
}

// Get retrieves an item from the cache, if it's still alive.
func (c *SyncWeakCache[K, V]) Get(key K) (*V, bool) {
	// Retrieve the weak pointer for the given key
	ptr, exists := c.items.Load(key)
	if !exists {
		return nil, false
	}

	// Attempt to dereference the weak pointer
	weakPtr := ptr.(weak.Pointer[V])
	val := (weakPtr).Value()
	if val == nil {
		// Object has been reclaimed by the garbage collector
		c.items.Delete(key)
		return nil, false
	}

	return val, true
}

// Set adds an item to the cache.
func (c *SyncWeakCache[K, V]) Set(key K, value V) {
	// Create a weak pointer to the value
	c.items.Store(key, weak.Make(&value))
}
