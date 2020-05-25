package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please type the year you were born\n")
	scanner.Scan()
	input_yob, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println("Please type the current year\n")
	scanner.Scan()
	current_year, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You are currently %d years old\n", current_year-input_yob)
}