package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	starOne(lines)
	starTwo(lines)
}

func starOne(lines []string) {
	highest := 0
	current := 0

	for _, line := range lines {
		if line == "" {
			if current > highest {
				highest = current
			}
			current = 0
			continue
		}
		cals, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		current += cals
	}

	fmt.Println(highest)
}

func starTwo(lines []string) {
	highest := []int{0, 0, 0}
	current := 0

	for _, line := range lines {
		if line == "" {
			higher := false
			for _, num := range highest {
				if current > num {
					higher = true
					break
				}
			}

			if higher {
				sort.Ints(highest)
				highest[0] = current
			}

			current = 0
			continue
		}
		cals, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		current += cals
	}

	total := 0
	for _, num := range highest {
		total += num
	}

	fmt.Println(total)

}
