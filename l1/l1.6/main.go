package main

import (
	"context"
	"fmt"
	// "log"
	"os"
	"os/signal"

	// "runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	intChan := make(chan int)

	fmt.Println("program start")

	// writer
	go func(ch chan<- int) {
		defer close(ch)
		for i := 0; ; i++ {
			ch <- i
			fmt.Println("write", i)
			time.Sleep(200 * time.Millisecond)
		}
	}(intChan)


	// №1. Выход из горутины по условию. Условие - если читаемое из канала значение четное,
	//  горутина прекращает свою работу.
	// for i := 0; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go func(id int, ch <-chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		res := <-ch
	// 		if res % 2 == 0 {
	// 			fmt.Printf("task %d interrupted\n", id)
	// 			return
	// 		}
	// 		fmt.Printf("task %d read %d\n", id, res)
	// 	}(i, intChan, wg)
	// 	time.Sleep(150 * time.Millisecond)
	// }


	// №2. Выход из горутины через канал уведомления.
	// stop := make(chan struct{})

	// for i := 0; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go func(id int, ch <-chan int, wg *sync.WaitGroup, stop chan struct{}) {
	// 		defer wg.Done()
	// 		select {
	// 		case <-stop:
	// 			fmt.Printf("task %d interrupted by channel\n", id)
	// 			return
	// 		case res := <-ch:
	// 			fmt.Printf("task %d read %d\n", id, res)
	// 		}
	// 	}(i, intChan, wg, stop)
	// 	time.Sleep(150 * time.Millisecond)
	// }

	// time.Sleep(time.Millisecond * 100)
	// close(stop)


	// №3. Выход из горутины c помощью контекста.
	// №3.1 context.WithCancel - программа ожидает сигнала отмены из консоли (CTRL + C)

	// ctx, cancel := context.WithCancel(context.Background())

	// for i := 0; i <= 100; i++ {
	// 	wg.Add(1)
	// 	go func(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Printf("task %d interrupted by cancel\n", id)
	// 			return
	// 		case res := <-ch:
	// 			fmt.Printf("task %d read %d\n", id, res)
	// 		}
	// 	}(ctx, i, intChan, wg)
	// 	time.Sleep(150 * time.Millisecond)
	// }


	// №3.2 context.WithTimeout - горутины завершатся спустя заданное время,
	// программа ожидает сигнала отмены из консоли (CTRL + C)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)

	// for i := 0; i <= 100; i++ {
	// 	wg.Add(1)
	// 	go func(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Printf("task %d interrupted by cancel\n", id)
	// 			return
	// 		case res := <-ch:
	// 			fmt.Printf("task %d read %d\n", id, res)
	// 		}
	// 	}(ctx, i, intChan, wg)
	// 	time.Sleep(150 * time.Millisecond)
	// }


	// №3.3 context.WithDeadline - горутины завершатся когда наступит deadline,
	// программа ожидает сигнала отмены из консоли (CTRL + C)

	// deadline := time.Now().Add(3 * time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), deadline)
	// for i := 0; i <= 100; i++ {
	// 	wg.Add(1)
	// 	go func(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Printf("task %d interrupted by cancel\n", id)
	// 			return
	// 		case res := <-ch:
	// 			fmt.Printf("task %d read %d\n", id, res)
	// 		}
	// 	}(ctx, i, intChan, wg)
	// 	time.Sleep(150 * time.Millisecond)
	// }


	// №4 прекращение работы горутины runtime.Goexit()
	// ctx, cancel := context.WithCancel(context.Background())

	// for i := 0; i <= 100; i++ {
	// 	wg.Add(1)
	// 	go func(id int, ch <-chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Printf("task %d interrupted by cancel\n", id)
	// 			return
	// 		case res := <-ch:
	// 			if res > 10 {
	// 				fmt.Printf("task %d terminates by runtime.Goexit()\n", id)
	// 				runtime.Goexit()
	// 			}
	// 			fmt.Printf("task %d read %d\n", id, res)
	// 		}
	// 	}(i, intChan, wg)
	// 	time.Sleep(150 * time.Millisecond)
	// }


	// №5 прекращение всей программы os.Exit() - это не остановит конкретную горутину,
	// но остановит всю программу
	// ctx, cancel := context.WithCancel(context.Background())

	// for i := 0; i <= 100; i++ {
	// 	wg.Add(1)
	// 	go func(id int, ch <-chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Printf("task %d interrupted by cancel\n", id)
	// 			return
	// 		case res := <-ch:
	// 			if res > 10 {
	// 				fmt.Printf("task %d terminates by os.Exit()\n", id)
	// 				os.Exit(1)
	// 			}
	// 			fmt.Printf("task %d read %d\n", id, res)
	// 		}
	// 	}(i, intChan, wg)
	// 	time.Sleep(150 * time.Millisecond)
	// }

	
	// №6 log.Fatalf, log.Fatal - они тоже не останавливают конкретную горутину,
	// а останавливают всю программу целиком
	// ctx, cancel := context.WithCancel(context.Background())

	// for i := 0; i <= 100; i++ {
	// 	wg.Add(1)
	// 	go func(id int, ch <-chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Printf("task %d interrupted by cancel\n", id)
	// 			return
	// 		case res := <-ch:
	// 			if res > 10 {
	// 				log.Fatalf("task %d terminates by log.Fatalf\n", id)
	// 			}
	// 			fmt.Printf("task %d read %d\n", id, res)
	// 		}
	// 	}(i, intChan, wg)
	// 	time.Sleep(150 * time.Millisecond)
	// }

	
	// №7 panic + defer с recover - завершает горутину.
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(id int, ch <-chan int, wg *sync.WaitGroup) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Паника поймана, горутина завершена:", r)
				}
			}()
			defer wg.Done()
			select {
			case <-ctx.Done():
				fmt.Printf("task %d interrupted by cancel\n", id)
				return
			case res := <-ch:
				if res > 10 {
					panic("Принудительное завершение")
				}
				fmt.Printf("task %d read %d\n", id, res)
			}
		}(i, intChan, wg)
		time.Sleep(150 * time.Millisecond)
	}


	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	<-sigChan
	cancel()

	wg.Wait()

	fmt.Println("program stopped")
}
