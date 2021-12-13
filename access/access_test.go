package access

import (
	"Go_Course/access/geo"
	"fmt"
	"testing"
)

func PrintArea(shape geo.Shape) {
	if shape == nil {
		return
	}
	fmt.Printf("Area is %f\n", shape.GetArea())
}

func TestAccess(t *testing.T){

	var r *geo.Rectangle

	r = &geo.Rectangle{
		Width: 3,
		Height: 4,
	}

	r = geo.NewRect(3, 4, "rect 2")

	PrintArea(r)
}
