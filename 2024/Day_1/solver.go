package main

import (
	"bufio"
	"fmt"
	"math"
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
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	var items1 []int
	var items2 []int

	for _, ln := range text {
		var items []string = strings.Split(ln, "   ")
		items1 = append(items1, Handle(strconv.Atoi(items[0])))
		items2 = append(items2, Handle(strconv.Atoi(items[1])))
	}
	slices.Sort(items1)
	slices.Sort(items2)
	ans := 0.0
	for i := 0; i < len(items1); i++ {
		ans += math.Abs(float64(items1[i] - items2[i]))
	}
	fmt.Printf("Part 1: %.1f\n", ans)

	score := 0
	var freq2 map[int]int = make(map[int]int)

	for i := 0; i < len(items2); i++ {
		freq2[items2[i]] += 1
	}
	for i := 0; i < len(items1); i++ {
		score += items1[i] * freq2[items1[i]]
	}
	fmt.Printf("Part 2: %d\n", score)
}
