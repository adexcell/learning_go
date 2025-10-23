package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numChan := make(chan int)
	doubleNumChan := make(chan int)

	go NumGenerator(ctx, nums, numChan)

	for range nums {
		wg.Add(2)
		go DoubleNum(ctx, wg, numChan, doubleNumChan)
		go PrintNum(ctx, wg, doubleNumChan)
	}

	go func() {
		<-stop
		fmt.Println("\nStopped")
		cancel()
	}()

	wg.Wait()

}

func NumGenerator(ctx context.Context, nums []int, ch chan<- int) {
	defer close(ch)
	for _, x := range nums {
		select {
		case <-ctx.Done():
			return
		case ch <- x:
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func DoubleNum(ctx context.Context, wg *sync.WaitGroup, numChan <-chan int, doubleNumChan chan<- int) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	case res, ok := <-numChan:
		if ok {
			doubleNumChan <- res * 2
			time.Sleep(time.Millisecond * 350)
		}
	}
}

func PrintNum(ctx context.Context, wg *sync.WaitGroup, doubleNumChan chan int) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	case res := <-doubleNumChan:
		fmt.Println(res)
	}
}
