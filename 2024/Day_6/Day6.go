package main

import (
	"bufio"
	"fmt"
	"os"
)

func Handle[T any](val T, err error) T {
	if err != nil {
		panic(err)
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

	var grid [][]int
	cx := -1
	cy := -1

	for i, ln := range text {
		row := make([]int, len(ln))
		for j, char := range ln {
			if char == '^' {
				row[j] = 0
				cx = i
				cy = j
			} else if char == '#' {
				row[j] = 1
			}
		}
		grid = append(grid, [][]int{row}...)
	}

	part1 := 0
	part2 := 0

	r := len(grid)
	c := len(grid[0])
	facing := 0
	d := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	grid[cx][cy] = 2
	nx := cx + d[facing][0]
	ny := cy + d[facing][1]
	for nx >= 0 && nx < r && ny >= 0 && ny < c {
		if grid[nx][ny] == 1 {
			facing = (facing + 1) % 4
		} else {
			if Check(&grid, nx, ny, cx, cy, facing, &d) {
				part2++
			}
			cx = nx
			cy = ny
		}
		grid[cx][cy] = 2
		nx = cx + d[facing][0]
		ny = cy + d[facing][1]
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 2 {
				part1++
			}
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func Check(grid *[][]int, x int, y int, cx int, cy int, facing int, d *[][]int) bool {
	r := len((*grid))
	c := len((*grid)[0])
	count := 0

	(*grid)[x][y] = 1
	facing = (facing + 1) % 4
	nx := cx + (*d)[facing][0]
	ny := cy + (*d)[facing][1]

	for nx < r && ny < c && nx >= 0 && ny >= 0 {
		count++
		if count > 50000 {
			(*grid)[x][y] = 2
			return true
		}
		if (*grid)[nx][ny] == 1 {
			facing = (facing + 1) % 4
		} else {
			cx = nx
			cy = ny
		}
		nx = cx + (*d)[facing][0]
		ny = cy + (*d)[facing][1]
	}

	(*grid)[x][y] = 2
	return false
}
