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


func TestLimiterMock(t *testing.T){

	tasks := make(chan string, 20)

	tick := time.Tick(500 * time.Millisecond)

	// Server
	go func(input chan string){
		for {
			task := <- input
			select {
				case <- tick:
					fmt.Println(task)
				default:
					fmt.Println("429 - Rate Limit Exceeded")
			}
		}
	}(tasks)

	// Client (Driver)
	for i := 0; i < 20; i++ {
		tasks <- fmt.Sprintf("This is task %d", i)
		time.Sleep(300 * time.Millisecond)
	}

	//time.Sleep(20 * time.Second)
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
