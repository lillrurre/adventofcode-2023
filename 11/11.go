package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"strings"
)

func main() {
	input := util.FileAsStringArr(11, "\n")

	util.RunBoth(func() (int, int) {
		galaxies := make([]util.Point, 0)
		emptyX := make(map[int]bool)
		emptyY := make(map[int]bool)
		for y, line := range input {
			if strings.Contains(line, "#") {
				emptyY[y] = true
			}
			for x, r := range line {
				if r == '#' {
					emptyX[x] = true
					galaxies = append(galaxies, util.Point{X: x, Y: y})
				}
			}
		}

		part1, part2 := 0, 0
		for y, galaxy := range galaxies {
			for x := y + 1; x < len(galaxies); x++ {
				other := galaxies[x]
				if galaxy.Equals(other) {
					continue
				}
				n := galaxy.ManhattanDistance(other)
				part1 += n
				part2 += n
				for i := min(galaxy.X, other.X); i < max(galaxy.X, other.X); i++ {
					if !emptyX[i] {
						part1++
						part2 += 999999
					}
				}
				for i := min(galaxy.Y, other.Y); i < max(galaxy.Y, other.Y); i++ {
					if !emptyY[i] {
						part1++
						part2 += 999999
					}
				}
			}
		}
		return part1, part2
	})
}
