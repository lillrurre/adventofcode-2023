package main

import (
	"fmt"
	"github.com/lillrurre/adventofcode-2023/util"
	"maps"
	"strings"
)

const (
	path       = '.'
	forest     = '#'
	slopeUp    = '^'
	slopeDown  = 'v'
	slopeLeft  = '<'
	slopeRight = '>'
)

var slopeMap = map[rune]util.Point{
	slopeDown:  util.North,
	slopeUp:    util.South,
	slopeLeft:  util.West,
	slopeRight: util.East,
}

func main() {
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(23, "\n")) })
	// util.Run(2, func() (sum int) { return part2(util.FileAsStringArr(23, "\n")) })
}

func part1(input []string) (sum int) {
	grid, start, end := parse(input)
	return solve(grid, start.Move(util.North), end, map[util.Point]bool{start: true}, 1) - 1
}

func solve(grid [][]rune, pos, end util.Point, seen map[util.Point]bool, sum int) int {
	seen[pos] = true
	if pos == end {
		return max(len(seen), sum)
	}
	tile := grid[pos.Y][pos.X]
	if tile == forest {
		return max(len(seen), sum)
	}
	if dir, ok := slopeMap[tile]; ok {
		next := pos.Move(dir)
		if !seen[next] && grid[next.Y][next.X] == path {
			sum = max(solve(grid, next, end, maps.Clone(seen), sum), len(seen))
		}
		return max(len(seen), sum)
	}
	for _, adj := range util.Adjacent {
		next := pos.Move(adj)
		if seen[next] {
			continue
		}
		tile = grid[next.Y][next.X]
		sum = max(solve(grid, next, end, maps.Clone(seen), sum), len(seen))
	}
	return max(len(seen), sum)
}

func parse(input []string) (grid [][]rune, start, end util.Point) {
	start.X = strings.IndexRune(input[0], '.')
	end.X = strings.IndexRune(input[len(input)-1], '.')
	end.Y = len(input) - 1
	for _, line := range input {
		grid = append(grid, []rune(line))
	}
	util.SetBounds(util.Point{
		X: len(grid[0]) - 1,
		Y: len(grid) - 1,
	})
	return grid, start, end
}

func part2(input []string) (sum int) {
	return sum
}

func parse2(input []string) (grid [][]rune, start, end util.Point) {
	start.X = strings.IndexRune(input[0], '.')
	end.X = strings.IndexRune(input[len(input)-1], '.')
	end.Y = len(input) - 1
	for _, line := range input {
		line = strings.NewReplacer(">", ".", "<", ".", "v", ".", "^", ".").Replace(line)
		grid = append(grid, []rune(line))
	}
	for _, line := range grid {
		fmt.Println(string(line))
	}
	return grid, start, end
}
