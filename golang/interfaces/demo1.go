package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, heigth float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.heigth * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("Şekil Alanı : ", s.area())
}

func Demo1() {
	r := rectangle{width: 10, heigth: 6}
	school(r)

	c := circle{radius: 10}
	school(c)
}
