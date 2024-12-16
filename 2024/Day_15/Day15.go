package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func Handle[T any](val T, e error) T {
	if e != nil {
		panic(e)
	}
	return val
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	input_path := "input.txt"
	test_path := "test.txt"
	test2_path := "test2.txt"
	test3_path := "test3.txt"
	modeFlag := flag.String("mode", "T", "Default: T, test. Insert I for input mode, T2 for test 2, T3 for test 3")
	flag.Parse()

	file_path := test_path
	if *modeFlag == "I" {
		file_path = input_path
	} else if *modeFlag == "T2" {
		file_path = test2_path
	} else if *modeFlag == "T3" {
		file_path = test3_path
	}

	file := Handle(os.Open(file_path))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	grid := make([][]int, len(text)-2)
	for i := range grid {
		grid[i] = make([]int, len(text[0]))
	}
	var commands string
	curr := ParseInput(&grid, &commands, &text)
	curr2 := make([]int, 2)

	grid2 := make([][]int, len(grid))
	for i := range grid2 {
		grid2[i] = make([]int, len(text[0])*2)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				grid2[i][j*2] = 1
				grid2[i][j*2+1] = 1
			} else if grid[i][j] == 2 {
				grid2[i][j*2] = 2
				grid2[i][j*2+1] = 4
			}
		}
	}
	curr2[0], curr2[1] = curr[0], curr[1]*2

	Simulate(&commands, &grid, curr[0], curr[1])
	Simulate2(&commands, &grid2, curr2[0], curr2[1])
	part1 := 0
	part2 := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				part1 += 100*i + j
			}
		}
	}

	for i := range grid2 {
		for j := range grid2[i] {
			if grid2[i][j] == 2 {
				part2 += 100*i + j
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func ParseInput(grid *[][]int, commands *string, text *[]string) []int {
	c := make([]int, 2)
	for i, ln := range *text {
		if ln == "" {
			*commands = (*text)[i+1]
			break
		}
		for j, r := range ln {
			if r == '#' {
				(*grid)[i][j] = 1
			} else if r == '@' {
				c[0], c[1] = i, j
			} else if r == 'O' {
				(*grid)[i][j] = 2
			}
		}
	}
	return c
}

func Simulate(commands *string, grid *[][]int, src_x int, src_y int) {
	d := make(map[rune][]int)
	d['^'] = []int{-1, 0}
	d['>'] = []int{0, 1}
	d['<'] = []int{0, -1}
	d['v'] = []int{1, 0}

	for _, char := range *commands {
		dx, dy := d[char][0], d[char][1]
		nx, ny := src_x+dx, src_y+dy
		c := 0
		for (*grid)[nx][ny] == 2 {
			nx, ny = nx+dx, ny+dy
			c++
		}

		if (*grid)[nx][ny] == 1 {
			continue
		}

		for c > 0 {
			(*grid)[nx][ny] = 2
			nx, ny = nx-dx, ny-dy
			c--
		}
		(*grid)[nx][ny] = 0
		src_x, src_y = nx, ny
	}
}

func Simulate2(commands *string, grid *[][]int, src_x int, src_y int) {
	d := make(map[rune][]int)
	d['^'] = []int{-1, 0}
	d['>'] = []int{0, 1}
	d['<'] = []int{0, -1}
	d['v'] = []int{1, 0}
	grid_copy := make([][]int, len(*grid))
	for i := range grid_copy {
		grid_copy[i] = make([]int, len((*grid)[0]))
	}

	for _, char := range *commands {
		for i := range *grid {
			copy(grid_copy[i], (*grid)[i])
		}

		dx, dy := d[char][0], d[char][1]
		if Push(grid, src_x+dx, src_y+dy, dx, dy) {
			//log.Println("Push success")
			src_x += dx
			src_y += dy
		} else {
			for i := range grid_copy {
				copy((*grid)[i], grid_copy[i])
			}
		}
	}
}

func Push(grid *[][]int, src_x int, src_y int, dx int, dy int) bool {
	if (*grid)[src_x][src_y] == 1 {
		return false
	}

	if (*grid)[src_x][src_y] == 0 {
		return true
	}

	if dx == 0 {
		if Push(grid, src_x+dx, src_y+dy, dx, dy) {
			(*grid)[src_x][src_y], (*grid)[src_x+dx][src_y+dy] = (*grid)[src_x+dx][src_y+dy], (*grid)[src_x][src_y]
			return true
		}
	} else {
		//log.Println("Push start")
		if (*grid)[src_x][src_y] == 2 {
			//log.Println("Pushing 2")
			push_left := Push(grid, src_x+dx, src_y+dy, dx, dy)
			//log.Println("push_left:", push_left)
			push_right := Push(grid, src_x+dx, src_y+1+dy, dx, dy)
			//log.Println("push_right:", push_right)
			if push_left && push_right {
				(*grid)[src_x][src_y], (*grid)[src_x+dx][src_y+dy] = (*grid)[src_x+dx][src_y+dy], (*grid)[src_x][src_y]
				(*grid)[src_x][src_y+1], (*grid)[src_x+dx][src_y+1+dy] = (*grid)[src_x+dx][src_y+1+dy], (*grid)[src_x][src_y+1]
				return true
			}
		} else {
			//log.Println("Pushing 4")
			push_right := Push(grid, src_x+dx, src_y+dy, dx, dy)
			//log.Println("push_right:", push_right)
			push_left := Push(grid, src_x+dx, src_y-1+dy, dx, dy)
			//log.Println("push_left:", push_left)
			if push_left && push_right {
				(*grid)[src_x][src_y], (*grid)[src_x+dx][src_y+dy] = (*grid)[src_x+dx][src_y+dy], (*grid)[src_x][src_y]
				(*grid)[src_x][src_y-1], (*grid)[src_x+dx][src_y-1+dy] = (*grid)[src_x+dx][src_y-1+dy], (*grid)[src_x][src_y-1]
				return true
			}
		}
	}
	return false
}
