package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"golang.org/x/exp/maps"
	"slices"
	"strings"
)

func main() {
	util.Run(1, func() (sum int) { return part1(util.FileAsStringArr(25, "\n")) })
}

func part1(input []string) (sum int) {
	var first, second, size int
	for size != 3 {
		first, second, size = cut(input)
	}
	return first * second
}

func cut(input []string) (first, second, size int) {
	components := parse(input)

	sizes := make(map[string]int)
	for k := range components {
		sizes[k] = 1
	}

	for len(components) > 2 {
		src := maps.Keys(components)[0]
		dest := components[src][1]

		for _, comp := range components[dest] {
			components[comp] = slices.DeleteFunc(components[comp], func(s string) bool { return s == dest })
			components[comp] = append(components[comp], src)
		}
		components[src] = slices.DeleteFunc(append(components[src], components[dest]...), func(s string) bool { return s == src })
		sizes[src] = sizes[src] + sizes[dest]
		delete(components, dest)
		delete(sizes, dest)
	}
	keys := maps.Keys(components)
	return sizes[keys[0]], sizes[keys[len(keys)-1]], len(components[keys[0]])
}

func parse(input []string) (components map[string][]string) {
	components = make(map[string][]string)

	addComponent := func(name, conn string) {
		if _, ok := components[name]; !ok {
			components[name] = make([]string, 0)
		}
		components[name] = append(components[name], conn)
	}

	for _, line := range input {
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ':' })
		for _, conn := range fields[1:] {
			addComponent(fields[0], conn)
			addComponent(conn, fields[0])
		}
	}

	return components
}
