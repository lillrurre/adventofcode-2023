package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"golang.org/x/exp/maps"
	"strings"
)

func main() {
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(21, "\n")) })
	util.Run(2, func() (sum int) { return part2(util.FileAsStringArr(21, "\n")) })
}

func part1(input []string) (sum int) {
	grid, start := parse(input)
	return solve(64, grid, map[util.Point]int{start: 0})
}

func part2(input []string) (sum int) {
	return
}

func solve(steps int, grid [][]rune, found map[util.Point]int) int {
	for step := 1; step <= steps; step++ {
		for k, v := range found {
			if v != step-1 {
				continue
			}
			for _, adj := range util.Adjacent {
				p := k.Move(adj)
				if grid[p.Y][p.X] != '.' {
					continue
				}
				found[p] = step
			}
		}
	}

	maps.DeleteFunc(found, func(point util.Point, i int) bool {
		return i%2 != 0
	})
	return len(found)
}

func parse(input []string) (grid [][]rune, start util.Point) {
	grid = make([][]rune, len(input))
	for y, line := range input {
		grid[y] = append(grid[y], []rune(line)...)
		if strings.Contains(line, "S") {
			start = util.Point{
				X: strings.IndexRune(line, 'S'),
				Y: y,
			}
		}
	}
	return grid, start
}
