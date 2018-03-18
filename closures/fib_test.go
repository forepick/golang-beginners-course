package closures

import (
	"testing"
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x2 := 1
	x1 := 0
	x := 0
	return func() int{
		ret := x
		x = x1 + x2
		x2 = x1
		x1 = x
		return ret
	}
}

func TestFib(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}