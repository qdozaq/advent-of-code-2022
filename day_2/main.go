package main

import (
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

var codeToShape = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
	"A": "R",
	"B": "P",
	"C": "S",
}

var shapeToScore = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

var winCondition = map[string]string{
	"R": "S",
	"P": "R",
	"S": "P",
}

func getBonus(elf string, me string) int {
	if elf == me {
		return 3
	}
	if winCondition[me] == elf {
		return 6
	}
	return 0
}

func starOne(lines []string) {
	total := 0
	for _, line := range lines {
		results := strings.Split(line, " ")

		elf := codeToShape[results[0]]
		me := codeToShape[results[1]]

		bonus := getBonus(elf, me)

		myScore := shapeToScore[me]

		total += myScore + bonus
	}
	fmt.Println(total)
}

func starTwo(lines []string) {
	total := 0
	for _, line := range lines {
		results := strings.Split(line, " ")

		elf := codeToShape[results[0]]
		status := results[1]

		me := ""

		switch status {
		case "X": //lose
			me = winCondition[elf]
		case "Y": //draw
			me = elf
		case "Z": //win
			me = winCondition[winCondition[elf]]
		}

		bonus := getBonus(elf, me)

		myScore := shapeToScore[me]

		total += myScore + bonus
	}
	fmt.Println(total)
}
