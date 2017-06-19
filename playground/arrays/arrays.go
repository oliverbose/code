package main

import (
	"fmt"
)

func main() {
	//emptyArray()
	//literalArray()
	//literalArrayWithSpecificElements()
	//arrayOfPointers()
	//copyArrayOfPointers()
	twoDimensionalArray()
}

func emptyArray() {
	var array [5]int
	printArray(array)
}

func literalArray() {
	array := [5]int{10, 20, 30, 40, 50}
	printArray(array)
}

func literalArrayWithSpecificElements() {
	array := [5]int{1: 20, 2: 30}
	array[3] = 40
	printArray(array)
}

func arrayOfPointers() {
	array := [5]*int{1: new(int), 2: new(int)}
	*array[1] = 20
	printPointerArray(array)
}

func copyArrayOfPointers() {
	array1 := [5]*int{1: new(int), 2: new(int)}
	var array2 [5]*int
	*array1[1] = 20
	printPointerArray(array1)
	array2 = array1
	*array2[1] = 30
	printPointerArray(array1)

}

func twoDimensionalArray() {
	var array [2][2]int
	array[0][1] = 10
	printTwoDimensionalArray(array)
}

func printArray(array [5]int) {
	for i, val := range array {
		fmt.Printf("index: %v, value: %v\n", i, val)
	}
}
func printPointerArray(array [5]*int) {
	for i, val := range array {
		if val != nil {
			fmt.Printf("index: %v, value: %v\n", i, *val)
		}
	}
}
func printTwoDimensionalArray(array [2][2]int) {
	for i, val := range array {
		for j, val2 := range array[i] {
			fmt.Printf("index: (%v,%v) value: (%v,%v)\n", i, j, val[i], val2)
		}
	}
}
