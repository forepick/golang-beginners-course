package main

import (
	"fmt"
	"time"
)

func main() {

	// *********** HTTP Handler Mock ************
	handler := getHandler(3, 1)


	// ************ Request simulation **********
	requestTicker := time.Tick(200 * time.Millisecond)
	for range requestTicker{
		fmt.Printf("Got %d status\n", handler("request Body"))
	}


}

func getHandler(burstSize int, maxRps float64) func(request string) int{

	fillCycle := time.Duration(1 / maxRps) * time.Second
	var rateBucket = make(chan time.Time, burstSize)
	for i:=0; i<burstSize; i++ {
		rateBucket <- time.Now()
	}
	rateTick := time.Tick(fillCycle)
	go func(){
		for {
			rateBucket <- <- rateTick
		}
	}()

	return func(request string) int {
		select {
		case <-rateBucket:
			return 200
		default:
			return 429
		}
	}

}


