package main

import (
	"fmt"
	"sync"

	cmap "github.com/orcaman/concurrent-map"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.v[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func main() {

	// вариант с mutex
	wgMutex := &sync.WaitGroup{}
	c := SafeCounter{v: make(map[string]int)}
	key := "somekey"
	for i := 0; i < 1000; i++ {
		wgMutex.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			c.Inc(key)
		}(wgMutex)
	}


	// вариант с concurrent-map
	wgSyncMap := &sync.WaitGroup{}
	concurrentMap := cmap.New()
	concurrentMap.Set(key, 0)

	for i := 0; i < 1000; i++ {
		wgSyncMap.Add(1)
		go func(concurrent *cmap.ConcurrentMap, wg *sync.WaitGroup, key string) {
			defer wg.Done()
			concurrentMap.Upsert(key, 1, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
				if exist {
					return valueInMap.(int) + 1
				}
				return newValue
			})
		}(&concurrentMap, wgSyncMap, key)
	}

	wgMutex.Wait()
	wgSyncMap.Wait()
	value, _ := concurrentMap.Get(key)

	fmt.Println("value from map with mutex -", c.Value(key))
	fmt.Println("value from syncMap -", value)
}
