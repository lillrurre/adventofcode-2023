package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
)

func main() {
	input := util.FileAsStringArr(17, "\n")
	grid := util.StrsToPointIntGrid(input...)
	end := util.Point{X: len(input[0]) - 1, Y: len(input) - 1}

	util.Run(1, func() (sum int) { return part1(grid, end) })
	util.Run(2, func() (sum int) { return part2(grid, end) })
}

func part1(grid map[util.Point]int, end util.Point) (sum int) {
	return solve(grid, end, 0, 3)
}

func part2(grid map[util.Point]int, end util.Point) (sum int) {
	return solve(grid, end, 4, 10)
}

type cart struct {
	straight int
	pos, dir util.Point
}

func solve(grid map[util.Point]int, end util.Point, minStraight, maxStraight int) int {
	visited := make(map[cart]int)

	queue := util.NewMinPriorityQueue[cart, int]()
	queue.Push(cart{pos: util.East, dir: util.East, straight: 1}, 0)
	queue.Push(cart{pos: util.South, dir: util.South, straight: 1}, 0)

	for {
		c, prio, _ := queue.Pop()

		heat, ok := grid[c.pos]
		if !ok {
			continue
		}

		heat += prio
		if c.pos == end {
			return heat
		}

		if h, ok := visited[c]; ok && h <= heat {
			continue
		}
		visited[c] = heat

		if c.straight < maxStraight {
			queue.Push(cart{straight: c.straight + 1, pos: c.pos.Move(c.dir), dir: c.dir}, heat)
		}

		if c.straight >= minStraight {
			queue.Push(cart{straight: 1, pos: c.pos.MoveLeft(c.dir), dir: c.dir.Left()}, heat)
			queue.Push(cart{straight: 1, pos: c.pos.MoveRight(c.dir), dir: c.dir.Right()}, heat)
		}

	}
}
