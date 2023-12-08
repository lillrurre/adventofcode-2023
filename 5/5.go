package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"regexp"
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
	input := util.FileAsStringArr(5, "\n\n")

	// part 1
	{
		a := new(almanac)
		a.seeds = parseSeeds(input[0])
		for i, in := range input[1:] {
			a.positions[i] = parse(in)
		}
		util.Run(1, func() any {
			return a.solve1()
		})
	}

	// part 2
	{
		a := new(almanac)
		a.seeds = parseSeeds(input[0])
		for i, in := range input[1:] {
			a.positions[i] = parse(in)
		}
		util.Run(2, func() any {
			return a.solve2()
		})
	}
}

func parseSeeds(seeds string) (s []int) {
	return util.StrsToIntSlice(regexp.MustCompile(`\d+`).FindAllString(seeds, -1)...)
}

func parse(m string) (p positions) {
	for _, line := range strings.Split(m, "\n") {
		values := util.StrsToIntSlice(regexp.MustCompile(`\d+`).FindAllString(line, -1)...)
		if len(values) != 3 {
			continue
		}
		p = append(p, position{dest: values[0], src: values[1], rng: values[2]})
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
