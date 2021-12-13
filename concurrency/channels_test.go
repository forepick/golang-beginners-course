package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestChannels(t *testing.T){

	c := make(chan int, 0)

	go func(cn *chan int) {
		for{
			time.Sleep(3 * time.Second)
			fmt.Println("pulling message")

			a := <- *cn
			fmt.Println(a)
		}
	}(&c)

	fmt.Println("pushing message...")

	c <- 1

}

func TestChannelIteration(t *testing.T){
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

	close(c)

	for val := range c {
		fmt.Printf("val is %d\n", val)
	}

}

func TestSelect(t *testing.T){
	c1 := make(chan int, 3)
	c2 := make(chan int, 3)

	defaultTick := time.Tick(500 * time.Millisecond)

	go func(){
		for {
			select {
				case a := <- c1:
					fmt.Printf("c1 gave %d\n", a)
				case a := <- c2:
					fmt.Printf("c2 gave %d\n", a)
				case <- defaultTick:
					fmt.Println("No values yet...")
			}
		}
	}()


	time.Sleep(2600 * time.Millisecond)
	c2 <- 2
	c1 <- 30
	time.Sleep(1 * time.Second)
}



func TestStopRoutine(t *testing.T){
	tasks := make(chan string, 10)
	stop := make(chan bool)

	go func(input chan string){
		for  {
			task := <- input
			<- stop
			fmt.Println(task)
		}
	}(tasks)
}


func TestTicker(t *testing.T){
	tick := time.Tick(1 * time.Second)

	go func(){
		for tm := range tick{
			fmt.Printf("time is: %s\n", tm.Format(time.Stamp))
		}
	}()

	time.Sleep(20 * time.Second)
}





func TestWorkers(t *testing.T){

	const NUM_OF_WORKERS = 10

	workerFunc := func(id int, input <-chan string, success chan<- bool){
		for task := range input {
			fmt.Printf("Worker %d started working...\n", id)
			//time.Sleep(1 * time.Second)
			fmt.Printf("Worker %d completed \"%s\"\n", id, task)
			success <- true
		}
	}

	tasks := make(chan string, 10)
	acks := make(chan bool)

	for i := 0; i < NUM_OF_WORKERS; i++ {
		go workerFunc(i, tasks, acks)
	}

	for i := 0; i < NUM_OF_WORKERS; i++ {
		tasks <- fmt.Sprintf("Task #%d", i)
	}
	for i := 0; i < NUM_OF_WORKERS; i++ {
		<- acks
	}

	fmt.Println("execution completed")
}


func TestWorkerControl(t *testing.T){
	worker := func(input <-chan string, termination <-chan bool){
		for {
			select {
				case <- termination:
					fmt.Println("Worker stopped")
					return
				case task := <- input:
					fmt.Printf("Got task %s\n", task)

			}
		}
	}

	tasks := make(chan string, 10)
	stop := make(chan bool)

	go worker(tasks, stop)

	// Driver
	go func(){
		for tm := range time.Tick(500 * time.Millisecond){
			tasks <- fmt.Sprintf("Task of %d", tm.UnixMilli())
		}
	}()

	// Control
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(1 * time.Second)
	fmt.Println("execution completed")
}

