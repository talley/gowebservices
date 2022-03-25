package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	var firstname *string = new(string)
	*firstname = "Abibou"

	fmt.Println(firstname)
}
func start() {
	fmt.Println("Hello World")
	var myint int = 34
	fmt.Println(myint)
	var myfloat float32 = 12.4
	fmt.Println(myfloat)

	name := "Abibou"
	fmt.Println(name)
	number := 100
	fmt.Println(number)

	cp := complex(3, 4)
	fmt.Println(cp)

	a, b := 3, 4
	fmt.Println(a, b)
}
