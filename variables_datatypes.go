package main

import "fmt"

func main() {
	var something string;
	var number uint8 = 255;
	number = number + 5
	something = "Hello World!";
	fmt.Println(something, number);
}

// The print in the terminal is 4 because of uint8 value overflow