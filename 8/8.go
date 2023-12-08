package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"regexp"
	"slices"
)

func main() {
	s := util.FileAsScanner(8)

	runeToInt := func(r rune) (n int) {
		if r == 'L' {
			return 1
		}
		return 2
	}

	// Directions as int
	s.Scan()
	line := s.Text()
	directions := make([]int, len(line))
	for i, r := range line {
		directions[i] = runeToInt(r)
	}

	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)
	network := make(map[string][2]string)
	for s.Scan() {
		line = s.Text()
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		network[matches[1]] = [2]string{matches[2], matches[3]}
	}

	util.Run(1, func() any {
		return part1(directions, network, "AAA", "ZZZ")
	})
	util.Run(2, func() any {
		return part2(directions, network)
	})
}

func part1(directions []int, network map[string][2]string, current string, targets ...string) (moves int) {
	for !slices.Contains(targets, current) {
		moves++
		move := directions[(moves-1)%len(directions)]
		current = network[current][move-1]
	}
	return moves
}

func part2(directions []int, network map[string][2]string) int {
	starts := make([]string, 0)
	targets := make([]string, 0)
	for key := range network {
		if key[2] == 'A' {
			starts = append(starts, key)
		}
		if key[2] == 'Z' {
			targets = append(targets, key)
		}
	}

	moves := make([]int, len(starts))
	for i, start := range starts {
		moves[i] = part1(directions, network, start, targets...)
	}
	return util.LCM(moves...)
}
