package main

import "fmt"

type MyType struct {
	Field1 string
	Field2 int
}

func main(){
	mt := MyType{}

	fmt.Printf("%s\n", mt)
	fmt.Printf("%+v\n", mt)
	fmt.Printf("%#v\n", mt)
	fmt.Printf("%T\n", mt)
	fmt.Printf("%s", false)
}
