package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var workersCount int
	mainChan := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	flag.IntVar(&workersCount, "w", 10, "enter workers count")
	flag.Parse()

	go producer(ctx, mainChan)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg, mainChan)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	<-sigChan
	cancel()

	wg.Wait()
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		log.Printf("Worker-%d was interrupted\n", id)
	case task, ok := <-ch:
		if !ok {
			return
		}
		fmt.Printf("Worker-%d printed \"%d\"\n", id, task)
	}

}

func producer(ctx context.Context, ch chan<- int) {
	defer close(ch)
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
			time.Sleep(time.Millisecond * 300)
		}
	}
}
