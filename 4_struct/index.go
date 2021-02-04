package main

import "fmt"

func main() {
	var dog animal
	dog.name = "Luffy"
	dog.legs = 4
	dog.color = "black"
	show(dog)

	// Not Pointer
	changName(dog, "Sabo")
	show(dog)

	// Not Pointer
	changName2(dog, "Sabo")
	show(dog)

	// Not Pointer
	dog = changName2(dog, "Sabo")
	show(dog)

	// Pointer
	pChangName(&dog, "Ace")
	show(dog)

	// Call struct method
	dog.run()
	dName := dog.getName()
	fmt.Printf("The dog name is %s.\n", dName)

	// เปลี่ยนชื่อแบบไม่ใช้ Pointer หรือ การ return
	dog.setName("Nami")
	show(dog)

	// เปลี่ยนชื่อแบบใช้ Pointer
	dog.setName2("Nami")
	show(dog)

	// เปลี่ยนชื่อแบบใช้ return
	dog = dog.setName3("Jinbe")
	show(dog)
}

type animal struct {
	name  string
	legs  int
	color string
}

func show(a animal) {
	fmt.Println(a)
	fmt.Printf("Name=%s Color=%s Leg(s)=%d\n", a.name, a.color, a.legs)
	fmt.Println("----------------------")
}

// Not Pointer, No Return
func changName(a animal, name string) {
	a.name = name
}

// Not Pointer
func changName2(a animal, name string) animal {
	a.name = name
	return a
}

// Pointer = a *animal
func pChangName(a *animal, name string) {
	a.name = name
}

// Struct method no return
func (a animal) run() {
	fmt.Printf("%s is running...\n", a.name)
}

// Struct method
func (a animal) getName() string {
	return a.name
}

// Struct method
func (a animal) setName(name string) {
	a.name = name
}

// Struct method
func (a *animal) setName2(name string) {
	a.name = name
}

// Struct method
func (a animal) setName3(name string) animal {
	a.name = name
	return a
}
