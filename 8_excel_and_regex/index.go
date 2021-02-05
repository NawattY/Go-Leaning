package main

import (
	"fmt"
	"math"
	"reflect"
	"regexp"

	"github.com/360EntSecGroup-Skylar/excelize"
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
	// c, ok := s.(circle)

	// if ok {
	// 	fmt.Println("Casting success.", c)
	// } else {
	// 	fmt.Println("Casting error.", ok)
	// }
}

func main() {
	xlsx, err := excelize.OpenFile("/Users/nawatty/Docker/www/ecs/cpm/cpm-api/storage/migrate/ecom.xlsx")
	if err != nil {
		fmt.Println("error2", err)
		return
	}

	rows, err2 := xlsx.GetRows("Sheet1")
	if err2 != nil {
		fmt.Println("error2", err2)
		return
	}

	var correctMobiles []string
	var duplicateMobiles map[string]int
	var nullMobiles int
	var wrongMobiles []string

	for _, row := range rows[1:] {
		mobile := row[5]

		if mobile == "NULL" || mobile == "" {
			nullMobiles++
			continue
		}

		if validateMobileFormat(mobile) {
			wrongMobiles = append(wrongMobiles, mobile)
			continue
		}

		_, existing := duplicateMobiles[mobile]
		if existing {
			duplicateMobiles[mobile]++
		}

		_, existing = correctMobiles[mobile]

		if existing {
			//duplicateMobiles = append(duplicateMobiles, [mobile: 2])
		} else {

		}
	}

	c := circle{redius: 30}
	// fmt.Println(getArea(c))

	s := squre{with: 30, heigh: 10}
	// fmt.Println(getArea(s))

	// showInfo(c)
	// showInfo(s)

	testHanldleCastingError(c)
	testHanldleCastingError(s)
}

func validateMobileFormat(m string) bool {
	var validdateMobile = regexp.MustCompile(`^([0][6|8|9][0-9]{8})+$`)
	return validdateMobile.MatchString(m)
}
