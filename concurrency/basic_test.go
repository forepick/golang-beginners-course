package concurrency

import (
	"testing"
	"fmt"
	"time"
)

func TestBasicConcurrency(t *testing.T){
	//messages := make(chan string, 10)

	/*
	go func() {
		fmt.Println("goroutine sleeping...")
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine woke up")
		messages <- "ping"
	}()*/

	cn := make(chan string, 10)

	for i := 0; i < 10; i++ {
		go func(i int, input chan string){
			for {
				fmt.Printf("%d: %s\n", i, <-input)
			}
		}(i, cn)
	}

	for i := 0; i < 10; i++ {
		cn <- "Hi"
	}

	time.Sleep(2 * time.Second)

/*
	fmt.Println("main sleeping...")
	time.Sleep(2 * time.Second)

	msg := <-messages
	fmt.Printf("main woke up: %s\n", msg)
*/
}
