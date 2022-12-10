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

var directionMap = map[string][2]int{
	"U": {1, 1},
	"D": {1, -1},
	"L": {0, -1},
	"R": {0, 1},
}

func parseVector(cmd string) (int, int, int) {
	args := strings.Split(cmd, " ")
	d := directionMap[args[0]]
	axis := d[0]
	direction := d[1]
	magnitude, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	return axis, direction, magnitude
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}

	return 1
}

func isAdjacent(head []int, tail []int) bool {
	xDiff := abs(head[0] - tail[0])
	yDiff := abs(head[1] - tail[1])

	return xDiff <= 1 && yDiff <= 1
}

func starOne(lines []string) {

	head_coords := []int{0, 0}
	tail_coords := []int{0, 0}
	visited := make(map[string]bool, 0)
	visited["0,0"] = true

	for _, line := range lines {
		axis, direction, magnitude := parseVector(line)

		for m := 1; m <= magnitude; m++ {
			head_coords[axis] += direction
			if !isAdjacent(head_coords, tail_coords) {

				var opposite int

				if axis == 0 {
					opposite = 1
				} else {
					opposite = 0
				}

				if tail_coords[opposite] == head_coords[opposite] { // both on same axis
					tail_coords[axis] += direction
				} else {
					tail_coords[opposite] = head_coords[opposite]
					tail_coords[axis] += direction
				}

				v := fmt.Sprintf("%d,%d", tail_coords[0], tail_coords[1])
				visited[v] = true

			}
		}

	}

	fmt.Println(len(visited))
}

func print(rope [10][]int) {

	occupied := make(map[string]int, 0)

	up, down, left, right := 10, -10, -10, 10

	for i, coord := range rope {
		v := fmt.Sprintf("%d,%d", coord[0], coord[1])
		if _, ok := occupied[v]; !ok {
			occupied[v] = i
		}

		if coord[0] > up {
			up = coord[0]
		}
		if coord[0] < down {
			down = coord[0]
		}
		if coord[1] > right {
			right = coord[1]
		}
		if coord[1] < left {
			left = coord[1]
		}

	}

	// fmt.Println(occupied)

	for y := right; y >= left; y-- {
		for x := down; x <= up; x++ {
			v := fmt.Sprintf("%d,%d", x, y)
			if val, ok := occupied[v]; ok {
				if val == 0 {
					fmt.Print("H")
				} else {
					fmt.Print(val)
				}
			} else if x == 0 && y == 0 {

				fmt.Print("s")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

	fmt.Println("")
}

func starTwo(lines []string) {

	rope_coords := [10][]int{}
	for i := range rope_coords {
		rope_coords[i] = []int{0, 0}
	}

	visited := make(map[string]bool, 0)
	visited["0,0"] = true

	for _, line := range lines {
		axis, direction, magnitude := parseVector(line)

		for m := 1; m <= magnitude; m++ {
			rope_coords[0][axis] += direction

			for i := 1; i < len(rope_coords); i++ {
				if !isAdjacent(rope_coords[i], rope_coords[i-1]) {

					cur_coords := rope_coords[i]
					prev_coords := rope_coords[i-1]

					if prev_coords[0] == cur_coords[0] { // both on x axis
						if cur_coords[1] > prev_coords[1] {
							cur_coords[1]--
						} else {
							cur_coords[1]++
						}
					} else if prev_coords[1] == cur_coords[1] { // both on y axis
						if cur_coords[0] > prev_coords[0] {
							cur_coords[0]--
						} else {
							cur_coords[0]++
						}
					} else {
						x_diff := prev_coords[0] - cur_coords[0]
						y_diff := prev_coords[1] - cur_coords[1]

						if abs(x_diff) > abs(y_diff) {
							cur_coords[1] += y_diff
							cur_coords[0] += sign(x_diff)
						} else {
							cur_coords[0] += x_diff
							cur_coords[1] += sign(y_diff)
						}
					}
				}

			}
			v := fmt.Sprintf("%d,%d", rope_coords[9][0], rope_coords[9][1])
			visited[v] = true
		}
		// fmt.Println(line)
		// print(rope_coords)

	}

	// print(rope_coords)

	fmt.Println(len(visited))

}
