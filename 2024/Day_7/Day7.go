package main

import (
	"bufio"
	"fmt"
	"math"
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

	for _, ln := range text {
		col_split := strings.Split(ln, ":")
		lhs := Handle(strconv.Atoi(col_split[0]))
		rhs_strings := strings.Split(col_split[1], " ")[1:]
		rhs := make([]int, len(rhs_strings))

		for idx, i := range rhs_strings {
			rhs[idx] = Handle(strconv.Atoi(i))
		}

		if CheckDFSUtil(1, len(rhs), rhs[0], &rhs, lhs) {
			part1 += lhs
		}

		if CheckDFSUtilConcat(1, len(rhs), rhs[0], &rhs, lhs) {
			part2 += lhs
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func CheckDFSUtil(idx int, lim int, local int, arr *[]int, lhs int) bool {
	if idx == lim {
		return local == lhs
	}

	mult := CheckDFSUtil(idx+1, lim, local*(*arr)[idx], arr, lhs)
	add := CheckDFSUtil(idx+1, lim, local+(*arr)[idx], arr, lhs)
	return mult || add
}

func CheckDFSUtilConcat(idx int, lim int, local int, arr *[]int, lhs int) bool {
	if idx == lim {
		return local == lhs
	}

	mult := CheckDFSUtilConcat(idx+1, lim, local*(*arr)[idx], arr, lhs)
	add := CheckDFSUtilConcat(idx+1, lim, local+(*arr)[idx], arr, lhs)

	tmp := (*arr)[idx]
	d := 0
	for tmp > 0 {
		tmp = tmp / 10
		d++
	}
	raiser := int(math.Pow(float64(10), float64(d)))
	conc := CheckDFSUtilConcat(idx+1, lim, local*raiser+(*arr)[idx], arr, lhs)

	return mult || add || conc
}
