package concurrency

import (
	"testing"
	"fmt"
	"time"
)

func TestChannels(t *testing.T){

	c := make(chan int, 1)

	go func(cn *chan int) {
		for{
			a := <- *cn
			fmt.Println(a)
		}
	}(&c)

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

	go func(){
		select{
		case a := <- c1:
			fmt.Printf("c1 gave %d\n", a)
		case a := <- c2:
			fmt.Printf("c2 gave %d\n", a)
		}
	}()

	time.Sleep(3 * time.Second)
	c2 <- 2
}

func TestTicker(t *testing.T){
	tick := time.Tick(1 * time.Second)

	go func(){
		for tm := range tick{
			fmt.Printf("time is: %s\n", tm.Format(time.RFC822))
		}
	}()

	time.Sleep(20 * time.Second)


}
