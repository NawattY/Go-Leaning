package main

import (
	"fmt"
)

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	var array2 []string = []string{"one", "two", "tree"}
	var array3 [3]bool

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	// array1[5] = 6 -> panic: runtime error: index out of range [5] with length 5
	// ไม่สามารถเพิ่มเข้าไปเกินกว่าตั้งต้นได้ ต้องไปใช้ slices
	for _, v := range array1 {
		fmt.Println(v)
	}

	for _, v := range array2 {
		fmt.Println(v)
	}

	array3[2] = true
	// array3[3] = true -> panic: runtime error: index out of range [3] with length 3
	for _, v := range array3 {
		fmt.Println(v)
	}
}
