package main

import (
	"fmt"

	"main/util"
)

func main() {
	grid := util.ReadInputRuneGrid()

	targets := []rune{'M', 'A', 'S'}

	cnt := 0

	for y, row := range grid {
		for x, r := range row {
			if r != 'X' {
				continue
			}

			for dx := -1; dx <= 1; dx += 1 {
				for dy := -1; dy <= 1; dy += 1 {
					if recurse(x+dx, y+dy, dx, dy, grid, targets) {
						cnt += 1
					}
				}
			}
		}
	}

	fmt.Printf("%d\n", cnt)
}

func recurse(x, y, dx, dy int, grid [][]rune, targets []rune) bool {
	if len(targets) == 0 {
		return true
	}

	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return false
	}

	if grid[y][x] != targets[0] {
		return false
	}

	return recurse(x+dx, y+dy, dx, dy, grid, targets[1:])
}
