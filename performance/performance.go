package main

import (
	"time"
	"fmt"
	"runtime"
)


func main(){
	N := 10000000

	fmt.Println(runtime.NumCPU())

	start := time.Now()

	implementation := 4

	switch implementation {
	case 0:
		SerialImpl(N)
	case 1:
		SimpleConcurrencyImpl(N)
	case 2:
		SingleWorkerImpl(N)
	case 3:
		WorkersImpl(N)
	case 4:
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

func SingleWorkerImpl(N int){
	input := make(chan int, N)
	exit := make(chan bool)
	for i := 0; i < N; i++ {
		input <- i
	}
	go func(){
		for{
			select {
				case v := <- input:
					Task(v)
				default:
					exit <- true

			}
		}
	}()
	<- exit
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
	input := make(chan int, N)
	for i := 0; i < N; i++ {
		input <- i
	}
	fmt.Println("Done submitting tasks")

	WorkerPool(input, func(a int) {
		Task(a)
	}, func(){
		exit <- true
	})

	<-exit
}


func WorkerPool(input chan int, task func(int), callback func()) chan int {
	poolSize := 4

	ack := make(chan bool, poolSize)

	fmt.Printf("goroutines spawned: %d\n", poolSize)
	for i := 0; i < poolSize; i++ {
		go func() {
			for {
				select {
					case v := <- input:
						task(v)
					default:
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
		fmt.Println("All workers completed")
		callback()
	}()

	return input
}

func BatchingImpl(N int){
	var batchSize = 2500000
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




