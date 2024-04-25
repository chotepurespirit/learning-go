package main

import "fmt"


func main() {
	var myArray [5]int // Declare an array of 3 integers
	myArray[0] = 10 // Assign the first element a value of 10
	myArray[1] = 15 // Assign the second element a value of 15
	myArray[2] = 39 // Assign the third element a value of 39
	myArray[3] = 25 // Assign the fourth element a value of 25
	myArray[4] = 50 // Assign the fifth element a value of 50

	for i :=0 ; i < len(myArray) ; i++ {
		fmt.Println(myArray[i])
		fmt.Printf("my size array is %d\n", len(myArray))
	}
}