package main

import (
	"errors"
	"fmt"
	"log"
	"os"
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

func getShared(items string) []rune {
	itemMap := make(map[rune]bool)
	shared := make(map[rune]bool)

	half := len(items) / 2

	for i, item := range items {
		if i >= half {
			if itemMap[item] {
				shared[item] = true
			}
			continue
		}

		itemMap[item] = true
	}

	keys := make([]rune, 0, len(shared))
	for key := range shared {
		keys = append(keys, key)
	}
	return keys
}

func getPriority(item rune) int {
	ascii := int(item)

	if ascii >= 65 && ascii <= 90 {
		return ascii - 38
	}

	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	}

	log.Fatal("please don't happen")
	return 0
}

func starOne(lines []string) {
	sharedItems := []rune{}
	for _, items := range lines {
		shared := getShared(items)
		sharedItems = append(sharedItems, shared...)
	}

	total := 0
	for _, item := range sharedItems {
		total += getPriority(item)
	}
	fmt.Println(total)
}

func getSharedBetweenElves(elves []string) (rune, error) {
	itemMap := make(map[rune]int)

	for _, elf := range elves {
		set := make(map[rune]bool)

		for _, item := range elf {
			set[item] = true
		}

		for item := range set {
			val, ok := itemMap[item]
			if ok {
				itemMap[item] = val + 1

				if val+1 == 3 {
					return item, nil
				}
			} else {
				itemMap[item] = 1
			}
		}
	}

	return ' ', errors.New("no shared items")
}

func starTwo(lines []string) {

	sharedItems := []rune{}

	for i := 0; i < len(lines); i += 3 {
		elves := []string{lines[i], lines[i+1], lines[i+2]}

		shared, err := getSharedBetweenElves(elves)
		if err != nil {
			log.Fatal(err)
		}

		sharedItems = append(sharedItems, shared)
	}

	total := 0
	for _, item := range sharedItems {
		total += getPriority(item)
	}

	fmt.Println(total)
}
