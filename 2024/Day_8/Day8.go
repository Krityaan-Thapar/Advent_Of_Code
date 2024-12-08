package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func Handle[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

type Point struct {
	x int
	y int
}

func main() {
	test_file := "test.txt"
	input_file := "input.txt"

	usePath := test_file
	modeFlag := flag.String("mode", "T", "Decides whether to use test file (T) or input file (I). Default (T)")
	flag.Parse()
	if *modeFlag == "I" {
		usePath = input_file
	}

	file := Handle(os.Open(usePath))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	antennae := make(map[rune][]Point)
	for i, ln := range text {
		for j, char := range ln {
			if char != rune('.') {
				antennae[char] = append(antennae[char], Point{x: i, y: j})
			}
		}
	}
	r := len(text)
	c := len(text[0])

	antinode := make(map[Point]bool)
	for _, ant_pos_arr := range antennae {
		for i := 0; i < len(ant_pos_arr); i++ {
			for j := i + 1; j < len(ant_pos_arr); j++ {
				a1 := FindAntinode(ant_pos_arr[i], ant_pos_arr[j])
				if BoundCheck(r, c, a1) {
					antinode[a1] = true
				}

				a2 := FindAntinode(ant_pos_arr[j], ant_pos_arr[i])
				if BoundCheck(r, c, a2) {
					antinode[a2] = true
				}
			}
		}
	}

	antinodePt2 := make(map[Point]bool)
	for _, ant_pos_arr := range antennae {
		for i := 0; i < len(ant_pos_arr); i++ {
			for j := i + 1; j < len(ant_pos_arr); j++ {
				s1 := ant_pos_arr[i]
				s2 := ant_pos_arr[j]
				antinodePt2[s1] = true
				antinodePt2[s2] = true

				tmps1 := s1
				tmps2 := s2
				a1 := FindAntinode(tmps1, tmps2)
				for BoundCheck(r, c, a1) {
					antinodePt2[a1] = true
					tmps1 = tmps2
					tmps2 = a1
					a1 = FindAntinode(tmps1, tmps2)
				}

				tmps1 = s1
				tmps2 = s2
				a2 := FindAntinode(tmps2, tmps1)
				for BoundCheck(r, c, a2) {
					antinodePt2[a2] = true
					tmps2 = tmps1
					tmps1 = a2
					a2 = FindAntinode(tmps2, tmps1)
				}
			}
		}
	}

	part1 := len(antinode)
	part2 := len(antinodePt2)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func FindAntinode(a Point, b Point) Point {
	x_diff := a.x - b.x
	y_diff := a.y - b.y

	return Point{x: a.x - 2*x_diff, y: a.y - 2*y_diff}
}

func BoundCheck(r int, c int, t Point) bool {
	if 0 <= t.x && t.x < r && t.y < c && 0 <= t.y {
		return true
	}
	return false
}
