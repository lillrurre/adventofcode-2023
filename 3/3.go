package main

import (
	"fmt"
	"github.com/lillrurre/adventofcode-2023/util"
	"regexp"
	"strings"
)

func main() {
	lines := util.FileAsStringArr(3, "\n")

	numRe := regexp.MustCompile(`\d+`)
	symRe := regexp.MustCompile(`[^0-9.]`)

	nums := make(map[util.Point]int)
	syms := make(map[util.Point]string)

	for y, line := range lines {
		for _, s := range numRe.FindAllString(line, -1) {
			nums[util.Point{X: strings.Index(line, s), Y: y}] = util.Atoi(s)
			line = strings.Replace(line, s, strings.Repeat(".", len(s)), 1)
		}
		for _, s := range symRe.FindAllString(line, -1) {
			syms[util.Point{X: strings.Index(line, s), Y: y}] = s
			line = strings.Replace(line, s, strings.Repeat(".", len(s)), 1)
		}
	}

	sum1, sum2 := 0, 0
	for symPoint, sym := range syms {
		var prev int
		for numPoint, num := range nums {
			if !numPoint.Adjacent(symPoint, util.AdjacentWithDiagonals, num) {
				continue
			}
			sum1 += num
			if sym != "*" {
				continue
			}
			if prev == 0 {
				prev = num
				continue
			}
			sum2 += prev * num
		}
		prev = 0
	}
	fmt.Printf("[1] Result: %d\n", sum1)
	fmt.Printf("[2] Result: %d\n", sum2)
}
