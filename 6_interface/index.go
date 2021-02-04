package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func getArea(s shape) float64 {
	return s.area()
}

type circle struct {
	redius int
}

type squre struct {
	with  int
	heigh int
}

// ถ้า struct มีฟังก์ชันตาม interface มันจะรู้อัตโนมัติว่า implement จาก interface อะไร
func (c circle) area() float64 {
	return math.Pi * float64(c.redius) * float64(c.redius)
}

func (s squre) area() float64 {
	return float64(s.heigh) * float64(s.with)
}

func main() {
	c := circle{redius: 30}
	fmt.Println(getArea(c))

	s := squre{with: 30, heigh: 10}
	fmt.Println(getArea(s))
}
