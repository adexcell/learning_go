package main

import (
	"fmt"
	"sync"
	"time"
)

func or[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	outputCh := make(chan T)
	for _, channel := range channels {
		go func() {
			defer wg.Done()
			select {
			case <-outputCh:
				return
			case value, ok := <-channel:
				if !ok {
					outputCh <-value
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
