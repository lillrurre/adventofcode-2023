package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"slices"
	"strings"
)

func main() {
	equal := func(a, b []string) bool { return slices.Equal(a, b) }
	mirror := func(grid []string, check func(a, b []string) bool) int {
		for pos := 1; pos < len(grid); pos++ {
			steps := min(pos, len(grid)-pos)
			top := slices.Clone(grid[pos-steps : pos])
			bottom := grid[pos : pos+steps]
			slices.Reverse(top)
			if check(top, bottom) {
				return pos
			}
		}
		return 0
	}
	smudge := func(top, bottom []string) bool {
		diffs := 0
		for i := range top {
			for j := range top[i] {
				if top[i][j] != bottom[i][j] {
					diffs++
				}
			}
		}
		return diffs == 1
	}
	util.RunBoth(func() (int, int) {
		part1, part2 := 0, 0
		for _, grid := range util.FileAsStringArr(13, "\n\n") {
			lines := strings.Split(grid, "\n")
			horizontal, vertical := make([]string, 0), make([]string, len(lines[0]))
			for _, line := range lines {
				horizontal = append(horizontal, line)
				for i, r := range line {
					vertical[i] += string(r)
				}
			}
			part1 += 100*mirror(horizontal, equal) + mirror(vertical, equal)
			part2 += 100*mirror(horizontal, smudge) + mirror(vertical, smudge)
		}
		return part1, part2
	})
}
