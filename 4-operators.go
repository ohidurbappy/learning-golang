package main

import "fmt"

func main() {

	// arithmetic operators
	a, b := 10, 20

	sum := a + b

	fmt.Println("Addition ", sum)

	result := sum - 7

	fmt.Println("Subtraction ", result)

	result = 9 * 3

	fmt.Println("Multiplication ", result)

	result = 69 / 3

	fmt.Println("Division ", result)

	result = 23 % 7

	fmt.Println("Modulas ", result)

	result++

	fmt.Println("Increment ", result)

	result--

	fmt.Println("Decrement ", result)

	// relational
	a = 10
	b = 20

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)

	// logical operator
	x := true
	y := false
	fmt.Println(x && y)
	fmt.Println(x || y)
	fmt.Println(!x)

	// bitwise operators

	d := 10
	e := 11
	fmt.Println(d & e)  //binary AND
	fmt.Println(d | e)  //binary OR
	fmt.Println(d ^ e)  //binary XOR
	fmt.Println(d << 2) //left shift
	fmt.Println(d >> 1) // right shift

	// assignment operator
	num1 := 21
	num2 := 22

	n3 := num1 + num2
	n3 += 10
	n3 -= 2
	n3 *= 1
	n3 /= 2
	n3 %= 5
	n3 <<= 2
	n3 >>= 1
	n3 &= 9
	n3 ^= 1
	n3 |= 4
	fmt.Println(n3)

}
