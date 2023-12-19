package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
)

type rule struct {
	part, operator uint8
	value          int
	next           string
}

type workflow []rule

func main() {
	input := util.FileAsStringArr(19, "\n\n")
	util.Run(1, func() (sum int) { return part1(input) })
	util.Run(2, func() (sum int) { return part2(input) })
}

func runWorkflow(w workflow, p map[uint8]int) (sum int, label string) {
	done := func(s string) (sum int, n string) {
		if s == "A" {
			return util.SliceSum(maps.Values(p)...), ""
		}
		return 0, ""
	}

	for _, r := range w {
		if r.part == 0 || r.operator == '>' && p[r.part] > r.value || r.operator == '<' && p[r.part] < r.value {
			if r.next == "A" || r.next == "R" {
				return done(r.next)
			}
			return 0, r.next
		}
	}

	return
}

func part1(input []string) (sum int) {
	workflows := parseWorkflows(input[0])
	for _, line := range strings.Fields(input[1]) {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return slices.Contains([]rune("{}=xmas,"), r)
		})
		p := make(map[uint8]int)
		for i, r := range []uint8("xmas") {
			p[r] = util.Atoi(fields[i])
		}
		label, res := "in", 0
		next := workflows[label]
		for next != nil {
			res, label = runWorkflow(next, p)
			next = workflows[label]
		}
		sum += res
	}
	return sum
}

type interval struct {
	low, high int
}

func countWorkflows(workflows map[string]workflow, label string, intervalMap map[uint8]*interval) (sum int) {
	// R, A -> the workflow path has come to an end, calculate sum and return
	switch label {
	case "R":
		return 0
	case "A":
		sum++
		for _, inter := range intervalMap {
			sum *= inter.high - inter.low + 1
		}
		return sum
	}

	for _, r := range workflows[label] {
		// No operator, so just use the same interval
		if r.part == 0 {
			sum += countWorkflows(workflows, r.next, intervalMap)
			continue
		}

		next, remainder := getIntervals(intervalMap[r.part], r.value, r.operator)

		// Count the workflows for the new interval with the new label
		// The map must be cloned because the next interval is not part of the current workflow
		if next.low < next.high {
			clonedIntervals := maps.Clone(intervalMap)
			clonedIntervals[r.part] = next
			sum += countWorkflows(workflows, r.next, clonedIntervals)
		}

		// The remainder that is not counted in the next could be accepted in the end of this interval or used by an else statement
		// Nice edge case 🔫
		intervalMap[r.part] = remainder
	}

	return sum
}

// getIntervals splits the interval into two.
// If the operator is <, the next to calculate is low < value. Else, the next is value < high.
func getIntervals(inter *interval, value int, operator uint8) (next *interval, remainder *interval) {
	if operator == '<' {
		return &interval{low: inter.low, high: value - 1}, &interval{low: value, high: inter.high}
	}
	return &interval{low: value + 1, high: inter.high}, &interval{low: inter.low, high: value}
}

func part2(input []string) (sum int) {
	workflows := parseWorkflows(input[0])
	valueMap := make(map[uint8]*interval)
	for _, r := range []uint8("xmas") {
		valueMap[r] = &interval{low: 1, high: 4000}
	}
	return countWorkflows(workflows, "in", valueMap)
}

func parseWorkflows(input string) map[string]workflow {
	workflows := make(map[string]workflow)
	for _, line := range strings.Fields(input) {
		w := make(workflow, 0)
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == '{' || r == '}' || r == ',' })
		label := fields[0]
		for _, field := range fields[1:] {
			if !strings.ContainsRune(field, ':') {
				w = append(w, rule{next: field})
				continue
			}
			colon := strings.FieldsFunc(field, func(r rune) bool { return r == ':' })
			w = append(w, rule{
				part:     field[0],
				operator: field[1],
				value:    util.Atoi(strings.FieldsFunc(colon[0], func(r rune) bool { return r == '<' || r == '>' })[1]),
				next:     colon[1],
			})
		}
		workflows[label] = w
	}
	return workflows
}
