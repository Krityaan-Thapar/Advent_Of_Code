package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test_file := "test.txt"
	input_file := "input.txt"

	modeFlag := flag.String("mode", "T", "Decides whether to use test file (T) or input file (I). Default (T)")
	flag.Parse()
	file_path := test_file
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

	items := strings.Split(text[0], " ")
	stones := make(map[int]int)

	for _, i := range items {
		stones[Handle(strconv.Atoi(i))] = 1
	}

	part1 := Simulate(25, stones)
	part2 := Simulate(75, stones)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func Simulate(iter int, stones map[int]int) int {
	local := make(map[int]int)
	for k, v := range stones {
		local[k] = v
	}

	for i := 0; i < iter; i++ {
		next_local := make(map[int]int)
		for k, v := range local {
			if k == 0 {
				next_local[1] += v
				continue
			}

			tmp := k
			d := 0
			for tmp > 0 {
				tmp = tmp / 10
				d++
			}

			if d%2 == 0 {
				raiser := int(math.Pow(float64(10), float64(d/2)))
				next_local[k/raiser] += v
				next_local[k%raiser] += v
			} else {
				next_local[k*2024] += v
			}
		}
		local = next_local
	}

	ans := 0
	for _, v := range local {
		ans += v
	}
	return ans
}
