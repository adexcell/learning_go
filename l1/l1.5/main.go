package main

import (
	"context"
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
	var timeout time.Duration
	mainChan := make(chan int, 1)
	timeout = time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	log.Println("start")

	go writer(ctx, mainChan)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go reader(ctx, mainChan, &wg, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <- sigChan:
		cancel()
	case <- ctx.Done():
		cancel()
	}

	wg.Wait()

}

func writer(ctx context.Context, ch chan int) {
	for i := 0; ; i++ {
		select {
		case ch <- i:
			fmt.Printf("передаем в канал значение %d\n", i)
			time.Sleep(time.Millisecond * 300)
		case <-ctx.Done():
			log.Println("ctx.done")
		}
	}
}

func reader(ctx context.Context, ch chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	var res int
	select {
	case res = <-ch:
		fmt.Printf("читаем из канала значение %d\n", res)
		time.Sleep(time.Millisecond * 300)
	case <-ctx.Done():
		log.Printf("ctx.done for reader  #%d\n", id)
	}

}
