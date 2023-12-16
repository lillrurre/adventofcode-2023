package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"math"
)

type rock uint8

const (
	round rock = 'O'
	cube  rock = '#'
	empty rock = '.'
)

func main() {
	input := util.FileAsStringArr(14, "\n")
	util.Run(1, func() int { return part1(input) })
	util.Run(2, func() int { return part2(input) })
}

func part1(input []string) int {
	grid := util.StrsToGrid[rock](input...)
	grid = tilt(grid)
	sum := 0
	for i, line := range grid {
		sum += util.SliceCount(line, round) * (len(line) - i)
	}
	return sum
}

func part2(input []string) int {
	grid := util.StrsToGrid[rock](input...)
	sum := 0
	for i, line := range cycleLoop(grid) {
		sum += util.SliceCount(line, round) * (len(line) - i)
	}
	return sum
}

func cycleLoop(grid [][]rock) [][]rock {
	cache := util.NewCache[string, string]()
	var key, val string
	for _, line := range grid {
		key += string(line)
	}

	first := math.MaxInt
	var loop []string
	for i := 0; i < 1000000000; i++ {
		if v, ok := cache.Get(key); ok {
			if first != math.MaxInt && v == loop[0] {
				break
			}
			if first == math.MaxInt {
				first = i
			}
			loop = append(loop, v)
			key = v
			continue
		}

		// Only tilt north and flip 90 degrees --> hacks
		for flip := 0; flip < 4; flip++ {
			grid = tilt(grid)
			grid = util.FlipGrid(grid)
		}

		val = ""
		for _, line := range grid {
			val += string(line)
		}
		cache.Set(key, val)
		key = val
	}

	val = loop[(1000000000-1-first)%len(loop)]
	start := 0
	for i := 0; i < len(grid); i++ {
		grid[i] = []rock(val[start : start+len(grid[0])])
		start += len(grid[0])
	}

	return grid
}

func tilt(grid [][]rock) [][]rock {
	for i := 0; i < len(grid[0]); i++ {
		index := -1
		for j := 0; j < len(grid); j++ {
			r := grid[j][i]
			switch {
			case r == cube:
				index = -1
			case r == empty && index == -1:
				index = j
			case r == round && index != -1:
				grid[index][i], grid[j][i] = round, empty
				index++
			}
		}
	}
	return grid
}

// String for debug
func (s rock) String() string {
	return string(s)
}
