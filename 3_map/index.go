package main

import "fmt"

func main() {
	var mNumber = map[string]int{}
	fmt.Println(mNumber)

	mNumber["a"] = 1
	mNumber["a"] = 2
	mNumber["b"] = 3
	// mNumber[1] = 1 -> key ต้องเป็น string
	// mNumber["b"] = "3" -> value ต้องเป็น int
	fmt.Println(mNumber)

	// แบบหลายมิติ
	var mMultiDimension = make(map[string]map[int]string)
	fmt.Println(mMultiDimension)

	// mMultiDimension["test"][30] = "Test:30" -> error เพราะต้องสร้างค่า mMultiDimension["test"] ก่อน
	mMultiDimension["test"] = map[int]string{29: "Test:29"}
	mMultiDimension["test"][30] = "Test:30"
	fmt.Println(mMultiDimension)
	fmt.Println(mMultiDimension["test"][30])
	fmt.Println("----------------------------------------------------------")
	fmt.Println(mMultiDimension["test"][28]) // ไม่ error แต่ return nil
	fmt.Println("----------------------------------------------------------")

	val, isExist := mMultiDimension["test"][28]
	if isExist {
		fmt.Printf("mMultiDimension[\"test\"][28] was existed = %s\n", val)
	} else {
		fmt.Println("mMultiDimension[\"test\"][28] not exist")
	}

	val, isExist = mMultiDimension["test"][30]
	if isExist {
		fmt.Printf("mMultiDimension[\"test\"][30] was existed = %s\n", val)
	} else {
		fmt.Println("mMultiDimension[\"test\"][30] not exist")
	}
}
