package geo

type (
	Shape interface {
		GetArea() float32
		GetCircumference() float32
		GetName() string
	}

	Rectangle struct {
		Height int
		Width int
		name string
	}
)

const SupportedShapes = 2

func NewRect(width, height int, name string) *Rectangle {
	r := &Rectangle{Width: width, Height: height, name: name}
	return r
}
func (r *Rectangle) GetName() string {
	return r.name
}
func (r *Rectangle) GetCircumference() float32 {
	return float32(2 * r.Height + 2 * r.Width)
}
func (r *Rectangle) GetArea() float32 {
	return float32(r.Height * r.Width)
}

