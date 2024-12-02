package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
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

	fmt.Println("Part 1:", check_reports(text, 0))
	fmt.Println("Part 2:", check_reports(text, 1))
}

func check_reports(text []string, dampener int) int {
	ans := 0
	for _, ln := range text {
		items_strings := strings.Split(ln, " ")
		var items []int
		for _, i := range items_strings {
			items = append(items, Handle(strconv.Atoi(i)))
		}

		if len(items) <= dampener+1 {
			ans++
			continue
		}

		cs := combin.Combinations(len(items), len(items)-dampener)
		for _, c := range cs {
			var selection []int
			for i := 0; i < len(c); i++ {
				selection = append(selection, items[c[i]])
			}

			if validate(selection) {
				ans++
				break
			}
		}
	}
	return ans
}

func validate(arr []int) bool {
	inc := true
	dec := true

	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] || arr[i]-arr[i-1] > 3 {
			inc = false
			break
		}
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] || arr[i-1]-arr[i] > 3 {
			dec = false
			break
		}
	}

	return inc || dec
}
