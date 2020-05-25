package main 

import (
	"fmt"
)

func main() {
	var num1 int8 = 8
	var num2 float64 = 3.2
	answer := num1 % int8(num2)
	fmt.Printf("The answer is: %d\n", answer)
}