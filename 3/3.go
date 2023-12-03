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
		for _, s := range numRe.FindAllString(line, -1) {
			n, _ := strconv.Atoi(s)
			nums[coordinate{x: strings.Index(line, s), y: y}] = n
			line = strings.Replace(line, s, strings.Repeat(".", len(s)), 1)
		}
		for _, s := range symRe.FindAllString(line, -1) {
			syms[coordinate{x: strings.Index(line, s), y: y}] = s
			line = strings.Replace(line, s, strings.Repeat(".", len(s)), 1)
		}
	}

	sum1, sum2 := 0, 0
	for symCoord, sym := range syms {
		var prev int
		for numCoord, num := range nums {
			if !numCoord.adjacent(num, symCoord) {
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

func (c *coordinate) adjacent(n int, other coordinate) bool {
	var digits int
	for n != 0 {
		n /= 10
		digits++
	}
	for _, coord := range adj {
		for i := 0; i < digits; i++ {
			if other.x == coord.x+c.x+i && other.y == coord.y+c.y {
				return true
			}
		}
	}
	return false
}
