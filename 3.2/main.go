package main

import (
	"fmt"
	"regexp"
	"strings"

	"main/util"
)

var mulRegex = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

func main() {
	input := util.ReadInputLines()[0]

	matches := mulRegex.FindAll([]byte(input), -1)

	strs := util.Map(matches, func(b []byte) string { return string(b) })

	enabled := true
	sum := 0

	for _, str := range strs {
		switch str {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				a, b := parseMul(str)
				sum += a * b
			}
		}
	}

	fmt.Printf("%d\n", sum)
}

func parseMul(mul string) (int, int) {
	mul = strings.Split(mul, "(")[1]
	mul = strings.Split(mul, ")")[0]
	pair := strings.Split(mul, ",")

	return util.AssertInt(pair[0]), util.AssertInt(pair[1])
}
