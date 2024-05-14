package main

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Perimeter() float64 {
	return c.radius * 3.14
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}
