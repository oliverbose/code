package main

import (
	"fmt"
)

func main() {
	//emptyMap()
	//literalMap()
	//addEntry()
	//doesEntryExist()
	//deleteEntry()
	printMapAddresses()
}

func emptyMap() {
	var myMap map[string]string
	myMap = make(map[string]string)
	printMap(myMap)
}

func literalMap() {
	myMap := map[string]string{"Red": "#FF0000", "Green": "#00FF00", "Blue": "#0000FF"}
	printMap(myMap)
}

func addEntry() {
	myMap := make(map[string]string)
	myMap["Red"] = "#FF0000"
	printMap(myMap)
}

func doesEntryExist() {
	myMap := map[string]string{"Red": "#FF0000"}
	printExists(myMap, "Red")
	printExists(myMap, "Blue")

}

func deleteEntry() {
	myMap := map[string]string{"Red": "#FF0000", "Green": "#00FF00", "Blue": "#0000FF"}
	printMap(myMap)
	delete(myMap, "Red")
	printMap(myMap)
	delete(myMap, "Yellow")
	printMap(myMap)
}

func printExists(aMap map[string]string, key string) {
	_, exists := aMap[key]
	fmt.Printf("%s exists: %v\n", key, exists)

}
func printMapAddresses() {

	myMap := make(map[string]string)
	red := "#FF0000"
	myMap["Red"] = red
	for key, value := range myMap {
		elem := myMap[key]
		elem2 := myMap[key]
		fmt.Printf("Key: %v, value: %v, valueAddr: %p, originalElemAddr: %p, elementInMapAddr: %p,  elementInMapAdd2r: %p\n", key, value, &value, &red, &elem, &elem2)
	}
}

func printMap(aMap map[string]string) {
	fmt.Printf("size: %v\n", len(aMap))
	for key, value := range aMap {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
