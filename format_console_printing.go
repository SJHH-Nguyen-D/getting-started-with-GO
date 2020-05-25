package main

import "fmt"

func main() {
	fmt.Printf("Hello %T %v", 10, 10)
	a := 5
	fmt.Printf("\n")
	var mystring string = fmt.Sprintf("Hello %s", a)
	fmt.Println(mystring)
	var num float64 = 14.3459235938458
	fmt.Printf("\n%e\n", num)
	fmt.Printf("\n%f\n", num)
	fmt.Printf("\n%g\n", num)
	saying := "Oh how are you doing Ariel"
	fmt.Printf("\n%q\n", saying)
	fmt.Printf("\n%s\n", saying)
	fmt.Printf("\n%.2f\n", num)
	fmt.Printf("\n%3.5f\n", num)
	fmt.Printf("\nHe goes: %-9q is cool\n", saying)
	var out string = fmt.Sprintf("Number: %07d is cool", 45)
	fmt.Println(out)
}