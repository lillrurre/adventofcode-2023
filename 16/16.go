package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"slices"
)

type tile rune

const (
	empty       tile = '.'
	mirrorRight tile = '/'
	mirrorLeft  tile = '\\'
	splitWidth  tile = '-'
	SplitHeight tile = '|'
)

func main() {
	util.SwitchPointPoles()
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(16, "\n")) })
	util.Run(1, func() (sum int) { return part2(util.FileAsStringArr(16, "\n")) })
}

func part1(input []string) (sum int) {
	grid := make([][]tile, 0)
	for _, line := range input {
		grid = append(grid, []tile(line))
	}
	m := move(grid, util.Point{}, util.East, make(map[util.Point][]util.Point))
	return len(m)
}

func part2(input []string) (sum int) {
	grid := make([][]tile, 0)
	for _, line := range input {
		grid = append(grid, []tile(line))
	}
	for i := range grid {
		m1 := move(grid, util.Point{Y: i}, util.East, make(map[util.Point][]util.Point))
		m2 := move(grid, util.Point{X: len(grid[0]) - 1, Y: i}, util.West, make(map[util.Point][]util.Point))
		sum = max(sum, len(m1), len(m2))
	}
	for i := range grid[0] {
		m1 := move(grid, util.Point{X: i}, util.South, make(map[util.Point][]util.Point))
		m2 := move(grid, util.Point{X: i, Y: len(grid) - 1}, util.North, make(map[util.Point][]util.Point))
		sum = max(sum, len(m1), len(m2))
	}
	return sum
}

func move(grid [][]tile, pos util.Point, dir util.Point, visited map[util.Point][]util.Point) map[util.Point][]util.Point {
	for {
		if pos.Y < 0 || pos.X < 0 || pos.Y >= len(grid) || pos.X >= len(grid[0]) {
			break
		}
		if l, ok := visited[pos]; ok && slices.Contains(l, dir) {
			break
		}
		visited[pos] = append(visited[pos], dir)
		t := grid[pos.Y][pos.X]
		switch t {
		case empty:
		case mirrorLeft, mirrorRight:
			dir = mirror(t, dir)
		case SplitHeight:
			if dir.Equals(util.West) || dir.Equals(util.East) {
				visited = move(grid, pos, util.North, visited)
				dir = util.South
			}
		case splitWidth:
			if dir.Equals(util.North) || dir.Equals(util.South) {
				visited = move(grid, pos, util.East, visited)
				dir = util.West
			}
		}
		pos = pos.Add(dir)
	}
	return visited
}

func mirror(t tile, dir util.Point) util.Point {
	if dir.Equals(util.North) {
		if t == mirrorLeft {
			return util.West
		}
		return util.East
	}
	if dir.Equals(util.West) {
		if t == mirrorLeft {
			return util.North
		}
		return util.South
	}
	if dir.Equals(util.South) {
		if t == mirrorLeft {
			return util.East
		}
		return util.West
	}
	if dir.Equals(util.East) {
		if t == mirrorLeft {
			return util.South
		}
		return util.North
	}
	return dir
}

func (t tile) String() string {
	return string(t)
}
