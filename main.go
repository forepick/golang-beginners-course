package main

import (
	"fmt"
)

func workerFunc(id int, input <-chan string, success chan<- bool){
	for task := range input {
		fmt.Printf("Worker %d started working...\n", id)
		//time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d completed \"%s\"\n", id, task)
		success <- true
	}
}
const NUM_OF_WORKERS = 10

func main(){
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