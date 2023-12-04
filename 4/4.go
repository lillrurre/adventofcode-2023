package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	b, _ := os.ReadFile("input/4")
	input := string(b)
	re := regexp.MustCompile(`\d+`)

	// part 1
	{
		sum := 0
		for _, line := range strings.Split(input, "\n") {
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
		fmt.Printf("[1] Result: %d\n", sum)
	}

	// part 2
	{
		copies := make(map[int]int)
		sum := 0
		for i, line := range strings.Split(input, "\n") {
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

		fmt.Printf("[2] Result: %d\n", sum)
	}
}
