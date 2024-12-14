package main

import (
	"fmt"

	"main/util"
)

func main() {
	grid := util.ReadInputRuneGrid()

	cnt := 0

	for y, row := range grid {
		for x, r := range row {
			if r != 'A' {
				continue
			}

			if !check(x, y, grid) {
				continue
			}

			cnt += 1
		}
	}

	fmt.Printf("%d\n", cnt)
}

func check(x, y int, grid [][]rune) bool {
	if x < 1 || y < 1 || y >= len(grid)-1 || x >= len(grid[y])-1 {
		return false
	}

	tl, tr, bl, br := grid[y-1][x-1], grid[y-1][x+1], grid[y+1][x-1], grid[y+1][x+1]

	as := len(util.Filter([]rune{tl, tr, bl, br}, func(x rune) bool { return x == 'S' }))
	ms := len(util.Filter([]rune{tl, tr, bl, br}, func(x rune) bool { return x == 'M' }))

	if as != 2 || ms != 2 {
		return false
	}

	return (tl == tr && bl == br) || (tr == br && bl == tl)
}
