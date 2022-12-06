package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	starOne(string(file))
	starTwo(string(file))
}

func doTheThing(size int, buffer string) {

	start := 0
	end := start + size

	for end < len(buffer) {
		slice := buffer[start:end]
		set := make(map[rune]bool)
		for _, s := range slice {
			set[s] = true
		}

		if len(set) == size {
			fmt.Println(end)
			break
		}
		start++
		end++
	}

}

func starOne(buffer string) {

	doTheThing(4, buffer)

}

func starTwo(buffer string) {
	doTheThing(14, buffer)
}
