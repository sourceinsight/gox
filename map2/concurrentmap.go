package map2

import (
	"fmt"
	"sync"
)

const defaultShardCount = 64

// ConcurrentMap a thread-safe map.
type ConcurrentMap[K comparable, V any] struct {
	shardCount uint64
	locks      []sync.RWMutex
	maps       []map[K]V
}

// NewConcurrentMap create with a specific number of shards
func NewConcurrentMap[K comparable, V any](shardCount int) *ConcurrentMap[K, V] {
	if shardCount <= 0 {
		shardCount = defaultShardCount
	}

	cm := &ConcurrentMap[K, V]{
		shardCount: uint64(shardCount),
		locks:      make([]sync.RWMutex, shardCount),
		maps:       make([]map[K]V, shardCount),
	}

	for i := range cm.maps {
		cm.maps[i] = make(map[K]V)
	}

	return cm
}

// Set key-value pair
func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
	shard := cm.getShard(key)

	cm.locks[shard].Lock()
	cm.maps[shard][key] = value

	cm.locks[shard].Unlock()
}

// Get value by key
func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	value, ok := cm.maps[shard][key]
	cm.locks[shard].RUnlock()

	return value, ok
}

// GetOrSet get value by key, if not exists, set it
func (cm *ConcurrentMap[K, V]) GetOrSet(key K, value V) (actual V, ok bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	if actual, ok := cm.maps[shard][key]; ok {
		cm.locks[shard].RUnlock()
		return actual, ok
	}
	cm.locks[shard].RUnlock()

	cm.locks[shard].Lock()
	if actual, ok = cm.maps[shard][key]; ok {
		cm.locks[shard].Unlock()
		return
	}

	cm.maps[shard][key] = value
	cm.locks[shard].Unlock()

	return value, ok
}

// Delete key-value pair
func (cm *ConcurrentMap[K, V]) Delete(key K) {
	shard := cm.getShard(key)

	cm.locks[shard].Lock()
	delete(cm.maps[shard], key)
	cm.locks[shard].Unlock()
}

// GetAndDelete get value by key and delete it
func (cm *ConcurrentMap[K, V]) GetAndDelete(key K) (actual V, ok bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	if actual, ok = cm.maps[shard][key]; ok {
		cm.locks[shard].RUnlock()
		cm.Delete(key)
		return
	}
	cm.locks[shard].RUnlock()

	return actual, false
}

// Has checked if the key exists
func (cm *ConcurrentMap[K, V]) Has(key K) bool {
	_, ok := cm.Get(key)
	return ok
}

// Range range all key-value pairs
func (cm *ConcurrentMap[K, V]) Range(iterator func(key K, value V) bool) {
	for shard := range cm.locks {
		cm.locks[shard].RLock()

		for k, v := range cm.maps[shard] {
			if !iterator(k, v) {
				cm.locks[shard].RUnlock()
				return
			}
		}
		cm.locks[shard].RUnlock()
	}
}

// getShard get the shard index by key
func (cm *ConcurrentMap[K, V]) getShard(key K) uint64 {
	hash := fnv32(fmt.Sprintf("%v", key))
	return uint64(hash) % cm.shardCount
}

// fnv32 hash function
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
