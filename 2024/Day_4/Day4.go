package main

import (
	"bufio"
	"fmt"
	"os"
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

	var grid [][]string
	for _, ln := range text {
		grid = append(grid, [][]string{strings.Split(ln, "")}...)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "X" {
				part1 += search(i, j, grid)
			}
		}
	}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == "A" {
				part2 += check_As(i, j, grid)
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func search(x int, y int, grid [][]string) int {
	found := 0
	target := "XMAS"

	r := len(grid)
	c := len(grid[0])
	d := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for curr_d_idx := 0; curr_d_idx < 8; curr_d_idx++ {
		dx := d[curr_d_idx][0]
		dy := d[curr_d_idx][1]
		cx := x
		cy := y

		comp := []string{"X"}
		for i := 0; i < 3; i++ {
			nx := cx + dx
			ny := cy + dy
			if nx < 0 || ny < 0 || nx >= r || ny >= c {
				break
			}

			comp = append(comp, grid[nx][ny])
			cx = nx
			cy = ny
		}

		generated := ""
		if len(comp) == 4 {
			generated = strings.Join(comp, "")
		}

		if generated == target {
			found++
		}
	}
	return found
}

func check_As(x int, y int, grid [][]string) int {
	ans := false
	left_diag := false
	right_diag := false

	if grid[x-1][y-1] == "M" && grid[x+1][y+1] == "S" {
		left_diag = left_diag || true
	}

	if grid[x-1][y-1] == "S" && grid[x+1][y+1] == "M" {
		left_diag = left_diag || true
	}

	if grid[x-1][y+1] == "S" && grid[x+1][y-1] == "M" {
		right_diag = right_diag || true
	}

	if grid[x-1][y+1] == "M" && grid[x+1][y-1] == "S" {
		right_diag = right_diag || true
	}

	ans = left_diag && right_diag
	if ans {
		return 1
	}
	return 0
}
