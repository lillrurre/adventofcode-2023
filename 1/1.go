package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input/1")
	input := string(b)

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

	solve := func(part int) {
		var sum int
		for _, line := range strings.Split(input, "\n") {
			f, l := firstLast(line)
			val, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(f), string(l)))
			sum += val
		}
		fmt.Printf("[%d] Result: %d\n", part, sum)
	}

	solve(1)

	input = strings.ReplaceAll(input, "one", "o1e")
	input = strings.ReplaceAll(input, "two", "t2o")
	input = strings.ReplaceAll(input, "three", "t3e")
	input = strings.ReplaceAll(input, "four", "f4r")
	input = strings.ReplaceAll(input, "five", "f5e")
	input = strings.ReplaceAll(input, "six", "s6x")
	input = strings.ReplaceAll(input, "seven", "s7n")
	input = strings.ReplaceAll(input, "eight", "e8t")
	input = strings.ReplaceAll(input, "nine", "n9e")

	solve(2)
}
