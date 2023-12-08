package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"strings"
)

func main() {
	input := util.FileAsString(1)

	firstLast := func(line string) (first rune, last rune) {
		for _, r := range line {
			if r >= '0' && r <= '9' {
				if first == 0 {
					first = r
				}
				last = r
			}
		}
		return first, last
	}

	solve := func() int {
		var sum int
		for _, line := range strings.Split(input, "\n") {
			f, l := firstLast(line)
			sum += util.ValuesToNum(f, l)
		}
		return sum
	}

	util.Run(1, func() any {
		return solve()
	})

	input = strings.ReplaceAll(input, "one", "o1e")
	input = strings.ReplaceAll(input, "two", "t2o")
	input = strings.ReplaceAll(input, "three", "t3e")
	input = strings.ReplaceAll(input, "four", "f4r")
	input = strings.ReplaceAll(input, "five", "f5e")
	input = strings.ReplaceAll(input, "six", "s6x")
	input = strings.ReplaceAll(input, "seven", "s7n")
	input = strings.ReplaceAll(input, "eight", "e8t")
	input = strings.ReplaceAll(input, "nine", "n9e")

	util.Run(2, func() any {
		return solve()
	})

}
