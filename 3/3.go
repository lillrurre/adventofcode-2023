package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

var adj = []coordinate{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {0, -1}, {1, 1}, {1, 0}, {1, -1}}

func main() {

	b, _ := os.ReadFile("input/3")
	lines := strings.Split(string(b), "\n")

	numRe := regexp.MustCompile(`\d+`)
	symRe := regexp.MustCompile(`[^0-9.]`)

	nums := make(map[coordinate]int)
	syms := make(map[coordinate]string)

	for y, line := range lines {
		matches := numRe.FindAllString(line, -1)
		for _, s := range matches {
			n, _ := strconv.Atoi(s)
			nums[coordinate{x: strings.Index(line, s), y: y}] = n
			line = strings.Replace(line, s, strings.Repeat(".", len(s)), 1)
		}
		matches = symRe.FindAllString(line, -1)
		for _, s := range matches {
			syms[coordinate{x: strings.Index(line, s), y: y}] = s
			line = strings.Replace(line, s, strings.Repeat(".", len(s)), 1)
		}
	}

	// part 1
	{
		sum := 0
		for numCoord, num := range nums {
			for symCoord := range syms {
				if numCoord.adjacent(len(fmt.Sprintf("%d", num)), symCoord) {
					sum += num
				}
			}
		}
		fmt.Printf("[1] Result: %d\n", sum)
	}

	// part 2
	{
		sum := 0
		for symCoord, sym := range syms {
			if sym != "*" {
				continue
			}
			var firstNum int
			for numCoord, num := range nums {
				if !numCoord.adjacent(len(fmt.Sprintf("%d", num)), symCoord) {
					continue
				}
				if firstNum == 0 {
					firstNum = num
					continue
				}
				sum += firstNum * num
			}
			firstNum = 0
		}
		fmt.Printf("[2] Result: %d\n", sum)
	}

}

func (c *coordinate) adjacent(n int, other coordinate) bool {
	for _, coord := range adj {
		for i := 0; i < n; i++ {
			if other.x == coord.x+c.x+i && other.y == coord.y+c.y {
				return true
			}
		}
	}
	return false
}
