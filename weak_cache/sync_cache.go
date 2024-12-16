package weak_cache

import "sync"

// SyncCache represents a thread-safe sync cache.
type SyncCache[K any, V any] struct {
	items sync.Map
}

// NewSyncCache creates a new generic SyncCache instance.
func NewSyncCache[K comparable, V any]() *SyncCache[K, V] {
	return &SyncCache[K, V]{
		items: sync.Map{},
	}
}

// Get retrieves an item from the cache, if it's still alive.
func (c *SyncCache[K, V]) Get(key K) (*V, bool) {
	val, exists := c.items.Load(key)
	if !exists {
		return nil, false
	}

	return val.(*V), true
}

// Set adds an item to the cache.
func (c *SyncCache[K, V]) Set(key K, value V) {
	c.items.Store(key, &value)
}
