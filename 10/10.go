package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
)

type step struct {
	steps int
	path  map[util.Point]bool
}

var next = map[rune][2]util.Point{
	'|': {util.North, util.South}, '-': {util.East, util.West}, 'F': {util.East, util.North},
	'7': {util.West, util.North}, 'J': {util.West, util.South}, 'L': {util.East, util.South},
}

func main() {
	input := util.FileAsStringArr(10, "\n")
	grid := make([][]rune, 0)
	var start util.Point
	for y, line := range input {
		grid = append(grid, []rune(line))
		for x, c := range line {
			if c == 'S' {
				start = util.Point{X: x, Y: y}
			}
		}
	}
	util.RunBoth(func() (int, int) {
		return solve(grid, start)
	})
}

func solve(grid [][]rune, start util.Point) (int, int) {
	res := new(step)
	for _, point := range util.Adjacent {
		if n := check(&step{path: make(map[util.Point]bool)}, grid, start.Add(point), util.Point{}); n.steps > res.steps {
			res = n
		}
	}
	return res.steps / 2, nest(grid, res.path)
}

func nest(grid [][]rune, path map[util.Point]bool) int {
	sum := 0
	last := '.'
	inside := false
	for y, line := range grid {
		for x, r := range line {
			if !path[util.Point{X: x, Y: y}] { // Point not part of the loop
				if inside { // If we are inside - add 1
					sum++
				}
				continue
			}
			if r == '|' || r == 'F' || r == 'L' || (last == 'L' && r == 'J') || (last == 'F' && r == '7') {
				inside, last = !inside, r
			}
		}
	}
	return sum
}

func check(s *step, grid [][]rune, point util.Point, last util.Point) *step {
	s.steps++
	if s.path[point] || grid[point.Y][point.X] == '.' {
		return s
	}
	s.path[point] = true
	points, ok := next[grid[point.Y][point.X]]
	if !ok {
		return s
	}
	p1, p2 := points[0], points[1]
	a, b := new(step), new(step)
	if nextPoint := point.Add(p1); !nextPoint.Equals(last) {
		a = check(s, grid, nextPoint, point)
	}
	if nextPoint := point.Add(p2); !nextPoint.Equals(last) {
		b = check(s, grid, nextPoint, point)
	}
	if a.steps > b.steps {
		return a
	}
	return b
}
