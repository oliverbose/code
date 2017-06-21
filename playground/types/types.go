package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	parseAndPrint("10.10.10.10")
	parseAndPrint("Bla")
}

func parseAndPrint(ipString string) {
	var ip IP
	parsedIP, error := ip.parse(ipString)
	if error != nil {
		fmt.Printf("Failed to parse ip. Reason: %s\n", error)
		return
	}
	fmt.Printf("Parsed ip = %v\n", parsedIP)
}

type IP []int

func (ip IP) parse(ipToParse string) ([]int, error) {
	var parsedIP IP
	parts := strings.Split(ipToParse, ".")
	for _, part := range parts {
		intPart, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		parsedIP = append(parsedIP, intPart)
	}
	return parsedIP, nil
}
