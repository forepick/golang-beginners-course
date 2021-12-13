package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T){

	x := 0

	workerFunction := func(index int){
		for {
			fmt.Printf("routine %d counts %d\n", index, x)
			x++
			time.Sleep(10 * time.Millisecond)
		}
	}

	for i := 0; i < 10; i++ {
		go workerFunction(i)
	}

	time.Sleep(1 * time.Second)
}

func TestBasicRoutine(t *testing.T) {

	for gr := 0; gr < 10; gr++ {
		go func(taskType int) {
			i := 0
			for {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Task type %d. Execution number: %d\n", taskType, i)
				i++
			}
		}(gr)
	}

	fmt.Println("workers are ready!")

	time.Sleep(20 * time.Second)


}
func TestWaitGroup(t *testing.T) {

	var worker = func(id int){
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}

	var wg sync.WaitGroup

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(index int) {
			defer wg.Done()
			worker(index)
		}(i)
	}
	fmt.Println("Waiting for all workers to complete...")

	wg.Wait()
}