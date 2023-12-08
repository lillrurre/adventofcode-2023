package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"regexp"
	"slices"
	"strings"
)

func main() {
	input := util.FileAsStringArr(4, "\n")
	re := regexp.MustCompile(`\d+`)

	// part 1
	{
		util.Run(1, func() any {
			sum := 0
			for _, line := range input {
				parts := strings.Split(line, "|")
				winning := re.FindAllString(strings.Split(parts[0], ":")[1], -1)
				have := re.FindAllString(parts[1], -1)
				n := 0
				for _, ch := range have {
					if slices.Contains(winning, ch) {
						if n == 0 {
							n++
							continue
						}
						n *= 2
					}
				}
				sum += n
			}
			return sum
		})
	}

	// part 2
	{
		util.Run(2, func() any {
			copies := make(map[int]int)
			sum := 0
			for i, line := range input {
				parts := strings.Split(line, "|")
				winning := re.FindAllString(strings.Split(parts[0], ":")[1], -1)
				have := re.FindAllString(parts[1], -1)

				c := copies[i] + 1
				n := i + 1
				for _, val := range have {
					if slices.Contains(winning, val) {
						copies[n] += c
						n++
					}
				}
				sum += c
			}
			return sum
		})
	}
}
