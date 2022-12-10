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

func process(line string) (int, int) {
	if line == "noop" {
		return 1, 0
	}

	val, err := strconv.Atoi(strings.Split(line, " ")[1])
	if err != nil {
		log.Fatal(err)
	}

	return 2, val

}

func starOne(lines []string) {
	cycles := 0
	register := 1

	strengths := 0
	for _, line := range lines {

		num_cycles, val := process(line)

		// always one cycle
		cycles++
		if (cycles-20)%40 == 0 {
			strength := cycles * register
			strengths += strength
			fmt.Printf("cycle: %d, register: %d strength: %d\n", cycles, register, strength)
		}

		if num_cycles == 2 {
			cycles++
			if (cycles-20)%40 == 0 {
				strength := cycles * register
				strengths += strength
				fmt.Printf("cycle: %d, register: %d strength: %d\n", cycles, register, strength)
			}
			register += val
		}

	}

	fmt.Println(strengths)
}

func starTwo(lines []string) {
	cycles := 0
	register := 1

	for _, line := range lines {

		num_cycles, val := process(line)

		// always one cycle
		cycles++

		pos := cycles % 40
		if pos == 0 && cycles != 0 {
			fmt.Println("")
		}
		if pos >= register-1 && pos <= register+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")

		}

		if num_cycles == 2 {
			cycles++
			pos := cycles % 40
			register += val
			if pos == 0 {
				fmt.Println("")
			}
			if pos >= register-1 && pos <= register+1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")

			}

		}

	}
}
