package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
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

	adj := make(map[int]map[int]bool)
	var orders [][]int
	parse_flag := true

	for _, ln := range text {
		if ln == "" {
			if !parse_flag {
				panic(errors.New("two blank new line partitions detected"))
			}
			parse_flag = false
			continue
		}

		if parse_flag {
			ParseRelations(ln, adj)
		} else {
			ParseOrdering(ln, &orders)
		}
	}
	part1 := 0
	part2 := 0

	for _, order := range orders {
		tmp := make([]int, len(order))
		copy(tmp, order)

		slices.SortStableFunc(tmp, func(a int, b int) int {
			_, ok := adj[b][a]
			if ok {
				return 1
			}

			_, ok2 := adj[a][b]
			if ok2 {
				return -1
			}

			return 0
		})

		flag_equal := true
		for i, v := range tmp {
			if v != order[i] {
				flag_equal = false
				break
			}
		}

		if flag_equal {
			part1 += order[len(order)/2]
		} else {
			part2 += tmp[len(tmp)/2]
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func ParseRelations(ln string, adj map[int]map[int]bool) {
	parts := strings.Split(ln, "|")
	if len(parts) != 2 {
		panic(errors.New("delimiter | part does not have 2 values, likely incorrect line input" + ln))
	}

	src := Handle(strconv.Atoi(parts[0]))
	dst := Handle(strconv.Atoi(parts[1]))
	_, ok := adj[src]
	if !ok {
		adj[src] = make(map[int]bool)
	}

	adj[src][dst] = true
}

func ParseOrdering(ln string, orders *[][]int) {
	parts := strings.Split(ln, ",")
	var line []int

	for _, i := range parts {
		line = append(line, Handle(strconv.Atoi(i)))
	}

	if len(line)%2 == 0 {
		panic(errors.New("incoming line has even length"))
	}

	*orders = append(*orders, [][]int{line}...)
}
