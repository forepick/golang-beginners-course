package data_structures

import (
	"testing"
	"fmt"
	"time"
)
func BenchmarkSomething(b *testing.B){
	for i := 0; i < b.N; i++ {
		time.Sleep(1000)
	}
}

func TestSlice(t *testing.T){
	arr := [4]int{1,2,3,4}
	fmt.Println(arr)

	slc := make([]int, 2, 4)
	fmt.Println(len(slc))
	fmt.Println(slc)

	slc2 := []int{1,2}
	slc3 := append(slc2, 3,5)
	fmt.Println(slc3)

	slc3 = append(slc2, slc...)
	fmt.Println(slc3)
}
