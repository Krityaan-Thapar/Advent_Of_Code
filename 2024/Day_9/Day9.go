package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Handle[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test_file := "test.txt"
	input_file := "input.txt"

	file_path := test_file
	modeFlag := flag.String("mode", "T", "Decides whether to use test file (T) or input file (I). Default (T)")
	flag.Parse()
	if *modeFlag == "I" {
		file_path = input_file
	}

	file := Handle(os.Open(file_path))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	src := make([]int, len(text[0]))
	src2 := make([]int, len(text[0]))
	FetchIntArray(text[0], &src)
	FetchIntArray(text[0], &src2)
	//log.Println("input assumptions valid:", SanityTestInput(&src))

	part1 := Defragment(&src)
	part2 := DefragmentPt2(&src2)

	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func FetchIntArray(src string, res *[]int) {
	for idx, i := range src {
		(*res)[idx] = Handle(strconv.Atoi(string(i)))
	}
}

func SanityTestInput(src *[]int) bool {
	for i := range *src {
		if i%2 == 0 && (*src)[i] == 0 {
			panic(errors.New("used file block has 0 memory assigned, unable to identify ID"))
		}
	}
	return true
}

func Defragment(src *[]int) int {
	ans := 0
	l := 0
	r := len(*src)
	id_left := 0
	id_right := 0

	for i := range *src {
		if i%2 == 0 {
			id_right++
		}
	}
	id_right--
	curr_idx := 0

	for l <= r {
		if l%2 == 0 {
			if (*src)[l] == 0 {
				//log.Println("l exhausted, moving up. l:", l, "id_left:", id_left)
				l++
				id_left++
			} else {
				ans += curr_idx * id_left
				//log.Println("Contribution to answer:", curr_idx*id_left, "l:", l, "id_left:", id_left, "curr_idx:", curr_idx)
				curr_idx++
				(*src)[l]--
			}
		} else {
			if (*src)[l] == 0 {
				//log.Println("free space at l:", l, "exhausted, moving up")
				l++
			} else if r%2 == 0 {
				if (*src)[r] == 0 {
					//log.Println("r exhausted, moving down. r:", r, "id_right:", id_right)
					r--
					id_right--
				} else {
					ans += curr_idx * id_right
					//log.Println("Contribution to answer:", curr_idx*id_right, "r:", r, "id_right:", id_right, "curr_idx:", curr_idx)
					curr_idx++
					(*src)[r]--
					(*src)[l]--
				}
			} else {
				//log.Println("r:", r, "is a free memory block. Moving down")
				r--
			}
		}
	}

	return ans
}

func DefragmentPt2(src *[]int) int {
	ans := 0
	id_left := 0
	id_right := 0
	r := len(*src)
	left_free_blocks := make(map[int]int)
	left_demand_blocks := make(map[int]int)

	for i := range *src {
		if i%2 == 0 {
			id_right++
			left_demand_blocks[i] = (*src)[i]
		} else {
			left_free_blocks[i] = (*src)[i]
		}
	}
	id_right--
	//log.Println("Sanity src:", (*src))
	//log.Println("id_left:", id_left, "id_right:", id_right)

	for r >= 0 {
		if r%2 == 0 {
			curr_idx := 0
			//log.Println("Checking for r:", r, "src[r]", (*src)[r], "left_free:", left_free_blocks)
			for l := 0; l < r; l++ {
				if l%2 == 1 && left_free_blocks[l] >= left_demand_blocks[r] {
					//log.Println("Found space at l:", l, "left_free_blocks:", left_free_blocks[l], "use demand:", (*src)[r])
					curr_idx += (*src)[l] - left_free_blocks[l]
					//log.Println("curr_idx:", curr_idx)
					for x := 0; x < (*src)[r]; x++ {
						//log.Println("ans contribution:", id_right, "*", curr_idx)
						ans += curr_idx * id_right
						curr_idx++
					}
					left_free_blocks[l] -= (*src)[r]
					left_demand_blocks[r] = 0
					break
				} else {
					curr_idx += (*src)[l]
				}
			}
			id_right--
			r -= 2
		} else {
			r--
		}
	}

	curr_idx := 0
	for l := 0; l < len(*src); l++ {
		if l%2 == 0 {
			if left_demand_blocks[l] == 0 {
				curr_idx += (*src)[l]
			} else {
				for x := 0; x < left_demand_blocks[l]; x++ {
					//log.Println("ans contribution:", id_left, "*", curr_idx)
					ans += curr_idx * id_left
					curr_idx += 1
				}
			}
			id_left++
		} else {
			curr_idx += (*src)[l]
		}
	}
	return ans
}
