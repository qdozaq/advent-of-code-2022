package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	sizes := process(lines)

	starOne(sizes)
	starTwo(sizes)
}

func parseLine(line string) (string, string) {
	cmdRegex := regexp.MustCompile(`\$ cd (.*)`)
	match := cmdRegex.FindStringSubmatch(line)
	if len(match) > 1 {
		return "cd", match[1]
	} else if strings.HasPrefix(line, "$ ls") {
		return "ls", ""
	}

	return "", ""

}

func folderSizes(cwd string, folders map[string][]string) map[string]int {
	sizes := make(map[string]int, 0)
	for k := range folders {
		folderSizesR(k, folders, sizes)
	}
	return sizes
}

func folderSizesR(cwd string, folders map[string][]string, sizes map[string]int) map[string]int {
	files := folders[cwd]
	size := 0
	for _, file := range files {
		args := strings.Split(file, " ")
		fileSize := args[0]
		fileName := args[1]
		if fileSize == "dir" {
			context := cwd + "/" + fileName
			if val, ok := sizes[context]; ok {
				size += val
			} else {
				size += folderSizesR(context, folders, sizes)[context]
			}
		} else {
			s, err := strconv.Atoi(fileSize)
			if err != nil {
				log.Fatal(err)
			}

			size += s
		}
	}

	sizes[cwd] = size
	return sizes
}

func process(lines []string) map[string]int {
	folderHistory := make([]string, 0)

	folders := make(map[string][]string, 0)

	context := "/"

	for _, line := range lines {
		cmd, dir := parseLine(line)
		switch cmd {
		case "cd":
			if dir == ".." {
				folderHistory = folderHistory[:len(folderHistory)-1]
			} else {
				folderHistory = append(folderHistory, dir)
			}
		case "ls":
			context = strings.Join(folderHistory, "/")
			if context != ".." {
				folders[context] = make([]string, 0)
			}
			continue
		default:
			folders[context] = append(folders[context], line)
		}
	}

	sizes := folderSizes("/", folders)

	return sizes

}

func starOne(sizes map[string]int) {

	total := 0
	for _, size := range sizes {
		if size <= 100000 {
			total += size
		}
	}

	// for k, v := range folders {
	// 	fmt.Println(k)
	// 	fmt.Println(sizes[k])
	// 	for _, vv := range v {
	// 		fmt.Println("  ", vv)
	// 	}
	// }

	fmt.Println(total)
}

func starTwo(sizes map[string]int) {
	totalSpace := 70000000
	needed := 30000000

	currentlyFree := totalSpace - sizes["/"]

	minimum := needed - currentlyFree

	lowest := totalSpace
	for _, size := range sizes {
		if size >= minimum && size < lowest {
			lowest = size
		}
	}

	fmt.Println(lowest)
}
