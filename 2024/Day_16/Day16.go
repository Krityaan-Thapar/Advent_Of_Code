package main

import (
	"bufio"
	"container/heap"
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

type HeapItem struct {
	cost   int
	x      int
	y      int
	facing int
	index  int
}

type PriorityQueue []*HeapItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*HeapItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	input_path := "input.txt"
	test_path := "test.txt"
	test2_path := "test2.txt"
	modeFlag := flag.String("mode", "T", "Chooses file path to run. Default T for test path. Enter I for input path. Enter T2 for test2 path")
	flag.Parse()

	file_path := test_path
	if *modeFlag == "I" {
		file_path = input_path
	} else if *modeFlag == "T2" {
		file_path = test2_path
	}

	file := Handle(os.Open(file_path))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	grid := make([][]int, len(text))
	for i := range grid {
		grid[i] = make([]int, len(text[i]))
	}

	visited := make([][]bool, len(text))
	for i := range visited {
		visited[i] = make([]bool, len(text[i]))
	}

	se := ParseInput(&text, &grid)
	src := se[0]
	dst := se[1]
	path := [][]int{}

	part1 := Dijsktra(&grid, &src, &dst)
	part2 := 0

	Traverse(part1, &grid, src[0], src[1], &dst, &visited, &path, 0, 0)
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				part2 += 1
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2+1)
}

func ParseInput(text *[]string, grid *[][]int) [][]int {
	se := make([][]int, 2)
	se[0] = make([]int, 2)
	se[1] = make([]int, 2)

	for i := range *text {
		for j := range (*text)[i] {
			if (*text)[i][j] == '#' {
				(*grid)[i][j] = 1
			} else if (*text)[i][j] == 'S' {
				se[0][0], se[0][1] = i, j
			} else if (*text)[i][j] == 'E' {
				se[1][0], se[1][1] = i, j
			}
		}
	}
	return se
}

func Dijsktra(grid *[][]int, src *[]int, dst *[]int) int {
	cx, cy := (*src)[0], (*src)[1]
	r, c := len(*grid), len((*grid)[0])
	d := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	dist := make([][]int, len(*grid))
	for i := range dist {
		dist[i] = make([]int, len((*grid)[i]))
		for j := range dist[i] {
			dist[i][j] = 1000000000000000000
		}
	}
	dist[cx][cy] = 0

	pq := make(PriorityQueue, 1)
	pq[0] = &HeapItem{
		cost:   0,
		x:      cx,
		y:      cy,
		facing: 0,
		index:  0,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*HeapItem)
		if item.x == (*dst)[0] && item.y == (*dst)[1] {
			return item.cost
		}

		if item.cost > dist[item.x][item.y] {
			continue
		}

		// Move facing
		next := &HeapItem{
			cost:   item.cost + 1,
			x:      item.x + d[item.facing][0],
			y:      item.y + d[item.facing][1],
			facing: item.facing,
		}
		if next.x >= 0 && next.x < r && next.y >= 0 && next.y < c && (*grid)[next.x][next.y] != 1 && next.cost < dist[next.x][next.y] {
			dist[next.x][next.y] = next.cost
			heap.Push(&pq, next)
		}

		// Rotate once
		next2 := &HeapItem{
			cost:   item.cost + 1001,
			x:      item.x + d[(item.facing+1)%4][0],
			y:      item.y + d[(item.facing+1)%4][1],
			facing: (item.facing + 1) % 4,
		}
		if next2.x >= 0 && next2.x < r && next2.y >= 0 && next2.y < c && (*grid)[next2.x][next2.y] != 1 && next2.cost < dist[next2.x][next2.y] {
			dist[next2.x][next2.y] = next2.cost
			heap.Push(&pq, next2)
		}

		next3 := &HeapItem{
			cost:   item.cost + 1001,
			x:      item.x + d[(item.facing+3)%4][0],
			y:      item.y + d[(item.facing+3)%4][1],
			facing: (item.facing + 3) % 4,
		}
		if next3.x >= 0 && next3.x < r && next3.y >= 0 && next3.y < c && (*grid)[next3.x][next3.y] != 1 && next3.cost < dist[next3.x][next3.y] {
			dist[next3.x][next3.y] = next3.cost
			heap.Push(&pq, next3)
		}
	}
	return dist[(*dst)[0]][(*dst)[1]]
}

func Traverse(best int, grid *[][]int, cx int, cy int, dst *[]int, visited *[][]bool, path *[][]int, cost int, facing int) {
	r, c := len(*grid), len((*grid)[0])
	d := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	if cx < 0 || cx >= r || cy < 0 || cy >= c {
		return
	}

	if cost > best {
		return
	}

	if (*grid)[cx][cy] == 1 {
		return
	}

	if cx == (*dst)[0] && cy == (*dst)[1] {
		for i := range *path {
			(*visited)[(*path)[i][0]][(*path)[i][1]] = true
		}
		return
	}

	*path = append(*path, [][]int{{cx, cy}}...)
	Traverse(best, grid, cx+d[facing][0], cy+d[facing][1], dst, visited, path, cost+1, facing)
	Traverse(best, grid, cx+d[(facing+1)%4][0], cy+d[(facing+1)%4][1], dst, visited, path, cost+1001, (facing+1)%4)
	Traverse(best, grid, cx+d[(facing+3)%4][0], cy+d[(facing+3)%4][1], dst, visited, path, cost+1001, (facing+3)%4)
	*path = (*path)[:len(*path)-1]
}
