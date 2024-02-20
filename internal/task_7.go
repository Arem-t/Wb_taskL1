package internal

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	data  map[string]int
	mutex sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[string]int)}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	value, ok := sm.data[key]
	return value, ok
}

func Task7() {
	safeMap := NewSafeMap()

	// Пример использования
	safeMap.Set("key1", 10)
	safeMap.Set("key2", 20)

	value1, ok1 := safeMap.Get("key1")
	value2, ok2 := safeMap.Get("key2")
	value3, ok3 := safeMap.Get("key3")

	fmt.Println("Value1:", value1, "Key1 exists?", ok1)
	fmt.Println("Value2:", value2, "Key2 exists?", ok2)
	fmt.Println("Value3:", value3, "Key3 exists?", ok3)
}
