package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//spawnMaxThreads()
	//thread(10)
	processParallel()
}

func processParallel() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines\n")

	go printAlphabet(&wg, 'A')
	go printAlphabet(&wg, 'a')

	fmt.Println("Waiting finish\n")
	wg.Wait()
	fmt.Println("Terminating\n")

}

func printAlphabet(wg *sync.WaitGroup, char int) {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		for c := char; c < char+26; c++ {
			fmt.Printf("%c ", c)
		}
	}
	fmt.Printf("\n")
}

func spawnMaxThreads() {
	for i := 0; i < 10000; i++ {
		go thread(i)
	}

	for {
	}
}

func thread(id int) {
	for {
		fmt.Printf("Thread %d\n", id)
	}
}
