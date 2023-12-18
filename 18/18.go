package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"strings"
)

var directionMap = map[uint8]util.Point{
	'U': util.North, 'D': util.South, 'L': util.West, 'R': util.East,
	'0': util.East, '1': util.South, '2': util.West, '3': util.North,
}

type dig struct {
	step int
	dir  util.Point
}

func main() {
	input := util.FileAsStringArr(18, "\n")
	util.Run(1, func() (sum int) { return part1(input) })
	util.Run(2, func() (sum int) { return part2(input) })
}

func part1(input []string) (sum int) {
	digs := make([]dig, 0)
	for _, line := range input {
		fields := strings.Fields(line)
		digs = append(digs, dig{step: util.Atoi(fields[1]), dir: directionMap[fields[0][0]]})
	}
	return calculateArea(digs)
}

func part2(input []string) int {
	digs := make([]dig, 0)
	for _, line := range input {
		field := strings.Fields(line)[2][2:]
		digs = append(digs, dig{step: util.ParseInt(field[:len(field)-2], 16), dir: directionMap[field[len(field)-2]]})
	}
	return calculateArea(digs)
}

func calculateArea(digs []dig) (area int) {
	var pos util.Point
	for _, d := range digs {
		next := pos.Move(d.dir.Multiply(d.step))
		area += (pos.Y*next.X - pos.X*next.Y) + d.step
		pos = next
	}
	return area/2 + 1
}
