package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleLimiter(t *testing.T) {

	tasks := make(chan string, 20)
	for i := 0; i < 20; i++ {
		tasks <- fmt.Sprintf("This is task %d", i)
	}
	close(tasks)

	rps := 2
	tick := time.Tick(time.Duration(1000 / rps) *  time.Millisecond)

	for task := range tasks {
		<- tick
		fmt.Printf("executing %s\n", task)
	}

}

func TestBurstLimiter(t *testing.T) {

	tasks := make(chan string, 20)
	for i := 0; i < 20; i++ {
		tasks <- fmt.Sprintf("This is task %d", i)
	}
	close(tasks)

	rps := 2
	tick := time.Tick(time.Duration(1000 / rps) *  time.Millisecond)

	burstLimit := 10
	bucket := make(chan time.Time, burstLimit)
	for i := 0; i < burstLimit; i++ {
		bucket <- time.Now()
	}

	go func(){
		for {
			bucket <- <- tick
		}
	}()

	for task := range tasks {
		<- bucket
		fmt.Printf("executing %s\n", task)
	}

}
