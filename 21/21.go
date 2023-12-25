package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"maps"
	"strings"
	"sync"
)

func main() {
	// Solution only works on real input.
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(21, "\n")) })
	// The second relies on a pattern in the real input data which is not present in the sample data.
	util.Run(2, func() (sum int) { return part2(util.FileAsStringArr(21, "\n")) }) // 605247138198755
}

func part1(input []string) (sum int) {
	grid, start := parse(input)
	return solve(64, grid, map[util.Point]int{start: 0})
}

func solve(steps int, grid [][]rune, found map[util.Point]int) int {
	for step := 1; step <= steps; step++ {
		for k, v := range found {
			if v != step-1 {
				continue
			}
			for _, adj := range util.Adjacent {
				p := k.Move(adj)
				if grid[p.Y][p.X] == '.' {
					found[p] = step
				}
			}
		}
	}
	var mod int
	for i := 2; mod == 0; i++ {
		if steps%i == 0 {
			mod = steps
		}
	}
	maps.DeleteFunc(found, func(point util.Point, i int) bool { return i%mod != 0 })
	return len(found)
}

func part2(input []string) (sum int) {
	grid, start := parse(input)
	size := len(grid)
	expansions := 5
	grid = expand(grid, expansions)
	start.X += size * (expansions / 2)
	start.Y += size * (expansions / 2)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	var p1, p2, p3 int
	go func() {
		defer wg.Done()
		p1 = solve(size/2, grid, map[util.Point]int{start: 0})
	}()
	go func() {
		defer wg.Done()
		p2 = solve(size+size/2, grid, map[util.Point]int{start: 0})
	}()
	go func() {
		defer wg.Done()
		p3 = solve(size*2+size/2, grid, map[util.Point]int{start: 0})
	}()
	wg.Wait()
	a := (p3 + p1 - 2*p2) / 2
	b := p2 - p1 - a
	c := p1
	n := 26501365 / size
	return a*n*n + b*n + c // a*n² + b*n + c
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
			grid[start.Y][start.X] = '.'
		}
	}
	return grid, start
}

func expand(grid [][]rune, expansions int) [][]rune {
	for y, line := range grid {
		for j := 0; j < expansions-1; j++ {
			grid[y] = append(grid[y], line...)
		}
	}
	for j := 0; j < expansions-1; j++ {
		for _, line := range grid {
			grid = append(grid, line)
		}
	}
	return grid
}
