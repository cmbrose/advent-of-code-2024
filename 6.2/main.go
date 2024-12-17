package main

import (
	"fmt"

	"main/util"
)

func main() {
	grid := util.ReadInputRuneGrid()

	x, y, dx, dy := locateGuard(grid)

	grid[y][x] = 'X'

	// Perform the initial basic walk to identify the path to add obstacles to
	walk(grid, x, y, dx, dy)

	cnt := 0

	for sy, row := range grid {
		for sx, r := range row {
			if sy == y && sx == x {
				// Skip this initial location
				continue
			}

			if r == 0 || r == '#' {
				// If the guard doesn't cross this point anyways, no point to test
				continue
			}

			// Get a fresh grid and make the test obstacle
			grid := util.ReadInputRuneGrid()
			grid[sy][sx] = '#'

			// Check if it loops
			if walk(grid, x, y, dx, dy) {
				cnt += 1
			}
		}
	}

	//util.PrintGrid(grid, "%c")

	fmt.Printf("%d\n", cnt)
}

func walk(grid [][]rune, x, y, dx, dy int) bool {
	m := mask(dx, dy)

	// clear out the grid so we can use bitmasks effectively
	for y, row := range grid {
		for x, r := range row {
			if r != '#' {
				grid[y][x] = 0
			}
		}
	}

	// mark the initial location and direction
	grid[y][x] |= m

	for y+dy >= 0 && y+dy < len(grid) && x+dx >= 0 && x+dx < len(grid[y]) {
		if grid[y+dy][x+dx] == '#' {
			// Hit an obstacle, rotate and try again
			dx, dy = -dy, dx
			m = mask(dx, dy)
			continue
		}

		x += dx
		y += dy

		// If we already walked on this tile in this direction, we have a loop!
		if (grid[y][x] & m) != 0 {
			return true
		}

		// Apply the bitmask to mark that we were at this spot, in this direction
		grid[y][x] |= m
	}

	return false
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

// mask returns a unique bitmask for the given direction
func mask(dx, dy int) rune {
	if dx == 0 && dy == -1 {
		return 0x0001
	} else if dx == 1 && dy == 0 {
		return 0x0010
	} else if dx == 0 && dy == 1 {
		return 0x0100
	} else if dx == -1 && dy == 0 {
		return 0x1000
	}

	panic("Unhandled direction to mask")
}
