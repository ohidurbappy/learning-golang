package main

import (
	"fmt"
)


func main() {
	a := 89

	if a < 90 {
		fmt.Println("You didn't get the highest mark")
	} else if a < 70 {
		fmt.Println("You haven't obtained letter marks")
	} else if a < 60 {
		fmt.Println("You couldn't get a B+")
	} else {
		fmt.Println("You have done bad!")
	}

}
