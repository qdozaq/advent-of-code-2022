package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	var depths []int

	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, depth)
	}

	increases := 0

	for i := range depths {
		if i == 0 {
			continue
		}

		if depths[i] > depths[i-1] {
			increases++
		}
	}

	fmt.Println(increases)
}
