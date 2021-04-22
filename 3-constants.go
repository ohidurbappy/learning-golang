package main

import (
	"fmt"
)

func main() {
	const PI float32 = 3.14

	radius := 11
	circumference := 2 * PI * float32(radius)

	fmt.Println("The circumference is ", circumference)
}
