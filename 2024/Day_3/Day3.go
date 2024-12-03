package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Handle[T any](val T, e error) T {
	if e != nil {
		panic(e)
	}
	return val
}

func main() {
	file := Handle(os.Open("input.txt"))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	part1 := 0
	part2 := 0
	pause := false
	dont := "don't()"
	do := "do()"

	for _, ln := range text {
		for _, match := range useRegex(ln) {
			if match == dont {
				pause = true
			} else if match == do {
				pause = false
			} else {
				consider := strings.Split(match[4:len(match)-1], ",")
				val1 := Handle(strconv.Atoi(consider[0]))
				val2 := Handle(strconv.Atoi(consider[1]))

				if !pause {
					part2 += val1 * val2
				}
				part1 += val1 * val2
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func useRegex(s string) []string {
	re := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)|do\\(\\)|don't\\(\\)")
	return re.FindAllString(s, -1)
}
