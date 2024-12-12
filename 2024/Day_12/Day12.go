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

	input_file := "input.txt"
	test_file := "test.txt"
	file_path := test_file

	modeFlag := flag.String("mode", "T", "Input file, I for input, default for test")
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

	grid := make([][]rune, len(text))
	for i := range grid {
		grid[i] = make([]rune, len(text[0]))
	}
	ParseGrid(&grid, &text)

	visited := make([][]bool, len(text))
	for i := range visited {
		visited[i] = make([]bool, len(text[i]))
	}

	part1 := 0
	part2 := 0

	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] {
				visited[i][j] = true
				part1 += CheckRegion(i, j, &grid, &visited)
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			visited[i][j] = false
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] {
				visited[i][j] = true
				part2 += CheckRegionCorners(i, j, &grid, &visited)
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func ParseGrid(grid *[][]rune, text *[]string) {
	for i, ln := range *text {
		for j, char := range ln {
			(*grid)[i][j] = char
		}
	}
}

func CheckRegion(x int, y int, grid *[][]rune, visited *[][]bool) int {
	area, perimeter, r, c := 0, 0, len(*grid), len((*grid)[0])
	d := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	q := [][]int{{x, y}}
	start_rune := (*grid)[x][y]

	for len(q) > 0 {
		curr := q[0]
		cx, cy := curr[0], curr[1]
		q = q[1:]
		area++

		for _, i := range d {
			nx, ny := cx+i[0], cy+i[1]
			if nx < 0 || nx >= r || ny < 0 || ny >= c {
				//log.Println("NX:", nx, "NY:", ny, "Out of bounds, perimeter + 1")
				perimeter++
				continue
			}

			if (*grid)[nx][ny] != start_rune {
				//log.Println("NX:", nx, "NY:", ny, "Different region, perimeter + 1")
				perimeter++
				continue
			}

			if (*visited)[nx][ny] {
				//log.Println("NX:", nx, "NY:", ny, "Visited, skipping")
				continue
			}

			//log.Println("NX:", nx, "NY:", ny, "Visiting region")
			(*visited)[nx][ny] = true
			q = append(q, []int{nx, ny})
		}
		//log.Println("X:", cx, "Y:", cy, "Area:", area, "Perimeter:", perimeter)
	}

	//log.Println("Start Rune:", string(start_rune), "X:", x, "Y:", y, "Area:", area, "Perimeter:", perimeter)
	return area * perimeter
}

func CheckRegionCorners(x int, y int, grid *[][]rune, visited *[][]bool) int {
	area, corners, r, c := 0, 0, len(*grid), len((*grid)[0])
	d := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	q := [][]int{{x, y}}
	start_rune := (*grid)[x][y]

	for len(q) > 0 {
		curr := q[0]
		cx, cy := curr[0], curr[1]
		q = q[1:]
		area++

		// Case: Outer Corners, if 90 degree adjacents are both not in region, corner found
		// Up and Right
		if cx-1 < 0 || cx-1 >= r || (*grid)[cx-1][cy] != start_rune {
			if cy+1 < 0 || cy+1 >= c || (*grid)[cx][cy+1] != start_rune {
				corners++
			}
		}

		// Right and Down
		if cy+1 < 0 || cy+1 >= c || (*grid)[cx][cy+1] != start_rune {
			if cx+1 < 0 || cx+1 >= r || (*grid)[cx+1][cy] != start_rune {
				corners++
			}
		}

		// Down and Left
		if cx+1 < 0 || cx+1 >= r || (*grid)[cx+1][cy] != start_rune {
			if cy-1 < 0 || cy-1 >= c || (*grid)[cx][cy-1] != start_rune {
				corners++
			}
		}

		// Left and Up
		if cy-1 < 0 || cy-1 >= c || (*grid)[cx][cy-1] != start_rune {
			if cx-1 < 0 || cx-1 >= r || (*grid)[cx-1][cy] != start_rune {
				corners++
			}
		}

		// Case: Inner Corners. If 90 degree adjacents are both in region, and the diagonal in their direction is not, corner found
		// Up and Right
		if cx-1 >= 0 && cx-1 < r && (*grid)[cx-1][cy] == start_rune {
			if cy+1 >= 0 && cy+1 < c && (*grid)[cx][cy+1] == start_rune {
				if (*grid)[cx-1][cy+1] != start_rune {
					corners++
				}
			}
		}

		// Right and Down
		if cy+1 >= 0 && cy+1 < c && (*grid)[cx][cy+1] == start_rune {
			if cx+1 >= 0 && cx+1 < r && (*grid)[cx+1][cy] == start_rune {
				if (*grid)[cx+1][cy+1] != start_rune {
					corners++
				}
			}
		}

		// Down and Left
		if cx+1 >= 0 && cx+1 < r && (*grid)[cx+1][cy] == start_rune {
			if cy-1 >= 0 && cy-1 < c && (*grid)[cx][cy-1] == start_rune {
				if (*grid)[cx+1][cy-1] != start_rune {
					corners++
				}
			}
		}

		// Left and Up
		if cy-1 >= 0 && cy-1 < c && (*grid)[cx][cy-1] == start_rune {
			if cx-1 >= 0 && cx-1 < r && (*grid)[cx-1][cy] == start_rune {
				if (*grid)[cx-1][cy-1] != start_rune {
					corners++
				}
			}
		}

		for _, i := range d {
			nx, ny := cx+i[0], cy+i[1]
			if nx < 0 || nx >= r || ny < 0 || ny >= c {
				continue
			}

			if (*grid)[nx][ny] != start_rune {
				continue
			}

			if (*visited)[nx][ny] {
				continue
			}

			(*visited)[nx][ny] = true
			q = append(q, []int{nx, ny})
		}
	}

	return area * corners
}
