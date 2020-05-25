package main

import "fmt"

func main() {
	var number = "2"
	num := 6
	y := true
	var x bool
	fmt.Println("%T", number)
	fmt.Printf("%T", num)
	fmt.Printf("%T", y)
	fmt.Println("%", x)
	x = true
	fmt.Println("%", x)
}