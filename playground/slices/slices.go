package main

import (
	"fmt"
)

func main() {
	//emptySlice()
	//literalSlice()
	//literalSliceWithSpecificElement()
	//sliceOfSlice()
	//ppendSliceWithCap()
	//appendSliceWithoutCap()
	//sliceOfSliceWithCapRestriction()
	//sliceElementIteration()
	accessSlicePartly()
}

func emptySlice() {
	slice := make([]int, 5)
	printSlice(slice)
	slice = make([]int, 3, 5)
	printSlice(slice)
}

func literalSlice() {
	slice := []int{10, 20, 30, 40, 50}
	printSlice(slice)
}

func literalSliceWithSpecificElement() {
	slice := []int{20: 10}
	printSlice(slice)
}

func sliceOfSlice() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	slice[2] = 99
	printSlice(newSlice)
}

func appendSliceWithCap() {
	slice := make([]int, 3, 4)
	printSlice(slice)
	newslice := append(slice, 40)
	printSlice(newslice)
}

func appendSliceWithoutCap() {
	slice := make([]int, 3, 3)
	printSlice(slice)
	newslice := append(slice, 40)
	printSlice(newslice)
}

func sliceOfSliceWithCapRestriction() {
	slice := []int{10, 20, 30, 40, 50}
	newslice := slice[1:3:3]
	printSlice(newslice)
}

func sliceElementIteration() {
	slice := []int{10, 20, 30, 40, 50}
	for _, val := range slice {
		val = val + 99
	}
	printSlice(slice)

	for i, val := range slice {
		slice[i] = val + 99
	}
	printSlice(slice)
	printSliceAddresses(slice)
}

func accessSlicePartly() {
	slice := []int{10, 20, 30, 40, 50}
	for i := 1; i < len(slice)-1; i++ {
		fmt.Printf("index: %v, value: %v\n", i, slice[i])
	}
}

func printSlice(slice []int) {
	fmt.Printf("addr: %p\n", &slice)
	fmt.Printf("len: %v\n", len(slice))
	fmt.Printf("cap: %v\n", cap(slice))

	for i, val := range slice {
		fmt.Printf("index: %v, value: %v\n", i, val)

	}
}

func printSliceAddresses(slice []int) {
	for i, val := range slice {
		fmt.Printf("Index: %v, value %v, valueaddr %p, elementaddr %p\n", i, val, &val, &slice[i])
	}
}
