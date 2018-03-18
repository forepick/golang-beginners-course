package concurrency

import (
	"testing"
	"time"
	"fmt"
)

func TestBasicRoutine(t *testing.T) {

	for gr := 0; gr < 10; gr++ {
		go func(taskType int) {
			i := 0
			for {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Task type %d: %d executed\n", taskType, i)
				i++
			}
		}(gr)
	}


	time.Sleep(20 * time.Second)


}
