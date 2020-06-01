package main

import ("fmt")

func main() {
	// when we specify an array, we have to specify it with the var keyword, the fixed number of elemens to be created for an array and the types of elements to be included in the array. This is fixed. Arrays have to be homogenous, in that every element in an array must be of the same type.
	var arr [4]int
	arr[3] = 3
	arr[0] = 100
	fmt.Println(arr)
	fmt.Println(arr[0])

	anotherArray := [5]int{1, 2, 6, 3, 22}
	fmt.Println(anotherArray[4])
	fmt.Println(anotherArray)
	// just like in python you can find out the length of the array with len()
	fmt.Println(len(anotherArray))

	sum := 0
	for i:=0; i<len(anotherArray); i++{
		sum+=anotherArray[i]
	}
	fmt.Println(sum)

	arr2D := [2][3]int{{6, 2, 80}, {10, 4, 100}}
	fmt.Println(arr2D)
}