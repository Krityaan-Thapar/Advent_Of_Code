package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	input_file := "input.txt"
	test_file := "test.txt"

	modeFlag := flag.String("mode", "T", "I for input mode, T for test mode")
	flag.Parse()
	file_path := test_file
	if *modeFlag == "I" {
		file_path = input_file
	}

	file := Handle(os.Open(file_path))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	var part1 int
	var part2 int
	if *modeFlag == "I" {
		part1 = CalcRobots(&text, 100, 103, 101)
	} else {
		part1 = CalcRobots(&text, 100, 7, 11)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func CalcRobots(text *[]string, time int, r int, c int) int {
	quadrants := [][]int{
		{0, 0},
		{0, 0},
	}
	mid_r, mid_c := (r-1)/2, (c-1)/2

	for _, ln := range *text {
		ln_parts := strings.Split(ln, " ")
		pos := ParseSection(ln_parts[0])
		velocity := ParseSection(ln_parts[1])

		y, x := pos[0], pos[1]
		step_y, step_x := velocity[0], velocity[1]
		nx := x + (time * step_x)
		ny := y + (time * step_y)

		nx = (nx + (r * 1000)) % r
		ny = (ny + (c * 1000)) % c

		if nx == mid_r || ny == mid_c {
			continue
		}

		var i int
		var j int
		if nx > mid_r {
			i = 1
		} else {
			i = 0
		}

		if ny > mid_c {
			j = 1
		} else {
			j = 0
		}

		quadrants[i][j] += 1
	}
	return quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1]
}

func ParseSection(s string) []int {
	result := make([]int, 2)
	s = s[2:]
	spl := strings.Split(s, ",")

	result[0] = Handle(strconv.Atoi(spl[0]))
	result[1] = Handle(strconv.Atoi(spl[1]))
	return result
}
