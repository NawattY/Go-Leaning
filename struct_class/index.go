package main

import (
	"fmt"
	"structClass/employee"
)

// run go mod init structClass ก่อน
// https://www.youtube.com/watch?v=bFeb64MSvvE&list=PLjPfp4Ph3gBrJ6jHPow7pZlOgMmfyQ7zK&index=35
func main() {
	d := employee.NewDeparture("Captain", 15)

	luffy := employee.New(
		"Monkey",
		"D Luffy",
		19,
		15000,
		14,
		d,
	)

	fmt.Println(luffy)
	luffy.LeavesRemining()

	// Pointer
	luffy.IncreseLeave()
	fmt.Println(luffy)
	luffy.LeavesRemining()
}
