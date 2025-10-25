package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterMutex struct {
	wg sync.WaitGroup
	mu sync.Mutex
	counter int
}

func (c *CounterMutex) Inc() {
	defer c.wg.Done()
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *CounterMutex) Value() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("counter on mutex", c.counter)
}

type CounterAtomic struct {
	wg sync.WaitGroup
	counter int64
}

func (c *CounterAtomic) Inc() {
	defer c.wg.Done()
	atomic.AddInt64(&c.counter, 1)
}

func (c *CounterAtomic) Value() {
	fmt.Println("counter on atomic", atomic.LoadInt64(&c.counter))
}

func main() {
	counterMutex := CounterMutex{counter: 0}
	for range 100 {
		counterMutex.wg.Add(1)
		go counterMutex.Inc()
	}
	counterMutex.wg.Wait()
	counterMutex.Value()

	counterAtomic := CounterAtomic{counter: 0}
	for range 100 {
		counterAtomic.wg.Add(1)
		go counterAtomic.Inc()
	}
	counterAtomic.wg.Wait()
	counterAtomic.Value()
}