package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	blocks := util.ReadInputBlocks()

	rules, pageOrders := blocks[0], blocks[1]

	successors := make(map[int][]int)

	for _, rule := range rules {
		pair := strings.Split(rule, "|")
		before := util.AssertInt(pair[0])
		after := util.AssertInt(pair[1])

		successors[before] = append(successors[before], after)
	}

	sum := 0

	for _, po := range pageOrders {
		order := util.Map(strings.Split(po, ","), util.AssertInt)

		order, changed := fix(order, successors)

		if !changed {
			continue
		}

		middle := order[len(order)/2]
		sum += middle
	}

	fmt.Printf("%d\n", sum)
}

func findSwap(order []int, successors map[int][]int) (int, int) {
	processed := make(map[int]int)

	for i, page := range order {
		for _, successes := range successors[page] {
			if j, ok := processed[successes]; ok {
				return i, j
			}
		}

		processed[page] = i
	}

	return -1, -1
}

func fix(order []int, successors map[int][]int) ([]int, bool) {
	i, j := findSwap(order, successors)
	if i == -1 {
		return order, false
	}

	order[i], order[j] = order[j], order[i]
	order, _ = fix(order, successors)

	return order, true
}
