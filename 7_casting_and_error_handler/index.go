package main

import (
	"fmt"
	"math"
	"reflect"
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

func showInfo(s shape) {
	t := reflect.TypeOf(s).Name()

	switch t {
	case "circle":
		c := s.(circle) // Casting เพราะไม่สามารถเข้าถึง property redius ของ circle ผ่าน s shape แบบ s.redius ได้
		fmt.Println(c.redius)
		break
	case "squre":
		sq := s.(squre) // Casting
		fmt.Println(sq.with, sq.heigh)
		break
	}
}

func testHanldleCastingError(s shape) {
	c, ok := s.(circle)

	if ok {
		fmt.Println("Casting success.", c)
	} else {
		fmt.Println("Casting error.", ok)
	}
}

func main() {
	c := circle{redius: 30}
	fmt.Println(getArea(c))

	s := squre{with: 30, heigh: 10}
	fmt.Println(getArea(s))

	showInfo(c)
	showInfo(s)

	testHanldleCastingError(c)
	testHanldleCastingError(s)
}
