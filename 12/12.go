package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"slices"
	"strings"
)

type spring rune

const (
	working spring = '.'
	broken  spring = '#'
	unknown spring = '?'
)

var cache = make(map[string]int)

func main() {
	input := util.FileAsStringArr(12, "\n")
	records := make([][]spring, 0)
	groups := make([][]uint8, 0)

	for _, line := range input {
		fields := strings.Fields(line)
		records = append(records, []spring(fields[0]))
		group := make([]uint8, 0)
		for _, n := range util.StrsToIntSlice(strings.Split(fields[1], ",")...) {
			group = append(group, uint8(n))
		}
		groups = append(groups, group)
	}

	util.Run(1, func() int {
		res := 0
		for i := range records {
			res += solve(records[i], groups[i])
		}
		return res
	})

	for i := 0; i < len(records); i++ {
		record := append([]spring{unknown}, records[i]...)
		group := groups[i]
		for j := 0; j < 4; j++ {
			records[i] = append(records[i], record...)
			groups[i] = append(groups[i], group...)
		}
	}

	util.Run(2, func() int {
		res := 0
		for i := range records {
			res += solve(records[i], groups[i])
		}
		return res
	})
}

func solve(record []spring, group []uint8) (res int) {
	if v, ok := cache[string(record)+string(group)]; ok {
		return v
	}
	if len(record) == 0 {
		if len(group) == 0 {
			return 1
		}
		return 0
	}

	if len(group) == 0 {
		if len(record) == 0 || !slices.Contains(record, broken) {
			return 1
		}
		return 0
	}

	defer func() {
		cache[string(record)+string(group)] = res
	}()

	firstSpring := record[0]
	firstNum := int(group[0])
	res = 0

	if firstSpring == working {
		return solve(record[1:], group)
	}

	if firstSpring == unknown {
		res += solve(record[1:], group)
	}

	if len(record) < firstNum || slices.Contains(record[:firstNum], working) || firstNum != len(record) && record[firstNum] == broken {
		return res
	}

	if firstNum == len(record) {
		return res + solve(nil, group[1:])
	}

	return res + solve(record[firstNum+1:], group[1:])
}

// String for debug
func (s spring) String() string {
	return string(s)
}
