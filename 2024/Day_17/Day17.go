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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	input_path := "input.txt"
	test_path := "test.txt"
	modeFlag := flag.String("mode", "T", "File path chooser. T is default, test. Insert I for input path")
	file_path := test_path
	if *modeFlag == "I" {
		file_path = input_path
	}

	file := Handle(os.Open(file_path))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, ln := range text {
		log.Println(ln)
	}

	part1 := 0
	part2 := 0

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
