package util

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Char interface {
	rune | byte
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInputLines() []string {
	content, err := os.ReadFile("./input.txt")
	Check(err)

	return strings.Split(string(content), "\n")
}

func ReadInputBlocks() [][]string {
	content, err := os.ReadFile("./input.txt")
	Check(err)

	blocks := strings.Split(string(content), "\n\n")

	return Map(blocks, func(block string) []string {
		return strings.Split(block, "\n")
	})
}

func ReadInputRuneGrid() [][]rune {
	lines := ReadInputLines()

	return Map(lines, func(line string) []rune { return []rune(line) })
}

func Repeat[T any](v T, cnt int) []T {
	arr := make([]T, cnt)

	for i := 0; i < cnt; i += 1 {
		arr[i] = v
	}

	return arr
}

func Map[X, Y any](xArr []X, f func(X) Y) []Y {
	yArr := make([]Y, len(xArr))

	for i, x := range xArr {
		y := f(x)
		yArr[i] = y
	}

	return yArr
}

func Zip[X, Y any](a, b []X, f func(X, X) Y) []Y {
	yArr := make([]Y, len(a))

	for i := range a {
		y := f(a[i], b[i])
		yArr[i] = y
	}

	return yArr
}

func Filter[X any](xArr []X, f func(X) bool) []X {
	var filtered []X

	for _, x := range xArr {
		if f(x) {
			filtered = append(filtered, x)
		}
	}

	return filtered
}

func All[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func Sum[T Number](a []T) T {
	var s T = 0
	for _, v := range a {
		s += v
	}
	return s
}

func AssertInt(str string) int {
	i, err := strconv.Atoi(str)
	Check(err)
	return i
}

func ParseBitString(str string) int {
	val := 0
	for _, c := range str {
		val <<= 1
		if c == '1' {
			val++
		}
	}

	return val
}

func Sign[T Number](x T) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Min[T Number](nums ...T) T {
	var min T

	for i, x := range nums {
		if i == 0 || x < min {
			min = x
		}
	}

	return min
}

func Max[T Number](nums ...T) T {
	var max T

	for i, x := range nums {
		if i == 0 || x > max {
			max = x
		}
	}

	return max
}

func SortStringChars(s string) string {
	rs := []rune(s)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})

	return string(rs)
}

func Keys[T comparable, U any](m map[T]U) []T {
	ret, i := make([]T, len(m)), 0
	for k := range m {
		ret[i] = k
		i += 1
	}
	return ret
}

func Values[T comparable, U any](m map[T]U) []U {
	ret, i := make([]U, len(m)), 0
	for _, v := range m {
		ret[i] = v
		i += 1
	}
	return ret
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func LCM(a ...int) int {
	res := 1
	for _, n := range a {
		res = res * n / GCD(res, n)
	}
	return res
}

// Must be pre-sorted!
func Intersect[T constraints.Ordered](a, b []T) []T {
	var res []T

	if len(a) == 0 || len(b) == 0 {
		return res
	}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		if a == b {
			res = append(res, a)
			i++
			j++
		} else if a > b { // a > b
			j++
		} else { // a < b
			i++
		}
	}

	return res
}

func IntersectAll[T constraints.Ordered](a ...[]T) []T {
	if len(a) == 0 {
		return nil
	}

	cur := a[0]

	if len(a) == 1 {
		return cur
	}

	for _, b := range a[1:] {
		cur = Intersect(cur, b)
	}

	return cur
}

// Must be pre-sorted!
func Except[T constraints.Ordered](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	var res []T

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		if a == b {
			i++
			j++
		} else if a > b { // a > b
			j++
		} else { // a < b
			res = append(res, a)
			i++
		}
	}

	for i < len(a) {
		res = append(res, a[i])
		i++
	}

	return res
}

// Must be pre-sorted!
func ExceptAll[T constraints.Ordered](a []T, b ...[]T) []T {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	cur := a

	for _, b := range b {
		cur = Except(cur, b)
	}

	return cur
}

func ParseIntGrid() [][]int {
	grid := [][]int{}

	for _, line := range ReadInputLines() {
		row := []int{}
		for _, cell := range line {
			row = append(row, int(cell-'0'))
		}

		grid = append(grid, row)
	}

	return grid
}

func PrintGrid[T any](grid [][]T, f string) {
	rows := make([]string, len(grid))

	if f == "" {
		f = "%v"
	}

	for i, row := range grid {
		rows[i] = ""
		for _, cell := range row {
			rows[i] += fmt.Sprintf(f, cell)
		}
	}

	fmt.Println(strings.Join(rows, "\n"))
}

func HexToBinary[T Char](char T) string {
	switch char {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	}

	panic(fmt.Sprintf("Unknown hex char: %c", char))
}

func IsUppercase[T Char](b T) bool {
	return b >= 'A' && b <= 'Z'
}

func IsLowercase[T Char](b T) bool {
	return b >= 'a' && b <= 'z'
}

func IsLetter[T Char](b T) bool {
	return IsUppercase(b) || IsLowercase(b)
}

func IsNumber[T Char](b T) bool {
	return b >= '0' && b <= '9'
}

func ReverseString(s string) string {
	r := ""
	for _, c := range s {
		r = fmt.Sprintf("%c%s", c, r)
	}

	return r
}

func Grid[T any](w, h int) [][]T {
	var grid [][]T

	for i := 0; i < h; i += 1 {
		grid = append(grid, make([]T, w))
	}

	return grid
}

func FillGrid[T any](w, h int, def T) [][]T {
	var grid [][]T

	for i := 0; i < h; i += 1 {
		row := make([]T, w)
		for i := range row {
			row[i] = def
		}
		grid = append(grid, row)
	}

	return grid
}

func Step(x1, y1, x2, y2 int, f func(x, y int)) {
	stepX := Sign(x2 - x1)
	stepY := Sign(y2 - y1)

	for x, y := x1, y1; x != x2 || y != y2; x, y = x+stepX, y+stepY {
		f(x, y)
	}

	// The last step is missed in the loop
	f(x2, y2)
}
