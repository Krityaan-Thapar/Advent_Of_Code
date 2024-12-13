package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/mat"
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

	test_file := "test.txt"
	input_file := "input.txt"
	file_path := test_file

	modeFlag := flag.String("mode", "T", "Default - T. Enter I for input file path")
	methodFlag := flag.String("method", "M", "Default - Math, Enter MI for inbuilt Matrix, G for Gauss matrix method")
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

	part1 := 0
	part2 := 0

	x := make([]int, 2)
	y := make([]int, 2)
	t := make([]int, 2)
	t_converted := make([]int, 2)
	i := 0
	for _, ln := range text {
		if i == 3 {
			i = 0
			continue
		}

		match := useRegex(ln)
		if i == 0 {
			x[0] = Handle(strconv.Atoi(match[0]))
			y[0] = Handle(strconv.Atoi(match[1]))
		} else if i == 1 {
			x[1] = Handle(strconv.Atoi(match[0]))
			y[1] = Handle(strconv.Atoi(match[1]))
		} else {
			t[0] = Handle(strconv.Atoi(match[0]))
			t[1] = Handle(strconv.Atoi(match[1]))
			t_converted[0] = t[0] + 10000000000000
			t_converted[1] = t[1] + 10000000000000

			if *methodFlag == "MI" {
				part1 += CalcMatrixInbuilt(&x, &y, &t)
				part2 += CalcMatrixInbuilt(&x, &y, &t_converted)
			} else {
				part1 += CalcMath(&x, &y, &t)
				part2 += CalcMath(&x, &y, &t_converted)
			}
		}
		i++
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func useRegex(s string) []string {
	re := regexp.MustCompile(`[0-9]+`)
	return re.FindAllString(s, -1)
}

func CalcMath(x *[]int, y *[]int, t *[]int) int {
	b := (((*y)[0] * (*t)[0]) - ((*t)[1] * (*x)[0])) / (((*x)[1] * (*y)[0]) - ((*x)[0] * (*y)[1]))
	a := ((*t)[0] - ((*x)[1] * b)) / (*x)[0]

	validate_t0 := (*x)[0]*a + (*x)[1]*b
	validate_t1 := (*y)[0]*a + (*y)[1]*b

	if validate_t0 == (*t)[0] && validate_t1 == (*t)[1] {
		return 3*a + b
	}
	return 0
}

func CalcMatrixInbuilt(x *[]int, y *[]int, t *[]int) int {
	A := mat.NewDense(2, 2, []float64{float64((*x)[0]), float64((*x)[1]), float64((*y)[0]), float64((*y)[1])})
	T := mat.NewVecDense(2, []float64{float64((*t)[0]), float64((*t)[1])})

	var X mat.VecDense
	Handle(1, X.SolveVec(A, T))

	a, b := int(math.Round(X.AtVec(0))), int(math.Round(X.AtVec(1)))
	validate_t0 := (*x)[0]*a + (*x)[1]*b
	validate_t1 := (*y)[0]*a + (*y)[1]*b

	if validate_t0 == (*t)[0] && validate_t1 == (*t)[1] {
		return 3*a + b
	}
	return 0
}

func CalcGauss(x *[]int, y *[]int, t *[]int) int {
	// doesn't work, need to research
	A := [][]int{{(*x)[0], (*x)[1], (*t)[0]}, {(*y)[0], (*y)[1], (*t)[1]}}
	n, m := 2, 2
	where := make([]int, m)
	for i := range where {
		where[i] = -1
	}

	for col, row := 0, 0; col < m && row < n; col++ {
		sel := row
		for i := row; i < n; i++ {
			if A[i][col] > A[sel][col] {
				sel = i
			}
		}

		if float64(A[sel][col]) < 0.000000001 {
			continue
		}

		for i := col; i <= m; i++ {
			A[sel][i], A[row][i] = A[row][i], A[sel][i]
		}
		where[col] = row

		for i := 0; i < n; i++ {
			if i != row {
				c := A[i][col] / A[row][col]
				for j := col; j <= m; j++ {
					A[i][j] -= A[row][j] * c
				}
			}
		}
		row++
	}

	ans := make([]int, m)
	for i := 0; i < m; i++ {
		if where[i] != -1 {
			ans[i] = A[where[i]][m] / A[where[i]][i]
		}
	}

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += ans[j] * A[i][j]
		}
		if math.Abs(float64(sum-A[i][m])) > 0.000000001 {
			return 0
		}
	}

	for i := 0; i < m; i++ {
		if where[i] == -1 {
			return 0
		}
	}

	a, b := ans[0], ans[1]
	validate_t0 := (*x)[0]*a + (*x)[1]*b
	validate_t1 := (*y)[0]*a + (*y)[1]*b
	log.Println("A:", a, "B:", b)
	log.Println("T:", t)
	log.Println("Validate:", validate_t0, validate_t1)

	if validate_t0 == (*t)[0] && validate_t1 == (*t)[1] {
		return 3*a + b
	}
	return 0
}
