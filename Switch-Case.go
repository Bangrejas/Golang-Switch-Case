package main

import (
	"fmt"
)

func main() {
	var point = 10

	switch point {
	case 10:
		fmt.Println("Perfect")
	case 9:
		fmt.Println("Awesome")
	case 8:
		fmt.Println("Good Work")
	default:
		fmt.Println("Not Bad")
	}
}