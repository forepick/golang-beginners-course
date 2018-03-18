package exceptions

import (
	"fmt"
	"testing"
)


func TestDeferStacking(t *testing.T){
	for i := 0; i < 10; i ++ {
		defer fmt.Println(i)
	}
}

func TestRecovery(t *testing.T){

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from a panic: %s\n", r)
		}
	}()
	panic("Panic Example")
}
