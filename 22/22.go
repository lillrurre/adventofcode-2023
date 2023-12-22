package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"slices"
	"strings"
)

type brick struct {
	start, end util.Cube
}

func main() {
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(22, "\n")) })
	util.Run(2, func() (sum int) { return part2(util.FileAsStringArr(22, "\n")) })
}

func part1(input []string) (sum int) {
	bricks := parse(input)
	ground := bricks[0].start.Z // ground is the lowest a brick can fall
	drop(bricks, ground)

	for i := 0; i < len(bricks); i++ {
		if supported(remove(bricks, i), ground) {
			sum++
		}
	}
	return sum
}

func part2(input []string) (sum int) {
	bricks := parse(input)
	ground := bricks[0].start.Z
	drop(bricks, ground)
	for i := 0; i < len(bricks); i++ {
		sum += chain(remove(bricks, i), ground)
	}
	return sum
}

func collision(i int, b brick, bricks []brick) bool {
	for j := i - 1; j > -1; j-- {
		if collide(b, bricks[j]) {
			return true
		}
	}
	return false
}

func collide(a, b brick) bool {
	if !(a.start.X <= b.end.X && a.end.X >= b.start.X) {
		return false
	}
	if !(a.start.Y <= b.end.Y && a.end.Y >= b.start.Y) {
		return false
	}
	if !(a.start.Z <= b.end.Z && a.end.Z >= b.start.Z) {
		return false
	}
	return true
}

func remove(bricks []brick, i int) []brick {
	return slices.Delete(slices.Clone(bricks), i, i+1)
}

// drop moves the Z-coordinate down until a collision, or at the ground.
func drop(bricks []brick, ground int) {
	for i, b := range bricks {
		for b.start.Z > ground {
			b = b.down()
			if collision(i, b, bricks) {
				break
			}
			bricks[i] = b
		}
	}
}

// supported tries to move a brick down and returns true if the brick cannot move.
func supported(bricks []brick, ground int) bool {
	for i, b := range bricks {
		for b.start.Z > ground {
			b = b.down()
			if collision(i, b, bricks) {
				break
			}
			return false
		}
	}
	return true
}

// chain returns the sum of moved bricks when a brick is removed.
func chain(bricks []brick, ground int) (reaction int) {
	for i, b := range bricks {
		for b.start.Z > ground {
			b = b.down()
			if !collision(i, b, bricks) {
				reaction++
				bricks[i] = b
			}
			break
		}
	}
	return reaction
}

func parse(input []string) (bricks []brick) {
	for _, line := range input {
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == '~' || r == ',' })
		bricks = append(bricks, brick{
			start: util.CubeFromNums([3]int(util.StrsToIntSlice(parts[0:3]...))),
			end:   util.CubeFromNums([3]int(util.StrsToIntSlice(parts[3:6]...))),
		})
	}
	slices.SortFunc(bricks, func(a, b brick) int { return min(a.start.Z, a.end.Z) - min(b.start.Z, b.end.Z) })
	return bricks
}

func (b brick) down() brick {
	b.start.Z--
	b.end.Z--
	return b
}
