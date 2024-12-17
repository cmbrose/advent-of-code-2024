package main

import (
	"fmt"

	"main/util"
)

func main() {
	grid := util.ReadInputRuneGrid()

	x, y, dx, dy := locateGuard(grid)

	grid[y][x] = 'X'
	cnt := 1

	for y+dy >= 0 && y+dy < len(grid) && x+dx >= 0 && x+dx < len(grid[y]) {
		if grid[y+dy][x+dx] == '#' {
			dx, dy = -dy, dx
			continue
		}

		x += dx
		y += dy

		if grid[y][x] != 'X' {
			grid[y][x] = 'X'
			cnt += 1
		}
	}

	//util.PrintGrid(grid, "%c")

	fmt.Printf("%d\n", cnt)
}

func locateGuard(grid [][]rune) (x, y, dx, dy int) {
	for y, row := range grid {
		for x, r := range row {
			if r != '.' && r != '#' {
				dx, dy := getDirection(r)
				return x, y, dx, dy
			}
		}
	}

	panic("Could not find guard")
}

func getDirection(r rune) (dx, dy int) {
	switch r {
	case '^':
		return 0, -1
	case '>':
		return 1, 0
	case 'v':
		return 0, 1
	case '<':
		return -1, 0
	default:
		panic("Unhandled guard symbol: " + string(r))
	}
}
