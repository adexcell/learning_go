package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(d time.Duration) {
	timer := time.NewTimer(d * time.Second)
	<- timer.C
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			start := time.Now()
			fmt.Printf("start working #%d\n", i)
			sleep(1)
			end := time.Now()
			fmt.Printf("end working, duration = %v\n", end.Sub(start))
		}
	}()

	wg.Wait()
}
