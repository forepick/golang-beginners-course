package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestBasicConcurrency(t *testing.T){

	cn := make(chan string, 10)

	for i := 0; i < 10; i++ {
		go func(index int, input <-chan string){
			//for{
				fmt.Printf("%d: %s\n", index, <-input)
			//}
		}(i, cn)
	}

	for i := 0; i < 22; i++ {
		cn <- fmt.Sprintf("Hi %d", i)
	}

	time.Sleep(2 * time.Second)

}
