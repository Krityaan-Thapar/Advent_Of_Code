package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Handle[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test_file := "test.txt"
	input_file := "input.txt"

	file_path := test_file
	modeFlag := flag.String("mode", "T", "Decides whether to use test file (T) or input file (I). Default (T)")
	flag.Parse()

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

	r := len(text)
	c := len(text[0])
	grid := make([][]int, r)
	for i := range grid {
		grid[i] = make([]int, c)
	}

	ParseGrid(&grid, text, r, c)
	d := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}

	part1 := 0
	part2 := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 0 {
				ResetVisited(&visited)
				part1 += dfs(i, j, &grid, &d, &visited, true)
				ResetVisited(&visited)
				part2 += dfs(i, j, &grid, &d, &visited, false)
				//LogVisited(&visited, x, y)
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func ParseGrid(grid *[][]int, text []string, r int, c int) {
	for i, ln := range text {
		if len(ln) != c {
			panic(errors.New("lines in grid not of uniform length"))
		}

		for j, ele := range ln {
			(*grid)[i][j] = Handle(strconv.Atoi(string(ele)))
		}
	}
}

func dfs(x int, y int, grid *[][]int, d *[][]int, visited *[][]bool, toggle bool) int {
	curr_val := (*grid)[x][y]
	if curr_val == 9 && !(*visited)[x][y] {
		if toggle {
			(*visited)[x][y] = true
		}
		return 1
	}

	sum := 0
	for i := 0; i < 4; i++ {
		nx := x + (*d)[i][0]
		ny := y + (*d)[i][1]

		if nx < 0 || ny < 0 || nx >= len((*grid)) || ny >= len((*grid)[0]) {
			continue
		}

		if (*grid)[nx][ny] == curr_val+1 {
			sum += dfs(nx, ny, grid, d, visited, toggle)
		}
	}
	return sum
}

func ResetVisited(visited *[][]bool) {
	for i := range *visited {
		for j := range *visited {
			(*visited)[i][j] = false
		}
	}
}

func LogVisited(visited *[][]bool, x int, y int) {
	log.Println("i:", x, "j:", y)
	for _, ln := range *visited {
		log.Println(ln)
	}
}
