package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"maps"
	"slices"
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

type node struct {
	point util.Point
	steps int
}

func main() {
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(23, "\n")) })
	util.Run(2, func() (sum int) { return part2(util.FileAsStringArr(23, "\n")) })
}

func part1(input []string) (sum int) {
	grid, start, end := parse(input)
	return solve(grid, start.Move(util.North), end, map[util.Point]bool{start: true}, 1) - 1
}

func part2(input []string) (sum int) {
	grid, start, end := parse(input)
	con := getConnections(grid, start, end)
	paths := getPaths(grid, con, true)
	return findLongestPath(grid, paths, start, end, 0, map[util.Point]bool{start: true})
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
	start.X = strings.IndexRune(input[0], path)
	end.X = strings.IndexRune(input[len(input)-1], path)
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

func getConnections(grid [][]rune, start, end util.Point) (connections []util.Point) {
	connections = append(connections, start, end)
	for y, line := range grid {
		for x, r := range line {
			if r == forest {
				continue
			}
			p := util.Point{X: x, Y: y}
			near := 0
			for _, adj := range util.Adjacent {
				next := p.Add(adj)
				if next.InBounds() && grid[next.Y][next.X] != forest {
					near++
				}
			}
			if near >= 3 {
				connections = append(connections, p)
			}
		}
	}
	return connections
}

func getPaths(grid [][]rune, connections []util.Point, part2 bool) (paths map[util.Point][]node) {
	paths = make(map[util.Point][]node)
	for _, pos := range connections {
		for _, adj := range util.Adjacent {
			next := pos.Move(adj)
			if next.InBounds() && grid[next.Y][next.X] != forest {
				n := getNode(grid, pos, next, adj, 1, connections)
				paths[pos] = append(paths[pos], n)
			}
		}
	}
	return paths
}

func getNode(grid [][]rune, pos, next, adj util.Point, steps int, connections []util.Point) node {
	for _, dir := range [3]util.Point{adj, adj.Left(), adj.Right()} {
		nextP := next.Move(dir)
		if grid[nextP.Y][nextP.X] != forest {
			if slices.Contains(connections, nextP) {
				return node{point: nextP, steps: steps + 1}
			}
			return getNode(grid, pos, nextP, dir, steps+1, connections)
		}
	}
	return node{point: util.Point{X: -1, Y: -1}, steps: 0}
}

func findLongestPath(grid [][]rune, paths map[util.Point][]node, start, end util.Point, step int, visited map[util.Point]bool) int {
	maxStep := 0
	for _, path := range paths[start] {
		if val, found := visited[path.point]; !found || !val {
			if path.point == end {
				return step + path.steps
			}
			visited[path.point] = true
			maxStep = max(maxStep, findLongestPath(grid, paths, path.point, end, step+path.steps, visited))
			visited[path.point] = false
		}
	}
	return maxStep
}
