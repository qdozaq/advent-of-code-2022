package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	boxes, moves := process(lines)
	starOne(boxes, moves)
	boxes, moves = process(lines)
	starTwo(boxes, moves)
}

func process(lines []string) ([][]string, [][]int) {
	// colWidth := 3

	// columns := (len(lines[0]) % colWidth) + 1
	columns := 9

	draw := true

	boxes := make([][]string, columns)
	for i := range boxes {
		boxes[i] = make([]string, 0)
	}

	moves := make([][]int, 0)

	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			draw = false
			continue
		}
		if draw {
			// spaces := len(line) % 3
			i := 1
			col := 0
			for i < len(line) {
				if unicode.IsLetter(rune(line[i])) {
					boxes[col] = append(boxes[col], string(line[i]))
				}

				i += 4
				col++
			}
		} else {
			m := r.FindStringSubmatch(line)

			// fmt.Println(m)
			mm := make([]int, 3)
			for i := 1; i < len(m); i++ {
				num, err := strconv.Atoi(m[i])
				if err != nil {
					log.Fatal(err)
					continue
				}

				mm[i-1] = num
			}

			moves = append(moves, mm)
		}
	}

	for i := 0; i < len(boxes); i++ {
		a := boxes[i]
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}

		boxes[i] = a
	}
	return boxes, moves
}

func starOne(boxes [][]string, moves [][]int) {

	for _, move := range moves {

		amount := move[0]
		from := move[1] - 1
		to := move[2] - 1

		fromBox := boxes[from]
		toBox := boxes[to]
		var letter string
		for i := 0; i < amount; i++ {
			//pop
			letter, fromBox = fromBox[len(fromBox)-1], fromBox[:len(fromBox)-1]

			toBox = append(toBox, letter)
		}

		boxes[from] = fromBox
		boxes[to] = toBox
	}

	final := ""

	for _, box := range boxes {
		final += box[len(box)-1]
	}
	fmt.Println(final)

}

func starTwo(boxes [][]string, moves [][]int) {

	for _, move := range moves {
		amount := move[0]
		from := move[1] - 1
		to := move[2] - 1

		fromBox := boxes[from]
		toBox := boxes[to]
		var letter []string
		letter, fromBox = fromBox[len(fromBox)-amount:], fromBox[:len(fromBox)-amount]

		toBox = append(toBox, letter...)

		boxes[from] = fromBox
		boxes[to] = toBox
	}

	final := ""

	for _, box := range boxes {
		final += box[len(box)-1]
	}
	fmt.Println(final)

}
