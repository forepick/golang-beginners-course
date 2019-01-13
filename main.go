package main

import (
	"time"
	"fmt"
)

func main() {

	done := make(chan bool)

	go func(complete chan bool){
		for i := 0; i < 10; i++ {
			println(i)
			time.Sleep(500 * time.Millisecond)
		}

		complete <- true

	}(done)

	<- done

	fmt.Println("Task completed")


}
