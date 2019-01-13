package closures

import (
	"fmt"
	"testing"
)


func getStepper(init int) func() int {
	return func() int {
		init++
		return init-1
	}
}

func TestStepper(t *testing.T){
	stepper1 := getStepper(6)
	stepper2 := getStepper(0)
	fmt.Println(stepper1())
	fmt.Println(stepper1())
	fmt.Println(stepper1())

	fmt.Println(stepper2())
	fmt.Println(stepper2())
	fmt.Println(stepper2())

	for i := 0; i < 10 ; i++ {

	}
}




