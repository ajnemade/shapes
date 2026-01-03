package shapes

type Shape interface {
	GetArea() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) GetArea() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) GetArea() float64 {
	return r.Length * r.Width
}
