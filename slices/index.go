package main

import "fmt"

/* Slice จะมี capasity เพิ่มขึ้น ถ้าหากว่าเกินจาก capasity ปัจจุบัน
 * โดยจะเพิ่มเป็น (capasity เริ่มต้น) * 2^(n-1)
 */
func main() {
	var sNumbers = make([]int, 3, 5)
	showSliceInfo(sNumbers)

	sNumbers = append(sNumbers, 1)
	sNumbers = append(sNumbers, 1)
	showSliceInfo(sNumbers)

	sNumbers = append(sNumbers, 1)
	showSliceInfo(sNumbers)

	sNumbers = append(sNumbers, 1)
	sNumbers = append(sNumbers, 1)
	sNumbers = append(sNumbers, 1)
	sNumbers = append(sNumbers, 1)
	showSliceInfo(sNumbers)

	sNumbers = append(sNumbers, 1)
	showSliceInfo(sNumbers)

	for x := 1; x <= 29; x++ {
		sNumbers = append(sNumbers, 1)
	}
	showSliceInfo(sNumbers)

	sNumbers = append(sNumbers, 1)
	showSliceInfo(sNumbers)

	var sNumbers2 []int
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)

	sNumbers2 = append(sNumbers2, 1)
	showSliceInfo(sNumbers2)
}

func showSliceInfo(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	fmt.Println("------------------")
}
