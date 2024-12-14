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
		middle := order[len(order)/2]

		processed := make(map[int]bool)

		for _, page := range order {
			for _, successes := range successors[page] {
				if processed[successes] {
					goto end
				}
			}

			processed[page] = true
		}

		sum += middle
	end:
	}

	fmt.Printf("%d\n", sum)
}
