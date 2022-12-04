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

	starOne(lines)
	starTwo(lines)
}

func getPairRange(sections string) []int {
	nums := strings.Split(sections, "-")

	start, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal(err)
	}

	return []int{start, end}
}

func starOne(lines []string) {
	overlaps := 0

	for _, line := range lines {
		pair := strings.Split(line, ",")
		r1 := getPairRange(pair[0])
		r2 := getPairRange(pair[1])

		if (r1[0] <= r2[0] && r1[1] >= r2[1]) || (r2[0] <= r1[0] && r2[1] >= r1[1]) {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}

func starTwo(lines []string) {
	overlaps := 0

outer:
	for _, line := range lines {
		pair := strings.Split(line, ",")
		r1 := getPairRange(pair[0])
		r2 := getPairRange(pair[1])

		same := make(map[int]bool)

		for i := r1[0]; i <= r1[1]; i++ {
			same[i] = true
		}

		for i := r2[0]; i <= r2[1]; i++ {
			if same[i] {
				overlaps++
				continue outer
			}
		}
	}

	fmt.Println(overlaps)

}
