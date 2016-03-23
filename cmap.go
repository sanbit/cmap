package cmap

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type ConcurrentMap struct {
	buckets []map[string]interface{}
	locks   []sync.RWMutex
}

func NewConcurrentMap(bucketSize int) *ConcurrentMap {
	m := &ConcurrentMap{
		buckets: make([]map[string]interface{}, bucketSize),
		locks:   make([]sync.RWMutex, bucketSize),
	}

	for i := range m.buckets {
		m.buckets[i] = make(map[string]interface{})
	}
	return m
}

func (m *ConcurrentMap) Len() int {
	length := 0
	for i := range m.locks {
		m.locks[i].RLock()
		length += len(m.buckets[i])
		m.locks[i].RUnlock()
	}
	return length
}

func (m *ConcurrentMap) Get(key string) interface{} {

	slot := m.hash(key) % len(m.buckets)
	m.locks[slot].RLock()
	value := m.buckets[slot][key]
	m.locks[slot].RUnlock()
	return value
}

func (m *ConcurrentMap) Set(key string, value interface{}) {
	slot := m.hash(key) % len(m.buckets)

	//fmt.Println(slot)

	m.locks[slot].Lock()
	m.buckets[slot][key] = value
	m.locks[slot].Unlock()
}

func (m *ConcurrentMap) Delete(key string) {
	slot := m.hash(key) % len(m.buckets)
	m.locks[slot].Lock()
	delete(m.buckets[slot], key)
	m.locks[slot].Unlock()
}

func (m *ConcurrentMap) hash(key string) int {
	hasher := fnv.New32()
	hasher.Write([]byte(key))

	return int(hasher.Sum32())
}

func main() {
	m := NewConcurrentMap(10)

	m.Set("aaa", 9)
	m.Set("bbb", 9)
	m.Set("ccc", 9)
	fmt.Println(m.Len())
}
