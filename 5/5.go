package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type position struct {
	dest, src, rng int
}

type positions []position

type almanac struct {
	seeds     []int
	positions [7]positions
}

func main() {
	b, _ := os.ReadFile("input/5")
	input := strings.Split(string(b), "\n\n")

	// part 1
	{
		a := new(almanac)
		a.seeds = parseSeeds(input[0])
		for i, in := range input[1:] {
			a.positions[i] = parse(in)
		}
		fmt.Printf("[1] Result: %d\n", a.solve1())
	}

	// part 2
	{
		a := new(almanac)
		a.seeds = parseSeeds(input[0])
		for i, in := range input[1:] {
			a.positions[i] = parse(in)
		}
		fmt.Printf("[2] Result: %d\n", a.solve2())
	}
}

func parseSeeds(seeds string) (s []int) {
	for _, match := range regexp.MustCompile(`\d+`).FindAllString(seeds, -1) {
		s = append(s, atoi(match))
	}
	return s
}

func parse(m string) (p positions) {
	for _, line := range strings.Split(m, "\n") {
		matches := regexp.MustCompile(`\d+`).FindAllString(line, -1)
		if len(matches) != 3 {
			continue
		}
		p = append(p, position{
			dest: atoi(matches[0]),
			src:  atoi(matches[1]),
			rng:  atoi(matches[2]),
		})
	}
	return p
}

func next(loc int, positions positions) int {
	for _, pos := range positions {
		if loc > pos.src && loc <= pos.src+pos.rng {
			return pos.dest + (loc - pos.src)
		}
	}
	return loc
}

func (a *almanac) solve1() int {
	lowest := 99999999999
	for _, seed := range a.seeds {
		lowest = min(lowest, a.seedToLocation(seed))
	}
	return lowest
}

func (a *almanac) seedToLocation(loc int) int {
	for _, pos := range a.positions {
		loc = next(loc, pos)
	}
	return loc
}

func (a *almanac) solve2() int {
	lowest := 99999999999
	for i := 0; i < len(a.seeds); i += 2 {
		seed := a.seeds[i]
		for j := seed; j < seed+a.seeds[i+1]; j++ {
			loc := a.seedToLocation(j)
			lowest = min(lowest, loc)
		}
	}
	return lowest - 1
}

func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}
