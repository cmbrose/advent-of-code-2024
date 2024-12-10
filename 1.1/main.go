package main

import (
	"fmt"
	"main/util"
	"sort"
	"strings"
)

func main() {
	var left, right []int

	for _, line := range util.ReadInputLines() {
		parts := strings.Split(line, " ")
		left = append(left, util.AssertInt(parts[0]))
		right = append(right, util.AssertInt(parts[len(parts)-1]))
	}

	sort.Ints(left)
	sort.Ints(right)

	diffs := util.Zip(left, right, func(l, r int) int { return util.Abs(l - r) })

	fmt.Printf("%d\n", util.Sum(diffs))
}
