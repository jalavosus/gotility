package maps

import (
	"sync"
)

// ConcurrentMap is a generic, concurrency-safe wrapper
// around the inbuilt map type.
type ConcurrentMap[K comparable, V any] struct {
	m   map[K]V
	mut sync.RWMutex
}

// NewConcurrentMap returns an initialized. ConcurrentMap with type constraints K and V,
// K being any valid map key type (satisfies comparable), V being any type whatsoever.
func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	c := new(ConcurrentMap[K, V])
	c.m = make(map[K]V)

	return c
}

// Get returns the value stored at the passed key (if it exists)
// and a boolean corresponding to key existence.
// Equivalent to `val, ok := map[key]`
func (c *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	return c.get(key)
}

// HasKey returns true if `key` is set in the underlying map.
// Equivalent to `_, ok := map[key]; ok`
func (c *ConcurrentMap[K, V]) HasKey(key K) bool {
	_, ok := c.get(key)
	return ok
}

// Set assigns `val` to `key` in the underlying map.
// Equivalent to `map[key] = val`
func (c *ConcurrentMap[K, V]) Set(key K, val V) {
	c.set(key, val)
}

// Unset removes `key` from the underlying map.
// Equivalent to `delete(map, key)`
func (c *ConcurrentMap[K, V]) Unset(key K) {
	c.unset(key)
}

func (c *ConcurrentMap[K, V]) get(key K) (V, bool) {
	c.mut.RLock()
	defer c.mut.RUnlock()

	var val V

	if v, ok := c.m[key]; ok {
		return v, ok
	}

	return val, false
}

func (c *ConcurrentMap[K, V]) set(key K, val V) {
	c.mut.Lock()
	c.m[key] = val
	c.mut.Unlock()
}

func (c *ConcurrentMap[K, V]) unset(key K) {
	c.mut.Lock()
	delete(c.m, key)
	c.mut.Unlock()
}
