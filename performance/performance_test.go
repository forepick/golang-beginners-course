package performance

import (
	"testing"
	"time"
	"fmt"
)


func TestPerformance(t *testing.T){
	N := 10000000
	start := time.Now()

	implementation := 0

	switch implementation {
	case 0:
		SerialImpl(N)
	case 1:
		SimpleConcurrencyImpl(N)
	case 2:
		WorkersImpl(N)
	case 3:
		BatchingImpl(N)
	}

	end := time.Since(start).Nanoseconds() / 1000000
	fmt.Println(end)

}

func Task(i int) {
	s := fmt.Sprintf("Task %d", i)
	s = s
}

func SerialImpl(N int){
	for i := 0; i < N; i++ {
		Task(i)
	}
}

func SimpleConcurrencyImpl(N int){

	ack := make(chan bool, N)
	for i := 0; i < N; i++ {
		go func(arg int) {
			Task(arg)
			ack <- true
		}(i)
	}

	for i := 0; i < N; i++ {
		<-ack                     // Point #2
	}
}
func WorkersImpl(N int){

	exit := make(chan bool)
	workers := WorkerPool(func(a int) {
		Task(a)
	}, func(){
		exit <- true
	})
	for i := 0; i < N; i++ {
		workers <- i
	}

}


func WorkerPool(task func(int), callback func()) chan int {
	poolSize := 100

	ack := make(chan bool, poolSize)

	fmt.Printf("goroutines spawned: %d\n", poolSize)
	input := make(chan int)
	for i := 0; i < poolSize; i++ {
		go func() {
			for {
				v, ok := <-input
				if ok {
					task(v)
				} else {
					ack <- true
					return
				}
			}
		}()
	}

	go func(){
		for i := 0; i < poolSize; i++ {
			<-ack
		}
		callback()
	}()

	return input
}

func BatchingImpl(N int){
	var batchSize = 100000
	var WorkersCount = N / batchSize
	fmt.Printf("goroutines spawned: %d\n", WorkersCount)
	var ack = make(chan bool, WorkersCount)
	j := 0
	for i := 0; i < WorkersCount; i++ {
		go func(start, batch int){
			for t := start; t < batch + start; t++ {
				Task(i)
			}
			ack <- true
		}(j, batchSize)
		j += batchSize
	}

	for i := 0; i < WorkersCount; i++ {
		<- ack
	}
}




