package structsmethodsinterface

import "math"

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base, height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	result := r.width * r.height
	return math.Round(result*100) / 100
}

func (c Circle) Area() float64 {
	result := math.Pi * c.radius * c.radius
	return math.Round(result*100) / 100
}

func (t Triangle) Area() float64 {
	result := (t.base * t.height) / 2
	return math.Round(result*100) / 100
}
