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

func TestArray(t *testing.T){
	var myArray [3]string

	myArray[0] = "Hello"
	myArray[1] = "World"
	myArray[2] = "Go-Course"

	mySlice := myArray[0:1]

	fmt.Printf("%+v\n", myArray)
	fmt.Printf("%+v\n", mySlice)

	mySlice[0] = "Hola!"
	fmt.Printf("%+v\n", myArray)

	var mySlice2 []string = make([]string, 10)
	mySlice2[0] = "Roy"

	println(len(mySlice2))
	println(cap(mySlice2))

}

func TestMaps(t *testing.T){
	myMap := make(map[string]int)
	myMap["roy"] = 21
	myMap["haim"] = 70

	fmt.Printf("%+v\n", myMap)

	if _, ok := myMap["yossi"]; ok {
		fmt.Println("yossi found!")
	}

	for k, v := range myMap {
		fmt.Printf("%s: %d\n", k, v)
	}



}

func TestSlice(t *testing.T){
	arr := [4]int{1,2,3,4}
	fmt.Println(arr)

	slc := make([]int, 2, 4)
	fmt.Println(len(slc))
	fmt.Println(slc)

	slc2 := []int{1,2}
	slc3 := append(slc2, 3, 5)
	fmt.Println(slc3)

	slc3 = append(slc2, slc...)
	fmt.Println(slc3)
}
