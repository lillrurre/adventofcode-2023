package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
)

type propagator interface {
	sendPulse(in input) (next []input)
	sendPulseTwo(in input, rxMap map[string]int, rxDest string, count int) (next []input)
}

type pulse bool

const (
	low  pulse = false
	high pulse = true
)

type flipFlop struct {
	destinations []string
	on           bool
}

type consecutive struct {
	destinations []string
	memory       map[string]pulse
}

type broadcast struct {
	destinations []string
}

type input struct {
	pulse pulse
	src   string
	dest  string
}

func main() {
	util.Run(1, func() (sum int) {
		return part1(util.FileAsStringArr(20, "\n"))
	})
	util.Run(2, func() (sum int) {
		return part2(util.FileAsStringArr(20, "\n"))
	})
}

func part1(in []string) int {
	modules, _, _ := parse(in)
	pulses := make(map[pulse]int)
	for i := 0; i < 1000; i++ {
		sendPulse(pulses, modules, []input{{pulse: low, dest: "broadcaster"}})
	}
	return pulses[high] * pulses[low]
}

func part2(in []string) int {
	modules, rxMap, rxDest := parse(in)
	count := 1
	for slices.Contains(maps.Values(rxMap), 0) {
		sendPulseTwo(modules, []input{{pulse: low, dest: "broadcaster"}}, rxMap, rxDest, count)
		count++
	}
	return util.LCM(maps.Values(rxMap)...)
}

func parse(input []string) (modules map[string]propagator, rxMap map[string]int, rxDest string) {
	modules = make(map[string]propagator)
	rxMap = make(map[string]int)
	for _, line := range input {
		parts := strings.Split(line, " -> ")

		var p propagator
		var label string
		destinations := strings.Split(parts[1], ", ")

		switch parts[0][0] {
		case '%':
			label = parts[0][1:]
			p = &flipFlop{destinations: destinations, on: false}
		case '&':
			label = parts[0][1:]
			memory := make(map[string]pulse)
			for _, memLine := range input {
				if slices.Contains(strings.Split(strings.Split(memLine, " -> ")[1], ", "), label) {
					memory[strings.Split(memLine, " -> ")[0][1:]] = low
				}
			}
			p = &consecutive{destinations: destinations, memory: memory}
			if slices.Contains(destinations, "rx") {
				rxDest = label
				for k := range memory {
					rxMap[k] = 0
				}
			}
		default:
			label = parts[0]
			p = &broadcast{destinations: destinations}
		}

		modules[label] = p
	}

	return modules, rxMap, rxDest
}

func sendPulse(pulses map[pulse]int, modules map[string]propagator, inputs []input) {
	next := make([]input, 0)
	for _, in := range inputs {
		pulses[in.pulse]++
		if m, ok := modules[in.dest]; ok {
			next = append(next, m.sendPulse(in)...)
		}
	}
	if len(next) != 0 {
		sendPulse(pulses, modules, next)
	}
}

func sendPulseTwo(modules map[string]propagator, inputs []input, rxMap map[string]int, rxDest string, count int) {
	next := make([]input, 0)
	for _, in := range inputs {
		if m, ok := modules[in.dest]; ok {
			next = append(next, m.sendPulseTwo(in, rxMap, rxDest, count)...)
		}
	}
	if len(next) != 0 {
		sendPulseTwo(modules, next, rxMap, rxDest, count)
	}
}

func (f *flipFlop) sendPulse(in input) (next []input) {
	if in.pulse == high {
		return make([]input, 0)
	}

	f.on = !f.on

	for _, dest := range f.destinations {
		next = append(next, input{pulse: pulse(f.on), src: in.dest, dest: dest})
	}
	return next
}

func (c *consecutive) sendPulse(in input) (next []input) {
	c.memory[in.src] = in.pulse

	p := low
	for _, mem := range c.memory {
		if mem == low {
			p = high
			break
		}
	}

	for _, dest := range c.destinations {
		next = append(next, input{pulse: p, src: in.dest, dest: dest})
	}
	return next
}

func (b *broadcast) sendPulse(in input) (next []input) {
	for _, dest := range b.destinations {
		next = append(next, input{pulse: in.pulse, src: in.dest, dest: dest})
	}

	return next
}

func (c *consecutive) sendPulseTwo(in input, rxMap map[string]int, rxDest string, count int) (next []input) {
	c.memory[in.src] = in.pulse

	if in.dest == rxDest && rxMap[in.src] == 0 && in.pulse == high {
		rxMap[in.src] = count
	}

	p := low
	for _, mem := range c.memory {
		if mem == low {
			p = high
			break
		}
	}

	for _, dest := range c.destinations {
		next = append(next, input{pulse: p, src: in.dest, dest: dest})
	}
	return next
}

func (f *flipFlop) sendPulseTwo(in input, _ map[string]int, _ string, _ int) (next []input) {
	return f.sendPulse(in)
}

func (b *broadcast) sendPulseTwo(in input, _ map[string]int, _ string, _ int) (next []input) {
	return b.sendPulse(in)
}
