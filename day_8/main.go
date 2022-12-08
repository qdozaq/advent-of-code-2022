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

	grid := process(lines)
	starOne(grid)
	starTwo(grid)
}

type Grid struct {
	height int
	width  int
	trees  []int
}

func (g *Grid) get(row int, col int) int {
	if row >= g.width || col >= g.height {
		log.Fatal("woah there that's out of range", row, col)
	}

	return g.trees[col*g.width+row]
}

func process(lines []string) Grid {
	trees := make([]int, 0)
	height := len(lines)
	width := len(lines[0])

	for _, line := range lines {
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			trees = append(trees, digit)
		}
	}

	return Grid{trees: trees, width: width, height: height}
}

func isVisible(grid Grid, row int, col int) bool {
	tree := grid.get(row, col)

	highest := 0
	// up
	for i := 0; i < col; i++ {
		other := grid.get(row, i)
		if other > highest {
			highest = other
		}
	}

	if tree > highest {
		return true
	}

	// down
	highest = 0
	for i := col + 1; i < grid.height; i++ {
		other := grid.get(row, i)
		if other > highest {
			highest = other
		}
	}

	if tree > highest {
		return true
	}

	// left
	highest = 0
	for i := 0; i < row; i++ {
		other := grid.get(i, col)
		if other > highest {
			highest = other
		}
	}

	if tree > highest {
		return true
	}

	// right
	highest = 0
	for i := row + 1; i < grid.width; i++ {
		other := grid.get(i, col)
		if other > highest {
			highest = other
		}
	}

	if tree > highest {
		return true
	}

	return false
}

func starOne(grid Grid) {
	// start with perimeter
	visible := 2*(grid.height+grid.width) - 4

	for col := 1; col < grid.height-1; col++ {
		for row := 1; row < grid.width-1; row++ {
			if isVisible(grid, row, col) {
				visible += 1
			}
		}
	}

	fmt.Println(visible)
}

func getScenicScore(grid Grid, row int, col int) int {
	tree := grid.get(row, col)
	up, down, left, right := 0, 0, 0, 0

	// up
	for i := col - 1; i >= 0; i-- {
		other := grid.get(row, i)
		up++
		if other >= tree {
			break
		}
	}

	// down
	for i := col + 1; i < grid.height; i++ {
		other := grid.get(row, i)
		down++
		if other >= tree {
			break
		}
	}

	// left
	for i := row - 1; i >= 0; i-- {
		other := grid.get(i, col)
		left++
		if other >= tree {
			break
		}
	}

	// right
	for i := row + 1; i < grid.width; i++ {
		other := grid.get(i, col)
		right++
		if other >= tree {
			break
		}
	}

	return up * down * left * right
}

func starTwo(grid Grid) {

	highestScore := 0

	for col := 0; col < grid.height; col++ {
		for row := 0; row < grid.width; row++ {
			s := getScenicScore(grid, row, col)
			if s > highestScore {
				highestScore = s
			}
		}
	}

	fmt.Println(highestScore)
}
