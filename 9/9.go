package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"regexp"
	"slices"
)

func main() {
	input := util.FileAsStringArr(9, "\n")
	util.RunBoth(func() (any, any) {
		p1, p2 := 0, 0
		for _, line := range input {
			values := util.StrsToIntSlice(regexp.MustCompile(`-?\d+`).FindAllString(line, -1)...)
			p1 += next(values)
			slices.Reverse(values)
			p2 += next(values)
		}
		return p1, p2
	})
}

func next(values []int) int {
	if len(values) == 0 || (slices.Min(values) == 0 && slices.Max(values) == 0) {
		return 0
	}
	size := len(values) - 1
	diff := make([]int, size)
	for i := 0; i < size; i++ {
		diff[i] = values[i+1] - values[i]
	}
	return values[size] + next(diff)
}
