package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"regexp"
	"slices"
)

func main() {
	input := util.FileAsStringArr(15, ",")
	util.Run(1, func() int { return part1(input) })
	util.Run(2, func() int { return part2(input) })
}

func hash(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
		sum *= 17
		sum %= 256
	}
	return sum
}

func part1(input []string) (sum int) {
	for _, s := range input {
		sum += hash(s)
	}
	return sum
}

type lens struct {
	label string
	focal int
}

func part2(input []string) (sum int) {
	re := regexp.MustCompile(`(\w+)([=-])(\d*)`)
	boxes := make([][]lens, 256)
	for _, s := range input {
		matches := re.FindStringSubmatch(s)
		label := matches[1]
		box := hash(label)
		pos := slices.IndexFunc(boxes[box], func(l lens) bool { return l.label == label })

		if pos >= 0 && matches[2] == "-" {
			boxes[box] = slices.Delete(boxes[box], pos, pos+1)
			continue
		}

		if matches[2] == "=" {
			if pos >= 0 {
				boxes[box][pos] = lens{label: label, focal: util.Atoi(matches[3])}
				continue
			}
			boxes[box] = append(boxes[box], lens{label: label, focal: util.Atoi(matches[3])})
		}
	}

	for i, box := range boxes {
		for j, l := range box {
			sum += (i + 1) * (j + 1) * l.focal
		}
	}
	return sum
}
