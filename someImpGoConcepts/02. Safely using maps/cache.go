package main

import (
	"crypto/sha1"
	"log"
	"sync"
)

// Maps in go are not concurrent safe
// maps are not safe to read and write from multiple go routines

// To make go apps more concurrent safe using sync.mutex and

// If more requests come in, to avoid lock contention we implement shard. This will avoid processes not to wait much on locks to get freed

type Shard struct {
	sync.RWMutex
	data map[string]any
}

type ShardMap []*Shard

func NewShardMap(n int) ShardMap {
	shards := make([]*Shard, n)
	for i := range n {
		shards[i] = &Shard{
			data: make(map[string]any),
		}
	}
	return shards
}

func (m ShardMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[0])
	i := hash % len(m)
	log.Printf("key: %v, index: %v", key, i)
	return i
}

func (m ShardMap) getShard(key string) *Shard {
	i := m.getShardIndex(key)
	return m[i]
}

func (m ShardMap) Get(key string) (any, bool) {
	shard := m.getShard(key)

	shard.RLock() // Make this a read lock. And Locking a shard instead of whole map
	defer shard.RUnlock()
	val := shard.data[key]
	return val, val != nil
}

func (m ShardMap) Set(key string, val any) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = val
}

func (m ShardMap) Delete(key string) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	delete(shard.data, key)
}

func (m ShardMap) Contains(key string) bool {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val := shard.data[key]
	return val != nil
}

func (m ShardMap) Keys() []string {
	keys := make([]string, 0)

	mutex := sync.Mutex{}

	wg := sync.WaitGroup{}

	wg.Add(len(m))

	for _, shard := range m {
		go func(s *Shard) {
			s.RLock()
			for k := range s.data {
				mutex.Lock()
				keys = append(keys, k)
				mutex.Unlock()
			}
			s.RUnlock()
			wg.Done()
		}(shard)
	}
	wg.Wait()

	return keys
}

func RunCacheExample() {
	sm := NewShardMap(3)

	sm.Set("a", 1)
	sm.Set("b", 2)
	sm.Set("c", 3)

	keys := sm.Keys()
	for k := range keys {
		log.Printf("Key: %v", k)
	}

	a, _ := sm.Get("a")
	log.Printf("a: %v", a)

	b, _ := sm.Get("b")
	log.Printf("b: %v", b)

	z, _ := sm.Get("z")
	log.Printf("z: %v", z)

	sm.Delete("a")
	sm.Delete("z")

	a, exists := sm.Get("a")
	log.Printf("a: %v, exists: %v", a, exists)

	keys = sm.Keys()

	for k := range keys {
		log.Printf("Key: %v", k)
	}
}
