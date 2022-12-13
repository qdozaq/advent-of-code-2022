package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

type Monkey struct {
	id        int
	items     []int
	operation string
	operand   int
	test      int
	testCase  map[bool]int
	inspected int
}

func (m *Monkey) print() {
	fmt.Println("id:", m.id)
	fmt.Println("inspected:", m.inspected)
	fmt.Println("items:", m.items)
	fmt.Println("operation:", m.operation, m.operand)
	fmt.Println("test: divisible by", m.test)
	fmt.Println("test cases:", m.testCase)
}

func (m *Monkey) operate(item int) int {

	operand := m.operand
	if operand == -1 {
		operand = item
	}

	switch m.operation {
	case "+":
		return item + operand
	case "*":
		return item * operand
	}
	log.Fatal("invalid operation")
	return -1
}

func (m *Monkey) testItem(item int) int {
	return m.testCase[item%m.test == 0]
}

func process(lines []string) map[int]Monkey {
	monkeys := make(map[int]Monkey)

	for i := 0; i < len(lines); i += 7 {
		match := regexp.MustCompile(`(\d+)`).FindStringSubmatch(lines[i])
		id, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("failed parsing id", err)
		}

		// items
		itemsRegex := regexp.MustCompile(`(\d+)`)

		itemsMatch := itemsRegex.FindAllStringSubmatch(lines[i+1], -1)

		startingItems := make([]int, 0)
		for m := 0; m < len(itemsMatch); m++ {
			item, err := strconv.Atoi(itemsMatch[m][1])
			if err != nil {
				log.Fatal("failed parsing item", err)
			}
			startingItems = append(startingItems, item)
		}

		// operation
		opRegex := regexp.MustCompile(`([+*]) (\w+)`)

		match = opRegex.FindStringSubmatch(lines[i+2])

		operation := match[1]

		var operand int

		if val, err := strconv.Atoi(match[2]); err == nil {
			operand = val
		} else {
			operand = -1
		}

		// test
		match = regexp.MustCompile(`(\d+)`).FindStringSubmatch(lines[i+3])
		test, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("failed parsing test", err)
		}

		match = regexp.MustCompile(`(\d+)`).FindStringSubmatch(lines[i+4])
		trueCase, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("failed parsing truecase", err)
		}

		match = regexp.MustCompile(`(\d+)`).FindStringSubmatch(lines[i+5])
		falseCase, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("failed parsing false case", err)
		}

		monkey := Monkey{
			id:        id,
			items:     startingItems,
			operation: operation,
			operand:   operand,
			test:      test,
			testCase:  map[bool]int{true: trueCase, false: falseCase},
			inspected: 0,
		}

		monkeys[id] = monkey
	}

	return monkeys

}

func starOne(lines []string) {

	monkeys := process(lines)

	rounds := 20

	for i := 0; i < rounds; i++ {
		for monkId := 0; monkId < len(monkeys); monkId++ {
			monkey := monkeys[monkId]

			for _, item := range monkey.items {
				monkey.inspected++
				item = monkey.operate(item)
				item = item / 3

				otherId := monkey.testItem(item)
				if other, ok := monkeys[otherId]; ok {
					other.items = append(other.items, item)
					monkeys[otherId] = other
				}
			}

			monkey.items = nil

			monkeys[monkId] = monkey

			// monkey.print()
		}
	}

	numInspections := make([]int, 0)
	for _, m := range monkeys {
		numInspections = append(numInspections, m.inspected)
	}

	sort.Slice(numInspections, func(i, j int) bool {
		return numInspections[i] > numInspections[j]
	})

	monkeyBusiness := numInspections[0] * numInspections[1]

	fmt.Println(monkeyBusiness)
}

func starTwo(lines []string) {
	monkeys := process(lines)

	rounds := 10000

	lcm := 1
	for _, m := range monkeys {
		lcm *= m.test
	}

	for i := 0; i < rounds; i++ {
		for monkId := 0; monkId < len(monkeys); monkId++ {
			monkey := monkeys[monkId]

			for _, item := range monkey.items {
				monkey.inspected++
				item = monkey.operate(item)

				item = item % lcm

				otherId := monkey.testItem(item)
				if other, ok := monkeys[otherId]; ok {
					other.items = append(other.items, item)
					monkeys[otherId] = other
				}
			}

			monkey.items = nil

			monkeys[monkId] = monkey
		}
	}

	numInspections := make([]int, 0)
	for _, m := range monkeys {
		numInspections = append(numInspections, m.inspected)
	}

	sort.Slice(numInspections, func(i, j int) bool {
		return numInspections[i] > numInspections[j]
	})

	monkeyBusiness := numInspections[0] * numInspections[1]

	fmt.Println(monkeyBusiness)

}
