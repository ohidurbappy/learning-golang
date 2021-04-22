package main

import "fmt"

func main() {
	var name string
	var age int
	var c byte
	var marks float32

	count := 1000
	c = 'a'
	name = "Ohidur"
	age = 25
	marks = 89.9
	fmt.Println(name, age, marks, c, count)

	fmt.Printf("name is of type is %T\n", name)

	var x, y, z = 10, 11, 12
	fmt.Println(x, y, z)

	var im complex64
	fmt.Printf("im is of type %T\n", im)

}
