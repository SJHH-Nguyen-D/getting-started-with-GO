package main 

import (
	"fmt"
)

func main() {
	val := 5 < 7 || false
	fmt.Printf("%t ", val)

	for x:=0; x<5; x++ {
		fmt.Println(x)
	}

	y:=10
	switch y {
	case 0:
		fmt.Println("Zero")
	case 10:
		fmt.Println("You did it")
	default:
		fmt.Println("This you're not there yet")
	}
}